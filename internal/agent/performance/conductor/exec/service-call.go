package conductorExec

import (
	"context"
	"github.com/aaronchen2k/deeptest/internal/agent/performance/conductor/dao"
	"github.com/aaronchen2k/deeptest/internal/agent/performance/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/agent/performance/pkg/domain"
	"github.com/aaronchen2k/deeptest/internal/agent/performance/pkg/log"
	"github.com/aaronchen2k/deeptest/internal/agent/performance/pkg/websocket"
	"github.com/aaronchen2k/deeptest/internal/agent/performance/proto"
	"github.com/aaronchen2k/deeptest/pkg/lib/log"
	_stringUtils "github.com/aaronchen2k/deeptest/pkg/lib/string"
	"github.com/facebookgo/inject"
	"github.com/goccy/go-json"
	"github.com/kataras/iris/v12/websocket"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"io"
	"log"
	"sync"
)

var (
	testInst *PerformanceTestService
)

func CreatePerformanceTestService() *PerformanceTestService {
	testInst = &PerformanceTestService{
		uuid: _stringUtils.Uuid(),
	}

	var g inject.Graph
	g.Logger = logrus.StandardLogger()

	if err := g.Provide(
		&inject.Object{Value: testInst},
	); err != nil {
		logrus.Fatalf("provide usecase objects to the Graph: %v", err)
	}

	err := g.Populate()
	if err != nil {
		logrus.Fatalf("populate the incomplete Objects: %v", err)
	}

	return testInst
}

type PerformanceTestService struct {
	uuid string

	TestReq      *ptdomain.PerformanceTestReq
	suspendWsMsg bool

	execCtx    context.Context
	execCancel context.CancelFunc

	logCtx    context.Context
	logCancel context.CancelFunc

	client *ptproto.PerformanceServiceClient

	PerformanceRemoteService *PerformanceRemoteService `inject:""`

	GrpcService     *GrpcService     `inject:"private"`
	ScheduleService *ScheduleService `inject:"private"`

	RemoteRunnerService *RemoteRunnerService `inject:"private"`
}

// performance test execution

func (s *PerformanceTestService) ExecStart(
	req ptdomain.PerformanceTestReq, wsMsg *websocket.Message) (err error) {

	// ignore if already a test is running
	if s.TestReq != nil {
		ptwebsocket.SendExecInstructionToClient(
			"performance testing is running on conductor", "", ptconsts.MsgInstructionAlreadyRunning, wsMsg)
		return
	}

	// cache request
	s.TestReq = &req

	// load test data from remote server
	data, err := s.PerformanceRemoteService.GetPlanToExec(req)
	if err != nil {
		return
	}

	data.Room = req.Room

	ptlog.Init(data.Room)

	if s.IsaRunnerBusy(data) {
		ptwebsocket.SendExecInstructionToClient(
			"performance testing is running on runner", "", ptconsts.MsgInstructionAlreadyRunning, wsMsg)
		return
	}

	item := AddTestItem(s.TestReq.Room, ptconsts.Conductor, s.TestReq, data.Runners, nil)

	ptwebsocket.SendExecInstructionToClient(
		"performance testing start", item, ptconsts.MsgInstructionStart, wsMsg)

	s.execCtx, s.execCancel = context.WithCancel(context.Background())

	err = dao.ClearData(data.Room)
	if err != nil {
		ptwebsocket.SendExecInstructionToClient(
			err.Error(), "", ptconsts.MsgInstructionException, wsMsg)
		return
	}

	err = dao.ResetInfluxdb(data.Room, data.InfluxdbAddress, data.InfluxdbOrg, data.InfluxdbToken)
	if err != nil {
		ptwebsocket.SendExecInstructionToClient(
			err.Error(), "", ptconsts.MsgInstructionException, wsMsg)
		return
	}

	s.GrpcService.ConductorClearAllGlobalVar(context.Background(), &ptproto.GlobalVarRequest{})

	// stop execution in 2 ways:
	// 1. call cancel in this method by websocket request OR after all runners completed
	// 2. sub cancel instruction from runner via grpc

	go s.ScheduleService.SendMetricsToClient(s.execCtx, s.execCancel, data, wsMsg)

	// start execution
	go func() {
		var wgRunners sync.WaitGroup
		for _, runner := range data.Runners {
			client := s.ConnectGrpc(runner)

			stream, err := s.CallRunnerExecStartByGrpc(client, data,
				runner.Id, runner.Name, runner.Weight, int32(req.EnvironmentId))
			if err != nil {
				ptlog.Logf("failed to call remote runner via grpc, err %s", err.Error())
				continue
			}

			wgRunners.Add(1)

			go func() {
				defer wgRunners.Done()

				s.handleGrpcMsg(stream)
				stream.CloseSend()
			}()
		}

		wgRunners.Wait()

		//s.StopAndClearScene(data.Room, wsMsg)
	}()

	return
}

func (s *PerformanceTestService) ExecStop(wsMsg *websocket.Message) (err error) {
	if s.TestReq == nil {
		return
	}

	// load test data from remote server
	data, err := s.PerformanceRemoteService.GetPlanToExec(*s.TestReq)
	if err != nil {
		return
	}

	// call remote runners to stop
	s.RemoteRunnerService.CallStop(data)

	s.StopAndClearScene(s.TestReq.Room, wsMsg)

	return
}

func (s *PerformanceTestService) StopAndClearScene(room string, wsMsg *websocket.Message) (err error) {
	// close exec and send msg
	if s.execCancel != nil {
		s.execCancel()
	}
	ptwebsocket.SendExecInstructionToClient("", "", ptconsts.MsgInstructionEnd, wsMsg)

	// remove from cache
	s.TestReq = nil
	RemoveTestItem(room)
	DeleteTestService(room)

	return
}

// call grpc client
func (s *PerformanceTestService) ConnectGrpc(runner *ptdomain.Runner) (client ptproto.PerformanceServiceClient) {
	connect, err := grpc.Dial(runner.GrpcAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}

	client = ptproto.NewPerformanceServiceClient(connect)

	return
}

func (s *PerformanceTestService) CallRunnerExecStartByGrpc(
	client ptproto.PerformanceServiceClient, req ptdomain.PerformanceTestData, runnerId int32, runnerName string, weight, environmentId int32) (stream ptproto.PerformanceService_RunnerExecStartClient, err error) {

	stream, err = client.RunnerExecStart(context.Background())
	if err != nil {
		ptlog.Logf(err.Error())
		return
	}

	runnerExecScenarios := s.getRunnerExecScenarios(req, runnerId)
	localVarsCacheRaw, _ := json.Marshal(req.LocalVarsCache)
	execSceneRaw, _ := json.Marshal(req.ExecScene)

	err = stream.Send(&ptproto.PerformanceExecStartReq{
		Room:       req.Room,
		RunnerId:   runnerId,
		RunnerName: runnerName,
		Title:      req.Title,

		Mode:              req.Mode.String(),
		Scenarios:         runnerExecScenarios,
		Weight:            weight,
		EnvironmentId:     environmentId,
		LocalVarsCacheRaw: localVarsCacheRaw,
		ExecSceneRaw:      execSceneRaw,

		WebServerUrl:   req.ServerUrl,
		WebServerToken: req.Token,

		ConductorGrpcAddress: req.ConductorGrpcAddress,
		InfluxdbAddress:      req.InfluxdbAddress,
		InfluxdbOrg:          req.InfluxdbOrg,
		InfluxdbToken:        req.InfluxdbToken,
	})

	if err != nil {
		return
	}

	return
}

func (s *PerformanceTestService) handleGrpcMsg(stream ptproto.PerformanceService_RunnerExecStartClient) (err error) {
	for true {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			continue
		}

		ptlog.Logf("get grpc msg from runner: %v", resp)

		// dealwith Instruction from agent
		if resp.Instruction == ptconsts.MsgInstructionRunnerFinish.String() {
			break
		}

		select {
		case <-s.execCtx.Done():
			_logUtils.Debug("<<<<<<< stop sendLog job")
			break

		default:
		}
	}

	return
}

// helper methods
func (s *PerformanceTestService) getRunnerExecScenarios(req ptdomain.PerformanceTestData, runnerId int32) (
	ret []*ptproto.Scenario) {

	notSet := false
	scenarioIdsMap := map[int32]bool{}
	for _, runner := range req.Runners {
		if runner.Id != runnerId {
			continue
		}

		if runner.Scenarios == nil {
			notSet = true

			break
		}

		for _, scId := range runner.Scenarios {
			scenarioIdsMap[scId] = true
		}

		break
	}

	for _, scenario := range req.Scenarios {
		if notSet || scenarioIdsMap[scenario.Id] {
			ret = append(ret, scenario)
		}
	}

	return
}

func (s *PerformanceTestService) IsaRunnerBusy(data ptdomain.PerformanceTestData) (ret bool) {
	for _, runner := range data.Runners {
		client := s.ConnectGrpc(runner)

		isBusy, _ := client.RunnerIsBusy(context.Background(), &wrapperspb.StringValue{Value: data.Room})
		if isBusy.Value {
			ret = isBusy.Value
			return
		}

	}

	return
}

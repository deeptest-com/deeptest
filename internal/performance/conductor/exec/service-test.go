package conductorExec

import (
	"context"
	"github.com/aaronchen2k/deeptest/internal/performance/conductor/dao"
	"github.com/aaronchen2k/deeptest/internal/performance/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/performance/pkg/domain"
	"github.com/aaronchen2k/deeptest/internal/performance/pkg/log"
	"github.com/aaronchen2k/deeptest/internal/performance/pkg/websocket"
	"github.com/aaronchen2k/deeptest/internal/performance/proto"
	"github.com/aaronchen2k/deeptest/pkg/lib/log"
	_stringUtils "github.com/aaronchen2k/deeptest/pkg/lib/string"
	"github.com/facebookgo/inject"
	"github.com/kataras/iris/v12/websocket"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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

	GrpcService         *GrpcService         `inject:"private"`
	ScheduleService     *ScheduleService     `inject:"private"`
	RemoteRunnerService *RemoteRunnerService `inject:"private"`
}

// performance test execution

func (s *PerformanceTestService) ExecStart(
	req ptdomain.PerformanceTestReq, wsMsg *websocket.Message) (err error) {

	if s.TestReq != nil {
		ptwebsocket.SendExecInstructionToClient(
			"performance testing is running on conductor", "", ptconsts.MsgInstructionAlreadyRunning, wsMsg)
		return
	}

	ptwebsocket.SendExecInstructionToClient(
		"performance testing start", nil, ptconsts.MsgInstructionStart, wsMsg)

	ptlog.Init(req.Room)

	s.TestReq = &req

	AddTestItem(s.TestReq.Room, ptconsts.Conductor, s.TestReq, nil)

	// start execution
	go func() {
		s.execCtx, s.execCancel = context.WithCancel(context.Background())

		dao.ClearData(req.Room)
		dao.ResetInfluxdb(req.Room, req.InfluxdbAddress, req.InfluxdbOrg, req.InfluxdbToken)
		s.GrpcService.ConductorClearAllGlobalVar(context.Background(), &ptproto.GlobalVarRequest{})

		// stop execution in 2 ways:
		// 1. call cancel in this method by websocket request OR after all runners completed
		// 2. sub cancel instruction from runner via grpc

		go s.ScheduleService.SendMetricsToClient(s.execCtx, s.execCancel, req, wsMsg)

		var wgRunners sync.WaitGroup
		for _, runner := range req.Runners {
			client := s.ConnectGrpc(runner)

			stream, err := s.CallRunnerExecStartByGrpc(client, req, runner.Id, runner.Name, runner.Weight)
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

		// wait all runners finished
		wgRunners.Wait()

		s.StopAndClearScene(req.Room, wsMsg)
	}()

	return
}

func (s *PerformanceTestService) ExecStop(wsMsg *websocket.Message) (err error) {
	if s.TestReq == nil {
		return
	}

	// call remote runners to stop
	s.RemoteRunnerService.CallStop(*s.TestReq)

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
func (s *PerformanceTestService) ConnectGrpc(runner *ptproto.Runner) (client ptproto.PerformanceServiceClient) {
	connect, err := grpc.Dial(runner.GrpcAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}

	client = ptproto.NewPerformanceServiceClient(connect)

	return
}

func (s *PerformanceTestService) CallRunnerExecStartByGrpc(
	client ptproto.PerformanceServiceClient, req ptdomain.PerformanceTestReq, runnerId int32, runnerName string, weight int32) (stream ptproto.PerformanceService_RunnerExecStartClient, err error) {

	stream, err = client.RunnerExecStart(context.Background())
	if err != nil {
		ptlog.Logf(err.Error())
		return
	}

	runnerExecScenarios := s.getRunnerExecScenarios(req, runnerId)

	err = stream.Send(&ptproto.PerformanceExecStartReq{
		Room:       req.Room,
		RunnerId:   runnerId,
		RunnerName: runnerName,
		Title:      req.Title,

		Mode:      req.Mode.String(),
		Scenarios: runnerExecScenarios,
		Weight:    weight,

		ServerAddress:   req.ServerAddress,
		InfluxdbAddress: req.InfluxdbAddress,
		InfluxdbOrg:     req.InfluxdbOrg,
		InfluxdbToken:   req.InfluxdbToken,
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
func (s *PerformanceTestService) getRunnerExecScenarios(req ptdomain.PerformanceTestReq, runnerId int32) (
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

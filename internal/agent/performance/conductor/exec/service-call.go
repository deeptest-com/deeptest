package conductorExec

import (
	"context"
	"errors"
	"fmt"
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
	"io"
	"net/url"
	"strings"
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

	suspendWsMsg bool

	execCtx    context.Context
	execCancel context.CancelFunc

	logCtx    context.Context
	logCancel context.CancelFunc

	client *ptproto.PerformanceServiceClient

	GrpcService     *GrpcService     `inject:"private"`
	ScheduleService *ScheduleService `inject:"private"`

	RemoteRunnerService *RemoteRunnerService `inject:"private"`
}

// performance test execution

func (s *PerformanceTestService) ExecStart(
	req ptdomain.PerformanceTestReq, wsMsg *websocket.Message) (err error) {

	conductorTask := GetConductorTask()

	// ignore if already a test is running
	if conductorTask != nil {
		ptwebsocket.SendExecInstructionToClient(
			"主控端有正在执行的性能测试", "", ptconsts.MsgInstructionAlreadyRunning, wsMsg)

		err = errors.New("主控端有正在执行的性能测试")

		RemoveTestTask(ptconsts.Conductor)
		DeleteTestService(req.Room)

		return
	}

	// load test data from remote server
	data, err := GetPlanToExec(req)
	if err != nil {
		return
	}

	data.Room = req.Room

	ptlog.Init(data.Room)

	if s.IsRunnerBusyWithGrpcAddressUpdated(&data) {
		ptwebsocket.SendExecInstructionToClient(
			"代理端有正在执行的性能测试", "", ptconsts.MsgInstructionAlreadyRunning, wsMsg)
		return
	}

	conductorTask = AddTestItem(req.Room, ptconsts.Conductor, &req, data.Runners, nil)

	ptwebsocket.SendExecInstructionToClient(
		"performance testing start", conductorTask, ptconsts.MsgInstructionStart, wsMsg)

	s.execCtx, s.execCancel = context.WithCancel(context.Background())

	err = dao.ClearData(data.Room)
	if err != nil {
		ptwebsocket.SendExecInstructionToClient(
			err.Error(), "dao.ClearData", ptconsts.MsgInstructionException, wsMsg)
		return
	}

	err = dao.ResetInfluxdb(data.Room, data.InfluxdbAddress, data.InfluxdbOrg, data.InfluxdbToken)
	if err != nil {
		ptwebsocket.SendExecInstructionToClient(
			err.Error(), "dao.ResetInfluxdb", ptconsts.MsgInstructionException, wsMsg)
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

			// async handle runner grpc msg
			go func() {
				defer wgRunners.Done()

				s.handleGrpcMsg(s.execCtx, stream, runner.Id, runner.Name) // finish loop in it if got a runnerFinish instruction

				stream.CloseSend()
			}()
		}

		// wait all async remote runner executions completed/canceled
		wgRunners.Wait()

		RemoveTestTask(ptconsts.Conductor)
		DeleteTestService(req.Room)

		ptlog.Logf("condutor: all runner stopped")
		ptwebsocket.SendExecInstructionToClient("", "", ptconsts.MsgInstructionEnd, wsMsg)
	}()

	return
}

func (s *PerformanceTestService) ExecStop(wsMsg *websocket.Message) (err error) {
	conductorTask := GetConductorTask()

	if conductorTask == nil {
		return
	}

	// call remote runners to stop
	s.RemoteRunnerService.CallStop(conductorTask.Room, conductorTask.Runners)

	// close exec and send msg
	if s.execCancel != nil {
		s.execCancel()
	}
	ptwebsocket.SendExecInstructionToClient("", "", ptconsts.MsgInstructionEnd, wsMsg)

	// remove from cache
	RemoveTestTask(ptconsts.Conductor)
	DeleteTestService(conductorTask.Room)

	return
}

// call grpc client
func (s *PerformanceTestService) ConnectGrpc(runner *ptdomain.Runner) (client ptproto.PerformanceServiceClient) {
	client = GetGrpcClient(runner.GrpcAddress)

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

		GoalLoop:     int32(req.Goal.Loop),
		GoalDuration: int32(req.Goal.Duration),
		// other goals will be controlled on conductor side

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

func (s *PerformanceTestService) handleGrpcMsg(ctx context.Context, stream ptproto.PerformanceService_RunnerExecStartClient, runnerId int32, runnerName string) (err error) {
	for true {
		// will hold-on to wait agent's msg
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			continue
		}

		ptlog.Logf("get grpc msg from runner %d-%s: %v", runnerId, runnerName, resp)

		// dealwith Instruction from agent
		if resp.Instruction == ptconsts.MsgInstructionRunnerFinish.String() {
			break

		} else if resp.Instruction == ptconsts.MsgInstructionConductorFinish.String() {
			ctx.Done()
			ptlog.Logf("stop conductor whole execution by runner %d-%s", runnerId, runnerName)

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

func (s *PerformanceTestService) IsRunnerBusyWithGrpcAddressUpdated(data *ptdomain.PerformanceTestData) (ret bool) {
	for index, runner := range data.Runners {
		mp, err := GetRunnerState(runner)
		if err != nil {
			return false
		}

		if mp["grpcPort"] != nil {
			u, _ := url.Parse(runner.WebAddress)
			arr := strings.Split(u.Host, ":")
			host := arr[0]

			data.Runners[index].GrpcAddress = fmt.Sprintf("%s:%s", host, mp["grpcPort"].(string))
		}

		if mp["isBusy"] != nil {
			isBusy := mp["isBusy"].(bool)
			if isBusy {
				ret = isBusy
				return
			}
		}
	}

	return
}

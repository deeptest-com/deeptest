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
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/websocket"
	"github.com/nxadm/tail"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"sync"
	"time"
)

var (
	once sync.Once
	inst *PerformanceTestService
)

func CreatePerformanceTestService() *PerformanceTestService {
	inst = &PerformanceTestService{
		uuid: _stringUtils.Uuid(),
	}

	var g inject.Graph
	g.Logger = logrus.StandardLogger()

	if err := g.Provide(
		&inject.Object{Value: inst},
	); err != nil {
		logrus.Fatalf("provide usecase objects to the Graph: %v", err)
	}

	err := g.Populate()
	if err != nil {
		logrus.Fatalf("populate the incomplete Objects: %v", err)
	}

	return inst
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

		// remove from cache
		RemoveTestItem(req.Room)
		DeleteTestService(req.Room)

		// close exec and send to web client
		s.execCancel()
		ptwebsocket.SendExecInstructionToClient("", "", ptconsts.MsgInstructionEnd, wsMsg)
	}()

	return
}

func (s *PerformanceTestService) ExecStop(wsMsg *websocket.Message) (err error) {
	if s.TestReq == nil {
		return
	}

	// stop server execution
	if s.execCancel != nil {
		s.execCancel()
	}

	// exec by runners
	s.RemoteRunnerService.CallStop(*s.TestReq)

	// send end msg to websocket client
	ptwebsocket.SendExecInstructionToClient("", "", ptconsts.MsgInstructionEnd, wsMsg)

	s.TestReq = nil

	return
}

func (s *PerformanceTestService) StartSendLog(req ptdomain.PerformanceTestReq, wsMsg *websocket.Message) (err error) {
	if s.logCtx != nil {
		return
	}

	s.logCtx, s.logCancel = context.WithCancel(context.Background())

	room := req.Room
	logPath := ptlog.GetLogPath(room)

	t, err := tail.TailFile(logPath, tail.Config{Follow: true, ReOpen: true})
	if err != nil {
		s.logCancel()
		s.logCtx = nil
		return
	}
	defer t.Cleanup()

	go func() {
		var arr []string

		for line := range t.Lines {
			arr = append(arr, line.Text)

			if len(arr) > 100 {
				data := iris.Map{
					"log": line.Text,
				}
				ptwebsocket.SendExecResultToClient(data, ptconsts.MsgResultRecord, req.Room, wsMsg)

				arr = make([]string, 0)
			}
		}

		s.logCancel()
		s.logCtx = nil
	}()

	go func() {
		for true {
			if s.logCtx == nil || IsWsMsgSuspend() {
				t.Stop()
				break
			}

			select {
			case <-s.logCtx.Done():
				_logUtils.Debug("<<<<<<< stop sendLog job by logCtx.Done")

				t.Stop()
				break

			default:
			}

			select {
			case <-s.execCtx.Done():
				_logUtils.Debug("<<<<<<< stop sendLog job by execCtx.Done")

				t.Stop()
				break

			default:
			}

			time.Sleep(3 * time.Second)
		}

		s.logCancel()
		s.logCtx = nil
	}()

	return
}

func (s *PerformanceTestService) StopSendLog() (err error) {
	s.logCancel()
	s.logCtx = nil

	return
}

// call grpc client
func (s *PerformanceTestService) ConnectGrpc(runner *ptproto.Runner) (client ptproto.PerformanceServiceClient) {
	connect, err := grpc.Dial(runner.Address, grpc.WithTransportCredentials(insecure.NewCredentials()))
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

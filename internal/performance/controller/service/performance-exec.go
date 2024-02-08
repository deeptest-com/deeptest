package controllerService

import (
	"context"
	"github.com/aaronchen2k/deeptest/internal/performance/controller/dao"
	"github.com/aaronchen2k/deeptest/internal/performance/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/performance/pkg/domain"
	websocketHelper "github.com/aaronchen2k/deeptest/internal/performance/pkg/websocket"
	ptProto "github.com/aaronchen2k/deeptest/internal/performance/proto"
	"github.com/facebookgo/inject"
	"github.com/kataras/iris/v12/websocket"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type PerformanceTestService struct {
	execCtx    context.Context
	execCancel context.CancelFunc
	client     *ptProto.PerformanceServiceClient

	GrpcService         *GrpcService         `inject:"private"`
	ScheduleService     *ScheduleService     `inject:"private"`
	RemoteRunnerService *RemoteRunnerService `inject:"private"`
}

func NewPerformanceTestService() *PerformanceTestService {
	service := &PerformanceTestService{}

	var g inject.Graph
	g.Logger = logrus.StandardLogger()

	if err := g.Provide(
		&inject.Object{Value: service},
	); err != nil {
		logrus.Fatalf("provide usecase objects to the Graph: %v", err)
	}
	err := g.Populate()
	if err != nil {
		logrus.Fatalf("populate the incomplete Objects: %v", err)
	}
	return service
}

func (s *PerformanceTestService) Connect(runner *ptProto.Runner) (client ptProto.PerformanceServiceClient) {
	connect, err := grpc.Dial(runner.Address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}

	client = ptProto.NewPerformanceServiceClient(connect)

	return
}

func (s *PerformanceTestService) ExecStart(req ptdomain.PerformanceTestReq, wsMsg *websocket.Message) (err error) {
	s.execCtx, s.execCancel = context.WithCancel(context.Background())

	dao.ClearData(req.Room)
	s.ScheduleService.Reset(req.Scenarios)
	s.GrpcService.ClearAllRendezvousVal(context.Background(), &ptProto.GlobalVarRequest{})

	// stop execution in 2 ways:
	// 1. call cancel in this method by websocket request
	// 2. sub cancel instruction from runner via grpc

	if req.NsqServerAddress != "" { // runner send logs via nsq MQ
		// check ctx.Done to stop
		go s.HandleRunnerNsqMsg(s.execCtx, req.Room, req.NsqLookupAddress, req.NsqServerAddress, wsMsg)

	} else { // runner send logs via grpc, server store msgs in queue
		// check ctx.Done to stop
		// cancel ctx by instruction from runner
		go s.HandleRunnerGrpcMsg(s.execCtx, s.execCancel, req.Room, wsMsg)
	}

	// independent job to summary metrics and send web client
	go s.ScheduleService.ScheduleJob(s.execCtx, s.execCancel, req, wsMsg)

	// exec by runners
	for _, runner := range req.Runners {
		client := s.Connect(runner)

		stream, err := s.CallRunnerExecStartByGrpc(client, req, runner.Id, runner.Weight)
		if err != nil {
			continue
		}

		s.HandleAndPubToQueueGrpcMsg(stream) // sync
		stream.CloseSend()
	}

	// send cancel signal
	s.execCancel()

	websocketHelper.SendExecInstructionToClient("", "", ptconsts.MsgInstructionEnd, req.Room, wsMsg)

	return
}

func (s *PerformanceTestService) ExecStop(req ptdomain.PerformanceTestReq, wsMsg *websocket.Message) (err error) {
	// stop server execution
	if s.execCancel != nil {
		s.execCancel()
	}

	// exec by runners
	s.RemoteRunnerService.CallStop(req)

	// send end msg to websocket client
	websocketHelper.SendExecInstructionToClient("", "", ptconsts.MsgInstructionEnd, req.Room, wsMsg)

	return
}

//func (s *PerformanceTestService) ExecReconnectListenMsg(req ptdomain.PerformanceTestReq, wsMsg *websocket.Message) (err error) {
//	s.execCtx, s.execCancel = context.WithCancel(context.Background())
//
//	if req.NsqServerAddress != "" {
//		go s.HandleRunnerNsqInstructionMsg(req, s.execCtx, req.Uuid, wsMsg)
//
//	} else {
//		go ptqueue.SubRunnerGrpcMsg(s.retrieveAndDealwithResult, s.execCtx, s.execCancel, req.Uuid, wsMsg)
//	}
//
//	return
//}

func (s *PerformanceTestService) CallRunnerExecStartByGrpc(
	client ptProto.PerformanceServiceClient, req ptdomain.PerformanceTestReq, runnerId, weight int32) (stream ptProto.PerformanceService_ExecStartClient, err error) {

	stream, err = client.ExecStart(context.Background())
	if err != nil {
		return
	}

	runnerExecScenarios := s.getRunnerExecScenarios(req, runnerId)

	err = stream.Send(&ptProto.PerformanceExecStartReq{
		Room:     req.Room,
		RunnerId: runnerId,
		Title:    req.Title,

		Mode:      req.Mode.String(),
		Scenarios: runnerExecScenarios,
		Weight:    weight,

		ServerAddress:    req.ServerAddress,
		NsqServerAddress: req.NsqServerAddress,
		NsqLookupAddress: req.NsqLookupAddress,
	})

	if err != nil {
		return
	}

	return
}

func (s *PerformanceTestService) getRunnerExecScenarios(req ptdomain.PerformanceTestReq, runnerId int32) (
	ret []*ptProto.Scenario) {

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

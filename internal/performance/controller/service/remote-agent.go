package controllerService

import (
	"context"
	"github.com/aaronchen2k/deeptest/internal/performance/pkg/domain"
	ptProto "github.com/aaronchen2k/deeptest/internal/performance/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type RemoteRunnerService struct {
	client *ptProto.PerformanceServiceClient
}

func (s *RemoteRunnerService) Connect(runner *ptProto.Runner) (client ptProto.PerformanceServiceClient) {
	connect, err := grpc.Dial(runner.Address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}

	client = ptProto.NewPerformanceServiceClient(connect)

	return
}

func (s *RemoteRunnerService) CallStop(req ptdomain.PerformanceTestReq) (err error) {
	for _, runner := range req.Runners {
		client := s.Connect(runner)

		stream, err := s.CallRunnerExecStopByGrpc(client, req.Room)
		if err != nil {
			continue
		}

		stream.CloseSend()
	}

	return
}

func (s *RemoteRunnerService) CallRunnerExecStopByGrpc(
	client ptProto.PerformanceServiceClient, room string) (
	stream ptProto.PerformanceService_ExecStopClient, err error) {

	stream, err = client.ExecStop(context.Background())
	if err != nil {
		return
	}

	err = stream.Send(&ptProto.PerformanceExecStopReq{
		Room: room,
	})

	return
}

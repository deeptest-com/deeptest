package conductorExec

import (
	"context"
	"github.com/aaronchen2k/deeptest/internal/agent/performance/pkg/domain"
	ptProto "github.com/aaronchen2k/deeptest/internal/agent/performance/proto"
)

type RemoteRunnerService struct {
	client *ptProto.PerformanceServiceClient
}

func (s *RemoteRunnerService) Connect(runner *ptdomain.Runner) (client ptProto.PerformanceServiceClient) {
	client = GetGrpcClient(runner.GrpcAddress)

	return
}

func (s *RemoteRunnerService) CallStop(room string, runners []*ptdomain.Runner) (err error) {
	for _, runner := range runners {
		client := s.Connect(runner)

		stream, err := s.callRunnerExecStopByGrpc(client, room)
		if err != nil {
			continue
		}

		stream.CloseSend()
	}

	return
}

func (s *RemoteRunnerService) callRunnerExecStopByGrpc(
	client ptProto.PerformanceServiceClient, room string) (
	stream ptProto.PerformanceService_RunnerExecStopClient, err error) {

	stream, err = client.RunnerExecStop(context.Background())
	if err != nil {
		return
	}

	err = stream.Send(&ptProto.PerformanceExecStopReq{
		Room: room,
	})

	return
}

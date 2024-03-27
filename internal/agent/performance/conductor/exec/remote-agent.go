package conductorExec

import (
	"context"
	"github.com/aaronchen2k/deeptest/internal/agent/performance/pkg/domain"
	ptlog "github.com/aaronchen2k/deeptest/internal/agent/performance/pkg/log"
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

		err := CallRunnerExecStopByGrpc(client, room)
		if err != nil {
			continue
		}

	}

	return
}

func CallRunnerExecStopByGrpc(
	client ptProto.PerformanceServiceClient, room string) (err error) {

	stream, err := client.RunnerExecStop(context.Background())
	if err != nil || stream == nil {
		ptlog.Logf("failed to get grpc stream of remote runner, err %s", err.Error())
		return
	}

	err = stream.Send(&ptProto.PerformanceExecStopReq{
		Room: room,
	})
	if err != nil {
		ptlog.Logf("failed to call STOP on remote runner via grpc, err %s", err.Error())
	}

	stream.CloseSend()

	return
}

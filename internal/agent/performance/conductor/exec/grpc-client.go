package conductorExec

import (
	ptProto "github.com/aaronchen2k/deeptest/internal/agent/performance/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func GetGrpcClient(address string) (client ptProto.PerformanceServiceClient) {
	connect, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}

	client = ptProto.NewPerformanceServiceClient(connect)

	return
}

package performance

import (
	controllerService "github.com/aaronchen2k/deeptest/internal/agent/performance/conductor/exec"
	ptProto "github.com/aaronchen2k/deeptest/internal/agent/performance/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

func StartGrpcServe() {
	server := grpc.NewServer()
	ptProto.RegisterPerformanceServiceServer(server, &controllerService.GrpcService{})

	lis, err := net.Listen("tcp", ":9528")
	if err != nil {
		log.Fatalf("grpc net.Listen err: %v", err)
	}
	server.Serve(lis)
}

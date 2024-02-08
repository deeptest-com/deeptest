package exec

import (
	"context"
	ptProto "github.com/aaronchen2k/deeptest/internal/performance/proto"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

var (
	serverAddress string

	conn   *grpc.ClientConn
	client ptProto.PerformanceServiceClient
)

func getConn(address string) (client ptProto.PerformanceServiceClient) {
	if conn == nil || client == nil || conn.GetState() != connectivity.Ready || address != serverAddress {
		var err error
		conn, err = grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))

		if err != nil {
			return
		}

		client = ptProto.NewPerformanceServiceClient(conn)

		serverAddress = address
	}

	return
}

func IncreaseRemoteVal(room, name, serverAddress string) (ret *wrapperspb.Int32Value) {
	req := ptProto.GlobalVarRequest{
		Room: room,
		Name: name,
	}

	client := getConn(serverAddress)

	ret, err := client.AddRendezvousVal(context.Background(), &req)
	if err != nil {
		logUtils.Debug(err.Error())

		ret = &wrapperspb.Int32Value{}
		ret.Value = 0

		return
	}

	return
}

func GetRemoteVal(room, name, serverAddress string) (ret *wrapperspb.Int32Value) {
	req := ptProto.GlobalVarRequest{
		Room: room,
		Name: name,
	}

	client := getConn(serverAddress)

	ret, err := client.GetRendezvousVal(context.Background(), &req)
	if err != nil {
		logUtils.Debug(err.Error())

		ret = &wrapperspb.Int32Value{}
		ret.Value = 0

		return
	}

	return
}

func ResetRemoteVal(room, name, serverAddress string) (ret *wrapperspb.BoolValue) {
	req := ptProto.GlobalVarRequest{
		Room: room,
		Name: name,
	}

	client := getConn(serverAddress)

	ret, err := client.ClearRendezvousVal(context.Background(), &req)
	if err != nil {
		logUtils.Debug(err.Error())

		ret = &wrapperspb.BoolValue{}

		return
	}

	return
}

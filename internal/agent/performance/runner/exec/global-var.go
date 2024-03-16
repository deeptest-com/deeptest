package runnerExec

import (
	"context"
	ptlog "github.com/aaronchen2k/deeptest/internal/agent/performance/pkg/log"
	"github.com/aaronchen2k/deeptest/internal/agent/performance/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

var (
	serverAddress string

	conn *grpc.ClientConn
)

func getConn(address string) (client ptproto.PerformanceServiceClient) {
	if conn == nil || client == nil || conn.GetState() != connectivity.Ready || address != serverAddress {
		var err error
		conn, err = grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))

		if err != nil {
			return
		}

		client = ptproto.NewPerformanceServiceClient(conn)

		serverAddress = address
	}

	return
}

func AddRemoteVal(room, name, serverAddress string) (ret *wrapperspb.Int32Value) {
	req := ptproto.GlobalVarRequest{
		Room: room,
		Name: name,
	}

	client := getConn(serverAddress)

	ret, err := client.ConductorAddGlobalVar(context.Background(), &req)
	if err != nil {
		ptlog.Logf(err.Error())

		ret = &wrapperspb.Int32Value{}
		ret.Value = 0

		return
	}

	return
}

func GetRemoteVal(room, name, serverAddress string) (ret *wrapperspb.Int32Value) {
	req := ptproto.GlobalVarRequest{
		Room: room,
		Name: name,
	}

	client := getConn(serverAddress)

	ret, err := client.ConductorGetGlobalVar(context.Background(), &req)
	if err != nil {
		ptlog.Logf(err.Error())

		ret = &wrapperspb.Int32Value{}
		ret.Value = 0

		return
	}

	return
}

func ResetRemoteVal(room, name, serverAddress string) (ret *wrapperspb.BoolValue) {
	req := ptproto.GlobalVarRequest{
		Room: room,
		Name: name,
	}

	client := getConn(serverAddress)

	ret, err := client.ConductorClearGlobalVar(context.Background(), &req)
	if err != nil {
		ptlog.Logf(err.Error())

		ret = &wrapperspb.BoolValue{}

		return
	}

	return
}

package rpc

import (
	"context"
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/agent/modules/arith"
	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/log"
	"testing"
)

func TestTcpClient(t *testing.T) {
	url := fmt.Sprintf("tcp@127.0.0.1:%d", 8086)
	d, _ := client.NewPeer2PeerDiscovery(url, "")

	xClient := client.NewXClient("arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xClient.Close()

	args := &arith.Request{
		A: 1,
		B: 2,
	}

	reply := &arith.Response{}

	err := xClient.Call(context.Background(), "Add", args, reply)
	if err != nil {
		log.Errorf("failed to call: %v", err)
	}

	log.Infof("%d + %d = %d", args.A, args.B, reply.C)
}

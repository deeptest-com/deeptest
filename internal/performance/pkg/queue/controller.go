package ptqueue

import (
	"context"
	"fmt"
	ptlog "github.com/aaronchen2k/deeptest/internal/performance/pkg/log"
	ptProto "github.com/aaronchen2k/deeptest/internal/performance/proto"
	"github.com/aaronchen2k/deeptest/pkg/core/mq"
	"github.com/kataras/iris/v12/websocket"
	"time"
)

var (
	queueTopicOfServer  = "QUEUE_TOPIC_OF_SERVER"
	queueClientOfServer *mq.Client
)

func InitControllerQueue() {
	queueClientOfServer = mq.NewClient()
	//defer queueClientOfServer.Close()
	queueClientOfServer.SetConditions(10000000)
}

func PubRunnerGrpcMsg(data ptProto.PerformanceExecResp) {
	err := queueClientOfServer.Publish(queueTopicOfServer, data)
	if err != nil {
		fmt.Println("pub mq message failed", err)
	}
}

func SubRunnerGrpcMsg(callback func(ptProto.PerformanceExecResp, string, *websocket.Message) error,
	ctx context.Context, cancel context.CancelFunc, execUuid string, wsMsg *websocket.Message) {

	ch, err := queueClientOfServer.Subscribe(queueTopicOfServer)
	if err != nil {
		fmt.Printf("sub mq topic %s failed, err: %s\n", queueTopicOfServer, err.Error())
		return
	}

	count := 0
	for {
		msg := queueClientOfServer.GetPayLoad(ch).(ptProto.PerformanceExecResp)
		fmt.Printf("get queue msg [%s]%s\n", queueTopicOfServer, msg.Instruction)

		callback(msg, execUuid, wsMsg)

		count += len(msg.Requests)
		ptlog.Logf("****** SERVER DEBUG: totally %d requests sub from queue and insert to sqlite", count)

		select {
		case <-ctx.Done():
			goto Label_END_SUB_RUNNER_GRPC_MSG

		default:
		}

		time.Sleep(time.Millisecond * 100)
	}

Label_END_SUB_RUNNER_GRPC_MSG:
	queueClientOfServer.Unsubscribe(queueTopicOfServer, ch)

	return
}

package ptwebsocket

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/pkg/core/mq"
	"github.com/aaronchen2k/deeptest/pkg/domain"
	"time"
)

var (
	mqTopicLog  = "MQ_TOPIC_TEST"
	mqClientLog *mq.Client
)

func InitLogMq() {
	mqClientLog = mq.NewClient()
	//defer mqClientLog.Close()

	mqClientLog.SetConditions(10000)

	go SubLogMsg()
}

func SubLogMsg() {
	ch, err := mqClientLog.Subscribe(mqTopicLog)
	if err != nil {
		fmt.Printf("sub test mq topic %s failed\n", mqTopicLog)
		return
	}

	for {
		msg := mqClientLog.GetPayLoad(ch).(_domain.MqMsg)
		fmt.Printf("%s get test mq msg '%#v'\n", mqTopicLog, msg.Content)

		if msg.Content == "exit" {
			mqClientLog.Unsubscribe(mqTopicLog, ch)
			break
		} else {
			BroadcastLog(msg.Namespace, msg.Room, msg.Event, msg.Content)
		}

		time.Sleep(time.Millisecond * 100)
	}
}

func PubLogMsg(data _domain.MqMsg) {
	err := mqClientLog.Publish(mqTopicLog, data)
	if err != nil {
		fmt.Println("pub test mq message failed")
	}
}

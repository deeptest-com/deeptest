package ptwebsocket

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/pkg/core/mq"
	"github.com/aaronchen2k/deeptest/pkg/domain"
	"time"
)

var (
	mqTopicTest  = "MQ_TOPIC_TEST"
	mqClientTest *mq.Client
)

func InitTestMq() {
	mqClientTest = mq.NewClient()
	//defer mqClientTest.Close()

	mqClientTest.SetConditions(10000)

	go SubTestMsg()
}

func SubTestMsg() {
	ch, err := mqClientTest.Subscribe(mqTopicTest)
	if err != nil {
		fmt.Printf("sub test mq topic %s failed\n", mqTopicTest)
		return
	}

	for {
		msg := mqClientTest.GetPayLoad(ch).(_domain.MqMsg)
		fmt.Printf("%s get test mq msg '%#v'\n", mqTopicTest, msg.Content)

		if msg.Content == "exit" {
			mqClientTest.Unsubscribe(mqTopicTest, ch)
			break
		} else {
			BroadcastTest(msg.Namespace, msg.Room, msg.Event, msg.Content)
		}

		time.Sleep(time.Millisecond * 100)
	}
}

func PubTestMsg(data _domain.MqMsg) {
	err := mqClientTest.Publish(mqTopicTest, data)
	if err != nil {
		fmt.Println("pub test mq message failed")
	}
}

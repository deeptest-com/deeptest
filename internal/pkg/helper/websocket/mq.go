package websocketHelper

import (
	"fmt"
	"github.com/deeptest-com/deeptest/pkg/core/mq"
	_domain "github.com/deeptest-com/deeptest/pkg/domain"
	"time"
)

var (
	mqTopic  = "MQ_WebsocketTopic"
	mqClient *mq.Client
)

func InitMq() {
	mqClient = mq.NewClient()
	//defer mqClient.Close()
	mqClient.SetConditions(10000)

	go SubMsg()
}

func SubMsg() {
	ch, err := mqClient.Subscribe(mqTopic)
	if err != nil {
		fmt.Printf("sub mq topic %s failed\n", mqTopic)
		return
	}

	for {
		msg := mqClient.GetPayLoad(ch).(_domain.MqMsg)
		fmt.Printf("%s get mq msg '%#v'\n", mqTopic, msg.Content)

		if msg.Content == "exit" {
			mqClient.Unsubscribe(mqTopic, ch)
			break
		} else {
			Broadcast(msg.Namespace, msg.Room, msg.Event, msg.Content)
		}

		time.Sleep(time.Millisecond * 100)
	}
}

func PubMsg(data _domain.MqMsg) {
	err := mqClient.Publish(mqTopic, data)
	if err != nil {
		fmt.Println("pub mq message failed")
	}
}

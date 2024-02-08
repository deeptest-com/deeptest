package ptwebsocket

import (
	"fmt"
	ptdomain "github.com/aaronchen2k/deeptest/internal/performance/pkg/domain"
	"github.com/aaronchen2k/deeptest/pkg/core/mq"
	"time"
)

var (
	queueTopicOfWebSocket  = "QUEUE_TOPIC_OF_WEBSOCKET"
	queueClientOfWebSocket *mq.Client
)

func InitWsMq() {
	queueClientOfWebSocket = mq.NewClient()
	//defer queueClientOfWebSocket.Close()
	queueClientOfWebSocket.SetConditions(10000)

	go SubWsMsg()
}

func SubWsMsg() {
	ch, err := queueClientOfWebSocket.Subscribe(queueTopicOfWebSocket)
	if err != nil {
		fmt.Printf("sub mq topic %s failed\n", queueTopicOfWebSocket)
		return
	}

	for {
		msg := queueClientOfWebSocket.GetPayLoad(ch).(ptdomain.MqMsg)
		fmt.Printf("%s get mq msg '%#v'\n", queueTopicOfWebSocket, msg.Content)

		if msg.Content == "exit" {
			queueClientOfWebSocket.Unsubscribe(queueTopicOfWebSocket, ch)
			break
		} else {
			Broadcast(msg.Namespace, msg.Room, msg.Event, msg.Content)
		}

		time.Sleep(time.Millisecond * 1)
	}
}

func PubWsMsg(data ptdomain.MqMsg) {
	err := queueClientOfWebSocket.Publish(queueTopicOfWebSocket, data)
	if err != nil {
		fmt.Println("pub mq message failed")
	}
}

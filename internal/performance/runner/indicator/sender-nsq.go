package indicator

import (
	"encoding/json"
	ptProto "github.com/aaronchen2k/deeptest/internal/performance/proto"
	"github.com/nsqio/go-nsq"
)

var (
	Instant *NsqSender
)

type NsqSender struct {
	NsqServerAddress string
	NsqLookupAddress string

	Topic    string
	producer *nsq.Producer
}

func GetNsqSenderInstant(topic string, nsqServerAddress, nsqLookupAddress string) MessageSender {
	if Instant != nil && Instant.producer != nil {
		return Instant
	}

	Instant = &NsqSender{
		NsqServerAddress: nsqServerAddress,
		NsqLookupAddress: nsqLookupAddress,

		Topic: topic,
	}

	var err error

	if Instant.producer == nil {
		Instant.producer, err = nsq.NewProducer(nsqServerAddress, nsq.NewConfig())
		if err != nil {
			return nil
		}
	}

	return Instant
}

func (s NsqSender) Send(result ptProto.PerformanceExecResp) (err error) {
	bytes, err := json.Marshal(result)
	if err != nil {
		return err
	}

	err = s.producer.Publish(s.Topic, bytes)
	if err != nil {
		return err
	}

	return
}

package controllerService

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/performance/controller/dao"
	ptlog "github.com/aaronchen2k/deeptest/internal/performance/pkg/log"
	ptqueue "github.com/aaronchen2k/deeptest/internal/performance/pkg/queue"
	ptProto "github.com/aaronchen2k/deeptest/internal/performance/proto"
	"github.com/kataras/iris/v12/websocket"
	"github.com/nsqio/go-nsq"
	"io"
	"log"
	"time"
)

var (
	count = 0
)

func (s *PerformanceTestService) HandleAndPubToQueueGrpcMsg(stream ptProto.PerformanceService_ExecStartClient) (err error) {
	count := 0

	for true {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			continue
		}

		// dealwith Instruction from runner
		if resp.Instruction != "" {
			continue
		}

		// for response msg, put it to queue
		ptqueue.PubRunnerGrpcMsg(*resp)

		count += len(resp.Requests)
		ptlog.Logf("****** SERVER DEBUG: totally %d requests pub to queue", count)
	}

	return
}

func (s *PerformanceTestService) HandleRunnerNsqMsg(ctx context.Context,
	room, nsqLookupAddress, nsqServerAddress string, wsMsg *websocket.Message) (err error) {

	channel := fmt.Sprintf("channel_%s", room)
	consumer, err := nsq.NewConsumer(room, channel, nsq.NewConfig())
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer consumer.Stop()

	consumer.AddHandler(newNsqMsgProcessor(s.nsqMsgCallback, room, wsMsg))

	nsqAddr := nsqServerAddress
	if nsqLookupAddress != "" {
		nsqAddr = nsqLookupAddress
	}
	err = consumer.ConnectToNSQD(nsqAddr)
	if err != nil {
		return
	}

	for true {
		select {
		case <-ctx.Done():
			return

		default:
			time.Sleep(3 * time.Second)
		}
	}

	return nil
}

func (s *PerformanceTestService) HandleRunnerGrpcMsg(execCtx context.Context, cancel context.CancelFunc,
	room string, wsMsg *websocket.Message) (err error) {

	ptqueue.SubRunnerGrpcMsg(s.retrieveAndDealwithResult, execCtx, cancel, room, wsMsg)

	return
}

func (s *PerformanceTestService) nsqMsgCallback(bytes []byte, execUUid string, wsMsg *websocket.Message) error {
	log.Println(fmt.Sprintf("receive msg: %s", bytes))

	result := ptProto.PerformanceExecResp{}
	json.Unmarshal(bytes, &result)

	result.ExecUUid = execUUid

	s.retrieveAndDealwithResult(result, execUUid, wsMsg)

	return nil
}

func (s *PerformanceTestService) retrieveAndDealwithResult(result ptProto.PerformanceExecResp, room string, wsMsg *websocket.Message) (err error) {
	count += len(result.Requests)
	ptlog.Logf("****** SERVER DEBUG: totally retrieve %d requests and insert to sqlite", count)

	dao.InsertRequestRecord(result.Requests, result.RunnerId, result.Room)
	dao.InsertMetricsRecord(result.Metrics, result.RunnerId, result.Timestamp)

	return
}

type NsqMsgProcessor struct {
	callback func(msg []byte, execUUid string, wsMsg *websocket.Message) error
	cancel   context.CancelFunc
	execUuid string
	wsMsg    *websocket.Message
}

func newNsqMsgProcessor(callback func(msg []byte, execUUid string, wsMsg *websocket.Message) error, execUUid string, wsMsg *websocket.Message) *NsqMsgProcessor {
	return &NsqMsgProcessor{
		execUuid: execUUid,
		callback: callback,
		wsMsg:    wsMsg,
	}
}

func (m *NsqMsgProcessor) HandleMessage(msg *nsq.Message) (err error) {
	err = m.callback(msg.Body, m.execUuid, m.wsMsg)
	if err != nil {
		return
	}

	msg.Finish()

	return nil
}

package metrics

import (
	ptProto "github.com/aaronchen2k/deeptest/internal/performance/proto"
)

type GrpcSender struct {
	Stream *ptProto.PerformanceService_RunnerExecStartServer
}

func NewGrpcSender(stream *ptProto.PerformanceService_RunnerExecStartServer) MessageSender {
	ret := GrpcSender{
		Stream: stream,
	}

	return ret
}

func (s GrpcSender) Send(result ptProto.PerformanceExecResp) (err error) {
	(*s.Stream).Send(&result)

	return
}

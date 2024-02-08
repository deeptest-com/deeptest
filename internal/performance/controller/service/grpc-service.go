package controllerService

import (
	"context"
	"fmt"
	ptdomain "github.com/aaronchen2k/deeptest/internal/performance/pkg/domain"
	ptProto "github.com/aaronchen2k/deeptest/internal/performance/proto"
	"github.com/aaronchen2k/deeptest/internal/performance/runner/exec"
	"github.com/aaronchen2k/deeptest/internal/performance/runner/indicator"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"io"
	"log"
	"sync"
)

type GrpcService struct {
	execCtx    context.Context
	execCancel context.CancelFunc

	variableMap sync.Map
}

// for controller
func (s *GrpcService) AddRendezvousVal(ctx context.Context, req *ptProto.GlobalVarRequest) (
	ret *wrapperspb.Int32Value, err error) {

	key := fmt.Sprintf("%s_%s", req.Room, req.Name)

	ret = &wrapperspb.Int32Value{}
	ret.Value = 0

	varRef, ok := s.variableMap.Load(key)
	if !ok {
		val := 1
		varRef = &val

		s.variableMap.Store(key, varRef)

		ret.Value = 1

		return
	}

	// change
	*varRef.(*int) = *varRef.(*int) + 1

	ret.Value = int32(*varRef.(*int))

	// test
	testValue, _ := s.GetRendezvousVal(ctx, &ptProto.GlobalVarRequest{
		Room: req.Room,
		Name: req.Name,
	})

	log.Println(testValue)

	return
}

func (s *GrpcService) GetRendezvousVal(ctx context.Context, req *ptProto.GlobalVarRequest) (
	ret *wrapperspb.Int32Value, err error) {

	key := fmt.Sprintf("%s_%s", req.Room, req.Name)

	ret = &wrapperspb.Int32Value{}
	ret.Value = 0

	varRef, ok := s.variableMap.Load(key)
	if !ok {
		val := 0
		varRef = &val

		s.variableMap.Store(key, varRef)

		return
	}

	ret.Value = int32(*varRef.(*int))

	return
}

func (s *GrpcService) ClearRendezvousVal(ctx context.Context, req *ptProto.GlobalVarRequest) (
	ret *wrapperspb.BoolValue, err error) {

	ret = &wrapperspb.BoolValue{}
	ret.Value = false

	key := fmt.Sprintf("%s_%s", req.Room, req.Name)

	s.variableMap.Delete(key)

	ret.Value = true

	return
}

func (s *GrpcService) ClearAllRendezvousVal(ctx context.Context, req *ptProto.GlobalVarRequest) (
	ret *wrapperspb.BoolValue, err error) {

	ret = &wrapperspb.BoolValue{}
	ret.Value = false

	s.variableMap.Range(func(key, value interface{}) bool {
		s.variableMap.Delete(key)

		return true
	})

	ret.Value = true

	return
}

// for runner
func (s *GrpcService) ExecStart(stream ptProto.PerformanceService_ExecStartServer) (err error) {
	if exec.IsRunnerTestRunning() {
		err = &ptdomain.ErrorAlreadyRunning{}

		return
	}

	exec.SetRunnerTestRunning(true)
	defer func() {
		exec.SetRunnerTestRunning(false)
	}()

	indicator.Init()

	req, err := stream.Recv()
	if err == io.EOF {
		err = nil
		return
	}
	if req == nil {
		return
	}

	// gen sender
	grpcSender := indicator.NewGrpcSender(&stream)
	msgSender := indicator.GetInfluxdbSenderInstant(req.Room,
		req.InfluxdbAddress, req.InfluxdbUsername, req.InfluxdbPassword)

	s.execCtx, s.execCancel = context.WithCancel(context.Background())

	// run interval job
	go indicator.ScheduleJob(s.execCtx, req.RunnerId, req.Room, grpcSender)

	exec.ExecProgram(s.execCtx, s.execCancel, req, msgSender) // sync

	return
}

func (s *GrpcService) ExecStop(stream ptProto.PerformanceService_ExecStopServer) (err error) {
	instruction, err := stream.Recv()
	if err == io.EOF {
		err = nil
		return
	}

	if instruction == nil {
		return
	}

	if s.execCancel != nil {
		s.execCancel()
	}

	return
}

func (s *GrpcService) ForwardResult(result ptProto.PerformanceExecResp,
	stream *ptProto.PerformanceService_ExecStartServer) (err error) {

	err = (*stream).Send(&result)

	return
}

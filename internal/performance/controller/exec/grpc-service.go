package controllerExec

import (
	"context"
	"fmt"
	ptconsts "github.com/aaronchen2k/deeptest/internal/performance/pkg/consts"
	ptdomain "github.com/aaronchen2k/deeptest/internal/performance/pkg/domain"
	ptlog "github.com/aaronchen2k/deeptest/internal/performance/pkg/log"
	ptProto "github.com/aaronchen2k/deeptest/internal/performance/proto"
	"github.com/aaronchen2k/deeptest/internal/performance/runner/exec"
	"github.com/aaronchen2k/deeptest/internal/performance/runner/indicator"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"io"
	"log"
	"sync"
	"time"
)

type GrpcService struct {
	execCtx    context.Context
	execCancel context.CancelFunc

	variableMap sync.Map
}

// controller call runner, executed on runner side
func (s *GrpcService) ExecStart(stream ptProto.PerformanceService_ExecStartServer) (err error) {
	if runnerExec.IsRunnerTestRunning() {
		err = &ptdomain.ErrorAlreadyRunning{}

		return
	}

	runnerExec.SetRunnerTestRunning(true)
	defer func() {
		runnerExec.SetRunnerTestRunning(false)
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
	msgSender := indicator.GetInfluxdbSenderInstant(req.Room, req.InfluxdbAddress, req.InfluxdbOrg, req.InfluxdbToken)
	if msgSender == nil {
		ptlog.Logf("stop to run since msgSender return nil")
		return
	}

	s.execCtx, s.execCancel = context.WithCancel(context.Background())

	// run interval job
	go indicator.ScheduleJob(s.execCtx, req.RunnerId, req.RunnerName, req.Room, msgSender)

	// sync exec testing
	runnerExec.ExecProgram(s.execCtx, s.execCancel, req, msgSender) //

	// send end signal to controller
	result := ptProto.PerformanceExecResp{
		Timestamp:   time.Now().UnixMilli(),
		RunnerId:    req.RunnerId,
		Room:        req.Room,
		Instruction: ptconsts.MsgInstructionRunnerFinish.String(),
	}
	grpcSender := indicator.NewGrpcSender(&stream)
	grpcSender.Send(result)

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

// runner call controller, executed on controller side
func (s *GrpcService) AddGlobalVar(ctx context.Context, req *ptProto.GlobalVarRequest) (
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
	testValue, _ := s.GetGlobalVar(ctx, &ptProto.GlobalVarRequest{
		Room: req.Room,
		Name: req.Name,
	})

	log.Println(testValue)

	return
}

func (s *GrpcService) GetGlobalVar(ctx context.Context, req *ptProto.GlobalVarRequest) (
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

func (s *GrpcService) ClearGlobalVar(ctx context.Context, req *ptProto.GlobalVarRequest) (
	ret *wrapperspb.BoolValue, err error) {

	ret = &wrapperspb.BoolValue{}
	ret.Value = false

	key := fmt.Sprintf("%s_%s", req.Room, req.Name)

	s.variableMap.Delete(key)

	ret.Value = true

	return
}

func (s *GrpcService) ClearAllGlobalVar(ctx context.Context, req *ptProto.GlobalVarRequest) (
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

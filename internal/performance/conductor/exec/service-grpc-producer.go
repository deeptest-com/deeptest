package conductorExec

import (
	"context"
	"fmt"
	ptconsts "github.com/aaronchen2k/deeptest/internal/performance/pkg/consts"
	ptdomain "github.com/aaronchen2k/deeptest/internal/performance/pkg/domain"
	ptlog "github.com/aaronchen2k/deeptest/internal/performance/pkg/log"
	ptProto "github.com/aaronchen2k/deeptest/internal/performance/proto"
	"github.com/aaronchen2k/deeptest/internal/performance/runner/exec"
	"github.com/aaronchen2k/deeptest/internal/performance/runner/metrics"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"io"
	"log"
	"sync"
	"time"
)

/**
	注意：本文件包括双向的GRPC调用方法的服务端，如conductor调用runner执行，runner调用conductor设置全局变量等。
         启动的Agent程序，同时包含conductor和runner端的代码和功能。
*/

type GrpcService struct {
	execCtx    context.Context
	execCancel context.CancelFunc

	variableMap sync.Map
}

// conductor call runner, executed on runner side
func (s *GrpcService) RunnerExecStart(stream ptProto.PerformanceService_RunnerExecStartServer) (err error) {
	if runnerExec.IsRunnerTestRunning() {
		err = &ptdomain.ErrorAlreadyRunning{}
		return
	}

	runnerExec.SetRunnerTestRunning(true)

	req, err := stream.Recv()
	if err == io.EOF {
		err = nil
		return
	}
	if req == nil {
		return
	}

	// init runner log
	ptlog.Init(req.Room)

	// runner add item to cache
	AddTestItem(req.Room, ptconsts.Runner, nil, nil, req)

	// gen influxdb sender
	influxdbSender := metrics.GetInfluxdbSenderInstant(req.InfluxdbAddress, req.InfluxdbOrg, req.InfluxdbToken)
	if influxdbSender == nil {
		ptlog.Logf("stop to run since msgSender return nil")
		return
	}

	// context
	s.execCtx, s.execCancel = context.WithCancel(context.Background())

	// schedule job to send metrics
	go metrics.ScheduleJob(s.execCtx, req.RunnerId, req.RunnerName, req.Room, influxdbSender)

	go func() {
		defer runnerExec.SetRunnerTestRunning(false)

		// SYNC exec testing
		runnerExec.ExecProgram(s.execCtx, s.execCancel, req, influxdbSender)

		// send end signal to conductor
		result := ptProto.PerformanceExecResp{
			Timestamp:   time.Now().UnixMilli(),
			RunnerId:    req.RunnerId,
			Room:        req.Room,
			Instruction: ptconsts.MsgInstructionRunnerFinish.String(),
		}
		grpcSender := metrics.NewGrpcSender(&stream)
		grpcSender.Send(result)
	}()

	return
}

func (s *GrpcService) RunnerExecStop(stream ptProto.PerformanceService_RunnerExecStopServer) (err error) {
	req, err := stream.Recv()
	if err == io.EOF {
		err = nil
		return
	}
	if req == nil {
		return
	}

	room := req.Room

	logService := GetLogService(room)
	if logService != nil {
		if logService.logCancel != nil {
			logService.logCancel()
		}
		DeleteLogService(room)
	}

	if s.execCancel != nil {
		s.execCancel()
	}

	return
}

// runner call conductor, executed on conductor side
func (s *GrpcService) ConductorAddGlobalVar(ctx context.Context, req *ptProto.GlobalVarRequest) (
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
	testValue, _ := s.ConductorGetGlobalVar(ctx, &ptProto.GlobalVarRequest{
		Room: req.Room,
		Name: req.Name,
	})

	log.Println(testValue)

	return
}

func (s *GrpcService) ConductorGetGlobalVar(ctx context.Context, req *ptProto.GlobalVarRequest) (
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

func (s *GrpcService) ConductorClearGlobalVar(ctx context.Context, req *ptProto.GlobalVarRequest) (
	ret *wrapperspb.BoolValue, err error) {

	ret = &wrapperspb.BoolValue{}
	ret.Value = false

	key := fmt.Sprintf("%s_%s", req.Room, req.Name)

	s.variableMap.Delete(key)

	ret.Value = true

	return
}

func (s *GrpcService) ConductorClearAllGlobalVar(ctx context.Context, req *ptProto.GlobalVarRequest) (
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

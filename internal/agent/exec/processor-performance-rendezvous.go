package agentExec

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/domain"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/utils/exec"
	performanceUtils "github.com/aaronchen2k/deeptest/internal/agent/exec/utils/performance"
	ptlog "github.com/aaronchen2k/deeptest/internal/agent/performance/pkg/log"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	commonUtils "github.com/aaronchen2k/deeptest/pkg/lib/comm"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	uuid "github.com/satori/go.uuid"
	"time"
)

type ProcessorPerformanceRendezvous struct {
	ID uint `json:"id" yaml:"id"`
	ProcessorEntityBase

	Target int `json:"target"`
}

func (entity ProcessorPerformanceRendezvous) Run(processor *Processor, session *ExecSession) (err error) {
	defer func() {
		if errX := recover(); errX != nil {
			processor.Error(session, errX)
		}
	}()
	logUtils.Infof("performance rendezvous entity")

	startTime := time.Now()
	processor.Result = &agentExecDomain.ScenarioExecResult{
		ID:                int(entity.ProcessorID),
		Name:              entity.Name,
		ProcessorCategory: entity.ProcessorCategory,
		ProcessorType:     entity.ProcessorType,
		StartTime:         &startTime,
		ParentId:          int(entity.ParentID),
		ScenarioId:        processor.ScenarioId,
		ProcessorId:       processor.ID,
		LogId:             uuid.NewV4(),
		ParentLogId:       processor.Parent.Result.LogId,
		Round:             processor.Round,
	}

	processor.Result.Summary = fmt.Sprintf("集合点%s开始。", entity.Name)
	processor.AddResultToParent()
	detail := map[string]interface{}{"name": entity.Name, "target": entity.Target}
	processor.Result.Detail = commonUtils.JsonEncode(detail)
	execUtils.SendExecMsg(*processor.Result, consts.Processor, session.WsMsg)

	endTime := time.Now()
	processor.Result.EndTime = &endTime

	// rendezvous control
	execParams := performanceUtils.GetExecParamsInCtx(session.Ctx)
	name := processor.Name

	newArrivedVal := performanceUtils.IncreaseRendezvousArrived(execParams.Room, name, execParams.ConductorGrpcAddress)
	ptlog.Logf("====== VU %d: rendezvous Arrived Value added, value is %d", session.VuNo, newArrivedVal)

	// wait condition util rendezvous ready
	var value int
	var ready bool
	var newPassedVal int

	for !ready {
		value, ready = performanceUtils.IsRendezvousReady(execParams.Room, name, execParams.ConductorGrpcAddress, entity.Target)

		ptlog.Logf("------ VU %d: rendezvous wait, Arrived Value is %d", session.VuNo, value)
		time.Sleep(1 * time.Second)

		select {
		case <-session.Ctx.Done():
			ptlog.Logf("****** VU %d: rendezvous processor exit scenario processors by signal", session.VuNo)
			goto Label_END_RENDEZVOUS_PROCESSOR

		default:
		}
	}

	ptlog.Logf("------ VU %d: rendezvous passed, Arrived Value is %d", session.VuNo, value)

	// reset if all vus passed
	newPassedVal = performanceUtils.IncreaseRendezvousPassed(execParams.Room, name, execParams.ConductorGrpcAddress)
	if newPassedVal >= entity.Target {
		ptlog.Logf("------ VU %d: before rendezvous reset, Arrived Value is %d", session.VuNo, value)

		value = performanceUtils.ResetRendezvous(execParams.Room, name, execParams.ConductorGrpcAddress)

		ptlog.Logf("====== VU %d: after rendezvous reset, Arrived Value is %d", session.VuNo, value)

		ptlog.Logf("\n")
	}

Label_END_RENDEZVOUS_PROCESSOR:

	return
}

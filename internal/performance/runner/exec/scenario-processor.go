package runnerExec

import (
	"context"
	"encoding/json"
	agentExec "github.com/aaronchen2k/deeptest/internal/agent/exec"
	ptconsts "github.com/aaronchen2k/deeptest/internal/performance/pkg/consts"
	ptlog "github.com/aaronchen2k/deeptest/internal/performance/pkg/log"
	"github.com/aaronchen2k/deeptest/internal/performance/runner/metrics"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	_logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
)

func ExecProcessors(timeoutCtx context.Context, sender metrics.MessageSender, runnerId int32, vuNo int) {
	execParams := getExecParamsInCtx(timeoutCtx)

	rootProcessor := agentExec.Processor{}
	json.Unmarshal(execParams.Scenario.ProcessorRaw, &rootProcessor)

	for index, processor := range rootProcessor.Children {
		if processor.EntityType == consts.ProcessorInterfaceDefault {
			ptlog.Logf("exec processor %v", processor)
			ExecInterfaceProcessor(processor, timeoutCtx, execParams.Room, runnerId, vuNo, index, sender)

		} else if execParams.Mode == ptconsts.Parallel && processor.EntityType == consts.ProcessorPerformanceRendezvousDefault {
			ptlog.Logf("exec processor %v", processor)
			ExecRendezvousProcessor(processor, timeoutCtx, execParams.Room, runnerId, vuNo, execParams.ConductorGrpcAddress, index)

		}

		// 每个场景处理器完成后，检测是否有终止执行的信号
		select {
		case <-timeoutCtx.Done():
			_logUtils.Debugf("vu %d exit scenario processors by signal", vuNo)
			goto Label_END_SCENARIO

		default:
		}
	}

	_logUtils.Debugf("vu %d complete scenario normally", vuNo)

Label_END_SCENARIO:

	return
}

package exec

import (
	"context"
	ptconsts "github.com/aaronchen2k/deeptest/internal/performance/pkg/consts"
	ptlog "github.com/aaronchen2k/deeptest/internal/performance/pkg/log"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	_logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
)

func ExecProcessors(timeoutCtx context.Context, vuNo int) {
	execParams := getExecParamsInCtx(timeoutCtx)

	for index, processor := range execParams.Scenario.Processors {
		if processor.Type == consts.ProcessorInterfaceDefault.ToString() {
			ptlog.Logf("exec processor %v", processor)
			ExecInterfaceProcessor(processor, execParams.Room, vuNo, index)

		} else if execParams.Mode == ptconsts.Parallel && processor.Type == ptconsts.Rendezvous.String() {
			ptlog.Logf("exec processor %v", processor)
			ExecRendezvousProcessor(timeoutCtx, processor, vuNo, index, execParams.Room, execParams.ServerAddress)

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

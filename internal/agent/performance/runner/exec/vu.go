package runnerExec

import (
	"context"
	performanceUtils "github.com/aaronchen2k/deeptest/internal/agent/exec/utils/performance"
)

func ExecScenarioWithVu(timeoutCtx context.Context, vuNo int) (err error) {
	execParams := performanceUtils.GetExecParamsInCtx(timeoutCtx)

	for index := 1; true; index++ {
		ExecProcessors(timeoutCtx, vuNo)

		// break if loop goal reach
		if execParams.Loop > 0 && index >= execParams.Loop {
			goto Label_END_TASK
		}

		// break if duration goal timeout
		select {
		case <-timeoutCtx.Done():
			goto Label_END_TASK

		default:
		}
	}

Label_END_TASK:

	return
}

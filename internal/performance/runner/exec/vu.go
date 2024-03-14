package runnerExec

import (
	"context"
	performanceUtils "github.com/aaronchen2k/deeptest/internal/agent/exec/utils/performance"
)

func ExecScenarioWithVu(timeoutCtx context.Context) (err error) {
	execParams := performanceUtils.GetExecParamsInCtx(timeoutCtx)

	for index := 0; execParams.Loop == 0 || index < execParams.Loop; index++ {
		ExecProcessors(timeoutCtx)

		// loop util stage duration end
		select {
		case <-timeoutCtx.Done():
			goto Label_END_TASK

		default:
		}
	}

Label_END_TASK:

	return
}

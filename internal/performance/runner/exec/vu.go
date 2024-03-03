package runnerExec

import (
	"context"
	"github.com/aaronchen2k/deeptest/internal/performance/runner/metrics"
)

func ExecScenarioWithVu(timeoutCtx context.Context, sender metrics.MessageSender, vuNo int) (err error) {
	execParams := getExecParamsInCtx(timeoutCtx)

	for index := 0; execParams.Loop == 0 || index < execParams.Loop; index++ {
		ExecProcessors(timeoutCtx, sender, execParams.RunnerId, vuNo)

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

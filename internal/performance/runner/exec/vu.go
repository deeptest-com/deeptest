package exec

import (
	"context"
)

func ExecScenarioWithVu(timeoutCtx context.Context, vuNo int) (err error) {
	execParams := getExecParamsInCtx(timeoutCtx)

	for index := 0; execParams.Loop == 0 || index < execParams.Loop; index++ {
		ExecProcessors(timeoutCtx, vuNo)

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

package runnerExec

import (
	"context"
	"github.com/aaronchen2k/deeptest/internal/performance/pkg/domain"
	"math"
)

var (
	isRunnerRunning = false
)

func IsRunnerTestRunning() bool {
	return isRunnerRunning
}

func SetRunnerTestRunning(val bool) {
	isRunnerRunning = val
}

func getVuNumbByWeight(target, weight int) (ret int) {
	runnerTarget := math.Round(float64(target) * float64(weight) / float64(100))

	ret = int(runnerTarget)

	return
}

func genExecParamsCtx(data *ptdomain.ExecParamsInCtx, parentCtx context.Context) (ret context.Context) {
	ret = context.WithValue(parentCtx, "execParams", data)

	return
}

func getExecParamsInCtx(ctx context.Context) (data *ptdomain.ExecParamsInCtx) {
	data = ctx.Value("execParams").(*ptdomain.ExecParamsInCtx)

	return
}

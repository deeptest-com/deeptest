package performanceUtils

import (
	"context"
	agentExecDomain "github.com/aaronchen2k/deeptest/internal/agent/exec/domain"
	"math"
)

func GetVuNumbByWeight(target, weight int) (ret int) {
	runnerTarget := math.Round(float64(target) * float64(weight) / float64(100))

	ret = int(runnerTarget)

	if ret == 0 {
		ret = 1
	}

	return
}

func GenExecParamsCtx(data *agentExecDomain.ExecParamsInCtx, parentCtx context.Context) (ret context.Context) {
	ret = context.WithValue(parentCtx, "execParams", data)

	return
}

func GetExecParamsInCtx(ctx context.Context) (data *agentExecDomain.ExecParamsInCtx) {
	data = ctx.Value("execParams").(*agentExecDomain.ExecParamsInCtx)

	return
}

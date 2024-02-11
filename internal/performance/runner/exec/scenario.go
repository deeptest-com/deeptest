package exec

import (
	"context"
	"github.com/aaronchen2k/deeptest/internal/performance/pkg/consts"
	ptdomain "github.com/aaronchen2k/deeptest/internal/performance/pkg/domain"
	ptProto "github.com/aaronchen2k/deeptest/internal/performance/proto"
	"github.com/aaronchen2k/deeptest/internal/performance/runner/indicator"
)

func ExecScenario(execCtx context.Context, mode ptconsts.ExecMode,
	scenario *ptProto.Scenario, weight int32, room, serverAddress string,
	runnerId int32, sender indicator.MessageSender) (result ptProto.PerformanceExecResp) {

	var generater VuGenerator

	var valueCtx context.Context

	if scenario.GenerateType == ptconsts.GeneratorConstant.String() {
		runDur := int(scenario.Stages[0].Duration)

		data := ptdomain.ExecParamsInCtx{
			Scenario: scenario,

			RunnerId: runnerId,
			Room:     room,
			Target:   int(scenario.Stages[0].Target),
			Weight:   int(weight),
			Mode:     mode,

			Duration: runDur,
			Loop:     int(scenario.Stages[0].Loop),

			ServerAddress: serverAddress,
		}

		valueCtx = genExecParamsCtx(&data, execCtx)

		generater = ConstantVuGenerator{}

	} else if scenario.GenerateType == ptconsts.GeneratorRamp.String() {
		data := ptdomain.ExecParamsInCtx{
			Stages:   scenario.Stages,
			Scenario: scenario,

			RunnerId: runnerId,
			Room:     room,
			Weight:   int(weight),
			Mode:     mode,

			// computer Duration and Loop in each stage
		}

		valueCtx = genExecParamsCtx(&data, execCtx)

		generater = RampVuGenerator{}
	}

	generater.Run(valueCtx, sender)

	return
}

package runnerExec

import (
	"context"
	"encoding/json"
	"github.com/aaronchen2k/deeptest/internal/performance/pkg/consts"
	ptdomain "github.com/aaronchen2k/deeptest/internal/performance/pkg/domain"
	ptproto "github.com/aaronchen2k/deeptest/internal/performance/proto"
	"github.com/aaronchen2k/deeptest/internal/performance/runner/metrics"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
)

func ExecScenario(execCtx context.Context, mode ptconsts.ExecMode,
	scenario *ptproto.Scenario, weight int32, environmentId int32, execSceneRaw []byte,
	room, serverAddress string,
	runnerId int32, runnerName string, sender metrics.MessageSender) (result ptproto.PerformanceExecResp) {

	var generater VuGenerator

	var valueCtx context.Context

	execScene := domain.ExecScene{}
	json.Unmarshal(execSceneRaw, &execScene)

	if scenario.GenerateType == ptconsts.GeneratorConstant.String() {
		runDur := int(scenario.Duration)

		data := ptdomain.ExecParamsInCtx{
			Scenario:      scenario,
			EnvironmentId: int(environmentId),

			RunnerId:   runnerId,
			RunnerName: runnerName,
			Room:       room,
			Target:     int(scenario.Target),
			Weight:     int(weight),
			Mode:       mode,
			ExecScene:  execScene,

			Sender: sender,

			Duration: runDur,
			Loop:     int(scenario.Stages[0].Loop),

			ConductorGrpcAddress: serverAddress,
		}

		valueCtx = genExecParamsCtx(&data, execCtx)

		generater = ConstantVuGenerator{}

	} else if scenario.GenerateType == ptconsts.GeneratorRamp.String() {
		data := ptdomain.ExecParamsInCtx{
			Stages:        scenario.Stages,
			Scenario:      scenario,
			EnvironmentId: int(environmentId),

			RunnerId:   runnerId,
			RunnerName: runnerName,
			Room:       room,
			Weight:     int(weight),
			Mode:       mode,
			ExecScene:  execScene,

			Sender: sender,

			// computer Duration and Loop in each stage
		}

		valueCtx = genExecParamsCtx(&data, execCtx)

		generater = RampVuGenerator{}
	}

	generater.Run(valueCtx)

	return
}

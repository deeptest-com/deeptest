package runnerExec

import (
	"context"
	"encoding/json"
	agentExecDomain "github.com/aaronchen2k/deeptest/internal/agent/exec/domain"
	performanceUtils "github.com/aaronchen2k/deeptest/internal/agent/exec/utils/performance"
	"github.com/aaronchen2k/deeptest/internal/performance/pkg/consts"
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

		data := agentExecDomain.ExecParamsInCtx{
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

		valueCtx = performanceUtils.GenExecParamsCtx(&data, execCtx)

		generater = ConstantVuGenerator{}

	} else if scenario.GenerateType == ptconsts.GeneratorRamp.String() {
		data := agentExecDomain.ExecParamsInCtx{
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

		valueCtx = performanceUtils.GenExecParamsCtx(&data, execCtx)

		generater = RampVuGenerator{}
	}

	generater.Run(valueCtx)

	return
}

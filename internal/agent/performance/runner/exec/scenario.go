package runnerExec

import (
	"context"
	"encoding/json"
	agentExec "github.com/aaronchen2k/deeptest/internal/agent/exec"
	agentExecDomain "github.com/aaronchen2k/deeptest/internal/agent/exec/domain"
	performanceUtils "github.com/aaronchen2k/deeptest/internal/agent/exec/utils/performance"
	"github.com/aaronchen2k/deeptest/internal/agent/performance/pkg/consts"
	ptproto "github.com/aaronchen2k/deeptest/internal/agent/performance/proto"
	"github.com/aaronchen2k/deeptest/internal/agent/performance/runner/metrics"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
)

func ExecScenario(execCtx context.Context, mode ptconsts.ExecMode,
	scenario *ptproto.Scenario, weight int32, environmentId int32, execSceneRaw []byte,
	room, webServerUrl, webServerToken, conductorGrpcAddress string,
	runnerId int32, runnerName string, sender metrics.MessageSender) (result ptproto.PerformanceExecResp) {

	var generater VuGenerator

	var valueCtx context.Context

	execScene := domain.ExecScene{}
	json.Unmarshal(execSceneRaw, &execScene)

	if scenario.GenerateType == ptconsts.GeneratorConstant.String() {
		runDur := int(scenario.Duration)

		//loop := 0
		//if scenario.GenerateType == ptconsts.GeneratorConstant {
		//	loop = scenario.Loop
		//}

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

			WebServerUrl:         webServerUrl,
			WebServerToken:       webServerToken,
			ConductorGrpcAddress: conductorGrpcAddress,
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

			WebServerUrl:         webServerUrl,
			ConductorGrpcAddress: conductorGrpcAddress,

			// computer Duration and Loop in each stage
		}

		valueCtx = performanceUtils.GenExecParamsCtx(&data, execCtx)

		generater = RampVuGenerator{}
	}

	agentExec.InitUserExecContext(room)

	generater.Run(valueCtx)

	agentExec.CloseUserExecCtx(room)

	return
}

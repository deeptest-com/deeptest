package runnerExec

import (
	"context"
	"encoding/json"
	agentExec "github.com/aaronchen2k/deeptest/internal/agent/exec"
	performanceUtils "github.com/aaronchen2k/deeptest/internal/agent/exec/utils/performance"
	ptlog "github.com/aaronchen2k/deeptest/internal/agent/performance/pkg/log"
	"github.com/aaronchen2k/deeptest/internal/agent/service"
)

func ExecProcessors(timeoutCtx context.Context, vuNo int) (err error) {
	execParams := performanceUtils.GetExecParamsInCtx(timeoutCtx)
	ptlog.Logf("1. perfomance ExecProcessors, param %v", execParams)

	rootProcessor := agentExec.Processor{}
	json.Unmarshal(execParams.Scenario.ProcessorRaw, &rootProcessor)

	scenarioExecObj := &agentExec.ScenarioExecObj{
		ScenarioExecObjBase: agentExec.ScenarioExecObjBase{
			ScenarioId: uint(execParams.Scenario.Id),
			Name:       execParams.Scenario.Name,

			ExecScene: execParams.ExecScene,
			ExecUuid:  execParams.Room,

			ServerUrl: execParams.WebServerUrl,
			Token:     execParams.WebServerToken,
		},
		RootProcessor: &rootProcessor,
	}
	ptlog.Logf("2. perfomance generate scenarioExecObj %v", scenarioExecObj)

	service.UpdateLocalValues(&scenarioExecObj.ExecScene, execParams.LocalVarsCache)

	session := agentExec.NewScenarioExecSession(timeoutCtx, execParams.RunnerId, vuNo, scenarioExecObj, execParams.EnvironmentId, nil)
	err = service.ExecScenario(session)

	return
}

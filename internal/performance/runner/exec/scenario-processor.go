package runnerExec

import (
	"context"
	"encoding/json"
	agentExec "github.com/aaronchen2k/deeptest/internal/agent/exec"
	"github.com/aaronchen2k/deeptest/internal/agent/service"
)

func ExecProcessors(timeoutCtx context.Context) (err error) {
	execParams := getExecParamsInCtx(timeoutCtx)

	rootProcessor := agentExec.Processor{}
	json.Unmarshal(execParams.Scenario.ProcessorRaw, &rootProcessor)

	scenarioExecObj := &agentExec.ScenarioExecObj{
		ScenarioExecObjBase: agentExec.ScenarioExecObjBase{
			ScenarioId: uint(execParams.Scenario.Id),
			Name:       execParams.Scenario.Name,

			BaseUrl:   "",
			ExecScene: execParams.ExecScene,
			ExecUuid:  execParams.Room,
		},
		RootProcessor: &rootProcessor,
	}
	service.UpdateLocalValues(&scenarioExecObj.ExecScene, execParams.LocalVarsCache)

	session := agentExec.NewScenarioExecSession(scenarioExecObj, execParams.EnvironmentId, nil)
	err = service.ExecScenario(session)

	//	for index, processor := range rootProcessor.Children {
	//		if processor.EntityType == consts.ProcessorInterfaceDefault {
	//			ptlog.Logf("exec processor %v", processor)
	//			ExecInterfaceProcessor(processor, timeoutCtx, execParams.Room, runnerId, vuNo, index, sender)
	//
	//		} else if execParams.Mode == ptconsts.Parallel && processor.EntityType == consts.ProcessorPerformanceRendezvousDefault {
	//			ptlog.Logf("exec processor %v", processor)
	//			ExecRendezvousProcessor(processor, timeoutCtx, execParams.Room, runnerId, vuNo, execParams.ConductorGrpcAddress, index)
	//
	//		}
	//
	//		// 每个场景处理器完成后，检测是否有终止执行的信号
	//		select {
	//		case <-timeoutCtx.Done():
	//			_logUtils.Debugf("vu %d exit scenario processors by signal", vuNo)
	//			goto Label_END_SCENARIO
	//
	//		default:
	//		}
	//	}
	//
	//	_logUtils.Debugf("vu %d complete scenario normally", vuNo)
	//
	//Label_END_SCENARIO:

	return
}

package service

import (
	"github.com/aaronchen2k/deeptest/internal/agent/exec"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/utils/exec"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/websocket"
)

func RunScenario(req *agentExec.ScenarioExecReq, localVarsCache iris.Map, wsMsg *websocket.Message) (err error) {
	logUtils.Infof("run scenario %d on environment %d", req.ScenarioId, req.EnvironmentId)

	// send start msg
	execUtils.SendStartMsg(wsMsg)

	//场景执行初始信息
	normalData := GetScenarioNormalData(req)
	execUtils.SendInitializeMsg(normalData, wsMsg)

	scenarioExecObj := GetScenarioToExec(req)
	if scenarioExecObj == nil {
		execUtils.SendEndMsg(wsMsg)
		return
	}
	updateLocalValues(&scenarioExecObj.ExecScene, localVarsCache)

	scenarioExecObj.ExecUuid = req.ExecUuid

	session := agentExec.NewScenarioExecSession(0, scenarioExecObj, req.EnvironmentId, wsMsg)
	err = ExecScenario(session)

	session.ScenarioDebug.RootProcessor.Result.Stat = *agentExec.GetInterfaceStat(session.ExecUuid)
	session.ScenarioDebug.RootProcessor.Result.EnvironmentId = req.EnvironmentId
	session.ScenarioDebug.RootProcessor.Result.ScenarioId = req.ScenarioId

	// submit result
	report, _ := SubmitScenarioResult(*session.ScenarioDebug.RootProcessor.Result, scenarioExecObj.RootProcessor.ScenarioId,
		session.ServerUrl, session.ServerToken, req.TenantId)

	execUtils.SendResultMsg(report, session.ScenarioDebug.WsMsg)

	// end msg
	execUtils.SendEndMsg(wsMsg)

	return
}

func ExecScenario(session *agentExec.ExecSession) (err error) {
	RestoreEntityFromRawAndSetParent(session.ScenarioDebug.RootProcessor)

	session.Run()

	if session.ScenarioDebug.RootProcessor.Result != nil {
		session.ScenarioDebug.RootProcessor.Result.ScenarioId = session.ScenarioDebug.ScenarioId
	}

	return
}

func RestoreEntityFromRawAndSetParent(root *agentExec.Processor) (err error) {
	processors := make([]*agentExec.Processor, 0)

	agentExec.GetProcessorList(root, &processors)

	processorMap := map[uint]*agentExec.Processor{}
	for _, processor := range processors {
		processor.RestoreEntity()
		processorMap[processor.ID] = processor
	}

	for _, obj := range processorMap {
		obj.Parent = processorMap[obj.ParentId]
	}

	return
}

package service

import (
	"github.com/aaronchen2k/deeptest/internal/agent/exec"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/utils/exec"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/websocket"
	"go.uber.org/zap"
)

func RunScenario(req *agentExec.ScenarioExecReq, localVarsCache iris.Map, wsMsg *websocket.Message) (err error) {
	logUtils.Infof("run scenario", zap.Int("ScenarioId", req.ScenarioId), zap.Int("environmentId", req.EnvironmentId))

	agentExec.ResetStat(req.ExecUuid)
	agentExec.SetForceStopExec(req.ExecUuid, false)

	agentExec.SetServerUrl(req.ExecUuid, req.ServerUrl)
	agentExec.SetServerToken(req.ExecUuid, req.Token)

	// start msg
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

	session, err := ExecScenario(scenarioExecObj, wsMsg)
	session.RootProcessor.Result.Stat = *agentExec.GetInterfaceStat(req.ExecUuid)
	session.RootProcessor.Result.EnvironmentId = req.EnvironmentId
	session.RootProcessor.Result.ScenarioId = uint(req.ScenarioId)

	// submit result
	report, _ := SubmitScenarioResult(*session.RootProcessor.Result, scenarioExecObj.RootProcessor.ScenarioId,
		agentExec.GetServerUrl(req.ExecUuid), agentExec.GetServerToken(req.ExecUuid))

	execUtils.SendResultMsg(report, session.WsMsg)
	//sendScenarioSubmitResult(session.RootProcessor.ID, session.WsMsg)

	// end msg
	execUtils.SendEndMsg(wsMsg)

	return
}

func ExecScenario(execObj *agentExec.ScenarioExecObj, wsMsg *websocket.Message) (
	session *agentExec.Session, err error) {
	// variables etc.
	agentExec.SetExecScene(execObj.ExecUuid, execObj.ExecScene)

	RestoreEntityFromRawAndSetParent(execObj.RootProcessor)

	agentExec.InitScenarioExecContext(execObj)
	agentExec.InitJsRuntime(execObj.RootProcessor.ProjectId, execObj.ExecUuid)

	// start msg
	//execUtils.SendStartMsg(wsMsg)

	// execution
	session = agentExec.NewSession(execObj, false, wsMsg)
	session.ExecUuid = execObj.ExecUuid

	session.Run()

	if session.RootProcessor.Result != nil {
		session.RootProcessor.Result.ScenarioId = execObj.ScenarioId
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

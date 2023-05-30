package service

import (
	"github.com/aaronchen2k/deeptest/internal/agent/exec"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/domain"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/utils/exec"
	"github.com/kataras/iris/v12/websocket"
)

func RunScenario(req *agentExec.ScenarioExecReq, wsMsg *websocket.Message) (err error) {
	agentExec.ServerUrl = req.ServerUrl
	agentExec.ServerToken = req.Token

	scenarioExecObj := GetScenarioToExec(req)

	// start msg
	execUtils.SendStartMsg(wsMsg)

	//场景执行初始信息
	normalData := GetScenarioNormalData(req)
	execUtils.SendInitializeMsg(normalData, wsMsg)

	session, err := ExecScenario(scenarioExecObj, wsMsg)
	session.RootProcessor.Result.EnvironmentId = req.EnvironmentId
	session.RootProcessor.Result.ScenarioId = uint(req.ScenarioId)

	// submit result
	report, _ := SubmitScenarioResult(*session.RootProcessor.Result, scenarioExecObj.RootProcessor.ScenarioId,
		agentExec.ServerUrl, agentExec.ServerToken)

	execUtils.SendResultMsg(report, session.WsMsg)
	//sendScenarioSubmitResult(session.RootProcessor.ID, session.WsMsg)

	// end msg
	execUtils.SendEndMsg(wsMsg)

	return
}

func ExecScenario(execObj *agentExec.ScenarioExecObj, wsMsg *websocket.Message) (
	session *agentExec.Session, err error) {

	// variables etc.
	agentExec.ExecScene = execObj.ExecScene

	RestoreEntityFromRawAndSetParent(execObj.RootProcessor)

	agentExec.InitExecContext(execObj)
	agentExec.InitJsRuntime()

	// start msg
	//execUtils.SendStartMsg(wsMsg)

	// execution
	session = agentExec.NewSession(execObj, false, wsMsg)
	session.Run()
	session.RootProcessor.Result.ScenarioId = execObj.ScenarioId
	return
}

func CancelAndSendMsg(scenarioId int, wsMsg websocket.Message) (err error) {
	execUtils.SendCancelMsg(wsMsg)
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

func sendScenarioSubmitResult(rootId uint, wsMsg *websocket.Message) (err error) {
	result := agentDomain.ScenarioExecResult{
		ID:       -3,
		ParentId: int(rootId),
		Name:     "提交执行结果成功",
		//Summary:  fmt.Sprintf("错误：%s", err.Error()),
	}

	execUtils.SendExecMsg(result, wsMsg)

	return
}

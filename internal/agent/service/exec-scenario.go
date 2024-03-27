package service

import (
	"context"
	"github.com/aaronchen2k/deeptest/internal/agent/exec"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/utils/exec"
	ptlog "github.com/aaronchen2k/deeptest/internal/agent/performance/pkg/log"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/websocket"
)

func RunScenario(ctx context.Context, req *agentExec.ScenarioExecReq, localVarsCache iris.Map, wsMsg *websocket.Message) (err error) {
	logUtils.Infof("run scenario %d on environment %d", req.ScenarioId, req.EnvironmentId)

	// send start msg
	execUtils.SendStartMsg(wsMsg)

	// send init data
	normalData := GetScenarioNormalData(req)
	execUtils.SendInitializeMsg(normalData, wsMsg)

	scenarioExecObj := GetScenarioToExec(req)
	UpdateLocalValues(&scenarioExecObj.ExecScene, localVarsCache)

	scenarioExecObj.ExecUuid = req.ExecUuid

	session := agentExec.NewScenarioExecSession(ctx, 0, 0, scenarioExecObj, req.EnvironmentId, wsMsg)
	err = ExecScenario(session)

	session.RootProcessor.Result.Stat = *agentExec.GetInterfaceStat(session.ExecUuid)
	session.RootProcessor.Result.EnvironmentId = req.EnvironmentId
	session.RootProcessor.Result.ScenarioId = uint(req.ScenarioId)

	// submit result
	report, _ := SubmitScenarioResult(*session.RootProcessor.Result, scenarioExecObj.RootProcessor.ScenarioId,
		session.ServerUrl, session.ServerToken)

	execUtils.SendResultMsg(report, session.WsMsg)

	// end msg
	execUtils.SendEndMsg(wsMsg)

	return
}

func ExecScenario(session *agentExec.ExecSession) (err error) {
	ptlog.Logf("3. perfomance exec session %v", session)

	RestoreEntityFromRawAndSetParent(session.RootProcessor)
	ptlog.Logf("4. perfomance exec session %v", session)

	session.Run()

	if session.RootProcessor.Result != nil {
		session.RootProcessor.Result.ScenarioId = session.ScenarioId
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

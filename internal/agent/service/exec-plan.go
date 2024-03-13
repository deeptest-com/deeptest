package service

import (
	"github.com/aaronchen2k/deeptest/internal/agent/exec"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/domain"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/utils/exec"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/websocket"
)

func RunPlan(req *agentExec.PlanExecReq, localVarsCache iris.Map, wsMsg *websocket.Message) (err error) {
	execUuid := req.ExecUuid

	planExecObj := GetPlanToExec(req)
	if planExecObj == nil || len(planExecObj.Scenarios) == 0 {
		execUtils.SendEndMsg(wsMsg)
		return
	}

	// start msg
	execUtils.SendStartMsg(wsMsg)

	normalData, err := GetPlanNormalData(req)
	if err != nil {
		return
	}

	execUtils.SendInitializeMsg(normalData, wsMsg)

	// execution
	var result = agentExecDomain.PlanExecResult{
		EnvironmentId: req.EnvironmentId,
		ID:            req.PlanId,
	}

	for _, scenarioExecObj := range planExecObj.Scenarios {
		scenarioExecObj.ExecUuid = execUuid
		scenarioExecObj.ServerUrl = req.ServerUrl
		scenarioExecObj.Token = req.Token
		updateLocalValues(&scenarioExecObj.ExecScene, localVarsCache)

		session := agentExec.NewScenarioExecSession(&scenarioExecObj, req.EnvironmentId, wsMsg)
		err = ExecScenario(session)

		scenarioReport, _ := SubmitScenarioResult(*session.RootProcessor.Result, session.RootProcessor.Result.ScenarioId,
			session.ServerUrl, session.ServerToken)
		session.RootProcessor.Result.EnvironmentId = req.EnvironmentId

		session.RootProcessor.Result.ScenarioReportId = uint(scenarioReport.ID)
		result.Scenarios = append(result.Scenarios, session.RootProcessor.Result)
		execUtils.SendResultMsg(scenarioReport, session.WsMsg)
	}

	// submit result
	result.Stat = *agentExec.GetInterfaceStat(execUuid)
	report, _ := SubmitPlanResult(result, req.PlanId, req.ServerUrl, req.Token)
	execUtils.SendResultMsg(report, wsMsg)
	//sendPlanSubmitResult(req.PlanId, wsMsg)

	// end msg
	execUtils.SendEndMsg(wsMsg)

	return
}

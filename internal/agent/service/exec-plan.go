package service

import (
	"github.com/aaronchen2k/deeptest/internal/agent/exec"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/domain"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/utils/exec"
	"github.com/kataras/iris/v12/websocket"
)

func RunPlan(req *agentExec.PlanExecReq, wsMsg *websocket.Message) (err error) {
	agentExec.ServerUrl = req.ServerUrl
	agentExec.ServerToken = req.Token

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
	_ = execUtils.SendResult(normalData, wsMsg)

	// execution
	var results = agentDomain.PlanExecResult{
		EnvironmentId: req.EnvironmentId,
		ID:            req.PlanId,
	}
	for _, scenario := range planExecObj.Scenarios {
		session, _ := ExecScenario(&scenario, wsMsg)
		scenarioReport, _ := SubmitScenarioResult(*session.RootProcessor.Result, session.RootProcessor.Result.ScenarioId,
			agentExec.ServerUrl, agentExec.ServerToken)
		session.RootProcessor.Result.EnvironmentId = req.EnvironmentId

		session.RootProcessor.Result.ScenarioReportId = uint(scenarioReport.ID)
		results.Scenarios = append(results.Scenarios, session.RootProcessor.Result)
		execUtils.SendResultMsg(scenarioReport, session.WsMsg)
	}

	// submit result
	report, _ := SubmitPlanResult(results, req.PlanId, req.ServerUrl, req.Token)
	execUtils.SendResultMsg(report, wsMsg)
	//sendPlanSubmitResult(req.PlanId, wsMsg)

	// end msg
	execUtils.SendEndMsg(wsMsg)

	return
}

func sendPlanSubmitResult(planId int, wsMsg *websocket.Message) (err error) {
	result := agentDomain.PlanExecResult{
		ID:   planId,
		Name: "提交计划执行结果成功",
		//Summary:  fmt.Sprintf("错误：%s", err.Error()),
	}
	execUtils.SendExecMsg(result, wsMsg)

	return
}

package service

import (
	"github.com/aaronchen2k/deeptest/internal/agent/exec"
	execUtils "github.com/aaronchen2k/deeptest/internal/agent/exec/utils/exec"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/websocket"
)

func RunCases(req *agentExec.CasesExecReq, wsMsg *websocket.Message) (err error) {
	logUtils.Infof("run cases %s on env %d", req.ExecUUid, req.EnvironmentId)

	// reset exec
	agentExec.ResetStat()
	agentExec.SetForceStopExec(req.ExecUuid, false)

	// start msg
	agentExec.SetIsRunning(req.ExecUuid, true)
	execUtils.SendStartMsg(wsMsg)

	// run case one by one
	for _, cs := range req.Cases {
		caseInterfaceExecObj := GetCaseToExec(
			req.ProjectId, req.BaseCaseId, cs, req.ServerUrl, req.Token, req.UsedBy)

		agentExec.SetCurrDebugInterfaceId(req.ExecUuid, caseInterfaceExecObj.DebugData.DebugInterfaceId)
		agentExec.SetCurrScenarioProcessorId(req.ExecUuid, 0) // not in a scenario

		agentExec.SetCurrRequest(req.ExecUuid, domain.BaseRequest{})
		agentExec.SetCurrResponse(req.ExecUuid, domain.DebugResponse{})
		agentExec.SetExecScene(req.ExecUuid, caseInterfaceExecObj.ExecScene)

		// init context
		agentExec.InitDebugExecContext(req.ExecUuid)
		agentExec.InitJsRuntime(req.ProjectId, req.ExecUuid)

		statusPreCondition, _ := agentExec.ExecPreConditions(caseInterfaceExecObj, req.ExecUuid) // must before PreRequest, since it will update the vari in script
		originalReqUri, _ := PreRequest(&caseInterfaceExecObj.DebugData, req.ExecUuid)

		agentExec.SetReqValueToGoja(caseInterfaceExecObj.DebugData.BaseRequest)
		agentExec.GetReqValueFromGoja(req.ExecUuid)

		// a new interface may not has a pre-script, which will not update agentExec.CurrRequest, need to skip
		if agentExec.GetCurrRequest(req.ExecUuid).Url != "" {
			caseInterfaceExecObj.DebugData.BaseRequest = agentExec.GetCurrRequest(req.ExecUuid) // update to the value changed in goja
		}

		resultResp, err1 := RequestInterface(&caseInterfaceExecObj.DebugData)
		if err1 != nil {
			execUtils.SendResult(err1, wsMsg)
			return err1
		}

		agentExec.SetRespValueToGoja(resultResp)
		statusPostCondition, _ := agentExec.ExecPostConditions(&caseInterfaceExecObj, resultResp, req.ExecUuid)
		agentExec.GetRespValueFromGoja(req.ExecUuid)
		PostRequest(originalReqUri, &caseInterfaceExecObj.DebugData)

		if agentExec.GetCurrResponse(req.ExecUuid).Data != nil {
			resultResp = agentExec.GetCurrResponse(req.ExecUuid)
		}

		status := consts.Pass
		if statusPreCondition == consts.Fail || statusPostCondition == consts.Fail {
			status = consts.Fail
		}

		result := iris.Map{
			"source": "execCases",

			"execUuid": req.ExecUUid,
			"caseUuid": cs.Key,
			"request":  caseInterfaceExecObj,
			"response": resultResp,

			"status": status,
		}

		// send result
		execUtils.SendExecMsg(result, consts.ProgressResult, wsMsg)

		// stop if needed
		if agentExec.GetForceStopExec(req.ExecUuid) {
			break
		}
	}

	// end msg
	agentExec.SetIsRunning(req.ExecUuid, false)
	execUtils.SendEndMsg(wsMsg)

	return
}

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
	agentExec.ForceStopExec = false

	// start msg
	execUtils.SendStartMsg(wsMsg)

	// run case one by one
	for _, cs := range req.Cases {
		caseInterfaceExecObj := GetCaseToExec(
			req.ProjectId, req.BaseCaseId, cs, req.ServerUrl, req.Token, req.UsedBy)

		agentExec.CurrDebugInterfaceId = caseInterfaceExecObj.DebugData.DebugInterfaceId
		agentExec.CurrScenarioProcessorId = 0 // not in a scenario

		agentExec.CurrRequest = domain.BaseRequest{}
		agentExec.CurrResponse = domain.DebugResponse{}
		agentExec.ExecScene = caseInterfaceExecObj.ExecScene

		// init context
		agentExec.InitDebugExecContext()
		agentExec.InitJsRuntime(req.ProjectId)

		statusPreCondition, _ := agentExec.ExecPreConditions(caseInterfaceExecObj) // must before PreRequest, since it will update the vari in script
		originalReqUri, _ := PreRequest(&caseInterfaceExecObj.DebugData)

		agentExec.SetReqValueToGoja(caseInterfaceExecObj.DebugData.BaseRequest)
		agentExec.GetReqValueFromGoja()

		// a new interface may not has a pre-script, which will not update agentExec.CurrRequest, need to skip
		if agentExec.CurrRequest.Url != "" {
			caseInterfaceExecObj.DebugData.BaseRequest = agentExec.CurrRequest // update to the value changed in goja
		}

		resultResp, err1 := RequestInterface(&caseInterfaceExecObj.DebugData)
		if err1 != nil {
			execUtils.SendResult(err1, wsMsg)
			return err1
		}

		agentExec.SetRespValueToGoja(resultResp)
		statusPostCondition, _ := agentExec.ExecPostConditions(&caseInterfaceExecObj, resultResp)
		agentExec.GetRespValueFromGoja()
		PostRequest(originalReqUri, &caseInterfaceExecObj.DebugData)

		if agentExec.CurrResponse.Data != nil {
			resultResp = agentExec.CurrResponse
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
		if agentExec.ForceStopExec {
			break
		}
	}

	// end msg
	execUtils.SendEndMsg(wsMsg)

	return
}

package service

import (
	"encoding/json"
	"github.com/aaronchen2k/deeptest/internal/agent/exec"
	execUtils "github.com/aaronchen2k/deeptest/internal/agent/exec/utils/exec"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/websocket"
)

func RunCases(req *agentExec.CasesExecReq, wsMsg *websocket.Message) (err error) {
	logUtils.Infof("run cases %s on env %d", req.ExecUuid, req.EnvironmentId)

	// reset exec
	agentExec.ResetStat()
	agentExec.SetForceStopExec(req.ExecUuid, false)

	// start msg
	execUtils.SendStartMsg(wsMsg)

	// run case one by one
	doExecCases(req, wsMsg, "")

	// end msg
	execUtils.SendEndMsg(wsMsg)

	return
}

func doExecCases(req *agentExec.CasesExecReq, wsMsg *websocket.Message, parentUuid string) (err error) {
	casesExecObj := GetCasesToExec(req)

	for _, cs := range casesExecObj.Children {
		doExecCase(cs, wsMsg, req.ExecUuid, parentUuid, req.ProjectId)

		// stop if needed
		if agentExec.GetForceStopExec(parentUuid) {
			break
		}
	}

	return
}

func doExecCase(cs *agentExec.CaseExecProcessor, wsMsg *websocket.Message, execUuid, parentUuid string, projectId uint) (err error) {
	if cs.Category != "case" {
		startMsg := iris.Map{
			"source":     "execCases",
			"execUuid":   execUuid,
			"caseUuid":   cs.Key,
			"category":   cs.Category,
			"title":      cs.Title,
			"parentUuid": parentUuid,
		}
		execUtils.SendExecMsg(startMsg, consts.ProgressResult, wsMsg)
	}

	for _, child := range cs.Children {
		doExecCase(child, wsMsg, execUuid, cs.Key, projectId)
	}

	if cs.Category != "case" {
		return
	}

	caseInterfaceExecObj := cs.Data

	// execution
	agentExec.SetCurrDebugInterfaceId(parentUuid, caseInterfaceExecObj.DebugData.DebugInterfaceId)
	agentExec.SetCurrScenarioProcessorId(parentUuid, 0) // not in a scenario

	agentExec.SetCurrRequest(parentUuid, domain.BaseRequest{})
	agentExec.SetCurrResponse(parentUuid, domain.DebugResponse{})
	agentExec.SetExecScene(parentUuid, caseInterfaceExecObj.ExecScene)

	// init context
	agentExec.InitDebugExecContext(execUuid)
	agentExec.InitJsRuntime(projectId, execUuid)

	statusPreCondition, _ := agentExec.ExecPreConditions(caseInterfaceExecObj, execUuid) // must before PreRequest, since it will update the vari in script
	originalReqUri, _ := PreRequest(&caseInterfaceExecObj.DebugData, execUuid)

	agentExec.SetReqValueToGoja(caseInterfaceExecObj.DebugData.BaseRequest)
	agentExec.GetReqValueFromGoja(execUuid)

	// a new interface may not has a pre-script, which will not update agentExec.CurrRequest, need to skip
	if agentExec.GetCurrRequest(execUuid).Url != "" {
		caseInterfaceExecObj.DebugData.BaseRequest = agentExec.GetCurrRequest(execUuid) // update to the value changed in goja
	}

	resultResp, err1 := RequestInterface(&caseInterfaceExecObj.DebugData)
	if err1 != nil {
		execUtils.SendResult(err1, wsMsg)
		return err1
	}

	agentExec.SetRespValueToGoja(resultResp)
	statusPostCondition, _ := agentExec.ExecPostConditions(caseInterfaceExecObj, resultResp, execUuid)
	agentExec.GetRespValueFromGoja(execUuid)
	PostRequest(originalReqUri, &caseInterfaceExecObj.DebugData)

	if agentExec.GetCurrResponse(execUuid).Data != nil {
		resultResp = agentExec.GetCurrResponse(execUuid)

		resultResp.ConsoleLogs = GenConditionLogsForCase(caseInterfaceExecObj) // only for cases
	}

	status := consts.Pass
	if statusPreCondition == consts.Fail || statusPostCondition == consts.Fail {
		status = consts.Fail
	}

	result := iris.Map{
		"source": "execCases",

		"execUuid": execUuid,
		"caseUuid": cs.Key,
		"request":  caseInterfaceExecObj,
		"response": resultResp,

		"status": status,

		"category":   cs.Category,
		"title":      cs.Title,
		"parentUuid": parentUuid,
	}

	// send result
	execUtils.SendExecMsg(result, consts.ProgressResult, wsMsg)

	return
}

func GenConditionLogsForCase(obj *agentExec.InterfaceExecObj) (ret []interface{}) {
	for _, pre := range obj.PreConditions {
		mp := map[string]interface{}{}
		json.Unmarshal(pre.Raw, &mp)

		ret = append(ret, mp)
	}

	for _, post := range obj.PostConditions {
		if post.Type == consts.ConditionTypeCheckpoint {
			continue
		}

		mp := map[string]interface{}{}
		json.Unmarshal(post.Raw, &mp)

		ret = append(ret, mp)
	}

	for _, post := range obj.PostConditions {
		if post.Type != consts.ConditionTypeCheckpoint {
			continue
		}

		mp := map[string]interface{}{}
		json.Unmarshal(post.Raw, &mp)

		ret = append(ret, mp)
	}

	return
}

package service

import (
	"encoding/json"
	"github.com/aaronchen2k/deeptest/internal/agent/exec"
	execUtils "github.com/aaronchen2k/deeptest/internal/agent/exec/utils/exec"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/websocket"
)

func RunCases(req *agentExec.CasesExecReq, localVarsCache iris.Map, wsMsg *websocket.Message) (err error) {
	logUtils.Infof("run cases %s on env %d", req.ExecUuid, req.EnvironmentId)

	// start msg
	execUtils.SendStartMsg(wsMsg)

	// run case one by one
	doExecCases(req, localVarsCache, wsMsg, "")

	// end msg
	execUtils.SendEndMsg(wsMsg)

	return
}

func doExecCases(req *agentExec.CasesExecReq, localVarsCache iris.Map, wsMsg *websocket.Message, parentUuid string) (err error) {
	userCtx := agentExec.GetUserExecContext(parentUuid)
	casesExecObj := GetCasesToExec(req)

	for _, cs := range casesExecObj.Children {
		doExecCase(cs, localVarsCache, wsMsg, req.TenantId, req.ExecUuid, parentUuid, req.ProjectId)

		select {
		case <-userCtx.ExecCtx.Done():
			break

		default:
		}
	}

	return
}

func doExecCase(cs *agentExec.CaseExecProcessor, localVarsCache iris.Map, wsMsg *websocket.Message, tenantId consts.TenantId, execUuid, parentUuid string, projectId uint) (err error) {
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
		doExecCase(child, localVarsCache, wsMsg, tenantId, execUuid, cs.Key, projectId)
	}

	if cs.Category != "case" {
		return
	}

	// init context
	call := agentExec.InterfaceExecReq{
		ExecUuid:  cs.ExecUUid,
		Data:      cs.Data.DebugData,
		ExecScene: cs.Data.ExecScene,
	}
	updateLocalValues(&call.ExecScene, localVarsCache)
	session := agentExec.NewInterfaceExecSession(call) // not in a scenario

	agentExec.ExecPreConditions(cs.Data, session) // must before PreRequest, since it will update the vari in script
	originalReqUri, _ := PreRequest(&cs.Data.DebugData, session)

	agentExec.SetReqValueToGoja(&cs.Data.DebugData.BaseRequest, session)
	agentExec.GetReqValueFromGoja(session)

	// a new interface may not has a pre-script, which will not update agentExec.CurrRequest, need to skip
	if session.GetCurrRequest().Url != "" {
		cs.Data.DebugData.BaseRequest = session.GetCurrRequest() // update to the value changed in goja
	}

	resultResp, err1 := RequestInterface(&cs.Data.DebugData, nil)
	if err1 != nil {
		execUtils.SendResult(err1, wsMsg)
		return err1
	}

	agentExec.SetRespValueToGoja(&resultResp, session)
	assertResultStatus, _ := agentExec.ExecPostConditions(cs.Data, resultResp, session)
	agentExec.GetRespValueFromGoja(session)
	PostRequest(originalReqUri, &cs.Data.DebugData)

	// get the response data updated by script post-condition
	if session.GetCurrResponse().Data != nil {
		resultResp = session.GetCurrResponse()
		resultResp.ConsoleLogs = GenConditionLogsForCase(cs.Data) // only for cases
	}

	status := consts.Pass
	if assertResultStatus == consts.Fail {
		status = consts.Fail
	}

	result := iris.Map{
		"source": "execCases",

		"execUuid": execUuid,
		"caseUuid": cs.Key,
		"request":  cs.Data,
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

package service

import (
	"context"
	"encoding/json"
	"github.com/aaronchen2k/deeptest/internal/agent/exec"
	execUtils "github.com/aaronchen2k/deeptest/internal/agent/exec/utils/exec"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/websocket"
)

func RunCases(ctx context.Context, req *agentExec.CasesExecReq, localVarsCache iris.Map, wsMsg *websocket.Message) (err error) {
	logUtils.Infof("run cases %s on env %d", req.ExecUuid, req.EnvironmentId)

	// start msg
	execUtils.SendStartMsg(wsMsg)

	// run case one by one
	doExecCases(ctx, req, localVarsCache, wsMsg, "")

	// end msg
	execUtils.SendEndMsg(wsMsg)

	return
}

func doExecCases(ctx context.Context, req *agentExec.CasesExecReq, localVarsCache iris.Map, wsMsg *websocket.Message, parentUuid string) (err error) {
	casesExecObj := GetCasesToExec(req)

	for _, cs := range casesExecObj.Children {
		doExecCase(ctx, cs, localVarsCache, wsMsg, req.ExecUuid, parentUuid, req.ProjectId)

		select {
		case <-ctx.Done():
			break

		default:
		}
	}

	return
}

func doExecCase(ctx context.Context, cs *agentExec.CaseExecProcessor, localVarsCache iris.Map, wsMsg *websocket.Message, execUuid, parentUuid string, projectId uint) (err error) {
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
		doExecCase(ctx, child, localVarsCache, wsMsg, execUuid, cs.Key, projectId)
	}

	if cs.Category != "case" {
		return
	}

	// init context
	call := domain.InterfaceCall{
		ExecUuid:  cs.ExecUUid,
		Data:      cs.Data.DebugData,
		ExecScene: cs.Data.ExecScene,
	}
	UpdateLocalValues(&call.ExecScene, localVarsCache)
	session := agentExec.NewInterfaceExecSession(call)

	// exec
	agentExec.ExecPreConditions(session, cs.Data) // must before PreRequest, since it will update the vari in script
	originalReqUri, _ := PreRequest(session, &cs.Data.DebugData)

	agentExec.SetReqValueToGoja(&cs.Data.DebugData.BaseRequest)
	agentExec.GetReqValueFromGoja(session)

	// a new interface may not has a pre-script, which will not update agentExec.CurrRequest, need to skip
	if session.CurrRequest.Url != "" {
		cs.Data.DebugData.BaseRequest = session.CurrRequest // update to the value changed in goja
	}

	resultResp, err1 := RequestInterface(&cs.Data.DebugData)
	if err1 != nil {
		execUtils.SendResult(err1, wsMsg)
		return err1
	}

	agentExec.SetRespValueToGoja(&resultResp)
	assertResultStatus, _ := agentExec.ExecPostConditions(session, cs.Data, resultResp)
	agentExec.GetRespValueFromGoja(session)
	PostRequest(originalReqUri, &cs.Data.DebugData)

	// get the response data updated by script post-condition
	if session.CurrResponse.Data != nil {
		resultResp = session.CurrResponse
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

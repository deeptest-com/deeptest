package service

import (
	agentExec "github.com/aaronchen2k/deeptest/internal/agent/exec"
	execUtils "github.com/aaronchen2k/deeptest/internal/agent/exec/utils/exec"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	_httpUtils "github.com/aaronchen2k/deeptest/pkg/lib/http"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/websocket"
)

func RunInterface(req *agentExec.InterfaceExecReq, localVarsCache iris.Map, wsMsg *websocket.Message) (err error) {
	logUtils.Infof("run debug interface %d on environment %d", req.Data.DebugInterfaceId, req.Data.EnvironmentId)

	// execution
	resultReq, resultResp, err := ExecInterface(req, localVarsCache, wsMsg)

	// send result
	result := iris.Map{
		"source":   "execInterface",
		"request":  resultReq,
		"response": resultResp,
	}
	execUtils.SendExecMsg(result, consts.ProgressResult, wsMsg)
	execUtils.SendExecMsg(iris.Map{"source": "execInterface"}, consts.ProgressEnd, wsMsg)

	return
}

func ExecInterface(req *agentExec.InterfaceExecReq, localVarsCache iris.Map, wsMsg *websocket.Message) (
	resultReq domain.DebugData, resultResp domain.DebugResponse, err error) {

	interfaceExecObj := GetInterfaceToExec(*req)
	req.ExecScene = interfaceExecObj.ExecScene
	updateLocalValues(&interfaceExecObj.ExecScene, localVarsCache)

	session := agentExec.NewInterfaceExecSession(*req)

	agentExec.SetReqValueToGoja(&interfaceExecObj.DebugData.BaseRequest, session)

	agentExec.ExecPreConditions(&interfaceExecObj, session) // must before PreRequest, since it will update the vari in script
	originalReqUri, _ := PreRequest(&interfaceExecObj.DebugData, session)

	agentExec.GetReqValueFromGoja(session)

	// A new interface may not has a pre-script, which will not update agentExec.CurrRequest, need to skip
	if session.GetCurrRequest().Url != "" {
		interfaceExecObj.DebugData.BaseRequest = session.GetCurrRequest() // update to the value changed in goja
	}

	resultResp, err = RequestInterface(&interfaceExecObj.DebugData, req.ExecUuid, wsMsg)

	agentExec.SetRespValueToGoja(&resultResp, session)
	assertResultStatusPost, _ := agentExec.ExecPostConditions(&interfaceExecObj, resultResp, session)

	agentExec.GetRespValueFromGoja(session)
	PostRequest(originalReqUri, &interfaceExecObj.DebugData)

	// get the response data updated by script post-condition
	if session.GetCurrResponse().Data != nil {
		resultResp = session.GetCurrResponse()
	}

	// submit result
	err = SubmitInterfaceResult(interfaceExecObj, resultResp, assertResultStatusPost, req.ServerUrl, req.Token)

	resultReq = interfaceExecObj.DebugData

	return
}

func PreRequest(req *domain.DebugData, session *agentExec.ExecSession) (originalReqUri string, err error) {
	// replace variables
	agentExec.ReplaceVariables(&req.BaseRequest, session)

	// gen url
	req.BaseRequest.Url, originalReqUri = UpdateUrl(*req)
	req.BaseRequest.FullUrlToDisplay = req.BaseRequest.Url
	logUtils.Info("requested url: " + req.BaseRequest.Url)

	// download form file item
	if req.BodyFormData != nil {
		for index, item := range *req.BodyFormData {
			if item.Type == consts.FormDataTypeFile {
				(*req.BodyFormData)[index].Value, err = agentExec.DownloadUploadedFile(item.Value, session)
			}
		}
	}

	return
}

func PostRequest(originalReqUri string, req *domain.DebugData) (err error) {
	// rollback for saved to db
	req.BaseRequest.Url = originalReqUri

	return
}

func RequestInterface(req *domain.DebugData, key string, wsMsg *websocket.Message) (ret domain.DebugResponse, err error) {
	ret, err = agentExec.Invoke(&req.BaseRequest, key, wsMsg)
	ret.Id = req.DebugInterfaceId

	return
}

func UpdateUrl(debugData domain.DebugData) (reqUrl, originalReqUri string) {
	// gen url
	if debugData.PathParams != nil {
		originalReqUri = agentExec.ReplacePathParams(debugData.Url, *debugData.PathParams)
	}

	notUseBaseUrl := execUtils.IsNotUseBaseUrl(debugData.UsedBy, debugData.ProcessorInterfaceSrc)
	if notUseBaseUrl {
		reqUrl = originalReqUri
	} else {
		reqUrl = _httpUtils.CombineUrls(debugData.BaseUrl, originalReqUri)
	}

	return
}

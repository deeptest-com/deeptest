package service

import (
	agentExec "github.com/aaronchen2k/deeptest/internal/agent/exec"
	execUtils "github.com/aaronchen2k/deeptest/internal/agent/exec/utils/exec"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	_httpUtils "github.com/aaronchen2k/deeptest/pkg/lib/http"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
)

func RunInterface(call domain.InterfaceCall) (resultReq domain.DebugData, resultResp domain.DebugResponse, err error) {
	req := GetInterfaceToExec(call)
	call.ExecScene = req.ExecScene
	updateLocalValues(&call.ExecScene, call.LocalVarsCache)
	session := agentExec.NewInterfaceExecSession(call)

	agentExec.SetReqValueToGoja(&req.DebugData.BaseRequest)

	agentExec.ExecPreConditions(session, &req) // must before PreRequest, since it will update the vari in script
	originalReqUri, _ := PreRequest(session, &req.DebugData)

	agentExec.GetReqValueFromGoja(session)

	// a new interface may not has a pre-script, which will not update agentExec.CurrRequest, need to skip
	if session.CurrRequest.Url != "" {
		req.DebugData.BaseRequest = session.CurrRequest // update to the value changed in goja
	}

	resultResp, err = RequestInterface(&req.DebugData)

	agentExec.SetRespValueToGoja(&resultResp)
	assertResultStatusPost, _ := agentExec.ExecPostConditions(session, &req, resultResp)

	agentExec.GetRespValueFromGoja(session)
	PostRequest(originalReqUri, &req.DebugData)

	// get the response data updated by script post-condition
	if session.CurrResponse.Data != nil {
		resultResp = session.CurrResponse
	}

	// submit result
	err = SubmitInterfaceResult(req, resultResp, assertResultStatusPost, call.ServerUrl, call.Token)

	resultReq = req.DebugData

	return
}

func PreRequest(session *agentExec.ExecSession, req *domain.DebugData) (originalReqUri string, err error) {
	// replace variables
	agentExec.ReplaceVariables(session, &req.BaseRequest)

	// gen url
	req.BaseRequest.Url, originalReqUri = UpdateUrl(*req)
	req.BaseRequest.FullUrlToDisplay = req.BaseRequest.Url
	logUtils.Info("requested url: " + req.BaseRequest.Url)

	// download form file item
	if req.BodyFormData != nil {
		for index, item := range *req.BodyFormData {
			if item.Type == consts.FormDataTypeFile {
				(*req.BodyFormData)[index].Value, err = agentExec.DownloadUploadedFile(session, item.Value)
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

func RequestInterface(req *domain.DebugData) (ret domain.DebugResponse, err error) {
	// send request
	ret, err = agentExec.Invoke(&req.BaseRequest)

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

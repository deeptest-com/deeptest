package service

import (
	agentDomain "github.com/aaronchen2k/deeptest/cmd/agent/v1/domain"
	agentExec "github.com/aaronchen2k/deeptest/internal/agent/exec"
	execUtils "github.com/aaronchen2k/deeptest/internal/agent/exec/utils/exec"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	_httpUtils "github.com/aaronchen2k/deeptest/pkg/lib/http"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
)

func RunInterface(call agentDomain.InterfaceCall) (resultReq domain.DebugData, resultResp domain.DebugResponse, err error) {
	req := GetInterfaceToExec(call)
	call.ExecScene = req.ExecScene
	updateLocalValues(&req.ExecScene, call.LocalVarsCache)
	session := agentExec.NewInterfaceExecSession(call)

	agentExec.SetReqValueToGoja(&req.DebugData.BaseRequest, session)

	agentExec.ExecPreConditions(&req, session) // must before PreRequest, since it will update the vari in script
	originalReqUri, _ := PreRequest(&req.DebugData, session)

	agentExec.GetReqValueFromGoja(session)

	// TODO: a new interface may not has a pre-script, which will not update agentExec.CurrRequest, need to skip
	if session.InterfaceDebug.CurrRequest.Url != "" {
		req.DebugData.BaseRequest = session.InterfaceDebug.CurrRequest // update to the value changed in goja
	}

	resultResp, err = RequestInterface(&req.DebugData)

	agentExec.SetRespValueToGoja(&resultResp, session)
	assertResultStatusPost, _ := agentExec.ExecPostConditions(&req, resultResp, session)

	agentExec.GetRespValueFromGoja(session)
	PostRequest(originalReqUri, &req.DebugData)

	// get the response data updated by script post-condition
	if session.InterfaceDebug.CurrResponse.Data != nil {
		resultResp = session.InterfaceDebug.CurrResponse
	}

	// submit result
	err = SubmitInterfaceResult(req, resultResp, assertResultStatusPost, call.ServerUrl, call.Token)

	resultReq = req.DebugData

	return
}

func PreRequest(req *domain.DebugData, session *Exec) (originalReqUri string, err error) {
	// replace variables
	agentExec.ReplaceVariables(&req.BaseRequest, tenantId, projectId, execUuid)

	// gen url
	req.BaseRequest.Url, originalReqUri = UpdateUrl(*req)
	req.BaseRequest.FullUrlToDisplay = req.BaseRequest.Url
	logUtils.Info("requested url: " + req.BaseRequest.Url)

	// download form file item
	if req.BodyFormData != nil {
		for index, item := range *req.BodyFormData {
			if item.Type == consts.FormDataTypeFile {
				(*req.BodyFormData)[index].Value, err = agentExec.DownloadUploadedFile(item.Value, execUuid)
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

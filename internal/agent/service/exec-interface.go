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
	agentExec.SetServerUrl(call.ExecUuid, call.ServerUrl)
	agentExec.SetServerToken(call.ExecUuid, call.Token)
	req := GetInterfaceToExec(call)

	agentExec.SetCurrDebugInterfaceId(call.ExecUuid, req.DebugData.DebugInterfaceId)
	agentExec.SetCurrScenarioProcessorId(call.ExecUuid, 0) // not in a scenario

	agentExec.SetCurrRequest(call.ExecUuid, domain.BaseRequest{})
	agentExec.SetCurrResponse(call.ExecUuid, domain.DebugResponse{})
	agentExec.SetExecScene(call.ExecUuid, req.ExecScene)

	// init context
	agentExec.InitDebugExecContext(call.ExecUuid)
	agentExec.InitJsRuntime(call.Data.ProjectId, call.ExecUuid)

	agentExec.SetReqValueToGoja(&req.DebugData.BaseRequest)

	agentExec.ExecPreConditions(&req, call.ExecUuid) // must before PreRequest, since it will update the vari in script
	originalReqUri, _ := PreRequest(&req.DebugData, call.ExecUuid)

	agentExec.GetReqValueFromGoja(call.ExecUuid)

	// TODO: a new interface may not has a pre-script, which will not update agentExec.CurrRequest, need to skip
	if agentExec.GetCurrRequest(call.ExecUuid).Url != "" {
		req.DebugData.BaseRequest = agentExec.GetCurrRequest(call.ExecUuid) // update to the value changed in goja
	}

	resultResp, err = RequestInterface(&req.DebugData)

	agentExec.SetRespValueToGoja(&resultResp)
	assertResultStatusPost, _ := agentExec.ExecPostConditions(&req, resultResp, call.ExecUuid)

	agentExec.GetRespValueFromGoja(call.ExecUuid)
	PostRequest(originalReqUri, &req.DebugData)

	// get the response data updated by script post-condition
	if agentExec.GetCurrResponse(call.ExecUuid).Data != nil {
		resultResp = agentExec.GetCurrResponse(call.ExecUuid)
	}

	// submit result
	err = SubmitInterfaceResult(req, resultResp, assertResultStatusPost, call.ServerUrl, call.Token)

	resultReq = req.DebugData

	return
}

func PreRequest(req *domain.DebugData, execUuid string) (originalReqUri string, err error) {
	// replace variables
	agentExec.ReplaceVariables(&req.BaseRequest, execUuid)

	// gen url
	if req.PathParams != nil {
		originalReqUri = agentExec.ReplacePathParams(req.Url, *req.PathParams)
	}

	notUseBaseUrl := execUtils.IsUseBaseUrl(req.UsedBy, req.ProcessorInterfaceSrc)
	if notUseBaseUrl {
		req.BaseRequest.Url = originalReqUri
	} else {
		req.BaseRequest.Url = _httpUtils.CombineUrls(req.BaseUrl, originalReqUri)
	}
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

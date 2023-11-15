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
	agentExec.ServerUrl = call.ServerUrl
	agentExec.ServerToken = call.Token
	req := GetInterfaceToExec(call)

	agentExec.CurrDebugInterfaceId = req.DebugData.DebugInterfaceId
	agentExec.CurrScenarioProcessorId = 0 // not in a scenario

	agentExec.CurrRequest = domain.BaseRequest{}
	agentExec.CurrResponse = domain.DebugResponse{}
	agentExec.ExecScene = req.ExecScene

	// init context
	agentExec.InitDebugExecContext()
	agentExec.InitJsRuntime(call.Data.ProjectId)

	agentExec.ExecPreConditions(req) // must before PreRequest, since it will update the vari in script
	originalReqUri, _ := PreRequest(&req.DebugData)

	agentExec.SetReqValueToGoja(req.DebugData.BaseRequest)
	agentExec.GetReqValueFromGoja()

	// a new interface may not has a pre-script, which will not update agentExec.CurrRequest, need to skip
	if agentExec.CurrRequest.Url != "" {
		req.DebugData.BaseRequest = agentExec.CurrRequest // update to the value changed in goja
	}

	resultResp, err = RequestInterface(&req.DebugData)

	agentExec.SetRespValueToGoja(resultResp)
	agentExec.ExecPostConditions(req, resultResp)
	agentExec.GetRespValueFromGoja()
	PostRequest(originalReqUri, &req.DebugData)

	if agentExec.CurrResponse.Data != nil {
		resultResp = agentExec.CurrResponse
	}

	// submit result
	err = SubmitInterfaceResult(req, resultResp, call.ServerUrl, call.Token)

	resultReq = req.DebugData

	return
}

func PreRequest(req *domain.DebugData) (originalReqUri string, err error) {
	// replace variables
	agentExec.ReplaceVariables(&req.BaseRequest, consts.InterfaceDebug)

	// gen url
	originalReqUri = agentExec.ReplacePathParams(req.Url, req.PathParams)

	notUseBaseUrl := execUtils.IsUseBaseUrl(req.UsedBy, req.ProcessorInterfaceSrc)
	if notUseBaseUrl {
		req.BaseRequest.Url = originalReqUri
	} else {
		req.BaseRequest.Url = _httpUtils.CombineUrls(req.BaseUrl, originalReqUri)
	}
	req.BaseRequest.FullUrlToDisplay = req.BaseRequest.Url
	logUtils.Info("requested url: " + req.BaseRequest.Url)

	return
}

func PostRequest(originalReqUri string, req *domain.DebugData) (err error) {
	req.BaseRequest.Url = originalReqUri // rollback for saved to db

	return
}

func RequestInterface(req *domain.DebugData) (ret domain.DebugResponse, err error) {
	// send request
	ret, err = agentExec.Invoke(&req.BaseRequest)

	ret.Id = req.DebugInterfaceId

	return
}

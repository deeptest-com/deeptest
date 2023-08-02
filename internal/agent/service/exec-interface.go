package service

import (
	agentDomain "github.com/aaronchen2k/deeptest/cmd/agent/v1/domain"
	agentExec "github.com/aaronchen2k/deeptest/internal/agent/exec"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	_httpUtils "github.com/aaronchen2k/deeptest/pkg/lib/http"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"strings"
)

func RunInterface(call agentDomain.InterfaceCall) (resultReq domain.DebugData, resultResp domain.DebugResponse, err error) {
	req := GetInterfaceToExec(call)

	agentExec.CurrDebugInterfaceId = req.DebugData.DebugInterfaceId
	agentExec.CurrScenarioProcessorId = 0 // not in a scenario

	agentExec.ExecScene = req.ExecScene

	// exec interface
	agentExec.ExecPreConditions(&req)
	resultResp, err = RequestInterface(req.DebugData)
	agentExec.ExecPostConditions(&req, resultResp)

	// submit result
	err = SubmitInterfaceResult(req, resultResp, call.ServerUrl, call.Token)

	resultReq = req.DebugData

	return
}

func RequestInterface(req domain.DebugData) (ret domain.DebugResponse, err error) {
	// replace variables
	agentExec.ReplaceVariables(&req.BaseRequest, consts.InterfaceDebug)

	// gen url
	reqUri := agentExec.ReplacePathParams(req.Url, req.PathParams)

	if req.ProcessorInterfaceSrc != consts.DiagnoseDebug {
		req.BaseRequest.Url = _httpUtils.CombineUrls(req.BaseUrl, reqUri)
	}
	logUtils.Info("url: " + req.BaseRequest.Url)

	// send request
	ret, err = agentExec.Invoke(&req.BaseRequest)

	req.BaseRequest.Url = reqUri // rollback for saved to db

	ret.Id = req.DebugInterfaceId

	return
}

func GetInterfaceContentProps(ret *domain.DebugResponse) {
	ret.ContentLang = consts.LangTEXT

	if ret.ContentLang == "" {
		return
	}

	arr := strings.Split(string(ret.ContentType), ";")
	arr1 := strings.Split(arr[0], "/")
	if len(arr1) == 1 {
		return
	}

	typeName := arr1[1]
	if typeName == "text" || typeName == "plain" {
		typeName = consts.LangTEXT.String()
	}
	ret.ContentLang = consts.HttpRespLangType(typeName)

	if len(arr) > 1 {
		arr2 := strings.Split(arr[1], "=")
		if len(arr2) > 1 {
			ret.ContentCharset = consts.HttpRespCharset(arr2[1])
		}
	}

	//ret.NodeContent = mockHelper.FormatXml(ret.NodeContent)

	return
}

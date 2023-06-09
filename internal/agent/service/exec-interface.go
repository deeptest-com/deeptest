package service

import (
	agentDomain "github.com/aaronchen2k/deeptest/cmd/agent/v1/domain"
	agentExec "github.com/aaronchen2k/deeptest/internal/agent/exec"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	_commUtils "github.com/aaronchen2k/deeptest/pkg/lib/comm"
	_httpUtils "github.com/aaronchen2k/deeptest/pkg/lib/http"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"strings"
)

func RunInterface(call agentDomain.InterfaceCall) (ret domain.DebugResponse, err error) {
	req := GetInterfaceToExec(call)

	agentExec.CurrInterfaceId = req.DebugData.EndpointInterfaceId
	agentExec.CurrProcessorId = 0 // not in a scenario

	agentExec.ExecScene = req.ExecScene

	logUtils.Info("DebugData:" + _commUtils.JsonEncode(req.DebugData))
	ret, err = RequestInterface(req.DebugData)
	logUtils.Info("DebugResponse:" + _commUtils.JsonEncode(ret))

	err = SubmitInterfaceResult(req.DebugData, ret, call.ServerUrl, call.Token)

	return
}

func RequestInterface(req domain.DebugData) (ret domain.DebugResponse, err error) {
	// exec pre-request script
	agentExec.ExecJs(req.PreRequestScript)

	// replace variables
	agentExec.ReplaceVariables(&req.BaseRequest, consts.InterfaceDebug)

	// gen url
	reqUri := agentExec.ReplacePathParams(req.Url, req.PathParams)
	req.BaseRequest.Url = _httpUtils.AddSepIfNeeded(req.BaseUrl) + reqUri
	logUtils.Info("reqUri:" + reqUri + "|" + "req.Url:" + req.Url + "|req.BaseRequest.Url:" + req.BaseRequest.Url)

	// send request
	ret, err = agentExec.Invoke(&req.BaseRequest)

	req.BaseRequest.Url = reqUri // rollback for saved to db

	ret.Id = req.EndpointInterfaceId

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

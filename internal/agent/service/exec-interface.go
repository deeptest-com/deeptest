package service

import (
	agentDomain "github.com/aaronchen2k/deeptest/cmd/agent/v1/domain"
	agentExec "github.com/aaronchen2k/deeptest/internal/agent/exec"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	_httpUtils "github.com/aaronchen2k/deeptest/pkg/lib/http"
	"strings"
)

type ExecInterfaceService struct {
	RemoteService *RemoteService `inject:""`
}

func (s *ExecInterfaceService) Run(call agentDomain.InterfaceCall) (ret domain.DebugResponse, err error) {
	req := s.RemoteService.GetInterfaceToExec(call)

	agentExec.CurrInterfaceId = req.DebugData.EndpointInterfaceId

	agentExec.ExecScene = req.ExecScene

	ret, err = s.Request(req.DebugData)
	err = s.RemoteService.SubmitInterfaceResult(req.DebugData, ret, call.ServerUrl, call.Token)

	return
}

func (s *ExecInterfaceService) Request(req domain.DebugData) (ret domain.DebugResponse, err error) {
	// exec pre-request script
	agentExec.ExecJs(req.PreRequestScript)

	// replace variables
	agentExec.DealwithVariables(&req.BaseRequest, consts.InterfaceDebug)

	// send request

	reqUrl := req.Url
	req.BaseRequest.Url = _httpUtils.AddSepIfNeeded(req.BaseUrl) + reqUrl

	ret, err = agentExec.Invoke(&req.BaseRequest)

	req.BaseRequest.Url = reqUrl // rollback for saved to db

	ret.Id = req.EndpointInterfaceId

	return
}

func (s *ExecInterfaceService) GetContentProps(ret *domain.DebugResponse) {
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

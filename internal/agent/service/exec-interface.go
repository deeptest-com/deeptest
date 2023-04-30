package service

import (
	"github.com/aaronchen2k/deeptest/cmd/agent/v1/domain"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	agentExec "github.com/aaronchen2k/deeptest/internal/agent/exec"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	_httpUtils "github.com/aaronchen2k/deeptest/pkg/lib/http"
	"strings"
)

type ExecInterfaceService struct {
	RemoteService *RemoteService `inject:""`
}

func (s *ExecInterfaceService) Run(call domain.InterfaceCall) (resp v1.DebugResponse, err error) {
	req := s.RemoteService.GetInterfaceToExec(call)

	resp, err = s.Request(req)
	err = s.RemoteService.SubmitInterfaceResult(req, resp, call.ServerUrl, call.Token)

	return
}

func (s *ExecInterfaceService) Request(req v1.DebugData) (ret v1.DebugResponse, err error) {
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

func (s *ExecInterfaceService) GetContentProps(ret *v1.DebugResponse) {
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

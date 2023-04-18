package service

import (
	"github.com/aaronchen2k/deeptest/cmd/agent/v1/domain"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	agentExec "github.com/aaronchen2k/deeptest/internal/agent/exec"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	_httpUtils "github.com/aaronchen2k/deeptest/pkg/lib/http"
	"strings"
)

type ExecService struct {
	RemoteService *RemoteService `inject:""`
}

func (s *ExecService) Run(call domain.InvokeCall) (resp v1.DebugResponse, err error) {
	req := s.RemoteService.GetInterfaceToExec(call)

	resp, err = s.Request(req)
	err = s.RemoteService.SubmitInterfaceResult(req, resp, call.ServerUrl, call.Token)

	/*
		if req.UsedBy == consts.InterfaceDebug {
			interfaceExecReq := s.RemoteService.GetInterfaceToExec(req)

			agentExec.EnvVars = interfaceExecReq.EnvVars
			agentExec.ShareVars = interfaceExecReq.ShareVars
			agentExec.DatapoolData = interfaceExecReq.Datapools

			resp, err = s.Request(interfaceExecReq)
			err = s.RemoteService.SubmitInterfaceResult(req, resp, req.ServerUrl, req.Token)

		} else if req.UsedBy == consts.ScenarioDebug {
			interfaceProcessorExecReq := s.RemoteService.GetProcessorInterfaceToExec(req)

			agentExec.EnvVars = interfaceProcessorExecReq.EnvVars
			agentExec.ShareVars = interfaceProcessorExecReq.ShareVars
			agentExec.DatapoolData = interfaceProcessorExecReq.Datapools

			resp, err = s.Request(interfaceProcessorExecReq)
			err = s.RemoteService.SubmitProcessorInterfaceResult(req, resp, req.ServerUrl, req.Token)

		}
	*/
	return
}

func (s *ExecService) Request(req v1.DebugRequest) (ret v1.DebugResponse, err error) {
	// exec pre-request script
	agentExec.ExecJs(req.PreRequestScript)

	// replace variables
	agentExec.ReplaceAll(&req.BaseRequest, agentExec.Environment, agentExec.Variables, agentExec.DatapoolData)

	// send request
	req.BaseRequest.Url = _httpUtils.AddSepIfNeeded(req.BaseUrl) + req.BaseRequest.Url
	ret, err = agentExec.Invoke(&req.BaseRequest)

	ret.Id = req.InterfaceId

	return
}

func (s *ExecService) GetContentProps(ret *v1.DebugResponse) {
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

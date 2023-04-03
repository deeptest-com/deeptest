package service

import (
	"github.com/aaronchen2k/deeptest/cmd/agent/v1/domain"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	agentExec "github.com/aaronchen2k/deeptest/internal/agent/exec"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"strings"
)

type InterfaceService struct {
	RemoteService *RemoteService `inject:""`
}

func (s *InterfaceService) Run(req domain.InvocationReq) (ret v1.DebugResponse, err error) {
	//单接口调试和多接口调试统一使用一套执行器表结构
	interfaceProcessorExecReq := s.RemoteService.GetProcessorInterfaceToExec(req)
	agentExec.Environment = interfaceProcessorExecReq.Environment
	agentExec.Variables = interfaceProcessorExecReq.Variables
	agentExec.DatapoolData = interfaceProcessorExecReq.Datapools
	ret, err = s.Request(interfaceProcessorExecReq)
	err = s.RemoteService.SubmitProcessorInterfaceResult(req, ret, req.ServerUrl, req.Token)

	/*
		if req.UsedBy == consts.UsedByInterface {
			interfaceExecReq := s.RemoteService.GetInterfaceToExec(req)

			agentExec.Environment = interfaceExecReq.Environment
			agentExec.Variables = interfaceExecReq.Variables
			agentExec.DatapoolData = interfaceExecReq.Datapools

			ret, err = s.Request(interfaceExecReq)
			err = s.RemoteService.SubmitInterfaceResult(req, ret, req.ServerUrl, req.Token)

		} else if req.UsedBy == consts.UsedByScenario {
			interfaceProcessorExecReq := s.RemoteService.GetProcessorInterfaceToExec(req)

			agentExec.Environment = interfaceProcessorExecReq.Environment
			agentExec.Variables = interfaceProcessorExecReq.Variables
			agentExec.DatapoolData = interfaceProcessorExecReq.Datapools

			ret, err = s.Request(interfaceProcessorExecReq)
			err = s.RemoteService.SubmitProcessorInterfaceResult(req, ret, req.ServerUrl, req.Token)

		}
	*/
	return
}

func (s *InterfaceService) Request(req v1.DebugRequest) (ret v1.DebugResponse, err error) {
	// exec pre-request script
	agentExec.ExecJs(req.PreRequestScript)

	// replace variables
	agentExec.ReplaceAll(&req.BaseRequest, agentExec.Environment, agentExec.Variables, agentExec.DatapoolData)

	// send request
	ret, err = agentExec.Invoke(&req.BaseRequest)

	ret.Id = req.InterfaceId

	return
}

func (s *InterfaceService) GetContentProps(ret *v1.DebugResponse) {
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

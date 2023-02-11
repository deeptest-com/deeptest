package service

import (
	"encoding/json"
	"fmt"
	"github.com/aaronchen2k/deeptest/cmd/agent/v1/domain"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	agentExec "github.com/aaronchen2k/deeptest/internal/agent/exec"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	httpHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/http"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/aaronchen2k/deeptest/pkg/lib/http"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"strings"
)

type InvocationService struct {
}

func (s *InvocationService) Invoke(req domain.InvocationReq) (ret v1.InvocationResponse, err error) {
	if req.UsedBy == consts.UsedByInterface {
		interfaceExecReq := s.getInterfaceToExec(req)

		agentExec.Environment = interfaceExecReq.Environment
		agentExec.Variables = interfaceExecReq.Variables
		agentExec.DatapoolData = interfaceExecReq.Datapools

		ret, err = s.Test(interfaceExecReq)
		err = s.SubmitInterfaceResult(req, ret, req.ServerUrl, req.Token)

	} else if req.UsedBy == consts.UsedByScenario {
		interfaceProcessorExecReq := s.getProcessorInterfaceToExec(req)

		agentExec.Environment = interfaceProcessorExecReq.Environment
		agentExec.Variables = interfaceProcessorExecReq.Variables
		agentExec.DatapoolData = interfaceProcessorExecReq.Datapools

		ret, err = s.Test(interfaceProcessorExecReq)
		err = s.SubmitProcessorInterfaceResult(req, ret, req.ServerUrl, req.Token)

	}

	return
}

func (s *InvocationService) getInterfaceToExec(req domain.InvocationReq) (ret v1.InvocationRequest) {
	url := fmt.Sprintf("invocations/loadInterfaceExecData")
	body, err := json.Marshal(req.Data)
	if err != nil {
		logUtils.Infof("marshal request data failed, error, %s", err.Error())
		return
	}

	httpReq := v1.BaseRequest{
		Url:               _httpUtils.AddSepIfNeeded(req.ServerUrl) + url,
		BodyType:          consts.ContentTypeJSON,
		Body:              string(body),
		AuthorizationType: consts.BearerToken,
		BearerToken: v1.BearerToken{
			Token: req.Token,
		},
	}

	resp, err := httpHelper.Post(httpReq)
	if err != nil {
		logUtils.Infof("get interface obj failed, error, %s", err.Error())
		return
	}

	if resp.StatusCode != consts.OK {
		logUtils.Infof("get interface obj failed, response %v", resp)
		return
	}

	respContent := _domain.Response{}
	json.Unmarshal([]byte(resp.Content), &respContent)

	if respContent.Code != 0 {
		logUtils.Infof("get interface obj failed, response %v", resp.Content)
		return
	}

	bytes, err := json.Marshal(respContent.Data)
	if respContent.Code != 0 {
		logUtils.Infof("get interface obj failed, response %v", resp.Content)
		return
	}

	json.Unmarshal(bytes, &ret)

	return
}

func (s *InvocationService) SubmitInterfaceResult(reqOjb domain.InvocationReq, repsObj v1.InvocationResponse, serverUrl, token string) (err error) {
	url := _httpUtils.AddSepIfNeeded(serverUrl) + fmt.Sprintf("invocations/submitInterfaceInvokeResult")

	data := v1.SubmitInvocationResultRequest{
		Request:  reqOjb.Data,
		Response: repsObj,
	}

	bodyBytes, _ := json.Marshal(data)

	req := v1.BaseRequest{
		Url:               url,
		BodyType:          consts.ContentTypeJSON,
		Body:              string(bodyBytes),
		AuthorizationType: consts.BearerToken,
		BearerToken: v1.BearerToken{
			Token: token,
		},
	}

	resp, err := httpHelper.Post(req)
	if err != nil {
		logUtils.Infof("submit result failed, error, %s", err.Error())
		return
	}

	if resp.StatusCode != consts.OK {
		logUtils.Infof("submit result failed, response %v", resp)
		return
	}

	ret := _domain.Response{}
	json.Unmarshal([]byte(resp.Content), &ret)

	if ret.Code != 0 {
		logUtils.Infof("submit result failed, response %v", resp.Content)
		return
	}

	return
}

func (s *InvocationService) getProcessorInterfaceToExec(req domain.InvocationReq) (ret v1.InvocationRequest) {
	url := fmt.Sprintf("processors/invocations/loadInterfaceExecData")
	body, err := json.Marshal(req.Data)
	if err != nil {
		logUtils.Infof("marshal request data failed, error, %s", err.Error())
		return
	}

	httpReq := v1.BaseRequest{
		Url:               _httpUtils.AddSepIfNeeded(req.ServerUrl) + url,
		BodyType:          consts.ContentTypeJSON,
		Body:              string(body),
		AuthorizationType: consts.BearerToken,
		BearerToken: v1.BearerToken{
			Token: req.Token,
		},
	}

	resp, err := httpHelper.Post(httpReq)
	if err != nil {
		logUtils.Infof("get interface obj failed, error, %s", err.Error())
		return
	}

	if resp.StatusCode != consts.OK {
		logUtils.Infof("get interface obj failed, response %v", resp)
		return
	}

	respContent := _domain.Response{}
	json.Unmarshal([]byte(resp.Content), &respContent)

	if respContent.Code != 0 {
		logUtils.Infof("get interface obj failed, response %v", resp.Content)
		return
	}

	bytes, err := json.Marshal(respContent.Data)
	if respContent.Code != 0 {
		logUtils.Infof("get interface obj failed, response %v", resp.Content)
		return
	}

	json.Unmarshal(bytes, &ret)

	return
}

func (s *InvocationService) SubmitProcessorInterfaceResult(reqOjb domain.InvocationReq, repsObj v1.InvocationResponse, serverUrl, token string) (err error) {
	url := _httpUtils.AddSepIfNeeded(serverUrl) + fmt.Sprintf("processors/invocations/submitInterfaceInvokeResult")

	data := v1.SubmitInvocationResultRequest{
		Request:  reqOjb.Data,
		Response: repsObj,
	}

	bodyBytes, _ := json.Marshal(data)

	req := v1.BaseRequest{
		Url:               url,
		BodyType:          consts.ContentTypeJSON,
		Body:              string(bodyBytes),
		AuthorizationType: consts.BearerToken,
		BearerToken: v1.BearerToken{
			Token: token,
		},
	}

	resp, err := httpHelper.Post(req)
	if err != nil {
		logUtils.Infof("submit result failed, error, %s", err.Error())
		return
	}

	if resp.StatusCode != consts.OK {
		logUtils.Infof("submit result failed, response %v", resp)
		return
	}

	ret := _domain.Response{}
	json.Unmarshal([]byte(resp.Content), &ret)

	if ret.Code != 0 {
		logUtils.Infof("submit result failed, response %v", resp.Content)
		return
	}

	return
}

func (s *InvocationService) Test(req v1.InvocationRequest) (ret v1.InvocationResponse, err error) {
	// exec pre-request script
	agentExec.ExecJs(req.PreRequestScript)

	// replace variables
	agentExec.ReplaceAll(&req.BaseRequest, agentExec.Environment, agentExec.Variables, agentExec.DatapoolData)

	// send request
	req.Url, err = _httpUtils.AddDefaultUrlSchema(req.Url)
	if err != nil {
		return
	}

	if req.Method == consts.GET {
		ret, err = httpHelper.Get(req.BaseRequest)
	} else if req.Method == consts.POST {
		ret, err = httpHelper.Post(req.BaseRequest)
	} else if req.Method == consts.PUT {
		ret, err = httpHelper.Put(req.BaseRequest)
	} else if req.Method == consts.DELETE {
		ret, err = httpHelper.Delete(req.BaseRequest)
	} else if req.Method == consts.PATCH {
		ret, err = httpHelper.Patch(req.BaseRequest)
	} else if req.Method == consts.HEAD {
		ret, err = httpHelper.Head(req.BaseRequest)
	} else if req.Method == consts.CONNECT {
		ret, err = httpHelper.Connect(req.BaseRequest)
	} else if req.Method == consts.OPTIONS {
		ret, err = httpHelper.Options(req.BaseRequest)
	} else if req.Method == consts.TRACE {
		ret, err = httpHelper.Trace(req.BaseRequest)
	}

	s.GetContentProps(&ret)

	ret.Id = req.Id

	return
}

func (s *InvocationService) GetContentProps(ret *v1.InvocationResponse) {
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

package service

import (
	"encoding/json"
	"fmt"
	"github.com/aaronchen2k/deeptest/cmd/agent/v1/domain"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	agentExec "github.com/aaronchen2k/deeptest/internal/agent/exec"
	agentDomain "github.com/aaronchen2k/deeptest/internal/agent/exec/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	httpHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/http"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/aaronchen2k/deeptest/pkg/lib/http"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
)

type RemoteService struct {
}

// for interface invocation in both endpoint and scenario
func (s *RemoteService) GetInterfaceToExec(req domain.InvokeCall) (ret v1.DebugData) {
	url := fmt.Sprintf("debugs/interface/load")
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
func (s *RemoteService) SubmitInterfaceResult(reqObj v1.DebugData, respObj v1.DebugResponse, serverUrl, token string) (err error) {
	url := fmt.Sprintf("debugs/invoke/submitResult")

	data := v1.SubmitDebugResultRequest{
		Request:  reqObj,
		Response: respObj,
	}

	bodyBytes, _ := json.Marshal(data)

	req := v1.BaseRequest{
		Url:               _httpUtils.AddSepIfNeeded(serverUrl) + url,
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

// for processor interface invocation
//func (s *RemoteService) GetProcessorInterfaceToExec(req domain.InvokeCall) (ret v1.DebugData) {
//	url := fmt.Sprintf("processors/invocations/loadInterfaceExecData")
//	body, err := json.Marshal(req.Data)
//	if err != nil {
//		logUtils.Infof("marshal request data failed, error, %s", err.Error())
//		return
//	}
//
//	httpReq := v1.BaseRequest{
//		Url:               _httpUtils.AddSepIfNeeded(req.ServerUrl) + url,
//		BodyType:          consts.ContentTypeJSON,
//		Body:              string(body),
//		AuthorizationType: consts.BearerToken,
//		BearerToken: v1.BearerToken{
//			Token: req.Token,
//		},
//	}
//
//	resp, err := httpHelper.Post(httpReq)
//	if err != nil {
//		logUtils.Infof("get interface obj failed, error, %s", err.Error())
//		return
//	}
//
//	if resp.StatusCode != consts.OK {
//		logUtils.Infof("get interface obj failed, response %v", resp)
//		return
//	}
//
//	respContent := _domain.Response{}
//	json.Unmarshal([]byte(resp.Content), &respContent)
//
//	if respContent.Code != 0 {
//		logUtils.Infof("get interface obj failed, response %v", resp.Content)
//		return
//	}
//
//	bytes, err := json.Marshal(respContent.Data)
//	if respContent.Code != 0 {
//		logUtils.Infof("get interface obj failed, response %v", resp.Content)
//		return
//	}
//
//	json.Unmarshal(bytes, &ret)
//
//	return
//}
//func (s *RemoteService) SubmitProcessorInterfaceResult(reqOjb domain.InvokeCall, repsObj v1.DebugResponse, serverUrl, token string) (err error) {
//	url := _httpUtils.AddSepIfNeeded(serverUrl) + fmt.Sprintf("processors/invocations/submitInterfaceInvokeResult")
//
//	data := v1.SubmitDebugResultRequest{
//		Request:  reqOjb.Data,
//		Response: repsObj,
//	}
//
//	bodyBytes, _ := json.Marshal(data)
//
//	req := v1.BaseRequest{
//		Url:               url,
//		BodyType:          consts.ContentTypeJSON,
//		Body:              string(bodyBytes),
//		AuthorizationType: consts.BearerToken,
//		BearerToken: v1.BearerToken{
//			Token: token,
//		},
//	}
//
//	resp, err := httpHelper.Post(req)
//	if err != nil {
//		logUtils.Infof("submit result failed, error, %s", err.Error())
//		return
//	}
//
//	if resp.StatusCode != consts.OK {
//		logUtils.Infof("submit result failed, response %v", resp)
//		return
//	}
//
//	ret := _domain.Response{}
//	json.Unmarshal([]byte(resp.Content), &ret)
//
//	if ret.Code != 0 {
//		logUtils.Infof("submit result failed, response %v", resp.Content)
//		return
//	}
//
//	return
//}

// for scenario exec
func (s *RemoteService) GetScenarioToExec(req *agentExec.ScenarioExecReq) (ret *agentExec.ScenarioExecObj) {
	url := "scenarios/exec/loadExecScenario"

	httpReq := v1.BaseRequest{
		Url:               _httpUtils.AddSepIfNeeded(req.ServerUrl) + url,
		AuthorizationType: consts.BearerToken,
		BearerToken: v1.BearerToken{
			Token: req.Token,
		},
		Params: []v1.Param{
			{
				Name:  "id",
				Value: fmt.Sprintf("%d", req.ScenarioId),
			},
		},
	}

	resp, err := httpHelper.Get(httpReq)
	if err != nil {
		logUtils.Infof("get exec obj failed, error, %s", err.Error())
		return
	}

	if resp.StatusCode != consts.OK {
		logUtils.Infof("get exec obj failed, response %v", resp)
		return
	}

	respContent := _domain.Response{}
	json.Unmarshal([]byte(resp.Content), &respContent)

	if respContent.Code != 0 {
		logUtils.Infof("get exec obj failed, response %v", resp.Content)
		return
	}

	bytes, err := json.Marshal(respContent.Data)
	if respContent.Code != 0 {
		logUtils.Infof("get exec obj failed, response %v", resp.Content)
		return
	}

	json.Unmarshal(bytes, &ret)

	ret.ServerUrl = req.ServerUrl
	ret.Token = req.Token

	return
}

func (s *RemoteService) SubmitScenarioResult(result agentDomain.ScenarioExecResult, scenarioId uint, serverUrl, token string) (
	report agentDomain.ReportSimple, err error) {

	bodyBytes, _ := json.Marshal(result)
	req := v1.BaseRequest{
		Url:               _httpUtils.AddSepIfNeeded(serverUrl) + fmt.Sprintf("scenarios/exec/submitResult/%d", scenarioId),
		Body:              string(bodyBytes),
		BodyType:          consts.ContentTypeJSON,
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

	reportContent, _ := json.Marshal(ret.Data)
	report = agentDomain.ReportSimple{}
	json.Unmarshal(reportContent, &report)

	return
}

// for plan exec
func (s *RemoteService) GetPlanToExec(req *agentExec.PlanExecReq) (ret *agentExec.PlanExecObj) {
	url := "plans/exec/loadExecPlan"

	httpReq := v1.BaseRequest{
		Url:               _httpUtils.AddSepIfNeeded(req.ServerUrl) + url,
		AuthorizationType: consts.BearerToken,
		BearerToken: v1.BearerToken{
			Token: req.Token,
		},
		Params: []v1.Param{
			{
				Name:  "id",
				Value: fmt.Sprintf("%d", req.PlanId),
			},
		},
	}

	resp, err := httpHelper.Get(httpReq)
	if err != nil {
		logUtils.Infof("get exec obj failed, error, %s", err.Error())
		return
	}

	if resp.StatusCode != consts.OK {
		logUtils.Infof("get exec obj failed, response %v", resp)
		return
	}

	respContent := _domain.Response{}
	json.Unmarshal([]byte(resp.Content), &respContent)

	if respContent.Code != 0 {
		logUtils.Infof("get exec obj failed, response %v", resp.Content)
		return
	}

	bytes, err := json.Marshal(respContent.Data)
	if respContent.Code != 0 {
		logUtils.Infof("get exec obj failed, response %v", resp.Content)
		return
	}

	json.Unmarshal(bytes, &ret)

	ret.ServerUrl = req.ServerUrl
	ret.Token = req.Token

	return
}

func (s *RemoteService) SubmitPlanResult(result agentDomain.PlanExecResult, planId uint, serverUrl, token string) (
	report agentDomain.ReportSimple, err error) {
	bodyBytes, _ := json.Marshal(result)
	req := v1.BaseRequest{
		Url:               _httpUtils.AddSepIfNeeded(serverUrl) + fmt.Sprintf("plans/exec/submitResult/%d", planId),
		Body:              string(bodyBytes),
		BodyType:          consts.ContentTypeJSON,
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

	reportContent, _ := json.Marshal(ret.Data)
	report = agentDomain.ReportSimple{}
	json.Unmarshal(reportContent, &report)

	return
}

func (s *RemoteService) GetMessageToExec(req *agentExec.MessageExecReq) (ret *agentExec.MessageExecObj) {
	url := "message/unreadCount"

	httpReq := v1.BaseRequest{
		Url:               _httpUtils.AddSepIfNeeded(req.ServerUrl) + url,
		AuthorizationType: consts.BearerToken,
		BearerToken: v1.BearerToken{
			Token: req.Token,
		},
	}

	resp, err := httpHelper.Get(httpReq)
	if err != nil {
		logUtils.Infof("get exec obj failed, error, %s", err.Error())
		return
	}

	if resp.StatusCode != consts.OK {
		logUtils.Infof("get exec obj failed, response %v", resp)
		return
	}

	respContent := _domain.Response{}
	json.Unmarshal([]byte(resp.Content), &respContent)

	if respContent.Code != 0 {
		logUtils.Infof("get exec obj failed, response %v", resp.Content)
		return
	}

	bytes, err := json.Marshal(respContent.Data)
	if respContent.Code != 0 {
		logUtils.Infof("get exec obj failed, response %v", resp.Content)
		return
	}

	json.Unmarshal(bytes, &ret)

	return
}

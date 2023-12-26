package service

import (
	"encoding/json"
	"fmt"
	v1 "github.com/aaronchen2k/deeptest/cmd/agent/v1/domain"
	agentExec "github.com/aaronchen2k/deeptest/internal/agent/exec"
	agentDomain "github.com/aaronchen2k/deeptest/internal/agent/exec/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	httpHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/http"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/aaronchen2k/deeptest/pkg/lib/http"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	_stringUtils "github.com/aaronchen2k/deeptest/pkg/lib/string"
)

// for interface invocation in both endpoint and scenario
func GetInterfaceToExec(req v1.InterfaceCall) (ret agentExec.InterfaceExecObj) {
	url := fmt.Sprintf("debugs/interface/loadForExec?currProjectId=%d", req.Data.ProjectId)
	body, err := json.Marshal(req.Data)
	if err != nil {
		logUtils.Infof("marshal request data failed, error, %s", err.Error())
		return
	}

	httpReq := domain.BaseRequest{
		Url:               _httpUtils.AddSepIfNeeded(req.ServerUrl) + url,
		BodyType:          consts.ContentTypeJSON,
		Body:              string(body),
		AuthorizationType: consts.BearerToken,
		BearerToken: domain.BearerToken{
			Token: req.Token,
		},
	}

	resp, err := httpHelper.Post(httpReq)
	if err != nil {
		logUtils.Infof("get interface obj failed, error, %s", err.Error())
		return
	}

	if resp.StatusCode != consts.OK.Int() {
		logUtils.Infof("get interface obj failed, response %v", resp)
		return
	}

	respContent := _domain.Response{}
	err = json.Unmarshal([]byte(resp.Content), &respContent)
	if err != nil {
		logUtils.Infof(err.Error())
	}

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

	// use the data from page if exist
	if req.Data.Method != "" {
		ret.DebugData = req.Data
	}

	return
}
func SubmitInterfaceResult(execObj agentExec.InterfaceExecObj, respObj *domain.DebugResponse, serverUrl, token string) (err error) {
	url := fmt.Sprintf("debugs/invoke/submitResult")

	data := domain.SubmitDebugResultRequest{
		Request:        execObj.DebugData,
		Response:       *respObj,
		PreConditions:  execObj.PreConditions,
		PostConditions: execObj.PostConditions,
	}

	bodyBytes, _ := json.Marshal(data)

	req := domain.BaseRequest{
		Url:               _httpUtils.AddSepIfNeeded(serverUrl) + url,
		BodyType:          consts.ContentTypeJSON,
		Body:              string(bodyBytes),
		AuthorizationType: consts.BearerToken,
		BearerToken: domain.BearerToken{
			Token: token,
		},
	}

	resp, err := httpHelper.Post(req)
	if err != nil {
		logUtils.Infof("submit result failed, error, %s", err.Error())
		return
	}

	if resp.StatusCode != consts.OK.Int() {
		logUtils.Infof("submit result failed, response %v", resp)
		return
	}

	ret := _domain.Response{}
	json.Unmarshal([]byte(resp.Content), &ret)

	if ret.Code != 0 {
		logUtils.Infof("submit result failed, response %v", resp.Content)
		return
	}
	respObj.InvokeId = uint(ret.Data.(float64))
	return
}

// for scenario exec
func GetScenarioToExec(req *agentExec.ScenarioExecReq) (ret *agentExec.ScenarioExecObj) {
	url := "scenarios/exec/loadExecScenario"

	httpReq := domain.BaseRequest{
		Url:               _httpUtils.AddSepIfNeeded(req.ServerUrl) + url,
		AuthorizationType: consts.BearerToken,
		BearerToken: domain.BearerToken{
			Token: req.Token,
		},
		QueryParams: &[]domain.Param{
			{
				Name:  "id",
				Value: fmt.Sprintf("%d", req.ScenarioId),
			},
			{
				Name:  "environmentId",
				Value: _stringUtils.IntToStr(req.EnvironmentId),
			},
		},
	}
	request, err := json.Marshal(httpReq)
	logUtils.Infof("get exec obj request, request: %s", string(request))

	resp, err := httpHelper.Get(httpReq)

	logUtils.Infof("get exec obj response, response: %s", resp.Content)

	if err != nil {
		logUtils.Infof("get exec obj failed, error, %s", err.Error())
		return
	}

	if resp.StatusCode != consts.OK.Int() {
		logUtils.Infof("get exec obj failed, response %v", resp)
		return
	}

	respContent := _domain.Response{}
	err = json.Unmarshal([]byte(resp.Content), &respContent)
	if err != nil {
		logUtils.Infof("get exec obj failed, err %v", err)
	}

	if respContent.Code != 0 {
		logUtils.Infof("get exec obj failed, response %v", resp.Content)
		return
	}

	bytes, err := json.Marshal(respContent.Data)
	if respContent.Code != 0 {
		logUtils.Infof("get exec obj failed, response %v", resp.Content)
		return
	}

	err = json.Unmarshal(bytes, &ret)
	if err != nil {
		logUtils.Infof("get exec obj failed,err:%v", err.Error())
	}
	response, _ := json.Marshal(ret)
	logUtils.Infof("get exec obj ret: %v", string(response))

	ret.ServerUrl = req.ServerUrl
	ret.Token = req.Token

	return
}

func GetScenarioNormalData(req *agentExec.ScenarioExecReq) (ret agentDomain.Report) {
	url := "scenarios/exec/getScenarioNormalData"

	httpReq := domain.BaseRequest{
		Url:               _httpUtils.AddSepIfNeeded(req.ServerUrl) + url,
		AuthorizationType: consts.BearerToken,
		BearerToken: domain.BearerToken{
			Token: req.Token,
		},
		QueryParams: &[]domain.Param{
			{
				Name:  "id",
				Value: fmt.Sprintf("%d", req.ScenarioId),
			},
			{
				Name:  "environmentId",
				Value: _stringUtils.IntToStr(req.EnvironmentId),
			},
		},
	}

	resp, err := httpHelper.Get(httpReq)
	if err != nil {
		logUtils.Infof("get exec obj failed, error, %s", err.Error())
		return
	}

	if resp.StatusCode != consts.OK.Int() {
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

func SubmitScenarioResult(result agentDomain.ScenarioExecResult, scenarioId uint, serverUrl, token string) (
	report agentDomain.Report, err error) {

	bodyBytes, _ := json.Marshal(result)
	req := domain.BaseRequest{
		Url:               _httpUtils.AddSepIfNeeded(serverUrl) + fmt.Sprintf("scenarios/exec/submitResult/%d", scenarioId),
		Body:              string(bodyBytes),
		BodyType:          consts.ContentTypeJSON,
		AuthorizationType: consts.BearerToken,
		BearerToken: domain.BearerToken{
			Token: token,
		},
	}

	resp, err := httpHelper.Post(req)
	if err != nil {
		logUtils.Infof("submit result failed, error, %s", err.Error())
		return
	}

	if resp.StatusCode != consts.OK.Int() {
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
	report = agentDomain.Report{}
	json.Unmarshal(reportContent, &report)

	return
}

// for plan exec
func GetPlanToExec(req *agentExec.PlanExecReq) (ret *agentExec.PlanExecObj) {
	ret = &agentExec.PlanExecObj{}
	url := "plans/exec/loadExecPlan"

	httpReq := domain.BaseRequest{
		Url:               _httpUtils.AddSepIfNeeded(req.ServerUrl) + url,
		AuthorizationType: consts.BearerToken,
		BearerToken: domain.BearerToken{
			Token: req.Token,
		},
		QueryParams: &[]domain.Param{
			{
				Name:  "id",
				Value: fmt.Sprintf("%d", req.PlanId),
			},
			{
				Name:  "environmentId",
				Value: _stringUtils.IntToStr(req.EnvironmentId),
			},
		},
	}

	resp, err := httpHelper.Get(httpReq)
	if err != nil {
		logUtils.Infof("get exec obj failed, error, %s", err.Error())
		return
	}

	if resp.StatusCode != consts.OK.Int() {
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

func SubmitPlanResult(result agentDomain.PlanExecResult, planId int, serverUrl, token string) (
	report agentDomain.Report, err error) {
	bodyBytes, _ := json.Marshal(result)
	req := domain.BaseRequest{
		Url:               _httpUtils.AddSepIfNeeded(serverUrl) + fmt.Sprintf("plans/exec/submitResult/%d", planId),
		Body:              string(bodyBytes),
		BodyType:          consts.ContentTypeJSON,
		AuthorizationType: consts.BearerToken,
		BearerToken: domain.BearerToken{
			Token: token,
		},
	}

	resp, err := httpHelper.Post(req)
	if err != nil {
		logUtils.Infof("submit result failed, error, %s", err.Error())
		return
	}

	if resp.StatusCode != consts.OK.Int() {
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
	report = agentDomain.Report{}
	json.Unmarshal(reportContent, &report)

	return
}

func GetMessageToExec(req *agentExec.MessageExecReq) (ret *agentExec.MessageExecObj) {
	url := "message/unreadCount"

	httpReq := domain.BaseRequest{
		Url:               _httpUtils.AddSepIfNeeded(req.ServerUrl) + url,
		AuthorizationType: consts.BearerToken,
		BearerToken: domain.BearerToken{
			Token: req.Token,
		},
	}

	resp, err := httpHelper.Get(httpReq)
	if err != nil {
		logUtils.Infof("get exec obj failed, error, %s", err.Error())
		return
	}

	if resp.StatusCode != consts.OK.Int() {
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

func GetPlanNormalData(req *agentExec.PlanExecReq) (ret agentDomain.Report, err error) {
	url := "plans/exec/getPlanReportNormalData"

	httpReq := domain.BaseRequest{
		Url:               _httpUtils.AddSepIfNeeded(req.ServerUrl) + url,
		AuthorizationType: consts.BearerToken,
		BearerToken: domain.BearerToken{
			Token: req.Token,
		},
		QueryParams: &[]domain.Param{
			{
				Name:  "id",
				Value: fmt.Sprintf("%d", req.PlanId),
			},
			{
				Name:  "environmentId",
				Value: _stringUtils.IntToStr(req.EnvironmentId),
			},
		},
	}

	resp, err := httpHelper.Get(httpReq)
	if err != nil {
		logUtils.Infof("get exec obj failed, error, %s", err.Error())
		return
	}

	if resp.StatusCode != consts.OK.Int() {
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

func GetCasesToExec(req *agentExec.CasesExecReq) (ret agentExec.CaseExecProcessor) {
	url := "endpoints/cases/alternatives/loadCasesForExec"

	body, err := json.Marshal(req)

	httpReq := domain.BaseRequest{
		Url:               _httpUtils.AddSepIfNeeded(req.ServerUrl) + url,
		AuthorizationType: consts.BearerToken,
		BearerToken: domain.BearerToken{
			Token: req.Token,
		},
		BodyType: consts.ContentTypeJSON,
		Body:     string(body),
	}
	request, err := json.Marshal(httpReq)
	logUtils.Infof("get case exec obj request, request: %s", string(request))

	resp, err := httpHelper.Post(httpReq)
	if err != nil {
		logUtils.Infof("get interface obj failed, error, %s", err.Error())
		return
	}

	if resp.StatusCode != consts.OK.Int() {
		logUtils.Infof("get interface obj failed, response %v", resp)
		return
	}

	respContent := _domain.Response{}
	err = json.Unmarshal([]byte(resp.Content), &respContent)
	if err != nil {
		logUtils.Infof(err.Error())
	}

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

package service

import (
	"encoding/json"
	"fmt"
	integrationDomain "github.com/aaronchen2k/deeptest/integration/domain"
	v1 "github.com/aaronchen2k/deeptest/integration/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	httpHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/http"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/jinzhu/copier"
)

type RemoteService struct {
}

func (s *RemoteService) GetUserInfoByToken(token, origin string) (user v1.UserInfo, err error) {
	url := fmt.Sprintf("%s/levault/usrsvr/Usr/GetUserInfoByToken", origin)
	body, _ := json.Marshal(map[string]string{"token": token})
	headers := make([]domain.Header, 0)
	headers = append(headers, []domain.Header{
		{
			Name:  "Tenant-Id",
			Value: "1632931640315338752", //TODO 做saas后可以考虑提到配置文件里
		},
		{
			Name:  "Ds-Tenant-Id",
			Value: "0",
		},
		{
			Name:  "Token",
			Value: token,
		},
		{
			Name:  "lang",
			Value: "zh_cn",
		}}...,
	)
	httpReq := domain.BaseRequest{
		Url:      url,
		BodyType: consts.ContentTypeJSON,
		Headers:  &headers,
		Body:     string(body),
	}

	var resp domain.DebugResponse
	resp, err = httpHelper.Post(httpReq)
	if err != nil {
		logUtils.Infof("meta get method detail failed, error, %s", err.Error())
		return
	}

	if resp.StatusCode != consts.OK.Int() {
		logUtils.Infof("meta get method detail failed, response %v", resp)
		err = fmt.Errorf("meta get method detail failed, response %v", resp)
		return
	}

	respContent := struct {
		Code int
		Data v1.OtherUserInfo
		Msg  string
	}{}
	err = json.Unmarshal([]byte(resp.Content), &respContent)
	if err != nil {
		logUtils.Infof(err.Error())
	}

	copier.CopyWithOption(&user, respContent.Data, copier.Option{DeepCopy: true})

	return
}

func (s *RemoteService) LoginByOauth(loginByOauthReq integrationDomain.LoginByOauthReq, baseUrl string) (ret integrationDomain.LoginByOauthResData) {
	req := struct {
		ThisObj integrationDomain.LoginByOauthReq `json:"thisObj"`
	}{}
	req.ThisObj = loginByOauthReq

	url := fmt.Sprintf("%s/levault/usrsvr/Usr/LoginByOauth", baseUrl)
	body, err := json.Marshal(req)
	if err != nil {
		logUtils.Infof("marshal request data failed, error, %s", err.Error())
		return
	}

	headers := s.getLcHeaders("")
	httpReq := domain.BaseRequest{
		Url:      url,
		BodyType: consts.ContentTypeJSON,
		Body:     string(body),
		Headers:  &headers,
	}

	resp, err := httpHelper.Post(httpReq)
	if err != nil {
		logUtils.Infof("login by oauth failed, error, %s", err.Error())
		return
	}

	if resp.StatusCode != consts.OK.Int() {
		logUtils.Infof("login by oauth failed, response %v", resp)
		return
	}

	respContent := integrationDomain.LoginByOauthRes{}
	err = json.Unmarshal([]byte(resp.Content), &respContent)
	if err != nil {
		logUtils.Infof(err.Error())
	}

	if respContent.Mfail != "0" {
		logUtils.Infof("login by oauth failed, response %v", resp.Content)
		return
	}

	ret = respContent.Data

	return
}

func (s *RemoteService) getLcHeaders(token string) (headers []domain.Header) {
	headers = []domain.Header{
		{
			Name:  "Tenant-Id",
			Value: "1632931640315338752", //TODO 做saas后可以考虑提到配置文件里
		},
		{
			Name:  "Ds-Tenant-Id",
			Value: "0",
		},
		{
			Name:  "Token",
			Value: token,
		},
		{
			Name:  "lang",
			Value: "zh_cn",
		},
	}

	return
}

func (s *RemoteService) GetTokenFromCode(req integrationDomain.GetTokenFromCodeReq, baseUrl string) (ret integrationDomain.GetTokenFromCodeResData) {
	url := fmt.Sprintf("%s/levault/usrsvr/Usr/GetTokenFromCode", baseUrl)
	body, err := json.Marshal(req)
	if err != nil {
		logUtils.Infof("marshal request data failed, error, %s", err.Error())
		return
	}

	headers := s.getLcHeaders("")
	httpReq := domain.BaseRequest{
		Url:      url,
		BodyType: consts.ContentTypeJSON,
		Body:     string(body),
		Headers:  &headers,
	}

	resp, err := httpHelper.Post(httpReq)
	if err != nil {
		logUtils.Infof("get token from code failed, error, %s", err.Error())
		return
	}

	if resp.StatusCode != consts.OK.Int() {
		logUtils.Infof("get token from code failed, response %v", resp)
		return
	}

	respContent := integrationDomain.GetTokenFromCodeRes{}
	err = json.Unmarshal([]byte(resp.Content), &respContent)
	if err != nil {
		logUtils.Infof(err.Error())
	}

	if respContent.Mfail != "0" {
		logUtils.Infof("get token from code failed, response %v", resp.Content)
		return
	}

	ret = respContent.Data

	return
}

func (s *RemoteService) FindClassByServiceCode(req integrationDomain.FindClassByServiceCodeReq, token string, baseUrl string) (ret []integrationDomain.FindClassByServiceCodeResData) {
	url := fmt.Sprintf("%s/levault/agentdesigner/classInfo/findByServiceCode", baseUrl)
	body, err := json.Marshal(req)
	if err != nil {
		logUtils.Infof("marshal request data failed, error, %s", err.Error())
		return
	}

	headers := s.getLcHeaders(token)

	httpReq := domain.BaseRequest{
		Url:      url,
		BodyType: consts.ContentTypeJSON,
		Headers:  &headers,
		Body:     string(body),
	}

	resp, err := httpHelper.Post(httpReq)
	if err != nil {
		logUtils.Infof("find class by serviceCode failed, error, %s", err.Error())
		return
	}

	if resp.StatusCode != consts.OK.Int() {
		logUtils.Infof("find class by serviceCode failed, response %v", resp)
		return
	}

	respContent := integrationDomain.FindClassByServiceCodeRes{}
	err = json.Unmarshal([]byte(resp.Content), &respContent)
	if err != nil {
		logUtils.Infof(err.Error())
	}

	if respContent.Mfail != "0" {
		logUtils.Infof("find class by serviceCode failed, response %v", resp.Content)
		return
	}

	ret = respContent.Data

	return
}

func (s *RemoteService) GetFunctionsByClass(req integrationDomain.GetFunctionsByClassReq, token string, baseUrl string) (ret []integrationDomain.GetFunctionsByClassResData) {
	url := fmt.Sprintf("%s/levault/agentdesigner/classMethod/listData", baseUrl)
	body, err := json.Marshal(req)
	if err != nil {
		logUtils.Infof("marshal request data failed, error, %s", err.Error())
		return
	}

	headers := s.getLcHeaders(token)

	httpReq := domain.BaseRequest{
		Url:      url,
		BodyType: consts.ContentTypeJSON,
		Headers:  &headers,
		Body:     string(body),
	}

	resp, err := httpHelper.Post(httpReq)
	if err != nil {
		logUtils.Infof("get functions by class failed, error, %s", err.Error())
		return
	}

	if resp.StatusCode != consts.OK.Int() {
		logUtils.Infof("get functions by class failed, response %v", resp)
		return
	}

	respContent := integrationDomain.GetFunctionsByClassRes{}
	err = json.Unmarshal([]byte(resp.Content), &respContent)
	if err != nil {
		logUtils.Infof(err.Error())
	}

	if respContent.Mfail != "0" {
		logUtils.Infof("get functions by class failed, response %v", resp.Content)
		return
	}

	ret = respContent.Data

	return
}

func (s *RemoteService) MetaGetMethodDetail(req integrationDomain.MetaGetMethodDetailReq, token string, baseUrl string) (ret integrationDomain.MetaGetMethodDetailResData) {
	url := fmt.Sprintf("%s/levault/meta/metaClassMethod/metaGetClassMethodDetail", baseUrl)
	body, err := json.Marshal(req)
	if err != nil {
		logUtils.Infof("marshal request data failed, error, %s", err.Error())
		return
	}

	headers := s.getLcHeaders(token)
	headers = append(headers, domain.Header{
		Name:  "Token",
		Value: token,
	})

	httpReq := domain.BaseRequest{
		Url:      url,
		BodyType: consts.ContentTypeJSON,
		Headers:  &headers,
		Body:     string(body),
	}

	resp, err := httpHelper.Post(httpReq)
	if err != nil {
		logUtils.Infof("meta get method detail failed, error, %s", err.Error())
		return
	}

	if resp.StatusCode != consts.OK.Int() {
		logUtils.Infof("meta get method detail failed, response %v", resp)
		return
	}

	respContent := integrationDomain.MetaGetMethodDetailRes{}
	err = json.Unmarshal([]byte(resp.Content), &respContent)
	if err != nil {
		logUtils.Infof(err.Error())
	}

	if respContent.Mfail != "0" {
		logUtils.Infof("meta get method detail failed, response %v", resp.Content)
		return
	}

	ret = respContent.Data

	return
}

func (s *RemoteService) GetFunctionDetailsByClass(classCode string, token string, baseUrl string) (ret []integrationDomain.GetFunctionDetailsByClassResData, err error) {
	url := fmt.Sprintf("%s/levault/meta/metaClassMethod/metaGetClassMessages", baseUrl)

	headers := s.getLcHeaders(token)
	httpReq := domain.BaseRequest{
		Url:      url,
		BodyType: consts.ContentTypeJSON,
		Headers:  &headers,
		QueryParams: &[]domain.Param{
			{
				Name:  "className",
				Value: classCode,
			},
		},
	}

	resp, err := httpHelper.Post(httpReq)
	if err != nil {
		logUtils.Infof("get functionDetails by class failed, error, %s", err.Error())
		return
	}

	if resp.StatusCode != consts.OK.Int() {
		logUtils.Infof("get functionDetails by class failed, response %v", resp)
		return
	}

	respContent := integrationDomain.GetFunctionDetailsByClassRes{}
	err = json.Unmarshal([]byte(resp.Content), &respContent)
	if err != nil {
		logUtils.Infof(err.Error())
		return
	}

	if respContent.Mfail != "0" {
		logUtils.Infof("get functions by class failed, response %v", resp.Content)
		return
	}

	ret = respContent.Data

	return
}

func (s *RemoteService) LcContainerQueryAgent(token, baseUrl string) (ret []integrationDomain.EngineeringItem) {
	url := fmt.Sprintf("%s/levault/acnsvr/Container/QueryAgent", baseUrl)
	req := s.getLcContainerQueryAgentRequest()
	body, err := json.Marshal(req)
	if err != nil {
		logUtils.Infof("marshal request data failed, error, %s", err.Error())
		return
	}

	headers := s.getLcHeaders(token)
	httpReq := domain.BaseRequest{
		Url:      url,
		BodyType: consts.ContentTypeJSON,
		Headers:  &headers,
		Body:     string(body),
	}

	resp, err := httpHelper.Post(httpReq)
	if err != nil {
		logUtils.Infof("LcContainerQueryAgent failed, error, %s", err.Error())
		return
	}

	if resp.StatusCode != consts.OK.Int() {
		logUtils.Infof("LcContainerQueryAgent failed, response %v", resp)
		return
	}

	respContent := integrationDomain.ContainerQueryAgentRes{}
	err = json.Unmarshal([]byte(resp.Content), &respContent)
	if err != nil {
		logUtils.Infof(err.Error())
	}

	if respContent.Mfail != "0" {
		logUtils.Infof("LcContainerQueryAgent failed, response %v", resp.Content)
		return
	}

	ret = respContent.Data

	return
}

func (s *RemoteService) getLcContainerQueryAgentRequest() (res integrationDomain.QueryAgentReq) {
	attrSet := []string{"displayClassName", "displayName", "code", "name", "projectManagerName", "lifeCycleState", "displayCreator", "createAt", "lastUpdate", "container", "objId", "className", "code", "tenantId"}
	conditionParam := integrationDomain.QueryAgentConditionParam{
		Key:     "containerRange",
		Compare: "LIKE",
		Value:   "all",
	}

	queryArgs := integrationDomain.QueryAgentQueryArgs{
		AttrSet:   attrSet,
		Condition: []integrationDomain.QueryAgentConditionParam{conditionParam},
	}

	queryArgs.Sort.SortBy = "createAt"
	queryArgs.Sort.SortOrder = "desc"

	res.ClassName = "Container"
	res.QueryArgs = queryArgs

	return
}

func (s *RemoteService) LcMlClassQueryAgent(serviceCode, token, baseUrl string) (ret []integrationDomain.FindClassByServiceCodeResData) {
	url := fmt.Sprintf("%s/levault/mdlsvr/MlClass/QueryAgent", baseUrl)
	req := s.getLcMlClassQueryAgentRequest(serviceCode)
	body, err := json.Marshal(req)
	if err != nil {
		logUtils.Infof("marshal request data failed, error, %s", err.Error())
		return
	}

	headers := s.getLcHeaders(token)
	httpReq := domain.BaseRequest{
		Url:      url,
		BodyType: consts.ContentTypeJSON,
		Headers:  &headers,
		Body:     string(body),
	}

	resp, err := httpHelper.Post(httpReq)
	if err != nil {
		logUtils.Infof("LcQueryAgent failed, error, %s", err.Error())
		return
	}

	if resp.StatusCode != consts.OK.Int() {
		logUtils.Infof("LcQueryAgent failed, response %v", resp)
		return
	}

	respContent := integrationDomain.MlClassQueryAgentRes{}
	err = json.Unmarshal([]byte(resp.Content), &respContent)
	if err != nil {
		logUtils.Infof(err.Error())
	}

	if respContent.Mfail != "0" {
		logUtils.Infof("LcQueryAgent failed, response %v", resp.Content)
		return
	}

	ret = respContent.Data.Data

	return
}

func (s *RemoteService) getLcMlClassQueryAgentRequest(serviceCode string) (res integrationDomain.QueryAgentReq) {
	attrSet := []string{"objId", "code", "parentCode", "parentCodes", "businessClassType", "container", "lastUpdate", "remark", "rightClassCode", "rightClassName", "rightRelationShip", "rightRelationShipName", "leftClassCode", "leftClassName", "leftRelationShip", "leftRelationShipName", "serviceId", "type", "dialogSource", "className", "classIcon", "serviceCode", "name", "displayName", "displayClassName", "displayCreator", "displayModifier"}
	conditionParam := integrationDomain.QueryAgentConditionParam{
		Key:     "serviceCode",
		Compare: "EQ",
		Value:   serviceCode,
	}

	queryArgs := integrationDomain.QueryAgentQueryArgs{
		AttrSet:   attrSet,
		Condition: []integrationDomain.QueryAgentConditionParam{conditionParam},
	}

	queryArgs.Sort.SortBy = "code"
	queryArgs.Sort.SortOrder = "asc"

	res.ClassName = "MlClass"
	res.QueryArgs = queryArgs

	return
}

func (s *RemoteService) LcQueryMsg(req integrationDomain.QueryMsgReq, token string, baseUrl string) (ret []integrationDomain.GetFunctionsByClassResData) {
	url := fmt.Sprintf("%s/levault/mdlsvr/ClsMsg/QueryMsg", baseUrl)
	body, err := json.Marshal(req)
	if err != nil {
		logUtils.Infof("marshal request data failed, error, %s", err.Error())
		return
	}

	headers := s.getLcHeaders(token)

	httpReq := domain.BaseRequest{
		Url:      url,
		BodyType: consts.ContentTypeJSON,
		Headers:  &headers,
		Body:     string(body),
	}

	resp, err := httpHelper.Post(httpReq)
	if err != nil {
		logUtils.Infof("LcQueryMsg failed, error, %s", err.Error())
		return
	}

	if resp.StatusCode != consts.OK.Int() {
		logUtils.Infof("LcQueryMsg failed, response %v", resp)
		return
	}

	respContent := integrationDomain.GetFunctionsByClassRes{}
	err = json.Unmarshal([]byte(resp.Content), &respContent)
	if err != nil {
		logUtils.Infof(err.Error())
	}

	if respContent.Mfail != "0" {
		logUtils.Infof("LcQueryMsg failed, response %v", resp.Content)
		return
	}

	ret = respContent.Data

	return
}

func (s *RemoteService) LcMlServiceQueryAgent(engineering, token, baseUrl string) (ret []integrationDomain.ServiceItem) {
	url := fmt.Sprintf("%s/levault/mdlsvr/MlService/QueryAgent", baseUrl)
	req := s.getLcMlServiceQueryAgentRequest(engineering)
	body, err := json.Marshal(req)
	if err != nil {
		logUtils.Infof("marshal request data failed, error, %s", err.Error())
		return
	}

	headers := s.getLcHeaders(token)
	httpReq := domain.BaseRequest{
		Url:      url,
		BodyType: consts.ContentTypeJSON,
		Headers:  &headers,
		Body:     string(body),
	}

	resp, err := httpHelper.Post(httpReq)
	if err != nil {
		logUtils.Infof("LcMlServiceQueryAgent failed, error, %s", err.Error())
		return
	}

	if resp.StatusCode != consts.OK.Int() {
		logUtils.Infof("LcMlServiceQueryAgent failed, response %v", resp)
		return
	}

	respContent := integrationDomain.MlServiceQueryAgentRes{}
	err = json.Unmarshal([]byte(resp.Content), &respContent)
	if err != nil {
		logUtils.Infof(err.Error())
	}

	if respContent.Mfail != "0" {
		logUtils.Infof("LcMlServiceQueryAgent failed, response %v", resp.Content)
		return
	}

	ret = respContent.Data

	return
}

func (s *RemoteService) LcAllServiceList(token, baseUrl string) (ret []integrationDomain.ServiceItem) {
	url := fmt.Sprintf("%s/levault/mdlsvr/MlMessage/QueryListGroupClass", baseUrl)
	req := struct {
		ThisObj struct {
			Content string `json:"content"`
		} `json:"thisObj"`
	}{}

	body, err := json.Marshal(req)
	if err != nil {
		logUtils.Infof("marshal request data failed, error, %s", err.Error())
		return
	}

	headers := s.getLcHeaders(token)
	httpReq := domain.BaseRequest{
		Url:      url,
		BodyType: consts.ContentTypeJSON,
		Headers:  &headers,
		Body:     string(body),
	}

	resp, err := httpHelper.Post(httpReq)
	if err != nil {
		logUtils.Infof("LcAllServiceList failed, error, %s", err.Error())
		return
	}

	if resp.StatusCode != consts.OK.Int() {
		logUtils.Infof("LcAllServiceList failed, response %v", resp)
		return
	}

	respContent := integrationDomain.MlServiceQueryAgentRes{}
	err = json.Unmarshal([]byte(resp.Content), &respContent)
	if err != nil {
		logUtils.Infof(err.Error())
	}

	if respContent.Mfail != "0" {
		logUtils.Infof("LcAllServiceList failed, response %v", resp.Content)
		return
	}

	ret = respContent.Data

	return
}

func (s *RemoteService) getLcMlServiceQueryAgentRequest(engineering string) (res integrationDomain.QueryAgentReq) {
	attrSet := []string{"className", "code", "langPrefix", "container", "codePrefix", "creator", "createAt", "modifier", "lastUpdate", "developMethod", "lifeCycleState", "objId", "tenantId", "name", "displayName", "displayClassName", "displayCreator", "displayModifier"}
	conditionParam := integrationDomain.QueryAgentConditionParam{
		Key:     "container",
		Compare: "EQ",
		Value:   engineering,
	}

	queryArgs := integrationDomain.QueryAgentQueryArgs{
		AttrSet:   attrSet,
		Condition: []integrationDomain.QueryAgentConditionParam{conditionParam},
	}

	queryArgs.Sort.SortBy = "lastUpdate"
	queryArgs.Sort.SortOrder = "desc"

	res.ClassName = "MlService"
	res.QueryArgs = queryArgs

	return
}

func (s *RemoteService) GetLovByCode() {
	
}

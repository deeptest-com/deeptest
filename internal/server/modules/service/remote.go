package service

import (
	"encoding/json"
	"fmt"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/config"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	httpHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/http"
	_commUtils "github.com/aaronchen2k/deeptest/pkg/lib/comm"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"strconv"
	"time"
)

type RemoteService struct {
}

func (s *RemoteService) LoginByOauth(req v1.LoginByOauthReq, baseUrl string) (ret v1.LoginByOauthResData) {
	url := fmt.Sprintf("%s/levault/usrsvr/Usr/LoginByOauth", baseUrl)
	body, err := json.Marshal(req)
	if err != nil {
		logUtils.Infof("marshal request data failed, error, %s", err.Error())
		return
	}

	httpReq := domain.BaseRequest{
		Url:      url,
		BodyType: consts.ContentTypeJSON,
		Body:     string(body),
		Headers: &[]domain.Header{
			{
				Name:  "Tenant-Id",
				Value: "0",
			},
		},
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

	respContent := v1.LoginByOauthRes{}
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

func (s *RemoteService) GetTokenFromCode(req v1.GetTokenFromCodeReq, baseUrl string) (ret v1.GetTokenFromCodeResData) {
	url := fmt.Sprintf("%s/levault/usrsvr/Usr/GetTokenFromCode", baseUrl)
	body, err := json.Marshal(req)
	if err != nil {
		logUtils.Infof("marshal request data failed, error, %s", err.Error())
		return
	}

	httpReq := domain.BaseRequest{
		Url:      url,
		BodyType: consts.ContentTypeJSON,
		Body:     string(body),
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

	respContent := v1.GetTokenFromCodeRes{}
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

func (s *RemoteService) FindClassByServiceCode(req v1.FindClassByServiceCodeReq, token string, baseUrl string) (ret []v1.FindClassByServiceCodeResData) {
	url := fmt.Sprintf("%s/levault/agentdesigner/classInfo/findByServiceCode", baseUrl)
	body, err := json.Marshal(req)
	if err != nil {
		logUtils.Infof("marshal request data failed, error, %s", err.Error())
		return
	}

	headers := make([]domain.Header, 0)
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
		logUtils.Infof("find class by serviceCode failed, error, %s", err.Error())
		return
	}

	if resp.StatusCode != consts.OK.Int() {
		logUtils.Infof("find class by serviceCode failed, response %v", resp)
		return
	}

	respContent := v1.FindClassByServiceCodeRes{}
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

func (s *RemoteService) GetFunctionsByClass(req v1.GetFunctionsByClassReq, token string, baseUrl string) (ret []v1.GetFunctionsByClassResData) {
	url := fmt.Sprintf("%s/levault/agentdesigner/classMethod/listData", baseUrl)
	body, err := json.Marshal(req)
	if err != nil {
		logUtils.Infof("marshal request data failed, error, %s", err.Error())
		return
	}

	headers := make([]domain.Header, 0)
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
		logUtils.Infof("get functions by class failed, error, %s", err.Error())
		return
	}

	if resp.StatusCode != consts.OK.Int() {
		logUtils.Infof("get functions by class failed, response %v", resp)
		return
	}

	respContent := v1.GetFunctionsByClassRes{}
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

func (s *RemoteService) MetaGetMethodDetail(req v1.MetaGetMethodDetailReq, token string, baseUrl string) (ret v1.MetaGetMethodDetailResData) {
	url := fmt.Sprintf("%s/levault/meta/metaClassMethod/metaGetMethodDetail", baseUrl)
	body, err := json.Marshal(req)
	if err != nil {
		logUtils.Infof("marshal request data failed, error, %s", err.Error())
		return
	}

	headers := make([]domain.Header, 0)
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

	respContent := v1.MetaGetMethodDetailRes{}
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

func (s *RemoteService) GetFunctionDetailsByClass(classCode string, token string, baseUrl string) (ret []v1.GetFunctionDetailsByClassResData, err error) {
	url := fmt.Sprintf("%s/levault/meta/metaClassMethod/metaGetClassMessages", baseUrl)

	headers := make([]domain.Header, 0)
	headers = append(headers, domain.Header{
		Name:  "Token",
		Value: token,
	})

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

	respContent := v1.GetFunctionDetailsByClassRes{}
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

func (s *RemoteService) GetUserInfoByToken(token string) (user v1.UserInfo, err error) {
	url := fmt.Sprintf("%s/api/v1/user/getUserInfo", config.CONFIG.ThirdParty.Url)

	headers := make([]domain.Header, 0)
	headers = append(headers, domain.Header{
		Name:  "X-Token",
		Value: token,
	})

	httpReq := domain.BaseRequest{
		Url:      url,
		BodyType: consts.ContentTypeJSON,
		Headers:  &headers,
	}

	var resp domain.DebugResponse
	resp, err = httpHelper.Get(httpReq)
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
		Data struct{ UserInfo v1.UserInfo }
		Msg  string
	}{}
	err = json.Unmarshal([]byte(resp.Content), &respContent)
	if err != nil {
		logUtils.Infof(err.Error())
	}

	user = respContent.Data.UserInfo

	return
}

func (s *RemoteService) GetProjectInfo(token, spaceCode string) (ret v1.ProjectInfo, err error) {
	url := fmt.Sprintf("%s/api/v1/project/info/%s", config.CONFIG.ThirdParty.Url, spaceCode)

	httpReq := domain.BaseRequest{
		Url:      url,
		BodyType: consts.ContentTypeJSON,
		Headers: &[]domain.Header{
			{
				Name:  "X-Token",
				Value: token,
			},
		},
	}

	resp, err := httpHelper.Get(httpReq)
	if err != nil {
		logUtils.Infof("get project info failed, error, %s", err.Error())
		return
	}

	if resp.StatusCode != consts.OK.Int() {
		logUtils.Infof("get project info failed, response %v", resp)
		err = fmt.Errorf("get project info failed, response %v", resp)
		return
	}

	respContent := struct {
		Code int
		Data struct{ v1.ProjectInfo }
		Msg  string
	}{}
	err = json.Unmarshal([]byte(resp.Content), &respContent)
	if err != nil {
		logUtils.Infof(err.Error())
	}

	if respContent.Code != 200 {
		logUtils.Infof("get project info failed, response %v", resp)
		err = fmt.Errorf("get project info failed, response %v", resp)
		return
	}

	ret = respContent.Data.ProjectInfo

	return
}

func (s *RemoteService) GetUserButtonPermissions(username, spaceCode string) (ret []string, err error) {
	url := fmt.Sprintf("%s/api/v1/openApi/getUserDynamicMenuPermission", config.CONFIG.ThirdParty.Url)

	xNancalTimestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
	xNancalNonceStr := _commUtils.RandStr(8)
	httpReq := domain.BaseRequest{
		Url:      url,
		BodyType: consts.ContentTypeJSON,
		Headers: &[]domain.Header{
			{
				Name:  "x-nancal-appkey",
				Value: config.CONFIG.ThirdParty.ApiSign.AppKey,
			},
			{
				Name:  "x-nancal-timestamp",
				Value: xNancalTimestamp,
			},
			{
				Name:  "x-nancal-nonce-str",
				Value: xNancalNonceStr,
			},
			{
				Name:  "x-nancal-sign",
				Value: _commUtils.GetSign(config.CONFIG.ThirdParty.ApiSign.AppKey, config.CONFIG.ThirdParty.ApiSign.AppSecret, xNancalNonceStr, xNancalTimestamp, ""),
			},
		},
		QueryParams: &[]domain.Param{
			{
				Name:  "typeStr",
				Value: "[20,30]",
			},
			{
				Name:  "nameEng",
				Value: spaceCode,
			},
			{
				Name:  "username",
				Value: username,
			},
		},
	}

	resp, err := httpHelper.Get(httpReq)
	if err != nil {
		logUtils.Infof("get UserButtonPermissions failed, error, %s", err.Error())
		return
	}

	if resp.StatusCode != consts.OK.Int() {
		logUtils.Infof("get UserButtonPermissions failed, response %v", resp)
		err = fmt.Errorf("get UserButtonPermissions failed, response %v", resp)
		return
	}

	respContent := struct {
		Code int
		Data []string
		Msg  string
	}{}
	err = json.Unmarshal([]byte(resp.Content), &respContent)
	if err != nil {
		logUtils.Infof(err.Error())
	}

	if respContent.Code != 200 {
		logUtils.Infof("getUserButtonPermissions failed, response %v", resp)
		err = fmt.Errorf("get UserButtonPermissions failed, response %v", resp)
		return
	}

	ret = respContent.Data

	return
}

func (s *RemoteService) GetUserMenuPermissions(username, spaceCode string) (ret []v1.UserMenuPermission, err error) {
	url := fmt.Sprintf("%s/api/v1/openApi/getUserDynamicMenu", config.CONFIG.ThirdParty.Url)

	xNancalTimestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
	xNancalNonceStr := _commUtils.RandStr(8)
	httpReq := domain.BaseRequest{
		Url:      url,
		BodyType: consts.ContentTypeJSON,
		Headers: &[]domain.Header{
			{
				Name:  "x-nancal-appkey",
				Value: config.CONFIG.ThirdParty.ApiSign.AppKey,
			},
			{
				Name:  "x-nancal-timestamp",
				Value: xNancalTimestamp,
			},
			{
				Name:  "x-nancal-nonce-str",
				Value: xNancalNonceStr,
			},
			{
				Name:  "x-nancal-sign",
				Value: _commUtils.GetSign(config.CONFIG.ThirdParty.ApiSign.AppKey, config.CONFIG.ThirdParty.ApiSign.AppSecret, xNancalNonceStr, xNancalTimestamp, ""),
			},
		},
		QueryParams: &[]domain.Param{
			{
				Name:  "typeStr",
				Value: "[10,20]",
			},
			{
				Name:  "nameEng",
				Value: spaceCode,
			},
			{
				Name:  "username",
				Value: username,
			},
		},
	}

	resp, err := httpHelper.Get(httpReq)
	if err != nil {
		logUtils.Infof("get GetUserMenuPermissions failed, error, %s", err.Error())
		return
	}

	if resp.StatusCode != consts.OK.Int() {
		logUtils.Infof("get GetUserMenuPermissions failed, response %v", resp)
		err = fmt.Errorf("get GetUserMenuPermissions failed, response %v", resp)
		return
	}

	respContent := struct {
		Code int
		Data []v1.UserMenuPermission
		Msg  string
	}{}
	err = json.Unmarshal([]byte(resp.Content), &respContent)
	if err != nil {
		logUtils.Infof(err.Error())
	}

	if respContent.Code != 200 {
		logUtils.Infof("getUserButtonPermissions failed, response %v", resp)
		err = fmt.Errorf("get UserButtonPermissions failed, response %v", resp)
		return
	}

	ret = respContent.Data

	return
}

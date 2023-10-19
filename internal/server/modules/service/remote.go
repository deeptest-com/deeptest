package service

import (
	"encoding/json"
	"fmt"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	httpHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/http"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
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
	}

	resp, err := httpHelper.Post(httpReq)
	if err != nil {
		logUtils.Infof("login by oauth failed, error, %s", err.Error())
		return
	}

	if resp.StatusCode != consts.OK {
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

	if resp.StatusCode != consts.OK {
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
		Headers:  headers,
		Body:     string(body),
	}

	resp, err := httpHelper.Post(httpReq)
	if err != nil {
		logUtils.Infof("find class by serviceCode failed, error, %s", err.Error())
		return
	}

	if resp.StatusCode != consts.OK {
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
		Headers:  headers,
		Body:     string(body),
	}

	resp, err := httpHelper.Post(httpReq)
	if err != nil {
		logUtils.Infof("get functions by class failed, error, %s", err.Error())
		return
	}

	if resp.StatusCode != consts.OK {
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
		Headers:  headers,
		Body:     string(body),
	}

	resp, err := httpHelper.Post(httpReq)
	if err != nil {
		logUtils.Infof("meta get method detail failed, error, %s", err.Error())
		return
	}

	if resp.StatusCode != consts.OK {
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

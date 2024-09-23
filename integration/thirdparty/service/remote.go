package service

import (
	"encoding/json"
	"fmt"
	integrationDomain "github.com/aaronchen2k/deeptest/integration/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/config"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	httpHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/http"
	commUtils "github.com/aaronchen2k/deeptest/internal/pkg/utils"
	_commUtils "github.com/aaronchen2k/deeptest/pkg/lib/comm"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"strconv"
	"time"
)

type RemoteService struct {
}

func (s *RemoteService) GetUserInfoByToken(tenantId consts.TenantId, token string) (user integrationDomain.UserInfo, err error) {

	url := fmt.Sprintf("%s/api/v1/user/getUserInfo", config.CONFIG.ThirdParty.Url)

	headers := []domain.Header{
		{
			Name:  "X-Token",
			Value: token,
		},
		{Name: "tenantId", Value: string(tenantId)},
	}

	httpReq := domain.BaseRequest{
		Url:      url,
		BodyType: consts.ContentTypeJSON,
		Headers:  &headers,
	}

	logUtils.Infof("thirdparty-getUserInfoByToken: %s", _commUtils.JsonEncode(httpReq))
	var resp domain.DebugResponse
	resp, err = httpHelper.Get(httpReq, nil)
	if err != nil {
		logUtils.Errorf("meta get method detail failed, error, %s", err.Error())
		return
	}

	if resp.StatusCode != consts.OK.Int() {
		logUtils.Errorf("meta get method detail failed, response %v", _commUtils.JsonEncode(resp))
		err = fmt.Errorf("meta get method detail failed, response %v", resp)
		return
	}

	respContent := struct {
		Code int
		Data struct{ UserInfo integrationDomain.UserInfo }
		Msg  string
	}{}
	err = json.Unmarshal([]byte(resp.Content), &respContent)
	if err != nil {
		logUtils.Errorf(err.Error())
	}

	if respContent.Data.UserInfo.Mail == "" {
		respContent.Data.UserInfo.Mail = respContent.Data.UserInfo.Username
	}

	user = respContent.Data.UserInfo

	return
}

func (s *RemoteService) GetProjectInfo(tenantId consts.TenantId, token, spaceCode string) (ret integrationDomain.ProjectInfo, err error) {
	url := fmt.Sprintf("%s/api/v1/project/info/%s", config.CONFIG.ThirdParty.Url, spaceCode)

	httpReq := domain.BaseRequest{
		Url:      url,
		BodyType: consts.ContentTypeJSON,
		Headers: &[]domain.Header{
			{
				Name:  "X-Token",
				Value: token,
			},
			{Name: "tenantId", Value: string(tenantId)},
		},
	}

	resp, err := httpHelper.Get(httpReq, nil)
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
		Data struct{ integrationDomain.ProjectInfo }
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

func (s *RemoteService) GetHeaders(tenantId consts.TenantId, body string) (headers []domain.Header) {
	xNancalTimestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
	xNancalNonceStr := _commUtils.RandStr(8)

	if body != "" {
		body = commUtils.CompressedJson(body)
	}

	headers = []domain.Header{
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
			Value: _commUtils.GetSign(config.CONFIG.ThirdParty.ApiSign.AppKey, config.CONFIG.ThirdParty.ApiSign.AppSecret, xNancalNonceStr, xNancalTimestamp, body),
		},
		{
			Name:  "tenantId",
			Value: string(tenantId),
		},
	}

	return
}

func (s *RemoteService) GetUserButtonPermissions(tenantId consts.TenantId, username string) (ret []string, err error) {
	url := fmt.Sprintf("%s/api/v1/openApi/getUserDynamicMenuPermission", config.CONFIG.ThirdParty.Url)

	headers := s.GetHeaders(tenantId, "")
	httpReq := domain.BaseRequest{
		Url:      url,
		BodyType: consts.ContentTypeJSON,
		Headers:  &headers,
		QueryParams: &[]domain.Param{
			{
				Name:  "typeStr",
				Value: "[20,30]",
			},
			{
				Name:  "username",
				Value: username,
			},
		},
	}
	logUtils.Infof("thirdparty-UserButtonPermissions,%s", _commUtils.JsonEncode(httpReq))
	resp, err := httpHelper.Get(httpReq, nil)
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

func (s *RemoteService) GetUserMenuPermissions(tenantId consts.TenantId, username, spaceCode string) (ret []integrationDomain.UserMenuPermission, err error) {
	url := fmt.Sprintf("%s/api/v1/openApi/getUserDynamicMenu", config.CONFIG.ThirdParty.Url)

	headers := s.GetHeaders(tenantId, "")
	httpReq := domain.BaseRequest{
		Url:      url,
		BodyType: consts.ContentTypeJSON,
		Headers:  &headers,
		QueryParams: &[]domain.Param{
			{
				Name:  "typeStr",
				Value: "[10,20]",
			},
			{
				Name:  "nameEngAbbr",
				Value: spaceCode,
			},
			{
				Name:  "username",
				Value: username,
			},
		},
	}
	logUtils.Infof("thirdparty-GetUserMenuPermissions %s", _commUtils.JsonEncode(httpReq))
	resp, err := httpHelper.Get(httpReq, nil)
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
		Data []integrationDomain.UserMenuPermission
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

func (s *RemoteService) GetSpaceRoles(tenantId consts.TenantId) (ret []integrationDomain.SpaceRole, err error) {
	url := fmt.Sprintf("%s/api/v1/openApi/getSpaceInitRole", config.CONFIG.ThirdParty.Url)

	headers := s.GetHeaders(tenantId, "")
	httpReq := domain.BaseRequest{
		Url:      url,
		BodyType: consts.ContentTypeJSON,
		Headers:  &headers,
	}

	resp, err := httpHelper.Get(httpReq, nil)
	if err != nil {
		logUtils.Infof("get SpaceRoles failed, error, %s", err.Error())
		return
	}

	if resp.StatusCode != consts.OK.Int() {
		logUtils.Infof("get SpaceRoles failed, response %v", resp)
		err = fmt.Errorf("get SpaceRoles failed, response %v", resp)
		return
	}

	respContent := struct {
		Code int
		Data []integrationDomain.SpaceRole
		Msg  string
	}{}
	err = json.Unmarshal([]byte(resp.Content), &respContent)
	if err != nil {
		logUtils.Infof(err.Error())
	}

	if respContent.Code != 200 {
		logUtils.Infof("get SpaceRoles failed, response %v", resp)
		err = fmt.Errorf("get SpaceRoles failed, response %v", resp)
		return
	}

	ret = respContent.Data

	return
}

func (s *RemoteService) GetRoleMenus(tenantId consts.TenantId, role string) (ret []string, err error) {
	url := fmt.Sprintf("%s/api/v1/openApi/getRoleMenus", config.CONFIG.ThirdParty.Url)

	headers := s.GetHeaders(tenantId, "")
	httpReq := domain.BaseRequest{
		Url:      url,
		BodyType: consts.ContentTypeJSON,
		Headers:  &headers,
		QueryParams: &[]domain.Param{
			{
				Name:  "roleValue",
				Value: role,
			},
		},
	}
	logUtils.Infof("thirdparty-GetRoleMenus,%s", _commUtils.JsonEncode(httpReq))
	resp, err := httpHelper.Get(httpReq, nil)
	if err != nil {
		logUtils.Infof("get RoleMenus failed, error, %s", err.Error())
		return
	}

	if resp.StatusCode != consts.OK.Int() {
		logUtils.Infof("get RoleMenus failed, response %v", resp)
		err = fmt.Errorf("get RoleMenus failed, response %v", resp)
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
		logUtils.Infof("get RoleMenus failed, response %v", resp)
		err = fmt.Errorf("get RoleMenus failed, response %v", resp)
		return
	}

	ret = respContent.Data

	return
}

func (s *RemoteService) GetUserProductList(tenantId consts.TenantId, page, pageSize int, username string) (ret []integrationDomain.ProductItem, err error) {
	url := fmt.Sprintf("%s/api/v1/openApi/listProductManageOptionSecrets", config.CONFIG.ThirdParty.Url)

	headers := s.GetHeaders(tenantId, "")
	httpReq := domain.BaseRequest{
		Url:      url,
		BodyType: consts.ContentTypeJSON,
		Headers:  &headers,
		QueryParams: &[]domain.Param{
			{
				Name:  "page",
				Value: strconv.Itoa(page),
			},
			{
				Name:  "pageSize",
				Value: strconv.Itoa(pageSize),
			},
			{
				Name:  "userName",
				Value: username,
			},
		},
	}

	resp, err := httpHelper.Get(httpReq, nil)
	logUtils.Infof("GetUserProductList userName:%+v, resp:%+v", username, resp)
	if err != nil {
		logUtils.Infof("GetUserProductList failed, error, %s", err.Error())
		return
	}

	if resp.StatusCode != consts.OK.Int() {
		logUtils.Infof("GetUserProductList failed, response %v", resp)
		err = fmt.Errorf("GetUserProductList failed, response %v", resp)
		return
	}

	respContent := struct {
		Code int
		Data []integrationDomain.ProductItem
		Msg  string
	}{}
	err = json.Unmarshal([]byte(resp.Content), &respContent)
	if err != nil {
		logUtils.Infof(err.Error())
	}

	if respContent.Code != 200 {
		logUtils.Infof("GetUserProductList failed, response %v", resp)
		err = fmt.Errorf("GetUserProductList failed, response %v", resp)
		return
	}

	ret = respContent.Data

	return
}

func (s *RemoteService) GetProductListById(tenantId consts.TenantId, productIds []uint) (ret []integrationDomain.ProductBaseItem, err error) {
	url := fmt.Sprintf("%s/api/v1/openApi/listProductManage", config.CONFIG.ThirdParty.Url)

	queryParams := make([]domain.Param, 0)
	for _, v := range productIds {
		paramTmp := domain.Param{
			//Name: fmt.Sprintf("productIds[%d]", k),
			Name:  "productIds[]",
			Value: strconv.Itoa(int(v)),
		}
		queryParams = append(queryParams, paramTmp)
	}

	headers := s.GetHeaders(tenantId, "")
	httpReq := domain.BaseRequest{
		Url:         url,
		BodyType:    consts.ContentTypeJSON,
		Headers:     &headers,
		QueryParams: &queryParams,
	}

	resp, err := httpHelper.Get(httpReq, nil)
	if err != nil {
		logUtils.Infof("GetProductListById failed, error, %s", err.Error())
		return
	}

	if resp.StatusCode != consts.OK.Int() {
		logUtils.Infof("GetProductListById failed, response %v", resp)
		err = fmt.Errorf("GetProductListById failed, response %v", resp)
		return
	}

	respContent := struct {
		Code int
		Data struct {
			List []integrationDomain.ProductBaseItem `json:"list"`
		} `json:"data"`
		Msg string
	}{}
	err = json.Unmarshal([]byte(resp.Content), &respContent)
	if err != nil {
		logUtils.Infof(err.Error())
	}

	if respContent.Code != 200 {
		logUtils.Infof("GetProductListById failed, response %v", resp)
		err = fmt.Errorf("GetProductListById failed, response %v", resp)
		return
	}

	ret = respContent.Data.List

	return
}

func (s *RemoteService) GetSpacesByUsername(tenantId consts.TenantId, username string) (ret []integrationDomain.SpaceItem, err error) {
	url := fmt.Sprintf("%s/api/v1/openApi/project/user", config.CONFIG.ThirdParty.Url)

	headers := s.GetHeaders(tenantId, "")
	httpReq := domain.BaseRequest{
		Url:      url,
		BodyType: consts.ContentTypeJSON,
		Headers:  &headers,
		QueryParams: &[]domain.Param{
			{
				Name:  "username",
				Value: username,
			},
		},
	}

	resp, err := httpHelper.Get(httpReq, nil)
	logUtils.Infof("GetSpacesByUsername username: +%v, resp:%+v", username, resp)
	if err != nil {
		logUtils.Infof("GetSpacesByUsername failed, error, %s", err.Error())
		return
	}

	if resp.StatusCode != consts.OK.Int() {
		logUtils.Infof("GetSpacesByUsername failed, response %v", resp)
		err = fmt.Errorf("GetSpacesByUsername failed, response %v", resp)
		return
	}

	respContent := struct {
		Code int
		Data []integrationDomain.SpaceItem
		Msg  string
	}{}
	err = json.Unmarshal([]byte(resp.Content), &respContent)
	if err != nil {
		logUtils.Infof(err.Error())
	}

	if respContent.Code != 200 {
		logUtils.Infof("GetSpacesByUsername failed, response %v", resp)
		err = fmt.Errorf("GetSpacesByUsername failed, response %v", resp)
		return
	}

	ret = respContent.Data

	return
}

func (s *RemoteService) BatchGetSpacesByCode(tenantId consts.TenantId, spaceCodes []string) (ret []integrationDomain.SpaceItem, err error) {
	url := fmt.Sprintf("%s/api/v1/openApi/project/abbr", config.CONFIG.ThirdParty.Url)

	queryParams := make([]domain.Param, 0)
	for _, v := range spaceCodes {
		paramTmp := domain.Param{
			Name:  "nameEngAbbrs",
			Value: v,
		}
		queryParams = append(queryParams, paramTmp)
	}

	headers := s.GetHeaders(tenantId, "")
	httpReq := domain.BaseRequest{
		Url:         url,
		BodyType:    consts.ContentTypeJSON,
		Headers:     &headers,
		QueryParams: &queryParams,
	}

	resp, err := httpHelper.Get(httpReq, nil)
	if err != nil {
		logUtils.Infof("BatchGetSpacesByCode failed, error, %s", err.Error())
		return
	}

	if resp.StatusCode != consts.OK.Int() {
		logUtils.Infof("BatchGetSpacesByCode failed, response %v", resp)
		err = fmt.Errorf("BatchGetSpacesByCode failed, response %v", resp)
		return
	}

	respContent := struct {
		Code int
		Data []integrationDomain.SpaceItem
		Msg  string
	}{}
	err = json.Unmarshal([]byte(resp.Content), &respContent)
	if err != nil {
		logUtils.Infof(err.Error())
	}

	if respContent.Code != 200 {
		logUtils.Infof("BatchGetSpacesByCode failed, response %v", resp)
		err = fmt.Errorf("BatchGetSpacesByCode failed, response %v", resp)
		return
	}

	ret = respContent.Data

	return
}

func (s *RemoteService) BatchGetMembersBySpaces(tenantId consts.TenantId, spaceCodes []string) (ret []integrationDomain.SpaceMembersAndRolesItem, err error) {
	url := fmt.Sprintf("%s/api/v1/openApi/project/member/abbrs", config.CONFIG.ThirdParty.Url)

	queryParams := make([]domain.Param, 0)
	for _, v := range spaceCodes {
		paramTmp := domain.Param{
			Name:  "projectEngAbbrs[]",
			Value: v,
		}
		queryParams = append(queryParams, paramTmp)
	}

	headers := s.GetHeaders(tenantId, "")
	httpReq := domain.BaseRequest{
		Url:         url,
		BodyType:    consts.ContentTypeJSON,
		Headers:     &headers,
		QueryParams: &queryParams,
	}

	resp, err := httpHelper.Get(httpReq, nil)
	if err != nil {
		logUtils.Infof("BatchGetMembersBySpaces failed, error, %s", err.Error())
		return
	}

	if resp.StatusCode != consts.OK.Int() {
		logUtils.Infof("BatchGetMembersBySpaces failed, response %v", resp)
		err = fmt.Errorf("BatchGetMembersBySpaces failed, response %v", resp)
		return
	}

	respContent := struct {
		Code int
		Data []integrationDomain.SpaceMembersAndRolesItem
		Msg  string
	}{}
	err = json.Unmarshal([]byte(resp.Content), &respContent)
	if err != nil {
		logUtils.Infof(err.Error())
	}

	if respContent.Code != 200 {
		logUtils.Infof("BatchGetMembersBySpaces failed, response %v", resp)
		err = fmt.Errorf("BatchGetMembersBySpaces failed, response %v", resp)
		return
	}

	ret = respContent.Data

	return
}

func (s *RemoteService) ApprovalAndMsg(tenantId consts.TenantId, req string) (ret string, err error) {
	url := fmt.Sprintf("%s/api/v1/openApi/approvalAndMsg", config.CONFIG.ThirdParty.Url)

	headers := s.GetHeaders(tenantId, req)
	httpReq := domain.BaseRequest{
		Url:      url,
		BodyType: consts.ContentTypeJSON,
		Headers:  &headers,
		Body:     req,
	}

	resp, err := httpHelper.Post(httpReq, nil)
	logUtils.Infof("ApprovalAndMsg url:%s, body:%+v,Headers:%+v , response %+v", url, req, headers, resp)
	if err != nil {
		logUtils.Infof("ApprovalAndMsg failed, error, %s", err.Error())
		return
	}

	if resp.StatusCode != consts.OK.Int() {
		logUtils.Infof("ApprovalAndMsg failed, response %+v", resp)
		err = fmt.Errorf("ApprovalAndMsg failed, response %v", resp)
		return
	}

	respContent := struct {
		Code int
		Data struct {
			InstanceId string `json:"instance_id"`
		} `json:"data"`
		Msg string
	}{}
	err = json.Unmarshal([]byte(resp.Content), &respContent)
	if err != nil {
		logUtils.Infof(err.Error())
	}

	if respContent.Code != 200 {
		logUtils.Infof("ApprovalAndMsg failed, response %+v", resp)
		err = fmt.Errorf("ApprovalAndMsg failed, response %v", resp)
		return
	}

	ret = respContent.Data.InstanceId

	return
}

func (s *RemoteService) GetUserOpenRoles(tenantId consts.TenantId, username string) (ret []integrationDomain.UserRoleItem, err error) {
	url := fmt.Sprintf("%s/api/v1/openApi/getUserOpenRole", config.CONFIG.ThirdParty.Url)

	headers := s.GetHeaders(tenantId, "")
	httpReq := domain.BaseRequest{
		Url:      url,
		BodyType: consts.ContentTypeJSON,
		Headers:  &headers,
		QueryParams: &[]domain.Param{
			{
				Name:  "username",
				Value: username,
			},
		},
	}

	resp, err := httpHelper.Get(httpReq, nil)
	logUtils.Infof("GetUserOpenRoles username: +%v, resp:%+v", username, resp)
	if err != nil {
		logUtils.Errorf("GetUserOpenRoles failed, username: +%v,error: %s", username, err.Error())
		return
	}

	if resp.StatusCode != consts.OK.Int() {
		logUtils.Infof("GetUserOpenRoles failed, response %v", resp)
		err = fmt.Errorf("GetUserOpenRoles failed, response %v", resp)
		return
	}

	respContent := struct {
		Code int
		Data []integrationDomain.UserRoleItem
		Msg  string
	}{}
	err = json.Unmarshal([]byte(resp.Content), &respContent)
	if err != nil {
		logUtils.Infof(err.Error())
	}

	if respContent.Code != 200 {
		logUtils.Infof("GetUserOpenRoles failed, response %v", resp)
		err = fmt.Errorf("GetUserOpenRoles failed, response %v", resp)
		return
	}

	ret = respContent.Data

	return
}

func (s *RemoteService) SaveReport(tenantId consts.TenantId, report integrationDomain.CreateReport) (err error) {
	url := fmt.Sprintf("%s/api/v1/openApi/saveReport", config.CONFIG.ThirdParty.Url)
	req := _commUtils.JsonEncode(report)
	headers := s.GetHeaders(tenantId, req)
	httpReq := domain.BaseRequest{
		Url:      url,
		BodyType: consts.ContentTypeJSON,
		Headers:  &headers,
		Body:     req,
	}
	logUtils.Infof("thirdparty-SaveReport %s", _commUtils.JsonEncode(httpReq))
	resp, err := httpHelper.Post(httpReq, nil)
	if err != nil {
		logUtils.Infof("get SaveReport failed, error, %s", err.Error())
		return
	}

	if resp.StatusCode != consts.OK.Int() {
		logUtils.Infof("get SaveReport failed, response %v", resp)
		err = fmt.Errorf("get SaveReport failed, response %v", resp)
		return
	}

	respContent := struct {
		Code int
		Data interface{}
		Msg  string
	}{}
	err = json.Unmarshal([]byte(resp.Content), &respContent)
	if err != nil {
		logUtils.Infof(err.Error())
	}

	if respContent.Code != 200 {
		logUtils.Infof("SaveReport failed, response %v", resp)
		err = fmt.Errorf("get SaveReport failed, response %v", resp)
		return
	}

	return
}

func (s *RemoteService) SyncPlan(tenantId consts.TenantId, data integrationDomain.SyncPlan) (err error) {
	url := fmt.Sprintf("%s/api/v1/openApi/syncPlan", config.CONFIG.ThirdParty.Url)
	req := _commUtils.JsonEncode(data)
	headers := s.GetHeaders(tenantId, req)
	httpReq := domain.BaseRequest{
		Url:      url,
		BodyType: consts.ContentTypeJSON,
		Headers:  &headers,
		Body:     req,
	}
	logUtils.Infof("thirdparty-SyncPlan %s", _commUtils.JsonEncode(httpReq))
	resp, err := httpHelper.Post(httpReq, nil)
	if err != nil {
		logUtils.Infof("get SyncPlan failed, error, %s", err.Error())
		return
	}

	if resp.StatusCode != consts.OK.Int() {
		logUtils.Infof("get SyncPlan failed, response %v", resp)
		err = fmt.Errorf("get SyncPlan failed, response %v", resp)
		return
	}

	respContent := struct {
		Code int
		Data interface{}
		Msg  string
	}{}
	err = json.Unmarshal([]byte(resp.Content), &respContent)
	if err != nil {
		logUtils.Infof(err.Error())
	}

	if respContent.Code != 200 {
		logUtils.Infof("SyncPlan failed, response %v", resp)
		err = fmt.Errorf("get SyncPlan failed, response %v", resp)
		return
	}

	return
}

func (s *RemoteService) DeleteReport(tenantId consts.TenantId, number string) (err error) {
	url := fmt.Sprintf("%s/api/v1/openApi/report/%s", config.CONFIG.ThirdParty.Url, number)
	headers := s.GetHeaders(tenantId, "")
	httpReq := domain.BaseRequest{
		Url:      url,
		BodyType: consts.ContentTypeJSON,
		Headers:  &headers,
	}
	logUtils.Infof("thirdparty-DeleteReport %s", _commUtils.JsonEncode(httpReq))
	resp, err := httpHelper.Get(httpReq, nil)
	if err != nil {
		logUtils.Infof("get DeleteReport failed, error, %s", err.Error())
		return
	}

	if resp.StatusCode != consts.OK.Int() {
		logUtils.Infof("get DeleteReport failed, response %v", resp)
		err = fmt.Errorf("get DeleteReport failed, response %v", resp)
		return
	}

	respContent := struct {
		Code int
		Data interface{}
		Msg  string
	}{}
	err = json.Unmarshal([]byte(resp.Content), &respContent)
	if err != nil {
		logUtils.Infof(err.Error())
	}

	if respContent.Code != 200 {
		logUtils.Infof("SaveReport failed, response %v", resp)
		err = fmt.Errorf("get SaveReport failed, response %v", resp)
		return
	}

	return
}

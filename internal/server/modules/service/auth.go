package service

import (
	"fmt"
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	model "github.com/deeptest-com/deeptest/internal/server/modules/model"
	repo2 "github.com/deeptest-com/deeptest/internal/server/modules/repo"
	_httpUtils "github.com/deeptest-com/deeptest/pkg/lib/http"
	"github.com/kataras/iris/v12"
	"net/url"
	"strings"
)

type AuthService struct {
	AuthRepo *repo2.AuthRepo `inject:""`
}

func (s AuthService) OAuth2Authorization(tenantId consts.TenantId, req model.DebugInterfaceOAuth20) (result iris.Map, err error) {
	//s.InterfaceRepo.UpdateOAuth20(req.EndpointInterfaceId, req)

	responseType := ""
	if req.GrantType == consts.AuthorizationCode {
		responseType = "code"
	}

	callbackUrl := req.CallbackUrl
	params := map[string]interface{}{
		"interfaceId":    req.InterfaceId,
		"accessTokenURL": req.AccessTokenURL,
		"clientId":       req.ClientID,
		"clientSecret":   req.ClientSecret,
		"name":           req.Name,
	}
	callbackUrl = _httpUtils.GenUrlWithParams("", params, callbackUrl)

	urlTempl := "%s" +
		"?response_type=%s" +
		"&client_id=%s" +
		"&state=state&scope=%s" +
		"&redirect_uri=%s"

	url := fmt.Sprintf(urlTempl,
		req.AuthURL, responseType,
		req.ClientID, req.Scope, url.QueryEscape(callbackUrl))

	result = iris.Map{
		"url": url,
	}

	return
}

func (s AuthService) GenOAuth2AccessToken(tenantId consts.TenantId, accessTokenURL, clientId, clientSecret, code string) (result iris.Map, err error) {
	url := fmt.Sprintf(
		"%s?client_id=%s&client_secret=%s&code=%s",
		accessTokenURL, clientId, clientSecret, code,
	)

	bytes, _, err := _httpUtils.Get(url, nil)

	result = iris.Map{}

	arr := strings.Split(string(bytes), "&")
	for _, pair := range arr {
		items := strings.Split(pair, "=")
		result[items[0]] = items[1]
	}

	return
}

func (s AuthService) AddToken(tenantId consts.TenantId, name, token, tokenType string, interfaceId, projectId int) (err error) {
	_, err = s.AuthRepo.CreateToken(tenantId, name, token, tokenType, projectId)
	if err != nil {
		return
	}

	//err = s.InterfaceRepo.SetOAuth2AccessToken(token, interfaceId)

	return
}

func (s AuthService) ListOAuth2Token(tenantId consts.TenantId, projectId int) (pos []model.Auth2Token, err error) {
	pos, err = s.AuthRepo.ListOAuth2Token(tenantId, projectId)

	return
}

func (s AuthService) RemoveToken(tenantId consts.TenantId, id int) (err error) {
	err = s.AuthRepo.RemoveToken(tenantId, uint(id))

	return
}

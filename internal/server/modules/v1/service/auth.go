package service

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/comm/consts"
	_httpUtils "github.com/aaronchen2k/deeptest/internal/pkg/lib/http"
	logUtils "github.com/aaronchen2k/deeptest/internal/pkg/lib/log"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/repo"
	"github.com/kataras/iris/v12"
	"net/url"
	"strings"
)

type AuthService struct {
	InvocationRepo *repo.InvocationRepo `inject:""`
	InterfaceRepo  *repo.InterfaceRepo  `inject:""`
}

func (s AuthService) OAuth2Authorization(req model.InterfaceOAuth20) (result iris.Map, err error) {
	s.InterfaceRepo.UpdateOAuth20(req.InterfaceId, req)

	responseType := ""
	if req.GrantType == consts.AuthorizationCode {
		responseType = "code"
	}

	callbackUrl := req.CallbackUrl
	params := map[string]interface{}{
		"accessTokenURL": req.AccessTokenURL,
		"clientId":       req.ClientID,
		"clientSecret":   req.ClientSecret,
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

func (s AuthService) GenOAuth2AccessToken(accessTokenURL, clientId, clientSecret, code string) (result iris.Map, err error) {
	url := fmt.Sprintf(
		"%s?client_id=%s&client_secret=%s&code=%s",
		accessTokenURL, clientId, clientSecret, code,
	)

	bytes, err := _httpUtils.Get(url)
	logUtils.Infof("#v", bytes)

	result = iris.Map{}

	arr := strings.Split(string(bytes), "&")
	for _, pair := range arr {
		items := strings.Split(pair, "=")
		result[items[0]] = items[1]
	}

	return
}

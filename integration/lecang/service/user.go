package service

import (
	"errors"
	integrationDomain "github.com/aaronchen2k/deeptest/integration/domain"
	v1 "github.com/aaronchen2k/deeptest/integration/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/config"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	_commUtils "github.com/aaronchen2k/deeptest/pkg/lib/comm"
)

type User struct {
}

func (s *User) GetUserInfoByToken(tenantId consts.TenantId, token, origin string) (user v1.UserInfo, err error) {
	user, err = new(RemoteService).GetUserInfoByToken(token, origin)
	return
}

func (s *User) GetToken(baseUrl string) (token string, err error) {
	loginByOauthReq := integrationDomain.LoginByOauthReq{
		LoginName: config.CONFIG.ThirdParty.Username,
		Password:  _commUtils.Sha256(config.CONFIG.ThirdParty.Password),
		TenantId:  "1632931640315338752",
	}

	loginByOauthResData := new(RemoteService).LoginByOauth(loginByOauthReq, baseUrl)
	if loginByOauthResData.Code == "" {
		return "", errors.New("login fail")
	}

	getTokenFromCodeReq := integrationDomain.GetTokenFromCodeReq{
		Code: loginByOauthResData.Code,
	}

	getTokenFromCodeResData := new(RemoteService).GetTokenFromCode(getTokenFromCodeReq, baseUrl)
	token = getTokenFromCodeResData.Token

	return
}

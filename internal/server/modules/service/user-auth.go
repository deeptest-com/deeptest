package service

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	integrationDomain "github.com/aaronchen2k/deeptest/integration/domain"
	"github.com/aaronchen2k/deeptest/integration/service"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	commonUtils "github.com/aaronchen2k/deeptest/pkg/lib/comm"
	"github.com/snowlyg/multi"
	"strconv"
	"time"
)

type UserAuthService struct {
	RemoteService *service.RemoteService `inject:""`
	UserRepo      *repo.UserRepo         `inject:""`
}

func (s *UserAuthService) Auth(tenantId consts.TenantId, token string) (user model.SysUser, err error) {
	var userInfo integrationDomain.UserInfo
	userInfo, err = s.RemoteService.GetUserInfoByToken(tenantId, token)
	if err != nil {
		req := v1.UserReq{UserBase: v1.UserBase{
			Username:  userInfo.Username,
			Email:     userInfo.Mail,
			ImAccount: userInfo.WxName,
			Name:      userInfo.RealName,
			Password:  commonUtils.RandStr(8),
		}}
		s.UserRepo.Create(tenantId, req)

	}
	user, err = s.UserRepo.GetByUserName(tenantId, userInfo.Username)
	return
}

func (s *UserAuthService) createSession(user model.SysUser) (token string, err error) {
	claims := &multi.CustomClaims{
		ID:            strconv.FormatUint(uint64(user.ID), 10),
		Username:      user.Username,
		AuthorityId:   "",
		AuthorityType: multi.AdminAuthority,
		LoginType:     multi.LoginTypeApp,
		AuthType:      multi.AuthPwd,
		CreationDate:  time.Now().Local().Unix(),
		ExpiresIn:     multi.RedisSessionTimeoutWeb.Milliseconds(),
	}

	token, _, err = multi.AuthDriver.GenerateToken(claims)
	if err != nil {
		return
	}
	return
}

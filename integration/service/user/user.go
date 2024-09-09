package user

import (
	v1 "github.com/aaronchen2k/deeptest/integration/domain"
	"github.com/aaronchen2k/deeptest/integration/enum"
	lecang "github.com/aaronchen2k/deeptest/integration/lecang/service"
	thirdparty "github.com/aaronchen2k/deeptest/integration/thirdparty/service"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
)

type user struct {
	u IUser
}

func NewUser(appName enum.AppName) *user {
	ret := new(user)
	ret.u = ret.getEntity(appName)

	return ret
}

func (s *user) GetUserInfoByToken(tenantId consts.TenantId, token, origin string) (v1.UserInfo, error) {
	return s.u.GetUserInfoByToken(tenantId, token, origin)
}

func (s *user) getEntity(appName enum.AppName) (u IUser) {
	switch appName {
	case enum.Lecang:
		u = new(lecang.User)
	case enum.Thirdparty:
		u = new(thirdparty.User)
	}
	return
}

package user

import (
	v1 "github.com/deeptest-com/deeptest/integration/domain"
	"github.com/deeptest-com/deeptest/integration/enum"
	lecang "github.com/deeptest-com/deeptest/integration/lecang/service"
	thirdparty "github.com/deeptest-com/deeptest/integration/thirdparty/service"
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
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

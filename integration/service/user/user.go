package user

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/integration/enum"
	lecang "github.com/aaronchen2k/deeptest/integration/lecang/service"
	leyan "github.com/aaronchen2k/deeptest/integration/leyan/service"
)

type user struct {
	u IUser
}

func NewUser(appName enum.AppName) *user {
	ret := new(user)
	ret.u = ret.getEntity(appName)

	return ret
}

func (s *user) GetUserInfoByToken(token, origin string) (v1.UserInfo, error) {
	return s.u.GetUserInfoByToken(token, origin)
}

func (s *user) getEntity(appName enum.AppName) (u IUser) {
	switch appName {
	case enum.Lecang:
		u = new(lecang.User)
	case enum.Leyan:
		u = new(leyan.User)
	}
	return
}

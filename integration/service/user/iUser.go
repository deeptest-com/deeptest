package user

import v1 "github.com/aaronchen2k/deeptest/integration/domain"

type IUser interface {
	GetUserInfoByToken(token, origin string) (v1.UserInfo, error)
}

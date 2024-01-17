package user

import v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"

type IUser interface {
	GetUserInfoByToken(token string) (v1.UserInfo, error)
}

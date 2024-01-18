package service

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
)

type User struct {
}

func (s *User) GetUserInfoByToken(token, origin string) (user v1.UserInfo, err error) {
	user, err = new(remote).GetUserInfoByToken(token, origin)
	return
}

package service

import (
	v1 "github.com/aaronchen2k/deeptest/integration/domain"
	"github.com/aaronchen2k/deeptest/integration/service"
)

type User struct {
}

func (s *User) GetUserInfoByToken(token, origin string) (user v1.UserInfo, err error) {
	user, err = new(service.RemoteService).GetUserInfoByToken(token)
	return
}

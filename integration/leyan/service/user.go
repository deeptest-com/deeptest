package service

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/service"
)

type User struct {
}

func (s *User) GetUserInfoByToken(token, origin string) (user v1.UserInfo, err error) {
	user, err = new(service.RemoteService).GetUserInfoByToken(token)
	return
}

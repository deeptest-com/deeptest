package service

import (
	v1 "github.com/aaronchen2k/deeptest/integration/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
)

type User struct {
}

func (s *User) GetUserInfoByToken(tenantId consts.TenantId, token, origin string) (user v1.UserInfo, err error) {
	user, err = new(remote).GetUserInfoByToken(token, origin)
	return
}

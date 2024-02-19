package user

import (
	v1 "github.com/aaronchen2k/deeptest/integration/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
)

type IUser interface {
	GetUserInfoByToken(tenantId consts.TenantId, token, origin string) (v1.UserInfo, error)
}

package user

import (
	v1 "github.com/deeptest-com/deeptest/integration/domain"
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
)

type IUser interface {
	GetUserInfoByToken(tenantId consts.TenantId, token, origin string) (v1.UserInfo, error)
}

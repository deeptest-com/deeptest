package common

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/kataras/iris/v12"
)

func GetTenantId(ctx iris.Context) consts.TenantId {
	ret := ctx.GetHeader("tenantId")
	return consts.TenantId(ret)
}

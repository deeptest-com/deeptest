package common

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/kataras/iris/v12"
)

func GetTenantId(ctx iris.Context) consts.TenantId {
	ret := ctx.GetHeader("tenantId")
	//ret = "1705374224174"
	return consts.TenantId(ret)
}

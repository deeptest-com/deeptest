package common

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/config"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/aaronchen2k/deeptest/saas/tenant"
	"github.com/kataras/iris/v12/context"
	"runtime/debug"
	"strings"
)

func GetTenantId(ctx *context.Context) consts.TenantId {
	ret := ctx.GetHeader("tenantId")
	//根据域名获取租户
	tenantId := consts.TenantId(ret)

	if config.CONFIG.Saas.Switch && ret == "" && strings.HasPrefix(ctx.Path(), "/mocks") {
		fmt.Println(ctx.Host(), ctx.Path(), "++++++++++++++++")
		domain := ctx.Host()
		array := strings.Split(domain, ".")
		prefix := strings.ReplaceAll(array[0], "dev-", "")
		info := tenant.NewTenant().GetInfo(tenantId, prefix)
		tenantId = info.Id
	}

	return tenantId
}

func AsyncCatchErrRun(f func()) {
	defer func() {
		if err := recover(); err != nil {
			logUtils.Info(fmt.Sprintf("%v", err))
			//panic(err)
			s := string(debug.Stack())
			fmt.Printf("err=%v, stack=%s\n", err, s)
		}
	}()

	f()
}

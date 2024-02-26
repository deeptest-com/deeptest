package common

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/kataras/iris/v12/context"
	"runtime/debug"
)

func GetTenantId(ctx *context.Context) consts.TenantId {
	ret := ctx.GetHeader("tenantId")
	//ret = "1705374224174"
	return consts.TenantId(ret)
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

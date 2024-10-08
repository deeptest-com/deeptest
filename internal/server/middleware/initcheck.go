package middleware

import (
	"fmt"
	"github.com/deeptest-com/deeptest/internal/pkg/config"
	"github.com/deeptest-com/deeptest/internal/server/core/cache"
	"github.com/deeptest-com/deeptest/internal/server/core/dao"
	"github.com/deeptest-com/deeptest/pkg/domain"
	"github.com/deeptest-com/deeptest/saas/common"
	"net/http"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

// InitCheck 初始化检测中间件
func InitCheck() iris.Handler {
	return func(ctx *context.Context) {
		tenantId := common.GetTenantId(ctx)
		if dao.GetDB(tenantId) == nil || (config.CONFIG.System.CacheType == "redis" && config.CACHE == nil) {
			ctx.StopWithJSON(http.StatusOK, _domain.Response{Code: _domain.NeedInitErr.Code, Data: nil, Msg: _domain.NeedInitErr.Msg})
		} else {
			host := ctx.Request().Header.Get("Origin")

			if ctx.GetHeader("X-Token") != "" {
				host = ctx.Request().Header.Get("X-API-Origin")
			}
			if host != "" {
				cache.SetCache(fmt.Sprintf("%s_host", tenantId), host, -1)
			}
			ctx.Next()
		}

		// 处理请求
	}
}

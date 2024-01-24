package middleware

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/config"
	"github.com/aaronchen2k/deeptest/internal/server/core/cache"
	"github.com/aaronchen2k/deeptest/internal/server/core/dao"
	"github.com/aaronchen2k/deeptest/pkg/domain"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"net/http"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

// InitCheck 初始化检测中间件
func InitCheck() iris.Handler {
	return func(ctx *context.Context) {
		if dao.GetDB() == nil || (config.CONFIG.System.CacheType == "redis" && config.CACHE == nil) {
			ctx.StopWithJSON(http.StatusOK, _domain.Response{Code: _domain.NeedInitErr.Code, Data: nil, Msg: _domain.NeedInitErr.Msg})
		} else {
			logUtils.Infof("initHostParam X-Token:%s,Origin:%s,X-API-Origin:%s", ctx.GetHeader("X-Token"), ctx.Request().Header.Get("Origin"), ctx.Request().Header.Get("X-API-Origin"))
			host := ctx.Request().Header.Get("Origin")
			thirdPartyHost := ctx.Request().Header.Get("Origin")

			if ctx.GetHeader("X-Token") != "" {
				host = ctx.Request().Header.Get("X-API-Origin")
				thirdPartyHost = ctx.Request().Header.Get("Origin")
			}
			if host != "" {
				cache.SetCache("host", host, -1)
			}
			if thirdPartyHost != "" {
				cache.SetCache("thirdPartyHost", thirdPartyHost, -1)
			}
			ctx.Next()
		}

		// 处理请求
	}
}

package middleware

import (
	serverConfig "github.com/aaronchen2k/deeptest/internal/server/config"
	"github.com/aaronchen2k/deeptest/internal/server/core/dao"
	"github.com/aaronchen2k/deeptest/pkg/domain"
	"net/http"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

// InitCheck 初始化检测中间件
func InitCheck() iris.Handler {
	return func(ctx *context.Context) {
		if dao.GetDB() == nil || (serverConfig.CONFIG.System.CacheType == "redis" && serverConfig.CACHE == nil) {
			ctx.StopWithJSON(http.StatusOK, _domain.Response{Code: _domain.NeedInitErr.Code, Data: nil, Msg: _domain.NeedInitErr.Msg})
		} else {
			ctx.Next()
		}

		// 处理请求
	}
}

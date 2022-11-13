package middleware

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/config"
	"github.com/aaronchen2k/deeptest/internal/server/core/dao"
	"github.com/aaronchen2k/deeptest/pkg/domain"
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
			ctx.Next()
		}

		// 处理请求
	}
}

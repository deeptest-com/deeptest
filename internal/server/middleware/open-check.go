package middleware

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/config"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

func OpenCheck() iris.Handler {
	return func(ctx *context.Context) {
		if ctx.Request().Header.Get("appsecret") != config.CONFIG.OpenApi.AppSecret {
			ctx.JSON(_domain.Response{Code: _domain.AuthActionErr.Code, Data: nil, Msg: "非法请求"})
			ctx.StopExecution()
			return
		}
		ctx.Next()
	}
}

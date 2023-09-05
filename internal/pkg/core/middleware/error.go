package middleware

import (
	"fmt"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

func Error() iris.Handler {

	return func(ctx *context.Context) {

		defer func(ctx *context.Context) {
			if err := recover(); err != nil {
				logUtils.Info(fmt.Sprintf("%v", err))
				panic(err)
				(*ctx).JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
			}
		}(ctx)

		ctx.Next()

	}
}

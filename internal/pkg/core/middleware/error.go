package middleware

import (
	"fmt"
	_domain "github.com/deeptest-com/deeptest/pkg/domain"
	logUtils "github.com/deeptest-com/deeptest/pkg/lib/log"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"runtime/debug"
)

func Error() iris.Handler {

	return func(ctx *context.Context) {
		defer func(ctx *context.Context) {
			if err := recover(); err != nil {
				logUtils.Info(fmt.Sprintf("%v", err))
				//panic(err)
				s := string(debug.Stack())
				fmt.Printf("err=%v, stack=%s\n", err, s)
				(*ctx).JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
			}
		}(ctx)

		ctx.Next()

	}
}

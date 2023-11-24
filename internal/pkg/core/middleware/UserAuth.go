package middleware

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

func UserAuth() iris.Handler {

	return func(ctx *context.Context) {
		token := ctx.GetHeader("X-Token")
		if token != "" {

		}

		ctx.Next()
	}
}

package middleware

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/server/core/dao"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"gorm.io/gorm"
	"runtime/debug"
)

func DBResolver() iris.Handler {

	return func(ctx *context.Context) {

		defer func(ctx *context.Context) {
			if err := recover(); err != nil {
				s := string(debug.Stack())
				fmt.Printf("err=%v, stack=%s\n", err, s)
				(*ctx).JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
			}
		}(ctx)

		dbname := ctx.URLParam("dbname")
		handler := func() (db *gorm.DB, err error) {
			return dao.InitSaasDBHandler(dbname)
		}
		db := dao.GetDBResolver().Apply(dbname, handler).GetConnPool(dbname)
		ctx.Values().Set("db", db)
		ctx.Next()
	}
}

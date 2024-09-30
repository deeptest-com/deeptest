package middleware

import (
	"fmt"
	"github.com/deeptest-com/deeptest/internal/server/core/dao"
	_domain "github.com/deeptest-com/deeptest/pkg/domain"
	"github.com/deeptest-com/deeptest/saas/common"
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

		dbname := common.GetTenantId(ctx)
		//logUtils.Infof("DBResolver,path:%s,dbname:%s", ctx.Path(), dbname)

		/*
			if config.CONFIG.Saas.Switch && dbname == "" {
				//panic(fmt.Errorf("the saas environment does not allow the tenant id to be empty"))
			}
		*/

		if dbname != "" {
			handler := func() (db *gorm.DB, err error) {
				return dao.InitSaasDBHandler(dbname)
			}
			dao.GetDBResolver().Apply(dbname, handler).GetConnPool(dbname)
		}

		ctx.Next()
	}
}

package middleware

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/helper/dbresolver"
	"github.com/aaronchen2k/deeptest/internal/server/core/dao"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

func DBResolver() iris.Handler {

	return func(ctx *context.Context) {
		////x := context2.WithValue(ctx, "dbanme", ctx.Path())
		x := new(dbresolver.DBResolver)
		x.Context = ctx
		dao.GetDB().Use(x)
		//fmt.Println(dao.GetDB().Plugins, "+++++")
		dao.GetDB().Plugins[x.Name()] = x
		//x.Context = ctx
		ctx.Next()
		//x.Context = nil
	}
}

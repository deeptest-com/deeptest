package router

import (
	"github.com/deeptest-com/deeptest/cmd/server/v1/handler"
	"github.com/deeptest-com/deeptest/internal/pkg/core/module"
	"github.com/deeptest-com/deeptest/internal/server/middleware"
	"github.com/kataras/iris/v12"
)

type DatabaseOptModule struct {
	DatabaseOptCtrl *handler.DatabaseOptCtrl `inject:""`
}

// Party 检查点
func (m *DatabaseOptModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin(), middleware.ProjectPerm())

		index.Get("/{id:uint}", m.DatabaseOptCtrl.Get).Name = "数据库操作后置条件详情"
		index.Put("/", m.DatabaseOptCtrl.Update).Name = "更新数据库操作后置条件"
	}

	return module.NewModule("/databaseOpts", handler)
}

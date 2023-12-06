package router

import (
	"github.com/aaronchen2k/deeptest/cmd/server/v1/handler"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/module"
	"github.com/aaronchen2k/deeptest/internal/server/middleware"
	"github.com/kataras/iris/v12"
)

type DatabaseConnModule struct {
	DatabaseConnCtrl *handler.DatabaseConnCtrl `inject:""`
}

// Party 项目
func (m *DatabaseConnModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin(), middleware.ProjectPerm())

		index.Get("/", m.DatabaseConnCtrl.List).Name = "数据库连接列表"
		index.Get("/{id:uint}", m.DatabaseConnCtrl.Get).Name = "数据库连接详情"
		index.Post("/", m.DatabaseConnCtrl.Save).Name = "保存数据库连接"
		index.Put("/updateName", m.DatabaseConnCtrl.UpdateName).Name = "修改数据库连接名称"
		index.Put("/{id:uint}/disable", m.DatabaseConnCtrl.Disable).Name = "禁用数据库连接"
		index.Delete("/{id:uint}", m.DatabaseConnCtrl.Delete).Name = "删除数据库连接"
	}
	return module.NewModule("/dbconns", handler)
}

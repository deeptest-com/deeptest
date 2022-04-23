package index

import (
	"github.com/aaronchen2k/deeptest/internal/server/core/module"
	"github.com/aaronchen2k/deeptest/internal/server/middleware"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/controller"
	"github.com/kataras/iris/v12"
)

type EnvironmentModule struct {
	EnvironmentCtrl *controller.EnvironmentCtrl `inject:""`
}

// Party 脚本
func (m *EnvironmentModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin())

		index.Get("/", m.EnvironmentCtrl.List).Name = "环境列表"
		index.Get("/{id:uint}", m.EnvironmentCtrl.Get).Name = "环境详情"
		index.Post("/changeEnvironment", m.EnvironmentCtrl.Change).Name = "修改环境"

		index.Post("/", m.EnvironmentCtrl.Create).Name = "新建接口"
		index.Put("/", m.EnvironmentCtrl.Update).Name = "更新接口"
		index.Delete("/{id:uint}", m.EnvironmentCtrl.Delete).Name = "删除环境"
	}
	return module.NewModule("/environments", handler)
}

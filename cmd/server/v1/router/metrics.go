package router

import (
	"github.com/deeptest-com/deeptest/cmd/server/v1/handler"
	"github.com/deeptest-com/deeptest/internal/pkg/core/module"
	"github.com/deeptest-com/deeptest/internal/server/middleware"
	"github.com/kataras/iris/v12"
)

type MetricsModule struct {
	MetricsCtrl *handler.MetricsCtrl `inject:""`
}

// Party 指标
func (m *MetricsModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin(), middleware.ProjectPerm())

		index.Get("", m.MetricsCtrl.List).Name = "指标列表"
		index.Post("/", m.MetricsCtrl.Create).Name = "新建指标"

		index.Delete("/{id:uint}", m.MetricsCtrl.Delete).Name = "删除指标"
		index.Post("/{id:uint}/disable", m.MetricsCtrl.Disable).Name = "禁用指标"
		index.Post("/disable", m.MetricsCtrl.Disable).Name = "移动指标"
		index.Post("/move", m.MetricsCtrl.Move).Name = "移动指标"

		index.Get("/{id:uint}/getEntity", m.MetricsCtrl.GetEntity).Name = "获取指标实体"
	}

	return module.NewModule("/metrics", handler)
}

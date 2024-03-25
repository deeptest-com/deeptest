package router

import (
	"github.com/aaronchen2k/deeptest/cmd/server/v1/handler"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/module"
	"github.com/aaronchen2k/deeptest/internal/server/middleware"
	"github.com/kataras/iris/v12"
)

type PerformanceRunnerModule struct {
	PerformanceRunnerCtrl *handler.PerformanceRunnerCtrl `inject:""`
}

// Party 场景
func (m *PerformanceRunnerModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin(), middleware.ProjectPerm())

		index.Get("/", m.PerformanceRunnerCtrl.List).Name = "性能代理列表"
		index.Get("/{id:uint}", m.PerformanceRunnerCtrl.Get).Name = "性能代理详情"
		index.Post("/", m.PerformanceRunnerCtrl.Select).Name = "添加性能代理"
		index.Delete("/{id:uint}", m.PerformanceRunnerCtrl.Delete).Name = "删除性能代理"
	}

	return module.NewModule("/performanceTestRunners", handler)
}

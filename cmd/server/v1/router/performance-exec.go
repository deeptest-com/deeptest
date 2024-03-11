package router

import (
	"github.com/aaronchen2k/deeptest/cmd/server/v1/handler"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/module"
	"github.com/aaronchen2k/deeptest/internal/server/middleware"
	"github.com/kataras/iris/v12"
)

type PerformanceExecModule struct {
	PerformanceExecCtrl *handler.PerformanceExecCtrl `inject:""`
}

// Party 场景
func (m *PerformanceExecModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin(), middleware.ProjectPerm())

		index.Get("/loadScenario", m.PerformanceExecCtrl.LoadExecData).Name = "加载执行场景"
	}

	return module.NewModule("/performanceTestPlans/exec/", handler)
}

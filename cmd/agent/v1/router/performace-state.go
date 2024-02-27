package router

import (
	"github.com/aaronchen2k/deeptest/cmd/agent/v1/handler"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/module"
	"github.com/kataras/iris/v12"
)

type PerformanceStateModule struct {
	PerformanceStateCtrl *handler.PerformanceStateCtrl `inject:""`
}

// Party 性能测试状态模块
func (m *PerformanceStateModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Get("/", m.PerformanceStateCtrl.Get).Name = "获取性能测试状态"
	}
	return module.NewModule("/performanceState", handler)
}

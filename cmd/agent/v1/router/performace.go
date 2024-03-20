package router

import (
	"github.com/aaronchen2k/deeptest/cmd/agent/v1/handler"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/module"
	"github.com/kataras/iris/v12"
)

type PerformanceModule struct {
	PerformanceStateCtrl *handler.PerformanceCtrl `inject:""`
}

// Party 性能测试状态模块
func (m *PerformanceModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Get("/getState", m.PerformanceStateCtrl.GetState).Name = "获取性能测试状态"
		index.Get("/forceStop", m.PerformanceStateCtrl.ForceStop).Name = "强制停止所有执行"
	}

	return module.NewModule("/performance", handler)
}

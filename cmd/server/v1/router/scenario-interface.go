package router

import (
	"github.com/aaronchen2k/deeptest/cmd/server/v1/handler"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/module"
	"github.com/aaronchen2k/deeptest/internal/server/middleware"
	"github.com/kataras/iris/v12"
)

type ScenarioInterfaceModule struct {
	ScenarioInterfaceCtrl *handler.ScenarioInterfaceCtrl `inject:""`
}

func NewScenarioInterfaceModule() *ScenarioInterfaceModule {
	return &ScenarioInterfaceModule{}
}

// Party 场景
func (m *ScenarioInterfaceModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin())

		index.Get("/getInterface", m.ScenarioInterfaceCtrl.GetInterface).Name = "获取场景接口处理器"
		index.Get("/istInvocation", m.ScenarioInterfaceCtrl.ListInvocation).Name = "获取场景接口调用历史"
	}

	return module.NewModule("/scenarios/interfaces", handler)
}

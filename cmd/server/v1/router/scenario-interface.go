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

func (m *ScenarioInterfaceModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin())
		index.Post("/saveDebugData", m.ScenarioInterfaceCtrl.SaveDebugData).Name = "保存场景调试接口"
		index.Post("/resetDebugData", m.ScenarioInterfaceCtrl.ResetDebugData).Name = "重置场景调试接口"
	}

	return module.NewModule("/scenarios/interface", handler)
}

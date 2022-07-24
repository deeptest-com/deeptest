package index

import (
	"github.com/aaronchen2k/deeptest/internal/server/core/module"
	"github.com/aaronchen2k/deeptest/internal/server/middleware"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/controller"
	"github.com/kataras/iris/v12"
)

type ScenarioExecModule struct {
	ScenarioExecCtrl *controller.ScenarioExecCtrl `inject:""`
}

func NewScenarioExecModule() *ScenarioExecModule {
	return &ScenarioExecModule{}
}

// Party 场景
func (m *ScenarioExecModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin())
		index.Get("/loadExecResult", m.ScenarioExecCtrl.LoadExecData).Name = "加载场景"
	}

	return module.NewModule("/scenarios/exec/", handler)
}

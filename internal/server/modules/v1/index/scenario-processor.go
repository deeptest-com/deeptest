package index

import (
	"github.com/aaronchen2k/deeptest/internal/server/core/module"
	"github.com/aaronchen2k/deeptest/internal/server/middleware"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/controller"
	"github.com/kataras/iris/v12"
)

type ScenarioProcessorModule struct {
	ScenarioProcessorCtrl *controller.ScenarioProcessorCtrl `inject:""`
}

func NewScenarioProcessorModule() *ScenarioProcessorModule {
	return &ScenarioProcessorModule{}
}

// Party 场景
func (m *ScenarioProcessorModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin())

		index.Get("/{id:uint}", m.ScenarioProcessorCtrl.Get).Name = "场景节点详情"

		index.Put("/updateName", m.ScenarioProcessorCtrl.UpdateName).Name = "更新名称"
		index.Put("/{category}/save", m.ScenarioProcessorCtrl.Save).Name = "保存"
	}

	return module.NewModule("/scenarios/processors", handler)
}

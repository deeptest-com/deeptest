package index

import (
	"github.com/aaronchen2k/deeptest/internal/server/core/module"
	"github.com/aaronchen2k/deeptest/internal/server/middleware"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/controller"
	"github.com/kataras/iris/v12"
)

type ScenarioModule struct {
	ScenarioCtrl     *controller.ScenarioCtrl     `inject:""`
	ScenarioNodeCtrl *controller.ScenarioNodeCtrl `inject:""`
}

func NewScenarioModule() *ScenarioModule {
	return &ScenarioModule{}
}

// Party 场景
func (m *ScenarioModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin())

		index.Get("/", m.ScenarioCtrl.List).Name = "场景列表"
		index.Get("/{id:uint}", m.ScenarioCtrl.Get).Name = "场景详情"
		index.Post("/", m.ScenarioCtrl.Create).Name = "新建场景"
		index.Put("/", m.ScenarioCtrl.Update).Name = "更新场景"
		index.Delete("/{id:uint}", m.ScenarioCtrl.Delete).Name = "删除场景"

		index.Get("/load", m.ScenarioNodeCtrl.LoadTree).Name = "场景树状数据"
	}

	return module.NewModule("/scenarios", handler)
}

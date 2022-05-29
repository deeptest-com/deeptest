package index

import (
	"github.com/aaronchen2k/deeptest/internal/server/core/module"
	"github.com/aaronchen2k/deeptest/internal/server/middleware"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/controller"
	"github.com/kataras/iris/v12"
)

type ScenarioNodeModule struct {
	ScenarioNodeCtrl *controller.ScenarioNodeCtrl `inject:""`
}

func NewScenarioNodeModule() *ScenarioNodeModule {
	return &ScenarioNodeModule{}
}

// Party 场景
func (m *ScenarioNodeModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin())

		//party.Get("/{id:uint}", m.ScenarioNodeNodeCtrl.Get).Name = "场景节点详情"
		//party.Post("/", m.ScenarioNodeNodeCtrl.Create).Name = "新建节点"
		//party.Put("/{id:uint}", m.ScenarioNodeNodeCtrl.Update).Name = "更新节点"
		index.Put("/{id:uint}/updateName", m.ScenarioNodeCtrl.UpdateName).Name = "更新节点名称"
		//party.Delete("/{id:uint}", m.ScenarioNodeNodeCtrl.Delete).Name = "删除节点"
		//party.Post("/move", m.ScenarioNodeNodeCtrl.Move).Name = "移动节点"
	}

	return module.NewModule("/scenarios/nodes", handler)
}

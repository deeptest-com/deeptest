package router

import (
	"github.com/deeptest-com/deeptest/cmd/server/v1/handler"
	"github.com/deeptest-com/deeptest/internal/pkg/core/module"
	"github.com/deeptest-com/deeptest/internal/server/middleware"
	"github.com/kataras/iris/v12"
)

type ScenarioModule struct {
	ScenarioCtrl     *handler.ScenarioCtrl     `inject:""`
	ScenarioNodeCtrl *handler.ScenarioNodeCtrl `inject:""`
}

// Party 场景
func (m *ScenarioModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin(), middleware.ProjectPerm())
		//index.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin())
		index.Get("/listByProject", m.ScenarioCtrl.ListByProject)
		index.Get("/", m.ScenarioCtrl.List).Name = "场景列表"
		index.Get("/{id:uint}", m.ScenarioCtrl.Get).Name = "场景详情"
		index.Post("/", m.ScenarioCtrl.Create).Name = "新建场景"
		index.Put("/", m.ScenarioCtrl.Update).Name = "更新场景"
		index.Delete("/{id:uint}", m.ScenarioCtrl.Delete).Name = "删除场景"
		index.Get("/load", m.ScenarioNodeCtrl.LoadTree).Name = "场景树状数据"
		index.Post("/{id:uint}/addPlans", m.ScenarioCtrl.AddPlans).Name = "关联计划"
		index.Post("/{id:uint}/plans", m.ScenarioCtrl.Plans).Name = "关联计划列表"
		index.Put("/{id:uint}/updateStatus", m.ScenarioCtrl.UpdateStatus).Name = "更新场景状态"
		index.Post("/{id:uint}/removePlans", m.ScenarioCtrl.RemovePlans).Name = "取消计划关联"
		index.Put("/{id:uint}/updatePriority", m.ScenarioCtrl.UpdatePriority).Name = "更新优先级"
	}

	return module.NewModule("/scenarios", handler)
}

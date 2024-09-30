package router

import (
	"github.com/deeptest-com/deeptest/cmd/server/v1/handler"
	"github.com/deeptest-com/deeptest/internal/pkg/core/module"
	"github.com/deeptest-com/deeptest/internal/server/middleware"
	"github.com/kataras/iris/v12"
)

type PlanModule struct {
	PlanCtrl *handler.PlanCtrl `inject:""`
}

// Party 计划
func (m *PlanModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin(), middleware.ProjectPerm())

		index.Get("/", m.PlanCtrl.List).Name = "计划列表"
		index.Get("/{id:uint}", m.PlanCtrl.Get).Name = "计划详情"
		index.Post("/", m.PlanCtrl.Create).Name = "新建计划"
		index.Put("/", m.PlanCtrl.Update).Name = "更新计划"
		index.Delete("/{id:uint}", m.PlanCtrl.Delete).Name = "删除计划"
		index.Post("/{id:uint}/clone", m.PlanCtrl.Clone).Name = "克隆计划"
		index.Get("/planScenariosList", m.PlanCtrl.PlanScenariosList).Name = "计划中的场景列表"
		index.Get("/notRelationScenarioList", m.PlanCtrl.NotRelationScenarioList).Name = "计划中未绑定的场景列表"

		index.Post("/{id:uint}/addScenarios", m.PlanCtrl.AddScenarios).Name = "添加场景"
		index.Post("/{id:uint}/removeScenario", m.PlanCtrl.RemoveScenario).Name = "移除场景"
		index.Post("/{id:uint}/removeScenarios", m.PlanCtrl.RemoveScenarios).Name = "批量移除场景"

		index.Get("/statusDropDownOptions", m.PlanCtrl.StatusDropDownOptions).Name = "计划状态下拉选项"
		index.Get("/testStageDropDownOptions", m.PlanCtrl.TestStageDropDownOptions).Name = "计划测试阶段下拉选项"
		index.Put("/moveScenario", m.PlanCtrl.Move).Name = "移动场景"
	}

	return module.NewModule("/plans", handler)
}

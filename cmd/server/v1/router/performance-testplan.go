package router

import (
	"github.com/aaronchen2k/deeptest/cmd/server/v1/handler"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/module"
	"github.com/aaronchen2k/deeptest/internal/server/middleware"
	"github.com/kataras/iris/v12"
)

type PerformanceTestPlanModule struct {
	PerformanceTestPlanCtrl *handler.PerformanceTestPlanCtrl `inject:""`
}

// Party 场景
func (m *PerformanceTestPlanModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin(), middleware.ProjectPerm())

		index.Get("/", m.PerformanceTestPlanCtrl.List).Name = "性能测试计划列表"
		index.Get("/{id:uint}", m.PerformanceTestPlanCtrl.Get).Name = "性能测试计划详情"
		index.Post("/", m.PerformanceTestPlanCtrl.Create).Name = "新建性能测试计划"
		index.Put("/", m.PerformanceTestPlanCtrl.Update).Name = "更新性能测试计划"
		index.Delete("/{id:uint}", m.PerformanceTestPlanCtrl.Delete).Name = "删除性能测试计划"
		index.Put("/{id:uint}/updateStatus", m.PerformanceTestPlanCtrl.UpdateStatus).Name = "更新性能测试计划状态"

		index.Get("/listRunner", m.PerformanceTestPlanCtrl.ListRunner).Name = "获取性能测试代理列表"
		index.Get("/getConductor", m.PerformanceTestPlanCtrl.GetConductor).Name = "获取性能测试的主控节点信息"
	}

	return module.NewModule("/performanceTestPlans", handler)
}

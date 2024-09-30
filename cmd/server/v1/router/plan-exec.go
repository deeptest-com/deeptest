package router

import (
	"github.com/deeptest-com/deeptest/cmd/server/v1/handler"
	"github.com/deeptest-com/deeptest/internal/pkg/core/module"
	"github.com/deeptest-com/deeptest/internal/server/middleware"
	"github.com/kataras/iris/v12"
)

type PlanExecModule struct {
	PlanExecCtrl *handler.PlanExecCtrl `inject:""`
}

// Party
func (m *PlanExecModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin(), middleware.ProjectPerm())

		index.Get("/loadExecPlan", m.PlanExecCtrl.LoadExecData).Name = "加载执行计划"
		index.Get("/loadExecResult", m.PlanExecCtrl.LoadExecResult).Name = "加载执行结果"

		index.Post("/submitResult/{id:uint}", m.PlanExecCtrl.SubmitResult).Name = "提交测试结果"
		index.Get("/getPlanReportNormalData", m.PlanExecCtrl.GetPlanReportNormalData).Name = "获取计划执行中的静态内容"
	}

	return module.NewModule("/plans/exec/", handler)
}

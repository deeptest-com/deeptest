package router

import (
	"github.com/deeptest-com/deeptest/cmd/server/v1/handler"
	"github.com/deeptest-com/deeptest/internal/pkg/core/module"
	"github.com/deeptest-com/deeptest/internal/server/middleware"
	"github.com/kataras/iris/v12"
)

type OpenModule struct {
	OpenCtrl       *handler.OpenCtrl       `inject:""`
	DataCtrl       *handler.DataCtrl       `inject:""`
	PlanCtrl       *handler.PlanCtrl       `inject:""`
	PlanReportCtrl *handler.PlanReportCtrl `inject:""`
}

func (m *OpenModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Use(middleware.OpenCheck())
		index.Get("/allProjectList", m.OpenCtrl.AllProjectList).Name = "所有项目列表"
		index.Get("/getProjectsBySpace", m.OpenCtrl.GetProjectsBySpace).Name = "获取空间关联的项目列表"
		index.Post("/saveSpaceRelatedProjects", m.OpenCtrl.SaveSpaceRelatedProjects).Name = "保存空间和项目的关系"
		index.Get("/getProjectRole", m.OpenCtrl.GetProjectRole).Name = "获取用户项目角色"
		index.Post("/initdb", m.DataCtrl.Init)
		index.Post("/createPlan", m.PlanCtrl.Create)
		index.Post("/updatePlan", m.PlanCtrl.Update)
		index.Post("/copyPlan/{id:uint}", m.PlanCtrl.Clone)
		index.Post("/getPlan/{id:uint}", m.PlanCtrl.Get)
		index.Get("/report/{id:uint}", m.PlanReportCtrl.Delete)
	}

	return module.NewModule("/openApi", handler)
}

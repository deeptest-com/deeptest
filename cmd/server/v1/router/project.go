package router

import (
	"github.com/deeptest-com/deeptest/cmd/server/v1/handler"
	"github.com/deeptest-com/deeptest/internal/pkg/core/module"
	"github.com/deeptest-com/deeptest/internal/server/middleware"
	"github.com/kataras/iris/v12"
)

type ProjectModule struct {
	ProjectCtrl *handler.ProjectCtrl `inject:""`
}

// Party 项目
func (m *ProjectModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin())

		index.Get("/", m.ProjectCtrl.List).Name = "项目列表"
		index.Get("/{id:uint}", m.ProjectCtrl.Get).Name = "项目详情"
		index.Post("/", m.ProjectCtrl.Create).Name = "新建项目"
		index.Put("/", m.ProjectCtrl.Update).Name = "更新项目"
		index.Delete("/{id:uint}", m.ProjectCtrl.Delete).Name = "删除项目"

		index.Post("/changeProject", m.ProjectCtrl.ChangeProject).Name = "切换用户默认项目"
		index.Get("/getByUser", m.ProjectCtrl.GetByUser).Name = "获取用户参与的项目"

		index.Get("/members", m.ProjectCtrl.Members).Name = "获取项目成员"
		index.Post("/removeMember", m.ProjectCtrl.RemoveMember).Name = "删除项目成员"
		index.Post("/changeUserRole", m.ProjectCtrl.ChangeUserRole).Name = "更新项目成员的角色"

		index.Post("/apply", m.ProjectCtrl.Apply).Name = "申请项目成员"
		index.Post("/auditList", m.ProjectCtrl.AuditList).Name = "申请加入审批列表"
		index.Post("/audit", m.ProjectCtrl.Audit).Name = "审批操作"
		index.Get("/auditUsers", m.ProjectCtrl.AuditUsers).Name = "审批人"

		index.Get("/checkProjectAndUser", m.ProjectCtrl.CheckProjectAndUser).Name = "校验项目和成员是否存在"
		index.Get("/integrationDetail", m.ProjectCtrl.GetIntegrationDetail).Name = "集成详情"
		index.Get("/userProducts", m.ProjectCtrl.GetUserProducts).Name = "获取用户的产品列表"
		index.Get("/userSpaces", m.ProjectCtrl.GetUserSpaces).Name = "获取用户的空间列表"
		index.Get("/userEngineering", m.ProjectCtrl.GetMyEngineeringList).Name = "获取用户工程列表"
		index.Get("/projectEngineering", m.ProjectCtrl.GetEngineeringByProject).Name = "获取项目工程列表"

	}
	return module.NewModule("/projects", handler)
}

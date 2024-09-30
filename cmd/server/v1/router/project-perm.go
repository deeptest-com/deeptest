package router

import (
	"github.com/deeptest-com/deeptest/cmd/server/v1/handler"
	"github.com/deeptest-com/deeptest/internal/pkg/core/module"
	"github.com/deeptest-com/deeptest/internal/server/middleware"
	"github.com/kataras/iris/v12"
)

type ProjectPermModule struct {
	ProjectRolePermCtrl *handler.ProjectRolePermCtrl `inject:""`
}

// Party 项目权限
func (m *ProjectPermModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin())
		index.Get("/rolesList", m.ProjectRolePermCtrl.AllRoleList).Name = "所有项目角色列表"
		index.Get("/userRole", m.ProjectRolePermCtrl.GetProjectUserRole).Name = "获取项目中用户的角色"
		index.Get("/rolePermList", m.ProjectRolePermCtrl.RolePermList).Name = "项目角色的权限列表"
		index.Get("/userPermList", m.ProjectRolePermCtrl.UserPermList).Name = "项目中用户的权限列表"
	}
	return module.NewModule("/projects/perms", handler)
}

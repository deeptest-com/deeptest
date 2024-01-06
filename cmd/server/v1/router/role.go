package router

import (
	"github.com/aaronchen2k/deeptest/cmd/server/v1/handler"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/module"
	"github.com/aaronchen2k/deeptest/internal/server/middleware"
	"github.com/kataras/iris/v12"
)

type RoleModule struct {
	RoleCtrl *handler.RoleCtrl `inject:""`
}

// Party 角色模块
func (m *RoleModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin())
		index.Get("/", m.RoleCtrl.GetAllRoles).Name = "角色列表"
		index.Get("/{id:uint}", m.RoleCtrl.GetRole).Name = "角色详情"
		index.Post("/", m.RoleCtrl.CreateRole).Name = "新建角色"
		index.Post("/{id:uint}", m.RoleCtrl.UpdateRole).Name = "编辑角色"
		index.Delete("/{id:uint}", m.RoleCtrl.DeleteRole).Name = "删除角色"
		index.Get("/all", m.RoleCtrl.AllRoleList).Name = "无分页的角色列表"
		index.Get("/getAuth", m.RoleCtrl.GetAuthByEnv).Name = "获取系统级别权限"

	}
	return module.NewModule("/roles", handler)
}

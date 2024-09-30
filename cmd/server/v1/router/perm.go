package router

import (
	"github.com/deeptest-com/deeptest/cmd/server/v1/handler"
	"github.com/deeptest-com/deeptest/internal/pkg/core/module"
	"github.com/deeptest-com/deeptest/internal/server/middleware"
	"github.com/kataras/iris/v12"
)

type PermModule struct {
	PermCtrl *handler.PermCtrl `inject:""`
}

// Party 权限模块
func (m *PermModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin())
		index.Get("/", m.PermCtrl.GetAllPerms).Name = "权限列表"
		index.Get("/{id:uint}", m.PermCtrl.GetPerm).Name = "权限详情"
		index.Post("/", m.PermCtrl.CreatePerm).Name = "新建权限"
		index.Post("/{id:uint}", m.PermCtrl.UpdatePerm).Name = "编辑权限"
		index.Delete("/{id:uint}", m.PermCtrl.DeletePerm).Name = "删除权限"
	}
	return module.NewModule("/perms", handler)
}

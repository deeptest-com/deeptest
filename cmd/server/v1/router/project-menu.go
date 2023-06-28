package router

import (
	"github.com/aaronchen2k/deeptest/cmd/server/v1/handler"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/module"
	"github.com/aaronchen2k/deeptest/internal/server/middleware"
	"github.com/kataras/iris/v12"
)

type ProjectMenuModule struct {
	ProjectMenuCtrl *handler.ProjectMenuCtrl `inject:""`
}

// Party 项目菜单
func (m *ProjectMenuModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin())
		index.Get("/userMenuList", m.ProjectMenuCtrl.UserMenuList).Name = "项目中用户的左侧菜单栏列表"
	}
	return module.NewModule("/projects/menus", handler)
}

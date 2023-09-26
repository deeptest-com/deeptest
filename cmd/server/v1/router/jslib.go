package router

import (
	"github.com/aaronchen2k/deeptest/cmd/server/v1/handler"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/module"
	"github.com/aaronchen2k/deeptest/internal/server/middleware"
	"github.com/kataras/iris/v12"
)

type JslibModule struct {
	JslibCtrl *handler.JslibCtrl `inject:""`
}

// Party 项目
func (m *JslibModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin(), middleware.ProjectPerm())

		index.Get("/", m.JslibCtrl.List).Name = "列表"
		index.Get("/{id:uint}", m.JslibCtrl.Get).Name = "详情"
		index.Post("/", m.JslibCtrl.Save).Name = "保存"
		index.Put("/updateName", m.JslibCtrl.UpdateName).Name = "修改名称"
		index.Put("/{id:uint}/disable", m.JslibCtrl.Disable).Name = "禁用"
		index.Delete("/{id:uint}", m.JslibCtrl.Delete).Name = "删除"
	}

	return module.NewModule("/jslib", handler)
}

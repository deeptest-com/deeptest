package router

import (
	"github.com/deeptest-com/deeptest/cmd/server/v1/handler"
	"github.com/deeptest-com/deeptest/internal/pkg/core/module"
	"github.com/kataras/iris/v12"
)

type ComponentModule struct {
	ComponentCtrl *handler.ComponentCtrl `inject:""`
}

// Party 注册模块
func (m *ComponentModule) Party() module.WebModule {
	handler := func(public iris.Party) {
		//public.Use(middleware.InitCheck())
		//public.Use(middleware.JwtHandler())
		public.Get("/detail", m.ComponentCtrl.Detail)
		public.Post("/save", m.ComponentCtrl.Save)

		//public.Use(middleware.JwtHandler(), middleware.Casbin(), middleware.OperationRecord())
	}
	return module.NewModule("/endpoint", handler)
}

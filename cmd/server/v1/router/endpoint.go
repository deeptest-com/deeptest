package router

import (
	"github.com/aaronchen2k/deeptest/cmd/server/v1/handler"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/module"
	"github.com/kataras/iris/v12"
)

type EndpointModule struct {
	EndpointCtrl *handler.EndpointCtrl `inject:""`
}

func NewDemoModule() *EndpointModule {
	return &EndpointModule{}
}

// Party 注册模块
func (m *EndpointModule) Party() module.WebModule {
	handler := func(public iris.Party) {
		//public.Use(middleware.InitCheck())
		//public.Use(middleware.JwtHandler())
		public.Post("/index", m.EndpointCtrl.Index)
		public.Post("/save", m.EndpointCtrl.Save)
		public.Get("/detail", m.EndpointCtrl.Detail)
		public.Delete("/delete", m.EndpointCtrl.Delete)
		public.Put("/expire", m.EndpointCtrl.Expire)

		//public.Use(middleware.JwtHandler(), middleware.Casbin(), middleware.OperationRecord())
	}
	return module.NewModule("/endpoint", handler)
}

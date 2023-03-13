package router

import (
	"github.com/aaronchen2k/deeptest/cmd/server/v1/handler"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/module"
	"github.com/aaronchen2k/deeptest/internal/server/middleware"
	"github.com/kataras/iris/v12"
)

type EndpointModule struct {
	EndpointCtrl *handler.EndpointCtrl `inject:""`
}

func NewEndpointModule() *EndpointModule {
	return &EndpointModule{}
}

// Party 注册模块
func (m *EndpointModule) Party() module.WebModule {
	handler := func(public iris.Party) {
		public.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin())
		public.Post("/index", m.EndpointCtrl.Index)
		public.Post("/save", m.EndpointCtrl.Save)
		public.Get("/detail", m.EndpointCtrl.Detail)
		public.Delete("/delete", m.EndpointCtrl.Delete)
		public.Put("/expire", m.EndpointCtrl.Expire)
		public.Put("/publish", m.EndpointCtrl.Publish)
		public.Put("/develop", m.EndpointCtrl.Develop)
		public.Get("/copy", m.EndpointCtrl.Copy)
		public.Post("/yaml", m.EndpointCtrl.Yaml)
		public.Put("/updateStatus", m.EndpointCtrl.UpdateStatus)
		public.Delete("/batchDelete", m.EndpointCtrl.BatchDelete)
		public.Post("/version/add", m.EndpointCtrl.AddVersion)
		public.Get("/version/list", m.EndpointCtrl.ListVersions)
	}
	return module.NewModule("/endpoint", handler)
}

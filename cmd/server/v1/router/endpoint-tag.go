package router

import (
	"github.com/aaronchen2k/deeptest/cmd/server/v1/handler"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/module"
	"github.com/aaronchen2k/deeptest/internal/server/middleware"
	"github.com/kataras/iris/v12"
)

type EndpointTagModule struct {
	EndpointTagCtrl *handler.EndpointTagCtrl `inject:""`
}

// Party 注册模块
func (m *EndpointTagModule) Party() module.WebModule {
	handler := func(public iris.Party) {
		public.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin())
		public.Get("/", m.EndpointTagCtrl.ListTags).Name = "获取标签列表"
	}
	return module.NewModule("/endpoint/tags", handler)
}

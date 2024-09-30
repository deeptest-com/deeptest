package router

import (
	"github.com/deeptest-com/deeptest/cmd/server/v1/handler"
	"github.com/deeptest-com/deeptest/internal/pkg/core/module"
	"github.com/deeptest-com/deeptest/internal/server/middleware"
	"github.com/kataras/iris/v12"
)

type EndpointTagModule struct {
	EndpointTagCtrl *handler.EndpointTagCtrl `inject:""`
}

// Party 注册模块
func (m *EndpointTagModule) Party() module.WebModule {
	handler := func(public iris.Party) {
		public.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin(), middleware.ProjectPerm())
		public.Get("/", m.EndpointTagCtrl.ListTags).Name = "获取标签列表"
	}
	return module.NewModule("/endpoint/tags", handler)
}

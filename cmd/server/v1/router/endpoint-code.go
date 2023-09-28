package router

import (
	"github.com/aaronchen2k/deeptest/cmd/server/v1/handler"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/module"
	"github.com/kataras/iris/v12"
)

type EndpointCodeModule struct {
	EndpointCodeCtrl *handler.EndpointCodeCtrl `inject:""`
}

// Party 注册模块
func (m *EndpointCodeModule) Party() module.WebModule {
	handler := func(public iris.Party) {
		//public.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin(), middleware.ProjectPerm())

		public.Post("/generate", m.EndpointCodeCtrl.Index).Name = "生成代码"
	}

	return module.NewModule("/endpoints/code", handler)
}

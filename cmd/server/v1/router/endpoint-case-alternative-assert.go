package router

import (
	"github.com/aaronchen2k/deeptest/cmd/server/v1/handler"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/module"
	"github.com/aaronchen2k/deeptest/internal/server/middleware"
	"github.com/kataras/iris/v12"
)

type EndpointCaseAlternativeAssertModule struct {
	EndpointCaseAlternativeAssertCtrl *handler.EndpointCaseAlternativeAssertCtrl `inject:""`
}

// Party 注册模块
func (m *EndpointCaseAlternativeAssertModule) Party() module.WebModule {
	handler := func(public iris.Party) {
		public.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin(), middleware.ProjectPerm())

		public.Get("/", m.EndpointCaseAlternativeAssertCtrl.List).Name = "备选用例断言列表"
		public.Post("/", m.EndpointCaseAlternativeAssertCtrl.Save).Name = "创建备选用例断言"
		public.Delete("/{id:uint}", m.EndpointCaseAlternativeAssertCtrl.Delete).Name = "删除备选用例断言"
		public.Post("/{id:uint}/disable", m.EndpointCaseAlternativeAssertCtrl.Disable).Name = "禁用备选用例断言"
		public.Post("/{id:uint}/move", m.EndpointCaseAlternativeAssertCtrl.Move).Name = "移动备选用例断言"
	}

	return module.NewModule("/endpoints/cases/alternatives/assertions", handler)
}

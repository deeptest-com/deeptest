package router

import (
	"github.com/aaronchen2k/deeptest/cmd/server/v1/handler"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/module"
	"github.com/aaronchen2k/deeptest/internal/server/middleware"
	"github.com/kataras/iris/v12"
)

type EndpointCaseAlternativeModule struct {
	EndpointCaseAlternativeCtrl *handler.EndpointCaseAlternativeCtrl `inject:""`
}

// Party 注册模块
func (m *EndpointCaseAlternativeModule) Party() module.WebModule {
	handler := func(public iris.Party) {
		public.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin(), middleware.ProjectPerm())

		public.Post("/loadAlternative", m.EndpointCaseAlternativeCtrl.LoadAlternative).Name = "加载备选用例"
		public.Post("/generate", m.EndpointCaseAlternativeCtrl.GenerateCases).Name = "生成测试用例"
	}

	return module.NewModule("/endpoints/cases", handler)
}

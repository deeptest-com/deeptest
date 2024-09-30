package router

import (
	"github.com/deeptest-com/deeptest/cmd/server/v1/handler"
	"github.com/deeptest-com/deeptest/internal/pkg/core/module"
	"github.com/deeptest-com/deeptest/internal/server/middleware"
	"github.com/kataras/iris/v12"
)

type EndpointCaseModule struct {
	EndpointCaseCtrl *handler.EndpointCaseCtrl `inject:""`
}

// Party 注册模块
func (m *EndpointCaseModule) Party() module.WebModule {
	handler := func(public iris.Party) {
		public.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin(), middleware.ProjectPerm())

		public.Post("/query", m.EndpointCaseCtrl.Paginate).Name = "用例列表"
		public.Get("/{id:uint}", m.EndpointCaseCtrl.Get).Name = "用例详情"
		public.Post("/{id:uint}", m.EndpointCaseCtrl.Create).Name = "保存用例"
		public.Post("/copy", m.EndpointCaseCtrl.Copy).Name = "复制用例"
		public.Put("/updateName", m.EndpointCaseCtrl.UpdateName).Name = "保存用例名称"
		public.Post("/saveDebugData", m.EndpointCaseCtrl.SaveDebugData).Name = "保存调试数据"
		public.Delete("/{id:uint}", m.EndpointCaseCtrl.Remove).Name = "删除用例"
		public.Post("/loadTree", m.EndpointCaseCtrl.LoadTree).Name = "分类接口用例树"
		public.Get("/listForBenchmark", m.EndpointCaseCtrl.ListForBenchmark).Name = "分类接口用例树"

	}

	return module.NewModule("/endpoints/cases", handler)
}

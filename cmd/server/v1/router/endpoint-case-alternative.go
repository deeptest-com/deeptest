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

		public.Get("/load", m.EndpointCaseAlternativeCtrl.LoadAlternative).Name = "用例生成因子树状结构"
		public.Post("/loadSaved", m.EndpointCaseAlternativeCtrl.LoadAlternativeSaved).Name = "已保存备选用例列表"

		public.Post("/createBenchmark", m.EndpointCaseAlternativeCtrl.CreateBenchmark).Name = "创建基准用例"
		public.Post("/saveBenchmarkConfig", m.EndpointCaseAlternativeCtrl.SaveAlternative).Name = "保存备选为独立用例"
		public.Post("/saveAlternative", m.EndpointCaseAlternativeCtrl.SaveAlternative).Name = "保存备选为独立用例"

		public.Post("/loadCaseForExec", m.EndpointCaseAlternativeCtrl.LoadCaseForExec).Name = "获取用例执行数据"
	}

	return module.NewModule("/endpoints/cases/alternatives", handler)
}

package router

import (
	"github.com/deeptest-com/deeptest/cmd/server/v1/handler"
	"github.com/deeptest-com/deeptest/internal/pkg/core/module"
	"github.com/deeptest-com/deeptest/internal/server/middleware"
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

		public.Get("/loadFactor", m.EndpointCaseAlternativeCtrl.LoadFactor).Name = "获取备选路径因子的列表"

		public.Post("/createBenchmark", m.EndpointCaseAlternativeCtrl.CreateBenchmark).Name = "创建基准用例"
		public.Post("/saveFactor", m.EndpointCaseAlternativeCtrl.SaveFactor).Name = "保存备选路径因子取值"
		public.Post("/saveCase", m.EndpointCaseAlternativeCtrl.SaveCase).Name = "保存备选路径为独立用例"

		public.Post("/loadCasesForExec", m.EndpointCaseAlternativeCtrl.LoadCasesForExec).Name = "获取用例执行数据"
	}

	return module.NewModule("/endpoints/cases/alternatives", handler)
}

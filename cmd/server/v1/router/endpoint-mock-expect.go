package router

import (
	"github.com/aaronchen2k/deeptest/cmd/server/v1/handler"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/module"
	"github.com/aaronchen2k/deeptest/internal/server/middleware"
	"github.com/kataras/iris/v12"
)

type EndpointMockExpectModule struct {
	EndpointMockExpectCtrl *handler.EndpointMockExpectCtrl `inject:""`
}

// Party 高级Mock-Mock期望模块
func (m *EndpointMockExpectModule) Party() module.WebModule {
	handler := func(public iris.Party) {
		public.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin(), middleware.ProjectPerm())

		public.Get("/list", m.EndpointMockExpectCtrl.List).Name = "期望列表"
		public.Get("/{id:uint}", m.EndpointMockExpectCtrl.Detail).Name = "期望详情"
		public.Post("/save", m.EndpointMockExpectCtrl.Save).Name = "保存期望"
		public.Get("/copy", m.EndpointMockExpectCtrl.Copy).Name = "复制期望"
		public.Delete("/{id:uint}", m.EndpointMockExpectCtrl.Delete).Name = "删除期望"
		public.Post("/order", m.EndpointMockExpectCtrl.Order).Name = "对期望排序"
		public.Post("/updateExpectDisabled", m.EndpointMockExpectCtrl.UpdateExpectDisabled).Name = "启用/禁用单个期望"
		public.Post("/updateName", m.EndpointMockExpectCtrl.UpdateExpectName).Name = "修改期望名字"
		public.Get("/requestOptions", m.EndpointMockExpectCtrl.GetExpectRequestOptions).Name = "获取请求参数下拉选项"

		public.Post("/{endpointId:uint}/disable", m.EndpointMockExpectCtrl.Disable).Name = "更新项目"
	}
	return module.NewModule("/mockExpect", handler)
}

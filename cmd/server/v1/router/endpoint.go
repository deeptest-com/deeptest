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

		public.Post("/index", m.EndpointCtrl.Index).Name = "设计器列表"
		public.Post("/save", m.EndpointCtrl.Save).Name = "保存设计器"
		public.Get("/detail", m.EndpointCtrl.Detail).Name = "设计器详情"
		public.Delete("/delete", m.EndpointCtrl.Delete).Name = "删除设计器"
		public.Put("/expire", m.EndpointCtrl.Expire).Name = "禁用设计器"
		public.Put("/publish", m.EndpointCtrl.Publish).Name = "发布设计器"
		public.Put("/develop", m.EndpointCtrl.Develop).Name = "开发设计器"
		public.Get("/copy", m.EndpointCtrl.Copy).Name = "复制设计器"
		public.Post("/yaml", m.EndpointCtrl.Yaml).Name = "设计器信息转yaml"
		public.Put("/updateStatus", m.EndpointCtrl.UpdateStatus).Name = "更新设计器状态"
		public.Delete("/batchDelete", m.EndpointCtrl.BatchDelete).Name = "批量删除"
		public.Post("/version/add", m.EndpointCtrl.AddVersion).Name = "添加版本"
		public.Get("/version/list", m.EndpointCtrl.ListVersions).Name = "版本列表"
	}
	return module.NewModule("/endpoint", handler)
}

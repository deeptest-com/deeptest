package router

import (
	"github.com/aaronchen2k/deeptest/cmd/server/v1/handler"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/module"
	"github.com/aaronchen2k/deeptest/internal/server/middleware"
	"github.com/kataras/iris/v12"
)

type ServeModule struct {
	ServeCtrl *handler.ServeCtrl `inject:""`
}

// Party 注册模块
func (m *ServeModule) Party() module.WebModule {
	handler := func(public iris.Party) {
		public.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin(), middleware.ProjectPerm())

		public.Post("/index", m.ServeCtrl.Index).Name = "服务列表"
		public.Post("/save", m.ServeCtrl.Save).Name = "保存服务"
		public.Get("/detail", m.ServeCtrl.Detail).Name = "服务详情"
		public.Delete("/delete", m.ServeCtrl.Delete).Name = "删除服务"
		public.Put("/expire", m.ServeCtrl.Expire).Name = "禁用服务"
		public.Get("/copy", m.ServeCtrl.Copy).Name = "复制服务"
		public.Post("/version/list", m.ServeCtrl.ListVersion).Name = "版本列表"
		public.Post("/version/save", m.ServeCtrl.SaveVersion).Name = "保存版本"
		public.Post("/version/bindEndpoint", m.ServeCtrl.BindEndpoint).Name = "关联接口"
		public.Delete("/version/delete", m.ServeCtrl.DeleteVersion).Name = "删除版本"
		public.Put("/version/expire", m.ServeCtrl.ExpireVersion).Name = "禁用版本"

		public.Post("/server/list", m.ServeCtrl.ListServer).Name = "环境列表"
		public.Post("/server/changeServer", m.ServeCtrl.ChangeServer).Name = "切换环境"

		public.Post("/schema/save", m.ServeCtrl.SaveSchema).Name = "保存Schema"
		public.Post("/schema/list", m.ServeCtrl.ListSchema).Name = "Schema列表"
		public.Post("/schema/detail", m.ServeCtrl.GetSchemaByRef).Name = "获取Schema"
		public.Delete("/schema/delete", m.ServeCtrl.DeleteSchema).Name = "删除Schema列表"
		public.Post("/security/save", m.ServeCtrl.SaveSecurity).Name = "保存授权"
		public.Post("/security/list", m.ServeCtrl.ListSecurity).Name = "授权列表"
		public.Delete("/security/delete", m.ServeCtrl.DeleteSecurity).Name = "删除授权"
		public.Put("/schema/copy", m.ServeCtrl.CopySchema).Name = "复制Schema"
		public.Post("/schema/example2schema", m.ServeCtrl.ExampleToSchema).Name = "example转schema"
		public.Post("/schema/schema2example", m.ServeCtrl.SchemaToExample).Name = "Schema生成Example"
		public.Post("/schema/schema2yaml", m.ServeCtrl.SchemaToYaml).Name = "schema转yaml"

		public.Get("/listByProject", m.ServeCtrl.ListByProject).Name = "获取项目下的服务"
		public.Post("/changeServe", m.ServeCtrl.ChangeServe).Name = "切换用户当前服务"
		public.Post("/addServerForHistory", m.ServeCtrl.AddServerForHistory).Name = "为历史项目和服务增加环境"
	}

	return module.NewModule("/serves", handler)
}

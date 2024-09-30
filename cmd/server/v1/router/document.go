package router

import (
	"github.com/deeptest-com/deeptest/cmd/server/v1/handler"
	"github.com/deeptest-com/deeptest/internal/pkg/core/module"
	"github.com/deeptest-com/deeptest/internal/server/middleware"
	"github.com/kataras/iris/v12"
)

type DocumentModule struct {
	DocumentCtrl *handler.DocumentCtrl `inject:""`
}

func (m *DocumentModule) Party() module.WebModule {
	handler := func(public iris.Party) {
		public.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin(), middleware.ProjectPerm())
		public.Post("/share", m.DocumentCtrl.GetShareLink).Name = "生成分享接口文档的链接"
		public.Get("/get_share_content", m.DocumentCtrl.GetContentsByShareLink).Name = "查看分享的文档"
		public.Get("/share_detail", m.DocumentCtrl.GetDocumentDetail).Name = "查看分享的文档详情"
		public.Post("/", m.DocumentCtrl.Index).Name = "接口文档"
		public.Post("/version_list", m.DocumentCtrl.DocumentVersionList).Name = "接口文档版本列表"
		public.Post("/publish", m.DocumentCtrl.Publish).Name = "发布接口文档"
		public.Delete("/delete", m.DocumentCtrl.DeleteSnapshot).Name = "删除接口文档"
		public.Post("/update_version", m.DocumentCtrl.UpdateDocument).Name = "更新文档版本信息"
		public.Get("/detail", m.DocumentCtrl.GetDocumentDetail).Name = "查看文档详情"

	}
	return module.NewModule("/document", handler)
}

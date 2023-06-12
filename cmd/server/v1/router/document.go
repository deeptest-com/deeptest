package router

import (
	"github.com/aaronchen2k/deeptest/cmd/server/v1/handler"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/module"
	"github.com/kataras/iris/v12"
)

type DocumentModule struct {
	DocumentCtrl *handler.DocumentCtrl `inject:""`
}

func (m *DocumentModule) Party() module.WebModule {
	handler := func(public iris.Party) {
		//public.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin())
		public.Post("/", m.DocumentCtrl.Index).Name = "接口文档"

	}
	return module.NewModule("/document", handler)
}

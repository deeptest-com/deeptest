package router

import (
	"github.com/aaronchen2k/deeptest/cmd/server/v1/handler"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/module"
	"github.com/aaronchen2k/deeptest/internal/server/middleware"
	"github.com/kataras/iris/v12"
)

type CommonModule struct {
	BaseCtrl *handler.BaseCtrl `inject:""`
}

// Party 注册模块
func (m *CommonModule) Party() module.WebModule {
	handler := func(public iris.Party) {
		public.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin())

		public.Post("/batchUpdateField", m.BaseCtrl.BatchUpdateField).Name = "批量更新字段内容"
	}
	return module.NewModule("/common", handler)
}

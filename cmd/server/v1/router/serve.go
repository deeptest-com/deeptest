package router

import (
	"github.com/aaronchen2k/deeptest/cmd/server/v1/handler"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/module"
	"github.com/kataras/iris/v12"
)

type ServeModule struct {
	ServeCtrl *handler.ServeCtrl `inject:""`
}

func NewServeModule() *ServeModule {
	return &ServeModule{}
}

// Party 注册模块
func (m *ServeModule) Party() module.WebModule {
	handler := func(public iris.Party) {
		//public.Use(middleware.InitCheck())
		//public.Use(middleware.JwtHandler())
		public.Post("/save", m.ServeCtrl.Save)
		public.Post("/index", m.ServeCtrl.Index)
		public.Get("/detail", m.ServeCtrl.Detail)
		public.Delete("/delete", m.ServeCtrl.Delete)
		public.Put("/expire", m.ServeCtrl.Expire)

		public.Get("/version/list", m.ServeCtrl.ListVersion)
		public.Post("/version/save", m.ServeCtrl.SaveVersion)
		public.Delete("/version/delete", m.ServeCtrl.DeleteVersion)
		public.Put("/version/expire", m.ServeCtrl.ExpireVersion)

		//public.Use(middleware.JwtHandler(), middleware.Casbin(), middleware.OperationRecord())
	}
	return module.NewModule("/serve", handler)
}

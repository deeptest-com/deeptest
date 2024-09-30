package router

import (
	"github.com/deeptest-com/deeptest/cmd/server/v1/handler"
	"github.com/deeptest-com/deeptest/internal/pkg/core/module"
	"github.com/deeptest-com/deeptest/internal/server/middleware"
	"github.com/kataras/iris/v12"
)

type AccountModule struct {
	AccountCtrl *handler.AccountCtrl `inject:""`
}

// Party 认证模块
func (m *AccountModule) Party() module.WebModule {
	handler := func(public iris.Party) {
		public.Use(middleware.InitCheck())

		public.Post("/login", m.AccountCtrl.Login)
		public.Get("/login", m.AccountCtrl.Login)
		public.Post("/register", m.AccountCtrl.Register)

		public.Post("/forgotPassword", m.AccountCtrl.ForgotPassword)
		public.Post("/resetPassword", m.AccountCtrl.ResetPassword)

		public.Use(middleware.JwtHandler(), middleware.Casbin(), middleware.OperationRecord())
	}
	return module.NewModule("/account", handler)
}

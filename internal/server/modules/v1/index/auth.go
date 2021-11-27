package index

import (
	"github.com/aaronchen2k/deeptest/internal/server/core/module"
	"github.com/aaronchen2k/deeptest/internal/server/middleware"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/controller"
	"github.com/kataras/iris/v12"
)

type AuthModule struct {
	AuthCtrl *controller.AuthCtrl `inject:""`
}

func NewAuthModule() *AuthModule {
	return &AuthModule{}
}

// Party 认证模块
func (m *AuthModule) Party() module.WebModule {
	handler := func(public iris.Party) {
		public.Use(middleware.InitCheck())
		public.Post("/login", m.AuthCtrl.Login)
		public.Use(middleware.JwtHandler(), middleware.Casbin(), middleware.OperationRecord())
	}
	return module.NewModule("/auth", handler)
}

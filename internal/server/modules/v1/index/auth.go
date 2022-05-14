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

// Party 脚本
func (m *AuthModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin())

		index.Post("/genOAuth2AccessToken", m.AuthCtrl.GenOAuth2AccessToken).Name = "生成OAuth认证信息"
	}
	return module.NewModule("/auth", handler)
}

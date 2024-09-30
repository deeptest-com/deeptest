package router

import (
	"github.com/deeptest-com/deeptest/cmd/server/v1/handler"
	"github.com/deeptest-com/deeptest/internal/pkg/core/module"
	"github.com/deeptest-com/deeptest/internal/server/middleware"
	"github.com/kataras/iris/v12"
)

type AuthModule struct {
	AuthCtrl *handler.AuthCtrl `inject:""`
}

// Party 脚本
func (m *AuthModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin())

		index.Post("/oauth2Authorization", m.AuthCtrl.OAuth2Authorization).Name = "生成OAuth认证信息"
		index.Post("/getOAuth2AccessToken", m.AuthCtrl.GetOAuth2AccessToken).Name = "调用认证服务生成访问令牌"
		index.Post("/useOAuth2AccessToken", m.AuthCtrl.UseOAuth2AccessToken).Name = "加载访问令牌到接口"

		index.Get("/listOAuth2Token", m.AuthCtrl.ListOAuth2Token).Name = "加载访问令牌到接口"
		index.Get("/removeToken", m.AuthCtrl.RemoveToken).Name = "加载访问令牌到接口"

	}
	return module.NewModule("/auth", handler)
}

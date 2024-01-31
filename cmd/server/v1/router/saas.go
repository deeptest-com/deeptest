package router

import (
	"github.com/aaronchen2k/deeptest/cmd/server/v1/handler"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/module"
	"github.com/kataras/iris/v12"
)

type SaaSModule struct {
	SaasCtrl *handler.SaasCtrl `inject:""`
}

// Party 角色模块
func (m *SaaSModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		//index.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin())
		index.Get("/", m.SaasCtrl.GetUserList).Name = "用户列表"

	}
	return module.NewModule("/saas", handler)
}

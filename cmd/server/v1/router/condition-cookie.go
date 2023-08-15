package router

import (
	"github.com/aaronchen2k/deeptest/cmd/server/v1/handler"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/module"
	"github.com/aaronchen2k/deeptest/internal/server/middleware"
	"github.com/kataras/iris/v12"
)

type CookieModule struct {
	CookieCtrl *handler.CookieCtrl `inject:""`
}

// Party Cookie
func (m *CookieModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin())

		index.Get("/{id:uint}", m.CookieCtrl.Get).Name = "Cookie详情"
		index.Put("/", m.CookieCtrl.Update).Name = "更新Cookie"
		index.Delete("/{id:uint}", m.CookieCtrl.Delete).Name = "删除Cookie"
	}

	return module.NewModule("/cookies", handler)
}

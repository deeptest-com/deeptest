package router

import (
	"github.com/aaronchen2k/deeptest/cmd/server/v1/handler"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/module"
	"github.com/aaronchen2k/deeptest/internal/server/middleware"
	"github.com/kataras/iris/v12"
)

type ScriptModule struct {
	ScriptCtrl *handler.ScriptCtrl `inject:""`
}

// Party 脚本
func (m *ScriptModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin(), middleware.ProjectPerm())

		index.Get("/{id:uint}", m.ScriptCtrl.Get).Name = "脚本详情"
		index.Put("/", m.ScriptCtrl.Update).Name = "更新脚本"
	}

	return module.NewModule("/scripts", handler)
}

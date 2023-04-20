package router

import (
	"github.com/aaronchen2k/deeptest/cmd/server/v1/handler"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/module"
	"github.com/aaronchen2k/deeptest/internal/server/middleware"
	"github.com/kataras/iris/v12"
)

type ShareVarModule struct {
	ShareVarCtrl *handler.ShareVarCtrl `inject:""`
}

// Party 提取器
func (m *ShareVarModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin())

		index.Post("/list", m.ShareVarCtrl.List).Name = "列出变量列表"
		index.Delete("/{id:uint}", m.ShareVarCtrl.Delete).Name = "删除共享变量"
		index.Post("/clear", m.ShareVarCtrl.Clear).Name = "清空共享变量"
	}

	return module.NewModule("/shareVars", handler)
}

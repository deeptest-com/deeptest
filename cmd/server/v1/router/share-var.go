package router

import (
	"github.com/deeptest-com/deeptest/cmd/server/v1/handler"
	"github.com/deeptest-com/deeptest/internal/pkg/core/module"
	"github.com/deeptest-com/deeptest/internal/server/middleware"
	"github.com/kataras/iris/v12"
)

type ShareVarModule struct {
	ShareVarCtrl *handler.ShareVarCtrl `inject:""`
}

// Party 提取器
func (m *ShareVarModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin(), middleware.ProjectPerm())

		index.Post("/list", m.ShareVarCtrl.List).Name = "列出变量列表"
		index.Delete("/{id:uint}", m.ShareVarCtrl.Delete).Name = "删除共享变量"
		index.Post("/clear", m.ShareVarCtrl.Clear).Name = "清空共享变量"
	}

	return module.NewModule("/shareVars", handler)
}

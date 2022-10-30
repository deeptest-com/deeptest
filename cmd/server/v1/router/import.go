package router

import (
	"github.com/aaronchen2k/deeptest/cmd/server/v1/handler"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/module"
	"github.com/aaronchen2k/deeptest/internal/server/middleware"
	"github.com/kataras/iris/v12"
)

type ImportModule struct {
	ImportCtrl *handler.ImportCtrl `inject:""`
}

// Party 脚本
func (m *ImportModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin())

		index.Post("/importSpec", m.ImportCtrl.ImportSpec).Name = "导入OpenApi文件"

	}
	return module.NewModule("/import", handler)
}

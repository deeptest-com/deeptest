package router

import (
	"github.com/deeptest-com/deeptest/cmd/server/v1/handler"
	"github.com/deeptest-com/deeptest/internal/pkg/core/module"
	"github.com/deeptest-com/deeptest/internal/server/middleware"
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
		index.Post("/importYapi", m.ImportCtrl.ImportYapi).Name = "导入yapi项目接口"

	}
	return module.NewModule("/import", handler)
}

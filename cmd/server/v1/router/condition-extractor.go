package router

import (
	"github.com/deeptest-com/deeptest/cmd/server/v1/handler"
	"github.com/deeptest-com/deeptest/internal/pkg/core/module"
	"github.com/deeptest-com/deeptest/internal/server/middleware"
	"github.com/kataras/iris/v12"
)

type ExtractorModule struct {
	ExtractorCtrl *handler.ExtractorCtrl `inject:""`
}

// Party 提取器
func (m *ExtractorModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin(), middleware.ProjectPerm())

		index.Get("/{id:uint}", m.ExtractorCtrl.Get).Name = "提取器详情"
		index.Put("/", m.ExtractorCtrl.Update).Name = "更新提取器"

		index.Post("/quickCreate", m.ExtractorCtrl.QuickCreate).Name = "从代码编辑器快速创建提取器"

		index.Post("/listExtractorVariableForCheckpoint", m.ExtractorCtrl.ListExtractorVariableForCheckpoint).Name = "提取器变量列表"
	}

	return module.NewModule("/extractors", handler)
}

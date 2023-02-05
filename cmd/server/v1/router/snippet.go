package router

import (
	"github.com/aaronchen2k/deeptest/cmd/server/v1/handler"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/module"
	"github.com/aaronchen2k/deeptest/internal/server/middleware"
	"github.com/kataras/iris/v12"
)

type SnippetModule struct {
	SnippetCtrl *handler.SnippetCtrl `inject:""`
}

// Party 脚本
func (m *SnippetModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin())

		index.Get("/", m.SnippetCtrl.Get).Name = "获取详情"
	}
	return module.NewModule("/snippets", handler)
}

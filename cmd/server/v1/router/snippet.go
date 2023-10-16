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

		index.Get("/", m.SnippetCtrl.Get).Name = "获取代码片段"

		index.Get("/listJslibNames", m.SnippetCtrl.ListJslibNames).Name = "获取所有自定义库名称"
		index.Get("/getJslibs", m.SnippetCtrl.GetJslibs).Name = "获取用户自定义脚本库"
		index.Post("/getJslibsForAgent", m.SnippetCtrl.GetJslibsForAgent).Name = "获取用户自定义脚本库"
	}
	return module.NewModule("/snippets", handler)
}

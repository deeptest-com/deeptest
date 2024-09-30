package router

import (
	"github.com/deeptest-com/deeptest/cmd/server/v1/handler"
	"github.com/deeptest-com/deeptest/internal/pkg/core/module"
	"github.com/deeptest-com/deeptest/internal/server/middleware"
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
		index.Get("/listVar", m.SnippetCtrl.ListVar).Name = "获取变量列表"
		index.Get("/listMock", m.SnippetCtrl.ListMock).Name = "获取mock规则"
		index.Get("/listSysFunc", m.SnippetCtrl.ListSysFunc).Name = "获取系统函数"
		index.Get("/ListCustomFunc", m.SnippetCtrl.ListCustomFunc).Name = "创建代码片段"

	}
	return module.NewModule("/snippets", handler)
}

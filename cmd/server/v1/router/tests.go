package router

import (
	"github.com/aaronchen2k/deeptest/cmd/server/v1/handler"
	"github.com/aaronchen2k/deeptest/internal/pkg/config"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/module"
	"github.com/kataras/iris/v12"
)

type TestsModule struct {
	TestsCtrl *handler.TestsCtrl `inject:""`
	FileCtrl  *handler.FileCtrl  `inject:""`
}

// Party 脚本
func (m *TestsModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Get("/", m.TestsCtrl.Gets).Name = "模拟Get"
		index.Post("/", m.TestsCtrl.Posts).Name = "模拟Post"
		index.Put("/", m.TestsCtrl.Posts).Name = "模拟Put"
		index.Delete("/", m.TestsCtrl.Posts).Name = "模拟Delete"
		index.Patch("/", m.TestsCtrl.Posts).Name = "模拟Patch"
		index.Head("/", m.TestsCtrl.Head).Name = "模拟Head"
		index.Connect("/", m.TestsCtrl.Connect).Name = "模拟Connect"
		index.Trace("/", m.TestsCtrl.Trace).Name = "模拟Trace"

		index.Post("/upload", iris.LimitRequestBodySize(config.CONFIG.MaxSize*iris.MB),
			m.FileCtrl.Upload).Name = "模拟上传文件"

		index.Post("/stream", iris.LimitRequestBodySize(config.CONFIG.MaxSize*iris.MB),
			m.TestsCtrl.Stream).Name = "模拟Stream流"
	}
	return module.NewModule("/test", handler)
}

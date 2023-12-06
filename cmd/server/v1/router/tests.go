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
		index.Get("/", m.TestsCtrl.Gets).Name = "模拟接口测试"
		index.Post("/", m.TestsCtrl.Posts).Name = "模拟接口测试"
		index.Put("/", m.TestsCtrl.Posts).Name = "模拟接口测试"
		index.Delete("/", m.TestsCtrl.Posts).Name = "模拟接口测试"

		index.Patch("/", m.TestsCtrl.Posts).Name = "模拟接口测试"
		index.Head("/", m.TestsCtrl.Head).Name = "模拟接口测试"

		index.Connect("/", m.TestsCtrl.Connect).Name = "模拟接口测试"
		index.Trace("/", m.TestsCtrl.Trace).Name = "模拟接口测试"

		index.Post("/upload", iris.LimitRequestBodySize(config.CONFIG.MaxSize*iris.MB),
			m.FileCtrl.Upload).Name = "上传文件测试"
	}
	return module.NewModule("/test", handler)
}

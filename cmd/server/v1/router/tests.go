package router

import (
	"github.com/aaronchen2k/deeptest/cmd/server/v1/handler"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/module"
	"github.com/kataras/iris/v12"
)

type TestsModule struct {
	TestsCtrl *handler.TestsCtrl `inject:""`
}

// Party 脚本
func (m *TestsModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Get("/", m.TestsCtrl.Test).Name = "测试用"
	}
	return module.NewModule("/tests", handler)
}

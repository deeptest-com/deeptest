package router

import (
	"github.com/aaronchen2k/deeptest/cmd/server/v1/handler"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/module"
	"github.com/kataras/iris/v12"
)

type MockModule struct {
	MockCtrl *handler.MockCtrl `inject:""`
}

func NewMockModule() *MockModule {
	return &MockModule{}
}

// Party 脚本
func (m *MockModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Get("/invoke", m.MockCtrl.Get).Name = "模拟接口测试"
		index.Post("/invoke", m.MockCtrl.Request).Name = "模拟接口测试"
		index.Put("/invoke", m.MockCtrl.Request).Name = "模拟接口测试"
		index.Delete("/invoke", m.MockCtrl.Request).Name = "模拟接口测试"

		index.Patch("/invoke", m.MockCtrl.Request).Name = "模拟接口测试"
		index.Head("/invoke", m.MockCtrl.Head).Name = "模拟接口测试"

		index.Connect("/invoke", m.MockCtrl.Connect).Name = "模拟接口测试"
		index.Trace("/invoke", m.MockCtrl.Trace).Name = "模拟接口测试"
	}
	return module.NewModule("/mock", handler)
}

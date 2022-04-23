package index

import (
	"github.com/aaronchen2k/deeptest/internal/server/core/module"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/controller"
	"github.com/kataras/iris/v12"
)

type MockModule struct {
	MockCtrl *controller.MockCtrl `inject:""`
}

func NewMockModule() *MockModule {
	return &MockModule{}
}

// Party 脚本
func (m *MockModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.PartyFunc("/json", func(party iris.Party) {
			party.Get("/invoke", m.MockCtrl.Get).Name = "模拟接口测试"
			party.Post("/invoke", m.MockCtrl.Request).Name = "模拟接口测试"
			party.Put("/invoke", m.MockCtrl.Request).Name = "模拟接口测试"
			party.Delete("/invoke", m.MockCtrl.Request).Name = "模拟接口测试"

			party.Patch("/invoke", m.MockCtrl.Request).Name = "模拟接口测试"
			party.Head("/invoke", m.MockCtrl.Head).Name = "模拟接口测试"

			party.Connect("/invoke", m.MockCtrl.Connect).Name = "模拟接口测试"
			party.Trace("/invoke", m.MockCtrl.Trace).Name = "模拟接口测试"
		})
	}
	return module.NewModule("/mock", handler)
}

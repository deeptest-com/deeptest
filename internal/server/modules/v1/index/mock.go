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
			party.Get("/request", m.MockCtrl.Request).Name = "模拟接口测试"
			party.Post("/request", m.MockCtrl.Request).Name = "模拟接口测试"
			party.Put("/request", m.MockCtrl.Request).Name = "模拟接口测试"
			party.Delete("/request", m.MockCtrl.Request).Name = "模拟接口测试"
		})
	}
	return module.NewModule("/mock", handler)
}

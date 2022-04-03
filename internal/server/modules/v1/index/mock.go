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
			party.Get("/get", m.MockCtrl.Get).Name = "用于接口GET测试"
			party.Post("/post", m.MockCtrl.Post).Name = "用于接口POST测试"
		})
	}
	return module.NewModule("/mock", handler)
}

package index

import (
	"github.com/aaronchen2k/deeptest/internal/server/core/module"
	"github.com/aaronchen2k/deeptest/internal/server/middleware"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/controller"
	"github.com/kataras/iris/v12"
)

type TestStepModule struct {
	TestStepCtrl *controller.TestStepCtrl `inject:""`
}

func NewTestStepModule() *TestStepModule {
	return &TestStepModule{}
}

// Party 步骤
func (m *TestStepModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin())

		index.Get("/{id:uint}", m.TestStepCtrl.Get).Name = "步骤详情"
		index.Post("/", m.TestStepCtrl.Create).Name = "新建步骤"
		index.Put("/{id:uint}", m.TestStepCtrl.Update).Name = "更新步骤"
		index.Delete("/{id:uint}", m.TestStepCtrl.Delete).Name = "删除步骤"
	}
	return module.NewModule("/steps", handler)
}

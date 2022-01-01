package index

import (
	"github.com/aaronchen2k/deeptest/internal/server/core/module"
	"github.com/aaronchen2k/deeptest/internal/server/middleware"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/controller"
	"github.com/kataras/iris/v12"
)

type TestCaseModule struct {
	TestCaseCtrl *controller.TestCaseCtrl `inject:""`
}

func NewTestCaseModule() *TestCaseModule {
	return &TestCaseModule{}
}

// Party 项目
func (m *TestCaseModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin())
		index.Get("/", m.TestCaseCtrl.Query).Name = "项目查询"
		index.Get("/{id:uint}", m.TestCaseCtrl.Get).Name = "项目详情"
		index.Post("/", m.TestCaseCtrl.Create).Name = "新建项目"
		index.Post("/{id:uint}", m.TestCaseCtrl.Update).Name = "编辑项目"
		index.Delete("/{id:uint}", m.TestCaseCtrl.Delete).Name = "删除项目"
	}
	return module.NewModule("/TestCases", handler)
}

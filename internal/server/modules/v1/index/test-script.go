package index

import (
	"github.com/kataras/iris/v12"
	"github.com/aaronchen2k/deeptest/internal/server/core/module"
	"github.com/aaronchen2k/deeptest/internal/server/middleware"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/controller"
)

type TestScriptModule struct {
	TestScriptCtrl *controller.TestScriptCtrl `inject:""`
}

func NewTestScriptModule() *TestScriptModule {
	return &TestScriptModule{}
}

// Party 项目
func (m *TestScriptModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin())
		index.Get("/", m.TestScriptCtrl.Query).Name = "项目查询"
		index.Get("/{id:uint}", m.TestScriptCtrl.Get).Name = "项目详情"
		index.Post("/", m.TestScriptCtrl.Create).Name = "创建项目"
		index.Post("/{id:uint}", m.TestScriptCtrl.Update).Name = "编辑项目"
		index.Delete("/{id:uint}", m.TestScriptCtrl.Delete).Name = "删除项目"
	}
	return module.NewModule("/TestScripts", handler)
}

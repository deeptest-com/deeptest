package router

import (
	"github.com/deeptest-com/deeptest/cmd/server/v1/handler"
	"github.com/deeptest-com/deeptest/internal/pkg/core/module"
	"github.com/deeptest-com/deeptest/internal/server/middleware"
	"github.com/kataras/iris/v12"
)

type ScenarioExecModule struct {
	ScenarioExecCtrl *handler.ScenarioExecCtrl `inject:""`
}

// Party 场景
func (m *ScenarioExecModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin(), middleware.ProjectPerm())

		index.Get("/loadExecScenario", m.ScenarioExecCtrl.LoadExecData).Name = "加载执行场景"
		index.Get("/loadExecResult", m.ScenarioExecCtrl.LoadExecResult).Name = "加载执行结果"

		index.Post("/submitResult/{id:uint}", m.ScenarioExecCtrl.SubmitResult).Name = "提交测试结果"
		index.Get("/getScenarioNormalData", m.ScenarioExecCtrl.GetScenarioNormalData).Name = "获取场景执行初始话信息"

	}

	return module.NewModule("/scenarios/exec/", handler)
}

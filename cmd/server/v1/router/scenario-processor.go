package router

import (
	"github.com/deeptest-com/deeptest/cmd/server/v1/handler"
	"github.com/deeptest-com/deeptest/internal/pkg/core/module"
	"github.com/deeptest-com/deeptest/internal/server/middleware"
	"github.com/kataras/iris/v12"
)

type ScenarioProcessorModule struct {
	ScenarioProcessorCtrl *handler.ScenarioProcessorCtrl `inject:""`
}

// Party 场景
func (m *ScenarioProcessorModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin(), middleware.ProjectPerm())

		index.Get("/{id:uint}", m.ScenarioProcessorCtrl.Get).Name = "场景节点详情"

		//index.Put("/updateName", m.ScenarioProcessorCtrl.UpdateName).Name = "更新名称"
		index.Put("/saveProcessorInfo", m.ScenarioProcessorCtrl.SaveBasicInfo).Name = "保存基本信息"
		index.Put("/{category}/save", m.ScenarioProcessorCtrl.Save).Name = "保存配置信息"
	}

	return module.NewModule("/scenarios/processors", handler)
}

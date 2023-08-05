package router

import (
	"github.com/aaronchen2k/deeptest/cmd/server/v1/handler"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/module"
	"github.com/aaronchen2k/deeptest/internal/server/middleware"
	"github.com/kataras/iris/v12"
)

type ScenarioNodeModule struct {
	ScenarioNodeCtrl *handler.ScenarioNodeCtrl `inject:""`
}

// Party 场景
func (m *ScenarioNodeModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin())
		index.Post("/addInterfacesFromDefine", m.ScenarioNodeCtrl.AddInterfacesFromDefine).Name = "添加定义接口"
		index.Post("/addInterfacesFromTest", m.ScenarioNodeCtrl.AddInterfacesFromTest).Name = "添加调试接口"
		index.Post("/addProcessor", m.ScenarioNodeCtrl.AddProcessor).Name = "新建处理器"

		index.Put("/{id:uint}/updateName", m.ScenarioNodeCtrl.UpdateName).Name = "更新节点名称"
		index.Delete("/{id:uint}", m.ScenarioNodeCtrl.Delete).Name = "删除节点"
		index.Post("/move", m.ScenarioNodeCtrl.Move).Name = "移动节点"
		index.Post("/importCurl", m.ScenarioNodeCtrl.ImportCurl).Name = "curl导入"
	}

	return module.NewModule("/scenarios/nodes", handler)
}

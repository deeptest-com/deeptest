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
		index.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin(), middleware.ProjectPerm())
		index.Post("/addInterfacesFromDefine", m.ScenarioNodeCtrl.AddInterfacesFromDefine).Name = "添加定义接口"
		index.Post("/addInterfacesFromDiagnose", m.ScenarioNodeCtrl.AddInterfacesFromDiagnose).Name = "添加调试接口"
		index.Post("/addInterfacesFromCase", m.ScenarioNodeCtrl.AddInterfacesFromCase).Name = "添加接口用例"
		index.Post("/addProcessor", m.ScenarioNodeCtrl.AddProcessor).Name = "新建处理器"
		index.Post("/copyProcessor", m.ScenarioNodeCtrl.CopyProcessor).Name = "复制处理器"

		index.Put("/{id:uint}/updateName", m.ScenarioNodeCtrl.UpdateName).Name = "更新节点名称"
		index.Delete("/{id:uint}", m.ScenarioNodeCtrl.Delete).Name = "删除节点"
		index.Post("/{id:uint}/disableOrNot", m.ScenarioNodeCtrl.DisableOrNot).Name = "删除节点"
		index.Post("/move", m.ScenarioNodeCtrl.Move).Name = "移动节点"
		index.Post("/importCurl", m.ScenarioNodeCtrl.ImportCurl).Name = "curl导入"
	}

	return module.NewModule("/scenarios/nodes", handler)
}

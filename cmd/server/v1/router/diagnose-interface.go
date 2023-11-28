package router

import (
	"github.com/aaronchen2k/deeptest/cmd/server/v1/handler"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/module"
	"github.com/aaronchen2k/deeptest/internal/server/middleware"
	"github.com/kataras/iris/v12"
)

type DiagnoseInterfaceModule struct {
	DiagnoseInterfaceCtrl *handler.DiagnoseInterfaceCtrl `inject:""`
}

// Party 脚本
func (m *DiagnoseInterfaceModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin(), middleware.ProjectPerm())

		index.Get("/", m.DiagnoseInterfaceCtrl.Load).Name = "获取测试接口"
		index.Get("/{id:uint}", m.DiagnoseInterfaceCtrl.Get).Name = "获取测试接口"
		index.Post("/", m.DiagnoseInterfaceCtrl.Save).Name = "新建测试接口"
		index.Put("/", m.DiagnoseInterfaceCtrl.Update).Name = "更新测试接口"
		index.Delete("/{id:uint}", m.DiagnoseInterfaceCtrl.Delete).Name = "删除测试接口"
		index.Post("/move", m.DiagnoseInterfaceCtrl.Move).Name = "移动节点"

		index.Post("/importInterfaces", m.DiagnoseInterfaceCtrl.ImportInterfaces).Name = "导入接口"
		index.Post("/importCurl", m.DiagnoseInterfaceCtrl.ImportCurl).Name = "导入cURL命令"
		index.Post("/importRecordData", m.DiagnoseInterfaceCtrl.ImportRecordData).Name = "导入录制的接口"

		index.Post("/saveDebugData", m.DiagnoseInterfaceCtrl.SaveDebugData).Name = "保存测试调试接口"
	}
	return module.NewModule("/diagnoseInterfaces", handler)
}

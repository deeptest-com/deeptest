package router

import (
	"github.com/deeptest-com/deeptest/cmd/server/v1/handler"
	"github.com/deeptest-com/deeptest/internal/pkg/core/module"
	"github.com/deeptest-com/deeptest/internal/server/middleware"
	"github.com/kataras/iris/v12"
)

type WebsocketInterfaceModule struct {
	WebsocketInterfaceCtrl *handler.WebsocketInterfaceCtrl `inject:""`
}

// Party 脚本
func (m *WebsocketInterfaceModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin(), middleware.ProjectPerm())

		index.Get("/getDebugData", m.WebsocketInterfaceCtrl.GetDebugData).Name = "获取Websocket测试接口"
		index.Put("/saveDebugData", m.WebsocketInterfaceCtrl.SaveDebugData).Name = "保存Websocket测试接口"
	}
	return module.NewModule("/websocketInterfaces", handler)
}

package router

import (
	"github.com/aaronchen2k/deeptest/cmd/server/v1/handler"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/module"
	"github.com/aaronchen2k/deeptest/internal/server/middleware"
	"github.com/kataras/iris/v12"
)

type GrpcInterfaceModule struct {
	GrpcInterfaceCtrl *handler.GrpcInterfaceCtrl `inject:""`
}

// Party 脚本
func (m *GrpcInterfaceModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin(), middleware.ProjectPerm())

		index.Get("/getDebugData", m.GrpcInterfaceCtrl.GetDebugData).Name = "获取gRPC测试接口"
		index.Put("/saveDebugData", m.GrpcInterfaceCtrl.SaveDebugData).Name = "保存gRPC测试接口"

		index.Post("/parseProto", m.GrpcInterfaceCtrl.ParseProto).Name = "解析proto"
		index.Post("/describeFunc", m.GrpcInterfaceCtrl.DescribeFunc).Name = "解析proto"

		index.Post("/listConn", m.GrpcInterfaceCtrl.ListConn).Name = "列出连接"
		index.Post("/deleteHandle", m.GrpcInterfaceCtrl.DeleteHandle).Name = "移除Handle"

		index.Post("/invokeFunc", m.GrpcInterfaceCtrl.InvokeFunc).Name = "调用方法"
	}

	return module.NewModule("/grpcInterfaces", handler)
}

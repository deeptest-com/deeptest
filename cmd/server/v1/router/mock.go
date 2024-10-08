package router

import (
	"github.com/deeptest-com/deeptest/cmd/server/v1/handler"
	"github.com/deeptest-com/deeptest/internal/pkg/core/module"
	"github.com/kataras/iris/v12"
)

type MockModule struct {
	MockCtrl *handler.MockCtrl `inject:""`
}

// Party 脚本
func (m *MockModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		//	index.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin(), middleware.ProjectPerm())
		index.Get("/{serveId:int}/{path:path}", m.MockCtrl.Mock).Name = "测试"
		index.Post("/{serveId:int}/{path:path}", m.MockCtrl.Mock).Name = "测试"
		index.Put("/{serveId:int}/{path:path}", m.MockCtrl.Mock).Name = "测试"
		index.Delete("/{serveId:int}/{path:path}", m.MockCtrl.Mock).Name = "测试"
		index.Patch("/{serveId:int}/{path:path}", m.MockCtrl.Mock).Name = "测试"
		index.Head("/{serveId:int}/{path:path}", m.MockCtrl.Mock).Name = "测试"
		index.Connect("/{serveId:int}/{path:path}", m.MockCtrl.Mock).Name = "测试"
		index.Trace("/{serveId:int}/{path:path}", m.MockCtrl.Mock).Name = "测试"
	}
	return module.NewModule("/", handler)
}

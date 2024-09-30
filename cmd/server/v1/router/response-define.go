package router

import (
	"github.com/deeptest-com/deeptest/cmd/server/v1/handler"
	"github.com/deeptest-com/deeptest/internal/pkg/core/module"
	"github.com/deeptest-com/deeptest/internal/server/middleware"
	"github.com/kataras/iris/v12"
)

type ResponseDefineModule struct {
	ResponseDefineCtrl *handler.ResponseDefineCtrl `inject:""`
}

// Party 检查点
func (m *ResponseDefineModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin())
		index.Put("/", m.ResponseDefineCtrl.Update).Name = "更新检查点"
	}
	return module.NewModule("/responseDefine", handler)
}

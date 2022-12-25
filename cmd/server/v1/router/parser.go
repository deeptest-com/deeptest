package router

import (
	"github.com/aaronchen2k/deeptest/cmd/server/v1/handler"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/module"
	"github.com/aaronchen2k/deeptest/internal/server/middleware"
	"github.com/kataras/iris/v12"
)

type ParserModule struct {
	ParserCtrl *handler.ParserCtrl `inject:""`
}

// Party 脚本
func (m *ParserModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin())

		index.Post("/parseHtml", m.ParserCtrl.ParseHtml).Name = "解析HTML"

		index.Post("/testXPath", m.ParserCtrl.TestXPath).Name = "测试XPath"
	}
	return module.NewModule("/parser", handler)
}

package router

import (
	"github.com/deeptest-com/deeptest/cmd/server/v1/handler"
	"github.com/deeptest-com/deeptest/internal/pkg/core/module"
	"github.com/deeptest-com/deeptest/internal/server/middleware"
	"github.com/kataras/iris/v12"
)

type ParserModule struct {
	ParserCtrl *handler.ParserCtrl `inject:""`
}

// Party 脚本
func (m *ParserModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin(), middleware.ProjectPerm())

		index.Post("/parseHtml", m.ParserCtrl.ParseHtml).Name = "解析HTML"
		index.Post("/parseXml", m.ParserCtrl.ParseXml).Name = "解析XML"
		index.Post("/parseJson", m.ParserCtrl.ParseJson).Name = "解析JSON"
		index.Post("/parseText", m.ParserCtrl.ParseText).Name = "解析TEXT"

		index.Post("/testExpr", m.ParserCtrl.TestExpr).Name = "测试XPath或正则表达式"
	}
	return module.NewModule("/parser", handler)
}

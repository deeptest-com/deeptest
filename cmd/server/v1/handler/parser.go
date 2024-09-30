package handler

import (
	"github.com/deeptest-com/deeptest/cmd/server/v1/domain"
	"github.com/deeptest-com/deeptest/internal/server/modules/service"
	"github.com/deeptest-com/deeptest/pkg/domain"
	"github.com/kataras/iris/v12"
)

type ParserCtrl struct {
	ParserService     *service.ParserService     `inject:""`
	ParserHtmlService *service.ParserHtmlService `inject:""`
	ParserXmlService  *service.ParserXmlService  `inject:""`
	ParserJsonService *service.ParserJsonService `inject:""`
	ParserRegxService *service.ParserRegxService `inject:""`
	BaseCtrl
}

// ParseHtml
// @Tags	解析模块
// @summary	解析HTML
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization	header	string						true	"Authentication header"
// @Param 	currProjectId	query	int							true	"当前项目ID"
// @Param 	ParserRequest 	body 	serverDomain.ParserRequest 	true 	"解析HTML的请求参数"
// @success	200	{object}	_domain.Response{data=serverDomain.ParserResponse}
// @Router	/api/v1/parser/parseHtml	[post]
func (c *ParserCtrl) ParseHtml(ctx iris.Context) {
	req := serverDomain.ParserRequest{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	resp, err := c.ParserHtmlService.ParseHtml(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.FailErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: resp, Msg: _domain.NoErr.Msg})
}

// ParseXml
// @Tags	解析模块
// @summary	解析XML
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization	header	string						true	"Authentication header"
// @Param 	currProjectId	query	int							true	"当前项目ID"
// @Param 	ParserRequest 	body 	serverDomain.ParserRequest 	true 	"解析XML的请求参数"
// @success	200	{object}	_domain.Response{data=serverDomain.ParserResponse}
// @Router	/api/v1/parser/parseXml	[post]
func (c *ParserCtrl) ParseXml(ctx iris.Context) {
	req := serverDomain.ParserRequest{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	resp, err := c.ParserXmlService.ParseXml(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.FailErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: resp, Msg: _domain.NoErr.Msg})
}

// ParseJson
// @Tags	解析模块
// @summary	解析JSON
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization	header	string						true	"Authentication header"
// @Param 	currProjectId	query	int							true	"当前项目ID"
// @Param 	ParserRequest 	body 	serverDomain.ParserRequest 	true 	"解析JSON的请求参数"
// @success	200	{object}	_domain.Response{data=serverDomain.ParserResponse}
// @Router	/api/v1/parser/parseJson	[post]
func (c *ParserCtrl) ParseJson(ctx iris.Context) {
	req := serverDomain.ParserRequest{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	resp, err := c.ParserJsonService.ParseJson(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.FailErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: resp, Msg: _domain.NoErr.Msg})
}

// ParseText
// @Tags	解析模块
// @summary	解析TEXT
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization	header	string						true	"Authentication header"
// @Param 	currProjectId	query	int							true	"当前项目ID"
// @Param 	ParserRequest 	body 	serverDomain.ParserRequest 	true 	"解析TEXT的请求参数"
// @success	200	{object}	_domain.Response{data=serverDomain.ParserResponse}
// @Router	/api/v1/parser/parseText	[post]
func (c *ParserCtrl) ParseText(ctx iris.Context) {
	req := serverDomain.ParserRequest{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	resp, err := c.ParserRegxService.ParseRegx(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.FailErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: resp, Msg: _domain.NoErr.Msg})
}

// TestExpr
// @Tags	解析模块
// @summary	测试XPath或正则表达式
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization	header	string							true	"Authentication header"
// @Param 	currProjectId	query	int								true	"当前项目ID"
// @Param 	TestExprRequest body 	serverDomain.TestExprRequest 	true 	"测试XPath或正则表达式的请求参数"
// @success	200	{object}	_domain.Response{data=serverDomain.TestExprResponse}
// @Router	/api/v1/parser/testExpr	[post]
func (c *ParserCtrl) TestExpr(ctx iris.Context) {
	req := serverDomain.TestExprRequest{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	resp, err := c.ParserService.TestExpr(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.FailErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: resp, Msg: _domain.NoErr.Msg})
}

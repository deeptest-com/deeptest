package handler

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	queryHelper "github.com/aaronchen2k/deeptest/internal/agent/exec/utils/query"
	"github.com/aaronchen2k/deeptest/internal/server/modules/service"
	"github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/kataras/iris/v12"
)

type ParserCtrl struct {
	ParserService     *service.ParserService     `inject:""`
	ParserHtmlService *service.ParserHtmlService `inject:""`
	ParserXmlService  *service.ParserXmlService  `inject:""`
	ParserJsonService *service.ParserJsonService `inject:""`
	BaseCtrl
}

// ParseHtml
func (c *ParserCtrl) ParseHtml(ctx iris.Context) {
	req := v1.ParserRequest{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
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
func (c *ParserCtrl) ParseXml(ctx iris.Context) {
	req := v1.ParserRequest{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
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
func (c *ParserCtrl) ParseJson(ctx iris.Context) {
	req := v1.ParserRequest{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	resp, err := c.ParserJsonService.ParseJson(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.FailErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: resp, Msg: _domain.NoErr.Msg})
}

// TestXPath
func (c *ParserCtrl) TestXPath(ctx iris.Context) {
	req := v1.TestXPathRequest{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	resp, err := c.ParserService.TestXPath(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.FailErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: resp, Msg: _domain.NoErr.Msg})
}

// TestRegx
func (c *ParserCtrl) TestRegx(ctx iris.Context) {
	req := v1.TestRegxRequest{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	resp := queryHelper.RegxQuery(req.Content, req.Expr)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.FailErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: resp, Msg: _domain.NoErr.Msg})
}

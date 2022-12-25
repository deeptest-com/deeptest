package handler

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	mockHelper "github.com/aaronchen2k/deeptest/internal/server/modules/helper/mock"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/service"
	"github.com/aaronchen2k/deeptest/pkg/domain"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/kataras/iris/v12"
)

type MockCtrl struct {
	MockService *service.MockService `inject:""`
	BaseCtrl
}

func (c *MockCtrl) OAuth2Callback(ctx iris.Context) {

}

func (c *MockCtrl) Get(ctx iris.Context) {
	respType := ctx.URLParam("respType")

	username, password, ok := ctx.Request().BasicAuth()
	if ok {
		logUtils.Infof("BasicAuth - username: %s, password: %s, ok: %t", username, password, ok)
	}

	authorization := ctx.GetHeader(consts.Authorization)
	logUtils.Infof("JWT Token - %s", authorization)

	value := ctx.GetHeader("k1")
	logUtils.Infof("API KEY - %s: %s", "k1", value)

	if respType == "html" {
		ctx.HTML(mockHelper.GetHtmlData())
	} else if respType == "xml" {
		ctx.XML(mockHelper.GetXmlData())
	} else if respType == "json" {
		ctx.JSON(mockHelper.GetJsonData())
	} else {
		ctx.HTML(mockHelper.GetTextData())
	}
}

func (c *MockCtrl) Request(ctx iris.Context) {
	var req model.Invocation
	err := ctx.ReadQuery(&req)
	if err != nil {
		logUtils.Errorf("参数获取失败", err.Error())
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	data, err := c.MockService.Exec(req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: err.Error()})
		return
	}

	reqBodyType := ctx.GetHeader(consts.ContentType)

	if reqBodyType == consts.ContentTypeJSON.String() {
		ctx.Header(consts.ContentType, consts.ContentTypeJSON.String()+";charset=utf-8")
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
	} else if reqBodyType == consts.ContentTypeXML.String() {
		ctx.Header(consts.ContentType, consts.ContentTypeXML.String()+";charset=utf-8")
		ctx.XML(_domain.Response{Code: _domain.NoErr.Code, Data: req, Msg: _domain.NoErr.Msg})
	} else if reqBodyType == consts.ContentTypeHTML.String() {
		ctx.Header(consts.ContentType, consts.ContentTypeHTML.String()+";charset=utf-8")
		ctx.HTML(mockHelper.GetHtmlData())
	} else {
		ctx.Header(consts.ContentType, consts.ContentTypeTEXT.String()+";charset=utf-8")
		ctx.Text(mockHelper.GetTextData())
	}
}

func (c *MockCtrl) Head(ctx iris.Context) {
	ctx.Header(consts.Server, "kataras iris v12")
}

func (c *MockCtrl) Connect(ctx iris.Context) {
	ctx.Header(consts.Server, "kataras iris v12")
}

func (c *MockCtrl) Options(ctx iris.Context) {
	ctx.Header(consts.Server, "kataras iris v12")
	ctx.Header(consts.Allow, "GET, POST, PUT, DELETE, PATCH, HEAD, CONNECT, OPTIONS, TRACE")
	ctx.Header(consts.ContentType, consts.ContentTypeUnixDir.String())
}

func (c *MockCtrl) Trace(ctx iris.Context) {
	ctx.Header(consts.Server, "kataras iris v12")
	ctx.Header(consts.Connection, "close")
	ctx.Header(consts.Host, "deeptest.com")
}

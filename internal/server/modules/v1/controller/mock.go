package controller

import (
	"github.com/aaronchen2k/deeptest/internal/comm/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	logUtils "github.com/aaronchen2k/deeptest/internal/pkg/lib/log"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/service"
	"github.com/kataras/iris/v12"
)

type MockCtrl struct {
	MockService *service.MockService `inject:""`
	BaseCtrl
}

func (c *MockCtrl) Get(ctx iris.Context) {
	ctx.Header(consts.ContentType, consts.ContentTypeJSON.String()+";charset=utf-8")

	data := iris.Map{}
	data["content"] = "test"

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
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
	} else {
		ctx.Header(consts.ContentType, consts.ContentTypeHTML.String()+";charset=utf-8")
		ctx.HTML("<html><p>Hello World!</p><html>")
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

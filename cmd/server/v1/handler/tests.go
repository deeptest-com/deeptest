package handler

import (
	serverDomain "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	mockGenerator "github.com/aaronchen2k/deeptest/internal/pkg/helper/openapi-mock/openapi/generator"
	"github.com/aaronchen2k/deeptest/internal/server/modules/service"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/kataras/iris/v12"
)

type TestsCtrl struct {
	BaseCtrl

	StreamTestService *service.StreamTestService `inject:""`
}

func (c *TestsCtrl) Gets(ctx iris.Context) {
	_, respStatusCode, respContentType := c.getParams(ctx)

	resp := mockGenerator.Response{
		StatusCode:  respStatusCode,
		ContentType: respContentType,
		Data: iris.Map{
			"key":    "key",
			"value":  "value",
			"params": ctx.URLParams(),

			"int":   1,
			"float": 6.6,
			"bool":  true,
			"obj": iris.Map{
				"name":  "aaron",
				"email": "462826@qq.com",
			},
			"arr": []interface{}{
				iris.Map{"name": "foo"},
				iris.Map{"name": "bar"},
			},
		},
	}

	username, password, ok := ctx.Request().BasicAuth()
	if ok {
		logUtils.Infof("BasicAuth - username: %s, password: %s, ok: %t", username, password, ok)
	}

	authorization := ctx.GetHeader(consts.Authorization)
	logUtils.Infof("JWT Token - %s", authorization)

	value := ctx.GetHeader("k1")
	logUtils.Infof("API KEY - %s: %s", "k1", value)

	co := ctx.GetCookie("cookie_from_client")

	ctx.SetCookieKV("cookie_from_client", "token_"+co)
	ctx.SetCookieKV("cookie_from_server", "value_from_server")

	c.WriteRespByContentType(resp, ctx)
}

func (c *TestsCtrl) Posts(ctx iris.Context) {
	reqContentType, respStatusCode, respContentType := c.getParams(ctx)

	resp := mockGenerator.Response{
		StatusCode:  respStatusCode,
		ContentType: respContentType,
	}

	if reqContentType == consts.ContentTypeJSON {
		var req serverDomain.MockReqJson
		ctx.ReadJSON(&req)

		resp.Data = iris.Map{"req": req}

	} else if reqContentType == consts.ContentTypeFormData {
		name := ctx.FormValue("name")
		password := ctx.FormValue("password")
		file, _, _ := ctx.FormFile("file")

		resp.Data = iris.Map{"name": name, "password": password, "file": file}
	}

	c.WriteRespByContentType(resp, ctx)
}

func (c *TestsCtrl) Head(ctx iris.Context) {
	ctx.Header(consts.Server, "kataras iris v12")
}

func (c *TestsCtrl) Connect(ctx iris.Context) {
	ctx.Header(consts.Server, "kataras iris v12")
}

func (c *TestsCtrl) Options(ctx iris.Context) {
	ctx.Header(consts.Server, "kataras iris v12")
	ctx.Header(consts.Allow, "GET, POST, PUT, DELETE, PATCH, HEAD, CONNECT, OPTIONS, TRACE")
	ctx.Header(consts.ContentType, consts.ContentTypeUnixDir.String())
}

func (c *TestsCtrl) Trace(ctx iris.Context) {
	ctx.Header(consts.Server, "kataras iris v12")
	ctx.Header(consts.Connection, "close")
	ctx.Header(consts.Host, "deeptest.com")
}

func (c *TestsCtrl) getParams(ctx iris.Context) (reqContentType consts.HttpContentType, respStatusCode consts.HttpRespCode, respContentType consts.HttpContentType) {
	reqContentTypeStr := ctx.URLParam("reqContentTypeStr")
	respStatusCodeInt, _ := ctx.URLParamInt("respStatusCodeInt")
	respContentTypeStr := ctx.URLParam("respContentTypeStr")

	if reqContentTypeStr != "" {
		reqContentType = consts.HttpContentType(reqContentTypeStr)
	} else {
		reqContentType = consts.ContentTypeJSON
	}

	if respStatusCodeInt > 0 {
		respStatusCode = consts.HttpRespCode(respStatusCodeInt)
	} else {
		respStatusCode = consts.OK
	}

	if respContentTypeStr != "" {
		respContentType = consts.HttpContentType(respContentTypeStr)
	} else {
		respContentType = consts.ContentTypeJSON
	}

	return
}

func (c *TestsCtrl) Stream(ctx iris.Context) {
	flusher, ok := ctx.ResponseWriter().Flusher()
	if !ok {
		ctx.StopWithText(iris.StatusHTTPVersionNotSupported, "Streaming unsupported!")
		return
	}

	//ctx.Header("content-type", "text/event-stream")
	ctx.ContentType("text/event-stream")
	ctx.Header("Cache-Control", "no-cache")
	ctx.SetCookieKV("cookie_from_server", "value123")

	req := serverDomain.StreamTestObj{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	c.StreamTestService.Chat(req, flusher, ctx)
}

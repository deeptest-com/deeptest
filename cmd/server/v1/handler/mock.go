package handler

import (
	serverDomain "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	mockHelper "github.com/aaronchen2k/deeptest/internal/server/modules/helper/mock"
	"github.com/aaronchen2k/deeptest/internal/server/modules/service"
	"github.com/aaronchen2k/deeptest/pkg/domain"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/kataras/iris/v12"
	"log"
)

type MockCtrl struct {
	MockService *service.MockService `inject:""`
	BaseCtrl
}

func (c *MockCtrl) OAuth2Callback(ctx iris.Context) {

}

func (c *MockCtrl) Mock(ctx iris.Context) {
	// http://127.0.0.1:8085/mocks/serve_id/json?id=44

	method := ctx.Method()
	serveId, _ := ctx.Params().GetInt("serveId")
	path := ctx.Params().Get("path")
	endpointInterfaceId := ctx.URLParamIntDefault("id", 0)

	logUtils.Infof("%s %d/%s", method, serveId, path)

	req := service.MockRequest{
		ServeId:             serveId,
		EndpointMethod:      consts.HttpMethod(method),
		EndpointPath:        path,
		EndpointInterfaceId: uint(endpointInterfaceId),
	}

	resp, err := c.MockService.ByRequest(&req, ctx)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	ctx.StatusCode(resp.StatusCode)
	ctx.ContentType(resp.ContentType)
	//ctx.Write(data)
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

	co := ctx.GetCookie("cookie_from_client")
	log.Print(co)

	ctx.SetCookieKV("cookie_from_client", "token_"+co)
	ctx.SetCookieKV("cookie_from_server", "value_from_server")

	if respType == "html" {
		ctx.HTML(mockHelper.GetHtmlData())
	} else if respType == "xml" {
		ctx.XML(mockHelper.GetXmlData())
	} else if respType == "json" {
		ctx.JSON(mockHelper.GetJsonData())
	} else {
		ctx.Text(mockHelper.GetTextData())
	}
}

func (c *MockCtrl) Posts(ctx iris.Context) {
	repType := ctx.URLParam("reqType")
	reqBodyType := ctx.GetHeader(consts.ContentType)

	var data interface{}
	var err error

	if repType == "json" {
		var req serverDomain.MockReqJson
		ctx.ReadJSON(&req)
		data = iris.Map{"req": req}
	} else if repType == "form" {
		name := ctx.FormValue("name")
		password := ctx.FormValue("password")
		data = iris.Map{"name": name, "password": password}
	} else if repType == "file" {
		name := ctx.FormValue("name")
		data, _, err = ctx.FormFile("myFile")
		data = iris.Map{"name": name, "data": data}
	}

	if reqBodyType == consts.ContentTypeJSON.String() {
		ctx.Header(consts.ContentType, consts.ContentTypeJSON.String()+";charset=utf-8")
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
	} else if reqBodyType == consts.ContentTypeXML.String() {
		ctx.Header(consts.ContentType, consts.ContentTypeXML.String()+";charset=utf-8")
		ctx.XML(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
	} else if reqBodyType == consts.ContentTypeHTML.String() {
		ctx.Header(consts.ContentType, consts.ContentTypeHTML.String()+";charset=utf-8")
		ctx.HTML(mockHelper.GetHtmlData())
	} else {
		ctx.Header(consts.ContentType, consts.ContentTypeTEXT.String()+";charset=utf-8")
		ctx.Text(mockHelper.GetTextData())
	}

	log.Println(err)
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

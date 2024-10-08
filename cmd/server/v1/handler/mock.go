package handler

import (
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	"github.com/deeptest-com/deeptest/internal/server/modules/service"
	"github.com/deeptest-com/deeptest/pkg/domain"
	logUtils "github.com/deeptest-com/deeptest/pkg/lib/log"
	"github.com/kataras/iris/v12"
)

type MockCtrl struct {
	MockService *service.MockService `inject:""`
	BaseCtrl
}

func (c *MockCtrl) OAuth2Callback(ctx iris.Context) {

}

func (c *MockCtrl) Mock(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	// http://127.0.0.1:8085/mocks/serve_id/json?id=44

	method := ctx.Method()
	serveId, _ := ctx.Params().GetInt("serveId")
	path := ctx.Params().Get("path")
	endpointInterfaceId := ctx.URLParamIntDefault("id", 0)
	code := ctx.URLParamDefault("code", "")

	logUtils.Infof("%s %d/%s", method, serveId, path)

	req := service.MockRequest{
		ServeId:             serveId,
		EndpointMethod:      consts.HttpMethod(method),
		EndpointPath:        "/" + path,
		EndpointInterfaceId: uint(endpointInterfaceId),
		Code:                code,
	}

	resp, err := c.MockService.ByRequest(tenantId, &req, ctx)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	c.WriteRespByContentType(resp, ctx)

}

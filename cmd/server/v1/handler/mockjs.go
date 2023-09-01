package handler

import (
	domain "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/service"
	"github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/kataras/iris/v12"
)

type MockJsCtrl struct {
	MockJsService *service.MockJsService `inject:""`
	BaseCtrl
}

func (c *MockJsCtrl) ListExpressions(ctx iris.Context) {
	data, err := c.MockJsService.ListExpressions()
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data})
}

func (c *MockJsCtrl) EvaluateExpression(ctx iris.Context) {
	expression := ctx.URLParam("expression")

	req := domain.MockJsExpression{
		Expression: expression,
	}

	data, err := c.MockJsService.EvaluateExpression(req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data})
}

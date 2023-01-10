package handler

import (
	"github.com/aaronchen2k/deeptest/cmd/agent/v1/domain"
	service "github.com/aaronchen2k/deeptest/internal/agent/service"
	"github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/kataras/iris/v12"
)

type InvocationCtrl struct {
	InvocationService *service.InvocationService `inject:""`
}

// InvokeInterface
func (c *InvocationCtrl) InvokeInterface(ctx iris.Context) {
	req := domain.InvocationReq{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	resp, err := c.InvocationService.Invoke(req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: resp})
}

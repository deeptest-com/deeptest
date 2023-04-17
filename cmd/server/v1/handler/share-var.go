package handler

import (
	"github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/service"
	"github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/kataras/iris/v12"
)

type ShareVarCtrl struct {
	ShareVarService *service.ShareVarService `inject:""`
	BaseCtrl
}

// List
func (c *ShareVarCtrl) List(ctx iris.Context) {
	req := domain.DebugCall{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	data := c.ShareVarService.List(req.InterfaceId, req.EndpointId, req.ProcessorId, req.UsedBy)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data})
}

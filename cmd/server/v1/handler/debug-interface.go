package handler

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	service "github.com/aaronchen2k/deeptest/internal/server/modules/service"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/kataras/iris/v12"
)

type DebugInterfaceCtrl struct {
	DebugInterfaceService *service.DebugInterfaceService `inject:""`
	ExtractorService      *service.ExtractorService      `inject:""`
	CheckpointService     *service.CheckpointService     `inject:""`
	BaseCtrl
}

// Load
func (c *DebugInterfaceCtrl) Load(ctx iris.Context) {
	req := domain.DebugReq{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	data, err := c.DebugInterfaceService.Load(req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
}

// Load
func (c *DebugInterfaceCtrl) LoadForExec(ctx iris.Context) {
	req := domain.DebugReq{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	data, err := c.DebugInterfaceService.LoadForExec(req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
}

// Save
func (c *DebugInterfaceCtrl) Save(ctx iris.Context) {
	req := domain.DebugData{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	po, err := c.DebugInterfaceService.Save(req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	loadReq := domain.DebugReq{
		EndpointInterfaceId: po.EndpointInterfaceId,
	}

	data, err := c.DebugInterfaceService.Load(loadReq)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data})
}

package handler

import (
	"github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	service2 "github.com/aaronchen2k/deeptest/internal/server/modules/service"
	"github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/kataras/iris/v12"
)

type InvocationCtrl struct {
	InvocationService *service2.InvocationService `inject:""`
	InterfaceService  *service2.InterfaceService  `inject:""`
	ExtractorService  *service2.ExtractorService  `inject:""`
	CheckpointService *service2.CheckpointService `inject:""`
	BaseCtrl
}

// LoadInterfaceExecData
func (c *InvocationCtrl) LoadInterfaceExecData(ctx iris.Context) {
	req := domain.InvocationRequest{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	data, err := c.InvocationService.LoadInterfaceExecData(req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
}

// SubmitInterfaceInvokeResult
func (c *InvocationCtrl) SubmitInterfaceInvokeResult(ctx iris.Context) {
	req := domain.SubmitInvocationResultRequest{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	err = c.InvocationService.SubmitInterfaceInvokeResult(req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code})
}

// List
func (c *InvocationCtrl) List(ctx iris.Context) {
	interfaceId, err := ctx.URLParamInt("interfaceId")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Data: nil, Msg: err.Error()})
		return
	}

	data, err := c.InvocationService.ListByInterface(interfaceId)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data})
}

func (c *InvocationCtrl) GetLastResp(ctx iris.Context) {
	id, err := ctx.URLParamInt("id")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	resp, err := c.InvocationService.GetLastResp(id)

	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: _domain.SystemErr.Msg})
		return
	}
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: resp})
}

// GetAsInterface 详情
func (c *InvocationCtrl) GetAsInterface(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Data: nil, Msg: _domain.ParamErr.Msg})
		return
	}

	invocation, err := c.InvocationService.GetAsInterface(id)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: _domain.SystemErr.Msg})
		return
	}
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: invocation})
}

// Created by interface test api

// Delete 删除
func (c *InvocationCtrl) Delete(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	err = c.InvocationService.Delete(uint(id))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

package controller

import (
	serverDomain "github.com/aaronchen2k/deeptest/internal/server/modules/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/service"
	"github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/kataras/iris/v12"
)

type InvocationCtrl struct {
	InvocationService *service.InvocationService `inject:""`
	InterfaceService  *service.InterfaceService  `inject:""`
	ExtractorService  *service.ExtractorService  `inject:""`
	CheckpointService *service.CheckpointService `inject:""`
	BaseCtrl
}

// Invoke
func (c *InvocationCtrl) Invoke(ctx iris.Context) {
	projectId, err := ctx.URLParamInt("currProjectId")
	if projectId == 0 {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: "projectId"})
		return
	}

	req := serverDomain.InvocationRequest{}
	err = ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	err = c.InterfaceService.UpdateByInvocation(req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: err.Error()})
		return
	}

	reqNew, err := c.InterfaceService.ReplaceEnvironmentAndExtractorVariables(req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: err.Error()})
		return
	}

	resp, err := c.InterfaceService.Test(reqNew)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: err.Error()})
		return
	}

	interf, _ := c.InterfaceService.Get(req.Id)
	c.ExtractorService.ExtractInterface(interf, resp, nil)
	c.CheckpointService.CheckInterface(interf, resp, nil)

	_, err = c.InvocationService.Create(req, resp, projectId)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: resp})
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

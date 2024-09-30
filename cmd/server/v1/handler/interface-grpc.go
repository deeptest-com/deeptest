package handler

import (
	serverDomain "github.com/deeptest-com/deeptest/cmd/server/v1/domain"
	"github.com/deeptest-com/deeptest/internal/pkg/domain"
	service "github.com/deeptest-com/deeptest/internal/server/modules/service"
	_domain "github.com/deeptest-com/deeptest/pkg/domain"
	"github.com/kataras/iris/v12"
)

type GrpcInterfaceCtrl struct {
	GrpcInterfaceService *service.GrpcInterfaceService `inject:""`

	BaseCtrl
}

func (c *GrpcInterfaceCtrl) GetDebugData(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	diagnoseInterfaceId, _ := ctx.URLParamInt("diagnoseInterfaceId")

	data, err := c.GrpcInterfaceService.GetDebugData(tenantId, diagnoseInterfaceId)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
}

func (c *GrpcInterfaceCtrl) SaveDebugData(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)

	req := domain.GrpcDebugData{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	err = c.GrpcInterfaceService.SaveDebugData(req, tenantId)

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

func (c *GrpcInterfaceCtrl) ParseProto(ctx iris.Context) {
	req := serverDomain.GrpcReq{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	data, err := c.GrpcInterfaceService.ParseProto(req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data})
}

func (c *GrpcInterfaceCtrl) DescribeFunc(ctx iris.Context) {
	req := serverDomain.GrpcReq{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	data, err := c.GrpcInterfaceService.DescribeFunction(req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data})
}

func (c *GrpcInterfaceCtrl) InvokeFunc(ctx iris.Context) {
	req := serverDomain.GrpcReq{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	results, err := c.GrpcInterfaceService.InvokeFunc(req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: results})
}

func (c *GrpcInterfaceCtrl) ListConn(ctx iris.Context) {
	req := serverDomain.GrpcReq{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	data, err := c.GrpcInterfaceService.ListActiveConn(req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Data: data})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code})
}

func (c *GrpcInterfaceCtrl) DeleteHandle(ctx iris.Context) {
	req := serverDomain.GrpcReq{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	err = c.GrpcInterfaceService.DeleteHandle(req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code})
}

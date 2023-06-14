package handler

import (
	"github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/service"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/kataras/iris/v12"
)

type EndpointInterfaceCtrl struct {
	EndpointInterfaceService *service.EndpointInterfaceService `inject:""`
}

func (c *EndpointInterfaceCtrl) ListForSelection(ctx iris.Context) {
	var req serverDomain.EndpointInterfaceReqPaginate
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	res, _ := c.EndpointInterfaceService.Paginate(req)
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: res})

	return
}

func (c *EndpointInterfaceCtrl) ImportEndpointData(ctx iris.Context) {
	var req serverDomain.ImportEndpointDataReq
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	projectId, err := ctx.URLParamInt("currProjectId")
	if projectId == 0 {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}
	req.ProjectId = uint(projectId)

	err = c.EndpointInterfaceService.ImportEndpointData(req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code})

	return
}

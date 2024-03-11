package handler

import (
	"github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/helper/openapi/convert"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/service"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/kataras/iris/v12"
	"github.com/snowlyg/multi"
)

type EndpointInterfaceCtrl struct {
	BaseCtrl
	EndpointInterfaceService *service.EndpointInterfaceService `inject:""`
	ThirdPartySyncService    *service.ThirdPartySyncService    `inject:""`
}

// ListForSelection
// @Tags	设计器/接口
// @summary	接口列表
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization					header	string										true	"Authentication header"
// @Param 	currProjectId					query	int											true	"当前项目ID"
// @Param 	EndpointInterfaceReqPaginate 	body 	serverDomain.EndpointInterfaceReqPaginate 	true 	"获取接口列表的请求参数"
// @success	200	{object}	_domain.Response{data=object{result=[]model.EndpointInterface}}
// @Router	/api/v1/endpoints/interfaces/listForSelection	[post]
func (c *EndpointInterfaceCtrl) ListForSelection(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	var req serverDomain.EndpointInterfaceReqPaginate
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	res, _ := c.EndpointInterfaceService.Paginate(tenantId, req)
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: res})

	return
}

// ImportEndpointData
// @Tags	设计器/接口
// @summary	导入接口数据
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization			header	string								true	"Authentication header"
// @Param 	currProjectId			query	int									true	"当前项目ID"
// @Param 	ImportEndpointDataReq 	body 	serverDomain.ImportEndpointDataReq 	true 	"导入接口数据的请求参数"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/endpoints/interfaces/importEndpointData	[post]
func (c *EndpointInterfaceCtrl) ImportEndpointData(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
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

	userId := multi.GetUserId(ctx)
	req.UserId = userId

	if req.DriverType == convert.LZOS {
		err = c.ThirdPartySyncService.ImportThirdPartyFunctions(tenantId, req)
	} else {
		err = c.EndpointInterfaceService.ImportEndpointData(tenantId, req)
	}

	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code})

	return
}

func (c *EndpointInterfaceCtrl) GenerateFromResponse(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	var req serverDomain.GenerateFromResponseReq
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}
	var data model.EndpointInterfaceResponseBody
	data, err = c.EndpointInterfaceService.GenerateFromResponse(tenantId, req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data})

	return
}

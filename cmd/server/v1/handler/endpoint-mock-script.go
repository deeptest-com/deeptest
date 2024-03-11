package handler

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/service"
	"github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/kataras/iris/v12"
)

type EndpointMockScriptCtrl struct {
	EndpointMockScriptService *service.EndpointMockScriptService `inject:""`
	BaseCtrl
}

// Get
// @Tags	Mock脚本管理
// @summary	脚本详情
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization	header	string	true	"Authentication header"
// @Param 	endpointId	path	int		true	"Endpoint ID"
// @success	200	{object}	_domain.Response{data=model.EndpointMockScript}
// @Router	/api/v1/mockScripts/{id}	[get]
func (c *EndpointMockScriptCtrl) Get(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	endpointId, err := ctx.Params().GetInt("endpointId")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	ret, err := c.EndpointMockScriptService.Get(tenantId, uint(endpointId))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: ret, Msg: _domain.NoErr.Msg})
}

// Update
// @Tags	Mock脚本管理
// @summary	更新脚本
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization	header	string					true	"Authentication header"
// @Param 	EndpointMockScript 		body 	model.EndpointMockScript true 	"update Mock Script Object"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/mockScripts	[put]
func (c *EndpointMockScriptCtrl) Update(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	var req model.EndpointMockScript
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	err = c.EndpointMockScriptService.Update(tenantId, req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

func (c *EndpointMockScriptCtrl) Disable(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	endpointId, err := ctx.Params().GetInt("endpointId")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	err = c.EndpointMockScriptService.Disable(tenantId, uint(endpointId))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

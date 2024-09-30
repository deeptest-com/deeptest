package handler

import (
	"github.com/deeptest-com/deeptest/internal/pkg/domain"
	"github.com/deeptest-com/deeptest/internal/server/modules/service"
	"github.com/deeptest-com/deeptest/pkg/domain"
	"github.com/kataras/iris/v12"
)

type ShareVarCtrl struct {
	ShareVarService *service.ShareVarService `inject:""`
	BaseCtrl
}

// List
// @Tags	共享变量
// @summary	列出变量列表
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string				true	"Authentication header"
// @Param 	currProjectId	query	int					true	"当前项目ID"
// @Param 	DebugInfo		body	domain.DebugInfo	true	"列出变量列表的请求参数"
// @success	200	{object}	_domain.Response{data=[]domain.GlobalVar}
// @Router	/api/v1/shareVars/list	[post]
func (c *ShareVarCtrl) List(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	req := domain.DebugInfo{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	data := c.ShareVarService.List(tenantId, req.DebugInterfaceId,
		req.EndpointInterfaceId, req.DiagnoseInterfaceId, req.CaseInterfaceId, req.ScenarioProcessorId,
		req.UsedBy)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data})
}

// Delete 删除
// @Tags	共享变量
// @summary	删除共享变量
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	id				path	int		true	"变量ID"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/shareVars/{id}	[delete]
func (c *ShareVarCtrl) Delete(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	err = c.ShareVarService.Delete(tenantId, id)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

// Clear 清除
// @Tags	共享变量
// @summary	清空共享变量
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string				true	"Authentication header"
// @Param 	currProjectId	query	int					true	"当前项目ID"
// @Param 	DebugInfo		body	domain.DebugInfo	true	"清空共享变量的请求参数"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/shareVars/clear	[post]
func (c *ShareVarCtrl) Clear(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	req := domain.DebugInfo{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	err = c.ShareVarService.Clear(tenantId, req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

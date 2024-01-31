package handler

import (
	serverDomain "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	service "github.com/aaronchen2k/deeptest/internal/server/modules/service"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/kataras/iris/v12"
	"github.com/snowlyg/multi"
)

type DebugInterfaceCtrl struct {
	EndpointCaseService   *service.EndpointCaseService   `inject:""`
	DebugInterfaceService *service.DebugInterfaceService `inject:""`
	ExtractorService      *service.ExtractorService      `inject:""`
	CheckpointService     *service.CheckpointService     `inject:""`

	BaseCtrl
}

// Load
// @Tags	接口调试
// @summary	获取调试接口请求
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization	header	string			true	"Authentication header"
// @Param 	currProjectId	query	int				true	"当前项目ID"
// @Param 	DebugReq 		body 	domain.DebugInfo true 	"获取调试接口请求的请求体"
// @success	200	{object}	_domain.Response{data=domain.DebugData}
// @Router	/api/v1/debugs/interface/load	[post]
func (c *DebugInterfaceCtrl) Load(ctx iris.Context) {
	req := domain.DebugInfo{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	req.UserId = multi.GetUserId(ctx)
	data, err := c.DebugInterfaceService.Load(req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
}

// LoadForExec
// @Tags	接口调试
// @summary	获取调试接口用于执行
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization	header	string			true	"Authentication header"
// @Param 	currProjectId	query	int				true	"当前项目ID"
// @Param 	DebugReq 		body 	domain.DebugInfo true 	"获取调试接口用于执行的请求体"
// @success	200	{object}	_domain.Response{data=agentExec.InterfaceExecObj}
// @Router	/api/v1/debugs/interface/loadForExec	[post]
func (c *DebugInterfaceCtrl) LoadForExec(ctx iris.Context) {
	req := domain.DebugInfo{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	req.UserId = multi.GetUserId(ctx)
	req.ProjectId, err = ctx.URLParamInt("currProjectId")

	data, err := c.DebugInterfaceService.LoadForExec(req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
}

// Save
// @Tags	接口调试
// @summary	保存调试接口
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization	header	string			true	"Authentication header"
// @Param 	currProjectId	query	int				true	"当前项目ID"
// @Param 	DebugReq 		body 	domain.DebugData true 	"保存调试接口的请求体"
// @success	200	{object}	_domain.Response{data=agentExec.InterfaceExecObj}
// @Router	/api/v1/debugs/interface/save	[post]
func (c *DebugInterfaceCtrl) Save(ctx iris.Context) {
	req := domain.DebugData{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	po, err := c.DebugInterfaceService.CreateOrUpdate(req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	loadReq := domain.DebugInfo{
		DebugInterfaceId:    po.ID,
		EndpointInterfaceId: po.EndpointInterfaceId,
		UsedBy:              req.UsedBy,
	}

	data, err := c.DebugInterfaceService.Load(loadReq)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data})
}

// SaveAsCase
func (c *DebugInterfaceCtrl) SaveAsCase(ctx iris.Context) {
	req := serverDomain.EndpointCaseSaveReq{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	req.CreateUserName = multi.GetUsername(ctx)
	req.CreateUserId = multi.GetUserId(ctx)

	c.EndpointCaseService.SaveFromDebugInterface(req)

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code})
}

// LoadCurl 导入cURL命令
// @Tags	快捷调试
// @summary	获取cURL命令
// @accept 	application/json
// @Produce application/json
// @Param	Authorization				header	string								true	"Authentication header"
// @Param 	currProjectId				query	int									true	"当前项目ID"
// @Param 	DiagnoseCurlLoadReq		body	serverDomain.DiagnoseCurlLoadReq	    true	"导入cURL命令的请求体"
// @success	200	{object}	_domain.Response{data=string}
// @Router	/api/v1/debugs/interface/loadCurl	[post]
func (c *DebugInterfaceCtrl) LoadCurl(ctx iris.Context) {
	req := serverDomain.DiagnoseCurlLoadReq{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	req.ProjectId, err = ctx.URLParamInt("currProjectId")
	req.UserId = multi.GetUserId(ctx)

	content, err := c.DebugInterfaceService.LoadCurl(req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: content})
}

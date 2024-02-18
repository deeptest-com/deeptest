package handler

import (
	"github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/service"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/kataras/iris/v12"
	"github.com/snowlyg/multi"
)

type EndpointCaseCtrl struct {
	BaseCtrl
	EndpointCaseService   *service.EndpointCaseService   `inject:""`
	DebugInterfaceService *service.DebugInterfaceService `inject:""`
}

// Paginate
// @Tags	设计器/接口用例
// @summary	用例列表
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	endpointId		query	int		true	"endpointId"
// @success	200	{object}	_domain.Response{data=[]model.EndpointCase}
// @Router	/api/v1/endpoints/cases/list	[get]
func (c *EndpointCaseCtrl) Paginate(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	var req serverDomain.EndpointCaseReqPaginate
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	req.ConvertParams()

	data, err := c.EndpointCaseService.Paginate(tenantId, req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
}

// Get
// @Tags	设计器/接口用例
// @summary	用例详情
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	id		path	int		true	"用例ID"
// @success	200	{object}	_domain.Response{data=model.EndpointCase}
// @Router	/api/v1/endpoints/cases/{id}	[get]
func (c *EndpointCaseCtrl) Get(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	id, _ := ctx.Params().GetInt("id")

	data, err := c.EndpointCaseService.Get(tenantId, uint(id))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
}

// Create
// @Tags	设计器/接口用例
// @summary	保存用例
// @accept 	application/json
// @Produce application/json
// @Param	Authorization			header	string								true	"Authentication header"
// @Param 	currProjectId			query	int									true	"当前项目ID"
// @Param 	id						path	int									true	"用例ID"
// @Param 	EndpointCaseSaveReq		body	serverDomain.EndpointCaseSaveReq	true	"保存用例的请求参数"
// @success	200	{object}	_domain.Response{data=[]model.EndpointCase}
// @Router	/api/v1/endpoints/cases/{id}	[post]
func (c *EndpointCaseCtrl) Create(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	req := serverDomain.EndpointCaseSaveReq{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	req.CreateUserName = multi.GetUsername(ctx)
	req.CreateUserId = multi.GetUserId(ctx)

	po, err := c.EndpointCaseService.Create(tenantId, req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	data, err := c.EndpointCaseService.List(tenantId, po.EndpointId)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data})
}

// Copy
// @Tags	设计器/接口用例
// @summary	用例复制
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	id		query	int		true	"用例ID"
// @success	200	{object}	_domain.Response{data=model.EndpointCase}
// @Router	/api/v1/endpoints/cases/copy	[post]
func (c *EndpointCaseCtrl) Copy(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	id, err := ctx.URLParamInt("id")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	userId := multi.GetUserId(ctx)
	userName := multi.GetUsername(ctx)

	po, err := c.EndpointCaseService.Copy(tenantId, id, "", 0, 0, userId, userName, "false")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: po})
}

// UpdateName
// @Tags	设计器/接口用例
// @summary	保存用例名称
// @accept 	application/json
// @Produce application/json
// @Param	Authorization			header	string								true	"Authentication header"
// @Param 	currProjectId			query	int									true	"当前项目ID"
// @Param 	id						path	int									true	"用例ID"
// @Param 	EndpointCaseSaveReq		body	serverDomain.EndpointCaseSaveReq	true	"保存用例名称的请求参数"
// @success	200	{object}	_domain.Response{data=[]model.EndpointCase}
// @Router	/api/v1/endpoints/cases/{id}	[put]
func (c *EndpointCaseCtrl) UpdateName(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	req := serverDomain.EndpointCaseSaveReq{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	err = c.EndpointCaseService.UpdateName(tenantId, req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	data, err := c.EndpointCaseService.List(tenantId, req.EndpointId)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data})
}

// SaveDebugData
// @Tags	设计器/接口用例
// @summary	保存调试数据
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string				true	"Authentication header"
// @Param 	currProjectId	query	int					true	"当前项目ID"
// @Param 	DebugData		body	domain.DebugData	true	"保存调试数据的请求参数"
// @success	200	{object}	_domain.Response{data=[]model.EndpointCase}
// @Router	/api/v1/endpoints/cases/saveDebugData	[post]
func (c *EndpointCaseCtrl) SaveDebugData(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	req := domain.DebugData{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	_, err = c.DebugInterfaceService.CreateOrUpdate(tenantId, req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	data, err := c.EndpointCaseService.List(tenantId, req.EndpointInterfaceId)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data})
}

// Remove
// @Tags	设计器/接口用例
// @summary	删除用例
// @accept 	application/json
// @Produce application/json
// @Param	Authorization			header	string								true	"Authentication header"
// @Param 	currProjectId			query	int									true	"当前项目ID"
// @Param 	id						path	int									true	"用例ID"
// @Param 	EndpointCaseSaveReq		body	serverDomain.EndpointCaseSaveReq	true	"删除用例的请求参数"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/endpoints/cases/{id}	[delete]
func (c *EndpointCaseCtrl) Remove(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	id, _ := ctx.Params().GetInt("id")

	req := serverDomain.EndpointCaseSaveReq{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	err = c.EndpointCaseService.Remove(tenantId, uint(id))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code})
}

// LoadTree
// @Tags	设计器/接口用例
// @summary	分类接口用例树
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	serveId			query	int		true	"服务ID"
// @success	200	{object}	_domain.Response{data=[]serverDomain.EndpointCaseTree}
// @Router	/api/v1/endpoints/cases/loadTree	[get]
func (c *EndpointCaseCtrl) LoadTree(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	projectId, _ := ctx.URLParamInt("currProjectId")
	var serveIds consts.Integers
	ctx.ReadJSON(&serveIds)

	data, err := c.EndpointCaseService.LoadTree(tenantId, uint(projectId), serveIds)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
}

// ListForBenchmark
// @Tags	设计器/接口用例
// @summary	自动生成用例-选择已有用例-用例列表
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	endpointId		query	int		true	"endpointId"
// @success	200	{object}	_domain.Response{data=[]serverDomain.EndpointCaseTree}
// @Router	/api/v1/endpoints/cases/listForBenchmark	[get]
func (c *EndpointCaseCtrl) ListForBenchmark(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	endpointId, err := ctx.URLParamInt("endpointId")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	data, err := c.EndpointCaseService.ListByCaseType(tenantId, uint(endpointId), consts.CaseDefault)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
}

package handler

import (
	"github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/service"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/kataras/iris/v12"
	"github.com/snowlyg/multi"
)

type EndpointCaseCtrl struct {
	EndpointCaseService *service.EndpointCaseService `inject:""`
}

// List
// @Tags	设计器/接口用例
// @summary	用例列表
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	endpointId		query	int		true	"endpointId"
// @success	200	{object}	_domain.Response{data=[]model.EndpointCase}
// @Router	/api/v1/endpoints/cases/list	[get]
func (c *EndpointCaseCtrl) List(ctx iris.Context) {
	endpointId, _ := ctx.URLParamInt("endpointId")

	data, err := c.EndpointCaseService.List(uint(endpointId))
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
	id, _ := ctx.Params().GetInt("id")

	data, err := c.EndpointCaseService.Get(id)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
}

// Save
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
func (c *EndpointCaseCtrl) Save(ctx iris.Context) {
	req := serverDomain.EndpointCaseSaveReq{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	req.CreateUserName = multi.GetUsername(ctx)
	req.CreateUserId = multi.GetUserId(ctx)

	po, err := c.EndpointCaseService.Save(req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	data, err := c.EndpointCaseService.List(po.EndpointId)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data})
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
	req := serverDomain.EndpointCaseSaveReq{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	err = c.EndpointCaseService.UpdateName(req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	data, err := c.EndpointCaseService.List(req.EndpointId)
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
	req := domain.DebugData{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	_, err = c.EndpointCaseService.SaveDebugData(req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	data, err := c.EndpointCaseService.List(req.EndpointInterfaceId)
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
	id, _ := ctx.Params().GetInt("id")

	req := serverDomain.EndpointCaseSaveReq{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	err = c.EndpointCaseService.Remove(uint(id))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code})
}

package handler

import (
	"github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/cron"
	"github.com/aaronchen2k/deeptest/internal/server/modules/service"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/kataras/iris/v12"
	"github.com/snowlyg/multi"
)

type ServeCtrl struct {
	Cron                     *cron.ServerCron                  `inject:""`
	ServeService             *service.ServeService             `inject:""`
	EndpointInterfaceService *service.EndpointInterfaceService `inject:""`
}

// ListByProject 项目服务列表
// @Tags	服务管理
// @summary	获取项目下的服务
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @success	200	{object}	_domain.Response{data=object{serves=[]model.Serve,currServe=model.Serve}}
// @Router	/api/v1/serves/listByProject	[get]
func (c *ServeCtrl) ListByProject(ctx iris.Context) {
	userId := multi.GetUserId(ctx)
	projectId, err := ctx.URLParamInt("currProjectId")

	if projectId == 0 {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	serves, currServe, err := c.ServeService.ListByProject(projectId, userId)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ret := iris.Map{"serves": serves, "currServe": currServe}
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: ret})

	return
}

// Index 服务列表
// @Tags	服务管理
// @summary	服务列表
// @accept 	application/json
// @Produce application/json
// @Param	Authorization		header	string							true	"Authentication header"
// @Param 	currProjectId		query	int								true	"当前项目ID"
// @Param 	ServeReqPaginate	body	serverDomain.ServeReqPaginate	true	"服务列表的请求参数"
// @success	200	{object}	_domain.Response{data=_domain.PageData{result=[]model.Serve}}
// @Router	/api/v1/serves/index	[post]
func (c *ServeCtrl) Index(ctx iris.Context) {
	var req serverDomain.ServeReqPaginate
	if err := ctx.ReadJSON(&req); err == nil {
		res, _ := c.ServeService.Paginate(req)
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: res})
	} else {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
	}
	return
}

// Save 保存服务
// @Tags	服务管理
// @summary	保存服务
// @accept 	application/json
// @Produce application/json
// @Param	Authorization		header	string							true	"Authentication header"
// @Param 	currProjectId		query	int								true	"当前项目ID"
// @Param 	ServeReq	body	serverDomain.ServeReq	true	"保存服务的请求参数"
// @success	200	{object}	_domain.Response{data=int}
// @Router	/api/v1/serves/save	[post]
func (c *ServeCtrl) Save(ctx iris.Context) {
	var req serverDomain.ServeReq
	if err := ctx.ReadJSON(&req); err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	req.CreateUser = multi.GetUsername(ctx)
	res, err := c.ServeService.Save(req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: res})
	return
}

// Detail 服务详情
// @Tags	服务管理
// @summary	服务详情
// @accept 	application/json
// @Produce application/json
// @Param	Authorization		header	string							true	"Authentication header"
// @Param 	currProjectId		query	int								true	"当前项目ID"
// @Param 	id		query	int								true	"服务ID"
// @success	200	{object}	_domain.Response{data=model.Serve}
// @Router	/api/v1/serves/detail	[get]
func (c *ServeCtrl) Detail(ctx iris.Context) {
	id := ctx.URLParamUint64("id")
	if id != 0 {
		res := c.ServeService.GetById(uint(id))
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: res})
	} else {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
	}
}

// Copy 克隆服务
// @Tags	服务管理
// @summary	复制服务
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	id				query	int		true	"服务ID"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/serves/copy	[get]
func (c *ServeCtrl) Copy(ctx iris.Context) {
	id := ctx.URLParamUint64("id")
	if id != 0 {
		res := c.ServeService.Copy(uint(id))
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: res})
	} else {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
	}
}

// Delete 删除服务
// @Tags	服务管理
// @summary	删除服务
// @accept 	application/json
// @Produce application/json
// @Param	Authorization		header	string							true	"Authentication header"
// @Param 	currProjectId		query	int								true	"当前项目ID"
// @Param 	id		query	int								true	"服务ID"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/serves/delete	[delete]
func (c *ServeCtrl) Delete(ctx iris.Context) {
	id := ctx.URLParamUint64("id")
	err := c.ServeService.DeleteById(uint(id))
	if err == nil {
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
	} else {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
	}
}

// Expire 禁用服务
// @Tags	服务管理
// @summary	禁用服务
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	id				query	int		true	"服务ID"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/serves/expire	[put]
func (c *ServeCtrl) Expire(ctx iris.Context) {
	id := ctx.URLParamUint64("id")
	err := c.ServeService.DisableById(uint(id))
	if err == nil {
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
	} else {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
	}
}

// SaveVersion 保存版本
// @Tags	服务管理/版本
// @summary	保存版本
// @accept 	application/json
// @Produce application/json
// @Param	Authorization			header	string							true	"Authentication header"
// @Param 	currProjectId			query	int								true	"当前项目ID"
// @Param 	ServeVersionReq			body	serverDomain.ServeVersionReq	true	"保存服务版本的请求参数"
// @success	200	{object}	_domain.Response{data=int}
// @Router	/api/v1/serves/version/save	[post]
func (c *ServeCtrl) SaveVersion(ctx iris.Context) {
	var req serverDomain.ServeVersionReq
	if err := ctx.ReadJSON(&req); err == nil {
		if res, err := c.ServeService.SaveVersion(req); err != nil {
			ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		} else {
			ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: res})
		}
	} else {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
	}
	return
}

// ListVersion 获取版本列表
// @Tags	服务管理/版本
// @summary	版本列表
// @accept 	application/json
// @Produce application/json
// @Param	Authorization			header	string								true	"Authentication header"
// @Param 	currProjectId			query	int									true	"当前项目ID"
// @Param 	ServeVersionPaginate	body	serverDomain.ServeVersionPaginate	true	"服务版本列表的请求参数"
// @success	200	{object}	_domain.Response{data=_domain.PageData{result=[]model.ServeVersion}}
// @Router	/api/v1/serves/version/list	[post]
func (c *ServeCtrl) ListVersion(ctx iris.Context) {
	var req serverDomain.ServeVersionPaginate
	if err := ctx.ReadJSON(&req); err == nil {
		res, _ := c.ServeService.PaginateVersion(req)
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg, Data: res})
	} else {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
	}
}

// DeleteVersion
// @Tags	服务管理/版本
// @summary	删除版本
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	id				query	int		true	"服务版本ID"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/serves/version/delete	[delete]
func (c *ServeCtrl) DeleteVersion(ctx iris.Context) {
	id := ctx.URLParamUint64("id")
	err := c.ServeService.DeleteVersionById(uint(id))
	if err == nil {
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
	} else {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
	}
}

// ExpireVersion
// @Tags	服务管理/版本
// @summary	禁用版本
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	id				query	int		true	"服务版本ID"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/serves/version/expire	[put]
func (c *ServeCtrl) ExpireVersion(ctx iris.Context) {
	id := ctx.URLParamUint64("id")
	err := c.ServeService.DisableVersionById(uint(id))
	if err == nil {
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
	} else {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
	}
}

// SaveSchema
// @Tags	服务管理/schema
// @summary	保存Schema
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string						true	"Authentication header"
// @Param 	currProjectId	query	int							true	"当前项目ID"
// @Param 	ServeSchemaReq	body	serverDomain.ServeSchemaReq	true	"保存Schema的请求参数"
// @success	200	{object}	_domain.Response{data=int}
// @Router	/api/v1/serves/schema/save [post]
func (c *ServeCtrl) SaveSchema(ctx iris.Context) {
	var req serverDomain.ServeSchemaReq
	if err := ctx.ReadJSON(&req); err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}
	projectId, _ := ctx.URLParamInt("currProjectId")
	req.ProjectId = uint(projectId)

	res, err := c.ServeService.SaveSchema(req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg, Data: res})
}

// SaveSecurity
// @Tags	服务管理/授权
// @summary	保存授权
// @accept 	application/json
// @Produce application/json
// @Param	Authorization		header	string							true	"Authentication header"
// @Param 	currProjectId		query	int								true	"当前项目ID"
// @Param 	ServeSecurityReq	body	serverDomain.ServeSecurityReq	true	"保存授权的请求参数"
// @success	200	{object}	_domain.Response{data=int}
// @Router	/api/v1/serves/security/save [post]
func (c *ServeCtrl) SaveSecurity(ctx iris.Context) {
	var req serverDomain.ServeSecurityReq
	if err := ctx.ReadJSON(&req); err == nil {
		res, _ := c.ServeService.SaveSecurity(req)
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg, Data: res})
	} else {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
	}
}

// ListSchema Schema列表
// @Tags	服务管理/schema
// @summary	Schema列表
// @accept 	application/json
// @Produce application/json
// @Param	Authorization		header	string								true	"Authentication header"
// @Param 	currProjectId		query	int									true	"当前项目ID"
// @Param 	ServeSchemaPaginate	body	serverDomain.ServeSchemaPaginate	true	"Schema列表的请求参数"
// @success	200	{object}	_domain.Response{data=_domain.PageData{result=[]model.ComponentSchema}}
// @Router	/api/v1/serves/schema/list	[post]
func (c *ServeCtrl) ListSchema(ctx iris.Context) {
	var req serverDomain.ServeSchemaPaginate
	if err := ctx.ReadJSON(&req); err == nil {
		res, _ := c.ServeService.PaginateSchema(req)
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg, Data: res})
	} else {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
	}
}

// GetSchemaByRef
// @Tags	服务管理/schema
// @summary	获取Schema
// @accept 	application/json
// @Produce application/json
// @Param	Authorization		header	string							true	"Authentication header"
// @Param 	currProjectId		query	int								true	"当前项目ID"
// @Param 	ServeSchemaRefReq	body	serverDomain.ServeSchemaRefReq	true	"获取Schema的请求参数"
// @success	200	{object}	_domain.Response{data=model.ComponentSchema}
// @Router	/api/v1/serves/schema/detail	[get]
func (c *ServeCtrl) GetSchemaByRef(ctx iris.Context) {
	id := ctx.URLParamUint64("id")
	if id == 0 {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
		return
	}

	res, _ := c.ServeService.GetSchema(uint(id))
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg, Data: res})

}

// DeleteSchema
// @Tags	服务管理/schema
// @summary	删除Schema
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	id				query	int		true	"schema ID"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/serves/schema/delete	[delete]
func (c *ServeCtrl) DeleteSchema(ctx iris.Context) {
	id := ctx.URLParamUint64("id")
	err := c.ServeService.DeleteSchemaById(uint(id))
	if err == nil {
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
	} else {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
	}
}

// ListServer
// @Tags	服务管理
// @summary	环境列表(不分页)
// @accept 	application/json
// @Produce application/json
// @Param	Authorization		header	string						true	"Authentication header"
// @Param 	currProjectId		query	int							true	"当前项目ID"
// @Param 	ServeServer			body	serverDomain.ServeServer	true	"环境列表的请求参数"
// @success	200	{object}	_domain.Response{data=object{servers=[]model.ServeServer, currServer=model.ServeServer}}
// @Router	/api/v1/serves/server/list	[post]
func (c *ServeCtrl) ListServer(ctx iris.Context) {
	var req serverDomain.ServeServer

	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}
	userId := multi.GetUserId(ctx)
	projectId, err := ctx.URLParamInt("currProjectId")
	if projectId == 0 {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	servers, currServer, err := c.ServeService.ListServer(req, uint(projectId), userId)

	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ret := iris.Map{"servers": servers, "currServer": currServer}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg, Data: ret})
}

// ChangeServer
// @Tags	服务管理
// @summary	切换环境
// @accept 	application/json
// @Produce application/json
// @Param	Authorization		header	string						true	"Authentication header"
// @Param 	currProjectId		query	int							true	"当前项目ID"
// @Param 	ServeServer			body	serverDomain.ServeServer	true	"服务列表的请求参数"
// @success	200	{object}	_domain.Response{data=model.ServeServer}
// @Router	/api/v1/serves/server/changeServer	[post]
func (c *ServeCtrl) ChangeServer(ctx iris.Context) {
	var req serverDomain.ServeServer

	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}
	userId := multi.GetUserId(ctx)
	projectId, err := ctx.URLParamInt("currProjectId")
	if projectId == 0 {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	currServer, err := c.ServeService.ChangeServer(uint(projectId), userId, req.ServeId, req.ServerId)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg, Data: currServer})
}

func (c *ServeCtrl) SaveServer(ctx iris.Context) {
	var req serverDomain.ServeServer
	if err := ctx.ReadJSON(&req); err == nil {
		res, _ := c.ServeService.SaveServer(req)
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg, Data: res})
	} else {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
	}
}

// ExampleToSchema
// @Tags	服务管理/schema
// @summary	example转schema
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string						true	"Authentication header"
// @Param 	currProjectId	query	int							true	"当前项目ID"
// @Param 	JsonContent		body	serverDomain.JsonContent	true	"example转schema的请求参数"
// @success	200	{object}	_domain.Response{data=openapi.Schema}
// @Router	/api/v1/serves/schema/example2schema [post]
func (c *ServeCtrl) ExampleToSchema(ctx iris.Context) {
	var req serverDomain.JsonContent
	if err := ctx.ReadJSON(&req); err == nil {
		res := c.ServeService.Example2Schema(req.Data)
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg, Data: res})
	} else {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
	}
}

// SchemaToExample
// @Tags	服务管理/schema
// @summary	Schema生成Example
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string						true	"Authentication header"
// @Param 	currProjectId	query	int							true	"当前项目ID"
// @Param 	JsonContent		body	serverDomain.JsonContent	true	"Schema生成Example的请求参数"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/serves/schema/schema2example [post]
func (c *ServeCtrl) SchemaToExample(ctx iris.Context) {
	var req serverDomain.JsonContent
	if err := ctx.ReadJSON(&req); err == nil {
		res := c.ServeService.Schema2Example(req.ServeId, req.Data)
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg, Data: res})
	} else {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
	}
}

// SchemaToYaml
// @Tags	服务管理/schema
// @summary	schema转yaml
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string						true	"Authentication header"
// @Param 	currProjectId	query	int							true	"当前项目ID"
// @Param 	JsonContent		body	serverDomain.JsonContent	true	"schema转yaml的请求参数"
// @success	200	{object}	_domain.Response{data=string}
// @Router	/api/v1/serves/schema/schema2yaml [post]
func (c *ServeCtrl) SchemaToYaml(ctx iris.Context) {
	var req serverDomain.JsonContent
	if err := ctx.ReadJSON(&req); err == nil {
		res := c.ServeService.Schema2Yaml(req.Data)
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg, Data: res})
	} else {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
	}
}

// CopySchema
// @Tags	服务管理/schema
// @summary	复制Schema
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	id				query	int		true	"Schema ID"
// @success	200	{object}	_domain.Response{data=int}
// @Router	/api/v1/serves/schema/copy [put]
func (c *ServeCtrl) CopySchema(ctx iris.Context) {
	id := ctx.URLParamUint64("id")
	if id != 0 {
		res, _ := c.ServeService.CopySchema(uint(id))
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: res.ID})
	} else {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
	}
}

// BindEndpoint
// @Tags	服务管理/版本
// @summary	关联接口
// @accept 	application/json
// @Produce application/json
// @Param	Authorization					header	string										true	"Authentication header"
// @Param 	currProjectId					query	int											true	"当前项目ID"
// @Param 	ServeVersionBindEndpointReq		body	serverDomain.ServeVersionBindEndpointReq	true	"服务版本关联接口的请求参数"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/serves/version/bindEndpoint	[post]
func (c *ServeCtrl) BindEndpoint(ctx iris.Context) {
	var req serverDomain.ServeVersionBindEndpointReq
	if err := ctx.ReadJSON(&req); err == nil {
		c.ServeService.BindEndpoint(req)
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
	} else {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
	}
}

// ChangeServe
// @Tags	服务管理
// @summary	切换用户当前服务
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string						true	"Authentication header"
// @Param 	currProjectId	query	int							true	"当前项目ID"
// @Param 	ChangeServeReq	body	serverDomain.ChangeServeReq	true	"切换用户当前服务的请求参数"
// @success	200	{object}	_domain.Response{data=model.Serve}
// @Router	/api/v1/serves/changeServe	[post]
func (c *ServeCtrl) ChangeServe(ctx iris.Context) {
	userId := multi.GetUserId(ctx)

	req := serverDomain.ChangeServeReq{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	currServe, err := c.ServeService.ChangeServe(req.Id, userId)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: currServe, Msg: _domain.NoErr.Msg})
}

// ListSecurity
// @Tags	服务管理/授权
// @summary	授权列表
// @accept 	application/json
// @Produce application/json
// @Param	Authorization			header	string								true	"Authentication header"
// @Param 	currProjectId			query	int									true	"当前项目ID"
// @Param 	ServeSecurityPaginate	body	serverDomain.ServeSecurityPaginate	true	"授权列表的请求参数"
// @success	200	{object}	_domain.Response{data=_domain.PageData{result=[]model.ComponentSchemaSecurity}}
// @Router	/api/v1/serves/security/list	[post]
func (c *ServeCtrl) ListSecurity(ctx iris.Context) {
	var req serverDomain.ServeSecurityPaginate
	if err := ctx.ReadJSON(&req); err == nil {
		res, _ := c.ServeService.PaginateSecurity(req)
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg, Data: res})
	} else {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
	}
}

// DeleteSecurity
// @Tags	服务管理/授权
// @summary	删除授权
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	id				query	int		true	"授权ID"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/serves/security/delete	[delete]
func (c *ServeCtrl) DeleteSecurity(ctx iris.Context) {
	id := ctx.URLParamUint64("id")
	err := c.ServeService.DeleteSecurityId(uint(id))
	if err == nil {
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
	} else {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
	}
}

func (c *ServeCtrl) AddServerForHistory(ctx iris.Context) {
	req := serverDomain.HistoryServeAddServesReq{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	err = c.ServeService.AddServerForHistory(req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code})
}

/*
// SwaggerSyncDetail
// @Tags	自动同步
// @summary	获取同步信息
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @success	200	{object}	_domain.Response{data=model.SwaggerSync}
// @Router	/api/v1/serves/swaggerSyncDetail	[get]
func (c *ServeCtrl) SwaggerSyncDetail(ctx iris.Context) {
	projectId := ctx.URLParamUint64("currProjectId")
	res, err := c.ServeService.SwaggerSyncDetail(uint(projectId))
	if err != nil {
		res.CategoryId = -1
		res.SyncType = consts.FullCover
		res.Cron = "23 * * * *"
	}
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg, Data: res})
}

func (c *ServeCtrl) InitSwaggerCron() {
	syncList, err := c.ServeService.SwaggerSyncList()
	if err != nil {
		return
	}
	for _, item := range syncList {
		c.ServeService.AddSwaggerCron(item)
	}

}
*/

package handler

import (
	"github.com/deeptest-com/deeptest/cmd/server/v1/domain"
	"github.com/deeptest-com/deeptest/internal/server/modules/model"
	"github.com/deeptest-com/deeptest/internal/server/modules/service"
	"github.com/deeptest-com/deeptest/pkg/domain"
	"github.com/kataras/iris/v12"
)

type EnvironmentCtrl struct {
	EnvironmentService *service.EnvironmentService `inject:""`
	BaseCtrl
}

// List
// @Tags	环境管理
// @summary	环境列表(List)
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @success	200	{object}	_domain.Response{data=[]model.Environment}
// @Router	/api/v1/environments	[get]
func (c *EnvironmentCtrl) List(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	projectId, err := ctx.URLParamInt("currProjectId")
	data, err := c.EnvironmentService.List(tenantId, projectId)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data})
}

// Get 详情
// @Tags	环境管理
// @summary	环境详情
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	id				path	int		true	"环境ID"
// @success	200	{object}	_domain.Response{data=model.Environment}
// @Router	/api/v1/environments{id}	[get]
func (c *EnvironmentCtrl) Get(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	projectId, err := ctx.URLParamInt("currProjectId")
	id, err := ctx.Params().GetInt("id")

	if id <= 0 && projectId <= 0 {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	env, err := c.EnvironmentService.Get(tenantId, uint(id), uint(projectId))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
		return
	}
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: env})
}

// Change 修改
// @Tags	环境管理
// @summary	修改环境
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	id 				query 	int 	true 	"环境id"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/environments/changeEnvironment	[post]
func (c *EnvironmentCtrl) Change(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	projectId, err := ctx.URLParamInt("currProjectId")
	id, err := ctx.URLParamInt("id")

	if projectId == 0 || id == 0 {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	err = c.EnvironmentService.Change(tenantId, id, projectId)
	if err != nil {
		ctx.JSON(_domain.Response{
			Code: _domain.SystemErr.Code,
		})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

// Create 添加
// @Tags	环境管理
// @summary	新建环境
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization	header	string						true	"Authentication header"
// @Param 	currProjectId	query	int							true	"当前项目ID"
// @Param 	Environment 	body 	model.Environment 			true 	"新建环境的请求参数"
// @success	200	{object}	_domain.Response{data=model.Environment}
// @Router	/api/v1/environments	[post]
func (c *EnvironmentCtrl) Create(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	projectId, err := ctx.URLParamInt("currProjectId")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	env := model.Environment{}
	err = ctx.ReadJSON(&env)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	err = c.EnvironmentService.Create(tenantId, &env, uint(projectId))
	if err != nil {
		ctx.JSON(_domain.Response{
			Code: _domain.SystemErr.Code,
		})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: env, Msg: _domain.NoErr.Msg})
}

// Update 更新
// @Tags	环境管理
// @summary	更新环境
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization	header	string						true	"Authentication header"
// @Param 	currProjectId	query	int							true	"当前项目ID"
// @Param 	Environment 	body 	model.Environment 			true 	"更新环境的请求参数"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/environments	[put]
func (c *EnvironmentCtrl) Update(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	var env model.Environment
	err := ctx.ReadJSON(&env)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	err = c.EnvironmentService.Update(tenantId, &env)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

// Copy 复制
// @Tags	环境管理
// @summary	复制环境
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	id 				query 	int 	true 	"环境id"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/environments/copyEnvironment	[post]
func (c *EnvironmentCtrl) Copy(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	id, err := ctx.URLParamInt("id")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	err = c.EnvironmentService.Copy(tenantId, id)
	if err != nil {
		ctx.JSON(_domain.Response{
			Code: _domain.SystemErr.Code,
		})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

// Delete 删除
// @Tags	环境管理
// @summary	删除环境(路径传参)
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	id 				path 	int 	true 	"环境ID"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/environments/{id}	[delete]
func (c *EnvironmentCtrl) Delete(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	err = c.EnvironmentService.Delete(tenantId, uint(id))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

// CreateVar 添加
// @Tags	环境管理/全局变量
// @summary	新建环境变量
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization	header	string					true	"Authentication header"
// @Param 	currProjectId	query	int						true	"当前项目ID"
// @Param 	EnvironmentVar 	body 	model.EnvironmentVar 	true 	"新建环境变量的请求参数"
// @success	200	{object}	_domain.Response{data=model.Environment}
// @Router	/api/v1/environments/vars	[post]
func (c *EnvironmentCtrl) CreateVar(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	envVar := model.EnvironmentVar{}
	err := ctx.ReadJSON(&envVar)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	err = c.EnvironmentService.CreateVar(tenantId, &envVar)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ErrNameExist.Code, Data: nil})
		return
	}

	env, err := c.EnvironmentService.Get(tenantId, envVar.EnvironmentId, 0)

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: env, Msg: _domain.NoErr.Msg})
}

// UpdateVar 更新
// @Tags	环境管理/全局变量
// @summary	更新环境变量
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization	header	string					true	"Authentication header"
// @Param 	currProjectId	query	int						true	"当前项目ID"
// @Param 	EnvironmentVar 	body 	model.EnvironmentVar 	true 	"更新环境变量的请求参数"
// @success	200	{object}	_domain.Response{data=model.Environment}
// @Router	/api/v1/environments/vars	[put]
func (c *EnvironmentCtrl) UpdateVar(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	var envVar model.EnvironmentVar
	err := ctx.ReadJSON(&envVar)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	err = c.EnvironmentService.UpdateVar(tenantId, &envVar)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	env, err := c.EnvironmentService.Get(tenantId, envVar.EnvironmentId, 0)
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: env, Msg: _domain.NoErr.Msg})
}

// DeleteVar 删除
func (c *EnvironmentCtrl) DeleteVar(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	err = c.EnvironmentService.DeleteVar(tenantId, uint(id))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	envVar, err := c.EnvironmentService.GetVar(tenantId, uint(id))
	env, err := c.EnvironmentService.Get(tenantId, envVar.EnvironmentId, 0)
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: env, Msg: _domain.NoErr.Msg})
}

// ClearVar 清除
// @Tags	环境管理/全局变量
// @summary	清空环境变量
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	environmentId 	query 	int 	true 	"环境ID"
// @success	200	{object}	_domain.Response{data=model.Environment}
// @Router	/api/v1/environments/vars/clear	[post]
func (c *EnvironmentCtrl) ClearVar(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	environmentId, err := ctx.URLParamInt("environmentId")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	err = c.EnvironmentService.ClearAllVar(tenantId, uint(environmentId))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	env, err := c.EnvironmentService.Get(tenantId, uint(environmentId), 0)
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: env, Msg: _domain.NoErr.Msg})
}

// DeleteShareVar 删除
// @Tags	环境管理/共享变量
// @summary	删除共享变量
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	id		path	int		true	"环境变量ID"
// @success	200	{object}	_domain.Response{data=model.Environment}
// @Router	/api/v1/environments/shareVars/{id}	[delete]
func (c *EnvironmentCtrl) DeleteShareVar(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	err = c.EnvironmentService.DisableShareVar(tenantId, uint(id))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	envVar, err := c.EnvironmentService.GetVar(tenantId, uint(id))
	env, err := c.EnvironmentService.Get(tenantId, envVar.EnvironmentId, 0)
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: env, Msg: _domain.NoErr.Msg})
}

// ClearShareVar 清除
// @Tags	环境管理/共享变量
// @summary	清空共享变量
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	interfaceId		query	int		true	"interfaceId"
// @success	200	{object}	_domain.Response{data=model.Environment}
// @Router	/api/v1/environments/shareVars/clear	[post]
func (c *EnvironmentCtrl) ClearShareVar(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	interfaceId, err := ctx.URLParamInt("interfaceId")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	err = c.EnvironmentService.DisableAllShareVar(uint(interfaceId))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	env, err := c.EnvironmentService.Get(tenantId, uint(interfaceId), 0)
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: env, Msg: _domain.NoErr.Msg})
}

// Save
// @Tags	环境管理
// @summary	保存环境
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization	header	string						true	"Authentication header"
// @Param 	currProjectId	query	int							true	"当前项目ID"
// @Param 	EnvironmentReq 	body 	serverDomain.EnvironmentReq true 	"保存环境的请求参数"
// @success	200	{object}	_domain.Response{data=int}
// @Router	/api/v1/environments/save	[post]
func (c *EnvironmentCtrl) Save(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	var req serverDomain.EnvironmentReq
	if err := ctx.ReadJSON(&req); err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	if id, err := c.EnvironmentService.Save(tenantId, req); err == nil {
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: id, Msg: _domain.NoErr.Msg})
	} else {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
	}
}

// Clone
// @Tags	环境管理
// @summary	复制环境
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	id 				query 	int 	false 	"环境ID"
// @success	200	{object}	_domain.Response{data=int}
// @Router	/api/v1/environments/copy	[get]
func (c *EnvironmentCtrl) Clone(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	id := ctx.URLParamIntDefault("id", 0)
	if id != 0 {
		if env, err := c.EnvironmentService.Clone(tenantId, uint(id)); err != nil {
			ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		} else {
			ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: env.ID, Msg: _domain.NoErr.Msg})
		}
	}
}

// DeleteEnvironment
// @Tags	环境管理
// @summary	删除环境(param传参)
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	id 				query 	int 	false 	"环境ID"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/environments/delete	[delete]
func (c *EnvironmentCtrl) DeleteEnvironment(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	id := ctx.URLParamIntDefault("id", 0)
	if id != 0 {
		if err := c.EnvironmentService.DeleteEnvironment(tenantId, uint(id)); err != nil {
			ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
			return
		}
	}
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

// ListAll
// @Tags	环境管理
// @summary	环境列表(ListAll)
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	projectId 		query 	int 	false 	"项目ID"
// @success	200	{object}	_domain.Response{data=[]model.Environment}
// @Router	/api/v1/environments/list	[get]
func (c *EnvironmentCtrl) ListAll(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	projectId := ctx.URLParamIntDefault("projectId", 0)
	if res, err := c.EnvironmentService.ListAll(tenantId, uint(projectId)); err == nil {
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: res, Msg: _domain.NoErr.Msg})
	} else {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
	}
}

// SaveGlobal
// @Tags	环境管理/全局变量
// @summary	保存全局变量
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization		header	string								true	"Authentication header"
// @Param 	currProjectId		query	int									true	"当前项目ID"
// @Param 	EnvironmentVariable body 	[]serverDomain.EnvironmentVariable 	true 	"保存全局变量的请求参数"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/environments/vars/global	[post]
func (c *EnvironmentCtrl) SaveGlobal(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	var req []serverDomain.EnvironmentVariable
	projectId := ctx.URLParamIntDefault("currProjectId", 0)

	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	err = c.EnvironmentService.SaveGlobal(tenantId, uint(projectId), req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

// ListGlobal
// @Tags	环境管理/全局变量
// @summary	列出全局变量
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @success	200	{object}	_domain.Response{data=[]model.EnvironmentVar}
// @Router	/api/v1/environments/vars/global	[get]
func (c *EnvironmentCtrl) ListGlobal(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	projectId := ctx.URLParamIntDefault("currProjectId", 0)
	res, err := c.EnvironmentService.ListGlobal(tenantId, uint(projectId))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: res, Msg: _domain.NoErr.Msg})
}

// SaveParams
// @Tags	环境管理/全局参数
// @summary	保存全局参数
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization			header	string								true	"Authentication header"
// @Param 	currProjectId			query	int									true	"当前项目ID"
// @Param 	EnvironmentParamsReq 	body 	serverDomain.EnvironmentParamsReq 	true 	"保保存全局参数的请求参数"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/environments/param	[post]
func (c *EnvironmentCtrl) SaveParams(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	var req serverDomain.EnvironmentParamsReq
	if err := ctx.ReadJSON(&req); err == nil {
		if err = c.EnvironmentService.SaveParams(tenantId, req); err == nil {
			ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
		} else {
			ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		}
	} else {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
	}
}

// ListParams
// @Tags	环境管理/全局参数
// @summary	全局参数列表
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	projectId		query	int		false	"项目ID"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/environments/param	[get]
func (c *EnvironmentCtrl) ListParams(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	projectId := ctx.URLParamIntDefault("projectId", 0)
	res, err := c.EnvironmentService.ListParams(tenantId, uint(projectId))
	if err == nil {
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: res, Msg: _domain.NoErr.Msg})
	} else {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
	}
}

// Order
// @Tags	环境管理
// @summary	修改环境的顺序
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization		header	string							true	"Authentication header"
// @Param 	currProjectId		query	int								true	"当前项目ID"
// @Param 	EnvironmentIdsReq 	body 	serverDomain.EnvironmentIdsReq 	true 	"修改环境顺序的请求参数"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/environments/order	[post]
func (c *EnvironmentCtrl) Order(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	var req serverDomain.EnvironmentIdsReq
	if err := ctx.ReadJSON(&req); err == nil {
		if err = c.EnvironmentService.SaveOrder(tenantId, req); err == nil {
			ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
		} else {
			ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		}
	} else {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
	}
}

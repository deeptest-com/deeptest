package handler

import (
	v1 "github.com/deeptest-com/deeptest/cmd/server/v1/domain"
	"github.com/deeptest-com/deeptest/internal/server/modules/model"
	"github.com/deeptest-com/deeptest/internal/server/modules/service"
	"github.com/deeptest-com/deeptest/pkg/domain"
	"github.com/kataras/iris/v12"
	"github.com/snowlyg/multi"
)

type DatabaseConnCtrl struct {
	DatabaseConnService *service.DatabaseConnService `inject:""`
	BaseCtrl
}

// List
// @Tags	数据库连接
// @summary	数据库连接列表
// @accept 	application/json
// @Produce application/json
// @Param	Authorization			header	string								true	"Authentication header"
// @Param 	envId			        query	int									true	"当前环境ID"
// @success	200	{object}	        _domain.Response{data=[]model.DatabaseConn}
// @Router	/api/v1/dbconns	[get]
func (c *DatabaseConnCtrl) List(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	projectId, err := ctx.URLParamInt("currProjectId")
	if projectId == 0 {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	keywords := ctx.URLParam("keywords")
	ignoreDisabled, err := ctx.URLParamBool("ignoreDisabled")

	res, err := c.DatabaseConnService.List(tenantId, keywords, projectId, ignoreDisabled)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: res})
}

// Get 详情
// @Tags	数据库连接
// @summary	数据库连接详情
// @accept 	application/json
// @Produce application/json
// @Param	Authorization		header	string	true	"Authentication header"
// @Param 	id					path	int		true	"数据库连接ID"
// @success	200	{object}	_domain.Response{data=model.DatabaseConn}
// @Router	/api/v1/dbconns/{id}	[get]
func (c *DatabaseConnCtrl) Get(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	po, err := c.DatabaseConnService.Get(tenantId, uint(id))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: po})
}

// Save 	保存
// @Tags	数据库连接
// @summary	保存数据库连接
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string								true	"Authentication header"
// @Param 	DatabaseConn    body	model.DatabaseConn 					true	"更新数据库连接的请求体"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/dbconns	[post]
func (c *DatabaseConnCtrl) Save(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	projectId, err := ctx.URLParamInt("currProjectId")

	req := model.DatabaseConn{}
	err = ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	req.ProjectId = uint(projectId)

	userName := multi.GetUsername(ctx)
	if req.ID > 0 {
		req.UpdateUser = userName
	} else {
		req.CreateUser = userName
	}

	err = c.DatabaseConnService.Save(tenantId, &req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ErrNameExist.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code})
}

func (c *DatabaseConnCtrl) UpdateName(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	projectId, err := ctx.URLParamInt("currProjectId")

	req := v1.DbConnReq{}
	err = ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	req.ProjectId = uint(projectId)
	req.UpdateUser = multi.GetUsername(ctx)

	err = c.DatabaseConnService.UpdateName(tenantId, req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ErrNameExist.Code, Data: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code})
}

// Delete 	删除
// @Tags	数据库连接
// @summary	删除数据库连接
// @accept 	application/json
// @Produce application/json
// @Param	Authorization		header	string	true	"Authentication header"
// @Param 	id					path	int		true	"数据库连接ID"
// @success	200	{object}	    _domain.Response
// @Router	/api/v1/dbconns/{id}	[delete]
func (c *DatabaseConnCtrl) Delete(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	err = c.DatabaseConnService.Delete(tenantId, uint(id))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

// Disable  禁用
// @Tags	数据库连接
// @summary	禁用数据库连接
// @accept 	application/json
// @Produce application/json
// @Param	Authorization		header	string	true	"Authentication header"
// @Param 	currProjectId		query	int		true	"当前项目ID"
// @Param 	id					path	int		true	"数据库连接ID"
// @success	200	{object}	    _domain.Response
// @Router	/api/v1/dbconns/{id}/disable	[put]
func (c *DatabaseConnCtrl) Disable(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	err = c.DatabaseConnService.Disable(tenantId, uint(id))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

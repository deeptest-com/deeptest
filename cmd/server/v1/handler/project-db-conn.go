package handler

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/service"
	"github.com/aaronchen2k/deeptest/pkg/domain"
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
	envId, _ := ctx.URLParamInt("envId")

	pos, _ := c.DatabaseConnService.List(uint(envId))
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: pos})
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
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	po, err := c.DatabaseConnService.Get(uint(id))
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
	req := model.DatabaseConn{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	err = c.DatabaseConnService.Save(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code})
}

func (c *DatabaseConnCtrl) UpdateName(ctx iris.Context) {
	req := v1.DbConnReq{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	req.UpdateUser = multi.GetUsername(ctx)

	err = c.DatabaseConnService.UpdateName(req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: err.Error()})
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
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	err = c.DatabaseConnService.Delete(uint(id))
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
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	err = c.DatabaseConnService.Disable(uint(id))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

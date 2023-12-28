package handler

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/service"
	"github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/kataras/iris/v12"
)

type DatabaseOptCtrl struct {
	DatabaseOptService *service.DatabaseOptService `inject:""`
	BaseCtrl
}

// Get 详情
// @Tags	数据库操作
// @summary	数据库操作详情
// @accept 	application/json
// @Produce application/json
// @Param	Authorization		header	string	true	"Authentication header"
// @Param 	id					path	int		true	"数据库操作ID"
// @success	200	{object}		_domain.Response{data=model.DebugConditionDatabaseOpt}
// @Router	/api/v1/databaseOpts/{id}	[get]
func (c *DatabaseOptCtrl) Get(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	opt, err := c.DatabaseOptService.Get(uint(id))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
		return
	}
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: opt})
}

// Update 更新
// @Tags	数据库连接
// @summary	更新数据库连接
// @accept 	application/json
// @Produce application/json
// @Param	Authorization				header	string								true	"Authentication header"
// @Param 	DebugInterfaceCheckpoint	body	model.DebugConditionDatabaseOpt		true	"更新数据库连接的请求体"
// @success	200	{object}				_domain.Response
// @Router	/api/v1/checkpoints	[put]
func (c *DatabaseOptCtrl) Update(ctx iris.Context) {
	var opt model.DebugConditionDatabaseOpt
	err := ctx.ReadJSON(&opt)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	err = c.DatabaseOptService.Update(&opt)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

// Delete 删除
// @Tags	数据库连接
// @summary	删除数据库连接
// @accept 	application/json
// @Produce application/json
// @Param	Authorization		header	string	true	"Authentication header"
// @Param 	id					path	int		true	"数据库连接ID"
// @success	200	{object}		_domain.Response
// @Router	/api/v1/checkpoints/{id}	[delete]
func (c *DatabaseOptCtrl) Delete(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	err = c.DatabaseOptService.Delete(uint(id))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

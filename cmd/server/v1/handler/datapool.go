package handler

import (
	"fmt"
	serverDomain "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/service"
	"github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/kataras/iris/v12"
	"github.com/snowlyg/multi"
)

type DatapoolCtrl struct {
	DatapoolService *service.DatapoolService `inject:""`
	BaseCtrl
}

// Index
// @Tags	数据池
// @summary	数据池列表
// @accept 	application/json
// @Produce application/json
// @Param	Authorization			header	string								true	"Authentication header"
// @Param 	currProjectId			query	int									true	"当前项目ID"
// @Param 	DatapoolReqPaginate		body	serverDomain.DatapoolReqPaginate	true	"获取数据池列表的请求体"
// @success	200	{object}	_domain.Response{data=_domain.PageData{result=[]model.Datapool}}
// @Router	/api/v1/datapools/index	[post]
func (c *DatapoolCtrl) Index(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	var req serverDomain.DatapoolReqPaginate

	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	if req.ProjectId == 0 {
		req.ProjectId, _ = ctx.URLParamInt64("currProjectId")
	}

	res, _ := c.DatapoolService.Paginate(tenantId, req)
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: res})
}

// Get
// @Tags	数据池
// @summary	数据池详情
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	id				path	int		true	"数据池ID"
// @success	200	{object}	_domain.Response{data=model.Datapool}
// @Router	/api/v1/datapools/{id}	[get]
func (c *DatapoolCtrl) Get(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	datapool, err := c.DatapoolService.Get(tenantId, uint(id))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: datapool, Msg: _domain.NoErr.Msg})
}

// Save
// @Tags	数据池
// @summary	保存数据池
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	Datapool				body	model.Datapool		true	"保存数据池的请求体"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/datapools/save	[post]
func (c *DatapoolCtrl) Save(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	userId := multi.GetUserId(ctx)
	projectId, err := ctx.URLParamInt("currProjectId")
	if projectId == 0 {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: "projectId"})
		return
	}

	req := model.Datapool{}
	err = ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}
	req.ProjectId = uint(projectId)

	// check name exist
	po, err := c.DatapoolService.GetByName(tenantId, req.Name, req.ProjectId)
	if po.ID > 0 && po.ID != req.ID {
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code,
			MsgKey: fmt.Sprintf("%v", _domain.ErrNameExist.Code)})
		return
	}

	err = c.DatapoolService.Save(tenantId, &req, userId)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code})
}

// Delete
// @Tags	数据池
// @summary	删除数据池
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	id				path	int		true	"数据池ID"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/datapools/{id}	[delete]
func (c *DatapoolCtrl) Delete(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	err = c.DatapoolService.Delete(tenantId, uint(id))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

// Disable
// @Tags	数据池
// @summary	禁用数据池
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	id				path	int		true	"数据池ID"
// @Param 	Datapool				body	model.Datapool		true	"保存数据池的请求体"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/datapools/{id}/disable	[put]
func (c *DatapoolCtrl) Disable(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	err = c.DatapoolService.Disable(tenantId, uint(id))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

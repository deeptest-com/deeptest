package handler

import (
	serverDomain "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/service"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/kataras/iris/v12"
	"github.com/snowlyg/multi"
)

type EndpointMockExpectCtrl struct {
	BaseCtrl
	EndpointMockExpectService *service.EndpointMockExpectService `inject:""`
}

// List
// @Tags	Mock期望
// @summary	期望列表
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	endpointId 		query 	int		true 	"endpointId"
// @success	200	{object}	_domain.Response{data=[]model.EndpointMockExpect}
// @Router	/api/v1/mockExpect/list	[get]
func (c *EndpointMockExpectCtrl) List(ctx iris.Context) {
	endpointId, err := ctx.URLParamInt("endpointId")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	res, err := c.EndpointMockExpectService.List(uint(endpointId))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: res, Msg: _domain.NoErr.Msg})

}

// Detail
// @Tags	Mock期望
// @summary	期望详情
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	id 		path 	int		true 	"期望ID"
// @success	200	{object}	_domain.Response{data=model.EndpointMockExpect}
// @Router	/api/v1/mockExpect	[get]
func (c *EndpointMockExpectCtrl) Detail(ctx iris.Context) {
	expectId, err := ctx.Params().GetInt("id")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	res, err := c.EndpointMockExpectService.GetDetail(uint(expectId))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: res, Msg: _domain.NoErr.Msg})
}

// Save
// @Tags	Mock期望
// @summary	保存期望
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization		header	string						true	"Authentication header"
// @Param 	currProjectId		query	int							true	"当前项目ID"
// @Param 	EndpointMockExpect 	body 	model.EndpointMockExpect	true 	"保存期望的请求体"
// @success	200	{object}	_domain.Response{data=int}
// @Router	/api/v1/mockExpect/save	[post]
func (c *EndpointMockExpectCtrl) Save(ctx iris.Context) {
	req := model.EndpointMockExpect{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	userName := multi.GetUsername(ctx)
	if req.ID == 0 {
		req.CreateUser = userName
	} else {
		req.UpdateUser = userName
	}
	expectId, err := c.EndpointMockExpectService.Save(req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: expectId, Msg: _domain.NoErr.Msg})
}

// Copy
// @Tags	Mock期望
// @summary	复制期望
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	id 				query 	int		true 	"期望ID"
// @success	200	{object}	_domain.Response{data=int}
// @Router	/api/v1/mockExpect/copy	[get]
func (c *EndpointMockExpectCtrl) Copy(ctx iris.Context) {
	expectId, err := ctx.URLParamInt("id")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	id, err := c.EndpointMockExpectService.Copy(uint(expectId))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: id, Msg: _domain.NoErr.Msg})
}

// Delete
// @Tags	Mock期望
// @summary	删除期望
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	id 				path 	int		true 	"期望ID"
// @success	200	{object}	_domain.Response{data=int}
// @Router	/api/v1/mockExpect	[delete]
func (c *EndpointMockExpectCtrl) Delete(ctx iris.Context) {
	expectId, err := ctx.Params().GetInt("id")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	if err = c.EndpointMockExpectService.DeleteById(uint(expectId)); err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

// Order
// @Tags	Mock期望
// @summary	对期望排序
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization		header	string							true	"Authentication header"
// @Param 	currProjectId		query	int								true	"当前项目ID"
// @Param 	MockExpectIdsReq 	body 	serverDomain.MockExpectIdsReq	true 	"对期望排序的请求体"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/mockExpect/order	[post]
func (c *EndpointMockExpectCtrl) Order(ctx iris.Context) {
	req := serverDomain.MockExpectIdsReq{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	if err = c.EndpointMockExpectService.SaveOrder(req); err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

// UpdateExpectDisabled
// @Tags	Mock期望
// @summary	启用或者禁用单个期望
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization		header	string							true	"Authentication header"
// @Param 	currProjectId		query	int								true	"当前项目ID"
// @Param 	id 	body 	int	true 	"期望ID"
// @Param 	disabled 	body 	bool	true 	"是否禁用"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/mockExpect/order	[post]
func (c *EndpointMockExpectCtrl) UpdateExpectDisabled(ctx iris.Context) {
	req := model.EndpointMockExpect{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	if err = c.EndpointMockExpectService.UpdateExpectDisabled(req.ID, req.Disabled); err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

package handler

import (
	"github.com/deeptest-com/deeptest/cmd/server/v1/domain"
	serverConsts "github.com/deeptest-com/deeptest/internal/server/consts"
	"github.com/deeptest-com/deeptest/internal/server/modules/service"
	"github.com/deeptest-com/deeptest/pkg/domain"
	logUtils "github.com/deeptest-com/deeptest/pkg/lib/log"
	"github.com/snowlyg/multi"

	"github.com/kataras/iris/v12"
)

type CategoryCtrl struct {
	CategoryService *service.CategoryService `inject:""`
	BaseCtrl
}

// LoadTree
// @Tags	分类管理
// @summary	分类树状数据
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string							true	"Authentication header"
// @Param 	currProjectId	query	int								true	"当前项目ID"
// @Param 	type			query	int								true	"类型"
// @Param 	serveId			query	int								true	"服务ID"
// @success	200	{object}	_domain.Response{data=serverDomain.Category}
// @Router	/api/v1/categories/load	[get]
func (c *CategoryCtrl) LoadTree(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	projectId, err := ctx.URLParamInt("currProjectId")
	nodeType := ctx.URLParamDefault("nodeType", "")
	if projectId == 0 {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}
	typ := ctx.URLParam("type")
	if typ == "" {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	data, err := c.CategoryService.GetTree(tenantId, serverConsts.CategoryDiscriminator(typ), projectId, serverConsts.NodeCreateType(nodeType))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
}

// Get 详情
// LoadTree
// @Tags	分类管理
// @summary	分类详情
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string							true	"Authentication header"
// @Param 	currProjectId	query	int								true	"当前项目ID"
// @Param 	id				path	int								true	"分类ID"
// @success	200	{object}	_domain.Response{data=model.Category}
// @Router	/api/v1/categories/{id}	[get]
func (c *CategoryCtrl) Get(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	po, err := c.CategoryService.Get(tenantId, id)

	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
		return
	}
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: po})
}

// Create 添加
// @Tags	分类管理
// @summary	新建分类
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization	header	string					true	"Authentication header"
// @Param 	currProjectId	query	int						true	"当前项目ID"
// @Param 	CategoryCreateReq 		body 	serverDomain.CategoryCreateReq true 	"新建分类的请求体"
// @success	200	{object}	_domain.Response{data=model.Category}
// @Router	/api/v1/categories	[post]
func (c *CategoryCtrl) Create(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	projectId, err := ctx.URLParamInt("currProjectId")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	req := serverDomain.CategoryCreateReq{}
	err = ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	req.ProjectId = uint(projectId)

	nodePo, bizErr := c.CategoryService.Create(tenantId, req)
	if bizErr != nil {
		ctx.JSON(_domain.Response{
			Code: _domain.SystemErr.Code,
		})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: nodePo})
}

// Update 更新
// @Tags	分类管理
// @summary	更新分类
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization	header	string						true	"Authentication header"
// @Param 	currProjectId	query	int							true	"当前项目ID"
// @Param 	CategoryReq 	body 	serverDomain.CategoryReq 	true 	"新建分类的请求体"
// @success	200	{object}	_domain.Response{data=serverDomain.CategoryReq}
// @Router	/api/v1/categories	[put]
func (c *CategoryCtrl) Update(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	req := serverDomain.CategoryReq{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	err = c.CategoryService.Update(tenantId, req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: req})
}

// UpdateName 更新
// @Tags	分类管理
// @summary	更新节点名称
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization	header	string						true	"Authentication header"
// @Param 	currProjectId	query	int							true	"当前项目ID"
// @Param 	id				path	int							true	"分类ID"
// @Param 	CategoryReq 	body 	serverDomain.CategoryReq 	true 	"更新节点名称的请求体"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/categories/{id}/updateName	[put]
func (c *CategoryCtrl) UpdateName(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	var req serverDomain.CategoryReq
	err := ctx.ReadJSON(&req)
	if err != nil {
		logUtils.Errorf("参数验证失败", err.Error())
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	err = c.CategoryService.UpdateName(tenantId, req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

// Delete 删除
// @Tags	分类管理
// @summary	删除节点
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization	header	string								true	"Authentication header"
// @Param 	currProjectId	query	int									true	"当前项目ID"
// @Param 	id				path	int									true	"分类ID"
// @Param 	type			query	serverConsts.CategoryDiscriminator	true	"类型"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/categories/{id}	[delete]
func (c *CategoryCtrl) Delete(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	projectId, err := ctx.URLParamInt("currProjectId")
	typ := ctx.URLParam("type")
	if typ == "" {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	err = c.CategoryService.Delete(tenantId, serverConsts.CategoryDiscriminator(typ), uint(projectId), uint(id))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

// Move 移动
// @Tags	分类管理
// @summary	移动节点
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization	header	string					true	"Authentication header"
// @Param 	currProjectId	query	int						true	"当前项目ID"
// @Param 	CategoryMoveReq 		body 	serverDomain.CategoryMoveReq true 	"移动节点的请求体"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/categories/move	[post]
func (c *CategoryCtrl) Move(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	projectId, _ := ctx.URLParamInt("currProjectId")

	var req serverDomain.CategoryMoveReq
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	_, err = c.CategoryService.Move(tenantId, uint(req.DragKey), uint(req.DropKey), req.DropPos, req.Type, uint(projectId))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

func (c *CategoryCtrl) BatchAddSchemaRoot(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	var req serverDomain.BatchAddSchemaRootReq
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	go c.CategoryService.BatchAddSchemaRoot(tenantId, req.ProjectIds)

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

// Copy 详情
// @Tags	分类管理
// @summary	复制分类
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string							true	"Authentication header"
// @Param 	id				path	int								true	"分类ID"
// @success	200	{object}	_domain.Response{}
// @Router	/api/v1/categories/copy/{id}	[get]
func (c *CategoryCtrl) Copy(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	targetId, err := ctx.Params().GetInt("id")

	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	userId := multi.GetUserId(ctx)
	userName := multi.GetUsername(ctx)

	nodePo, err := c.CategoryService.Copy(tenantId, uint(targetId), 0, userId, userName)

	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})

	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg, Data: nodePo})
}

func (c *CategoryCtrl) LoadChildren(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	projectId, err := ctx.URLParamInt("currProjectId")
	nodeType := ctx.URLParamDefault("nodeType", "")
	if projectId == 0 {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}
	typ := ctx.URLParam("type")
	if typ == "" {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}
	categoryId := ctx.URLParamIntDefault("categoryId", 0)
	data, err := c.CategoryService.GetChildrenNodes(tenantId, serverConsts.CategoryDiscriminator(typ), projectId, categoryId, serverConsts.NodeCreateType(nodeType))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
}

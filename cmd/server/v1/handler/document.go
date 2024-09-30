package handler

import (
	domain "github.com/deeptest-com/deeptest/cmd/server/v1/domain"
	"github.com/deeptest-com/deeptest/internal/server/modules/service"
	_domain "github.com/deeptest-com/deeptest/pkg/domain"
	"github.com/kataras/iris/v12"
)

type DocumentCtrl struct {
	BaseCtrl
	DocumentService *service.DocumentService `inject:""`
}

// Index
// @Tags	接口文档
// @summary	接口文档列表
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization	header	string						true	"Authentication header"
// @Param 	currProjectId	query	int							true	"当前项目ID"
// @Param 	DocumentReq 	body 	serverDomain.DocumentReq 	true 	"查看接口文档列表的请求参数"
// @success	200	{object}	_domain.Response{data=serverDomain.DocumentRep}
// @Router	/api/v1/document	[post]
func (c *DocumentCtrl) Index(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	var req domain.DocumentReq
	if err := ctx.ReadJSON(&req); err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	res, _ := c.DocumentService.Content(tenantId, req)
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: res})

	return
}

// DocumentVersionList
// @Tags	接口文档
// @summary	接口文档版本列表
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization			header	string								true	"Authentication header"
// @Param 	currProjectId			query	int									true	"当前项目ID"
// @Param 	DocumentVersionListReq 	body 	serverDomain.DocumentVersionListReq true 	"接口文档版本列表的请求参数"
// @success	200	{object}	_domain.Response{data=[]model.EndpointDocument}
// @Router	/api/v1/document/version_list	[post]
func (c *DocumentCtrl) DocumentVersionList(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	projectId, _ := ctx.URLParamInt("currProjectId")

	var req domain.DocumentVersionListReq
	if err := ctx.ReadJSON(&req); err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	res, err := c.DocumentService.GetDocumentVersionList(tenantId, uint(projectId), req.NeedLatest)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg, Data: res})
	return
}

// Publish
// @Tags	接口文档
// @summary	发布接口文档
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization		header	string							true	"Authentication header"
// @Param 	currProjectId		query	int								true	"当前项目ID"
// @Param 	DocumentVersionReq 	body 	serverDomain.DocumentVersionReq true 	"发布接口文档的请求参数"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/document/publish	[post]
func (c *DocumentCtrl) Publish(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	projectId, _ := ctx.URLParamInt("currProjectId")

	var req domain.DocumentVersionReq
	if err := ctx.ReadJSON(&req); err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	documentId, err := c.DocumentService.Publish(tenantId, req, uint(projectId))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: documentId, Msg: _domain.NoErr.Msg})
	return
}

// DeleteSnapshot
// @Tags	接口文档
// @summary	删除接口文档
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	id 				query 	int 	true 	"文档ID"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/document/delete	[delete]
func (c *DocumentCtrl) DeleteSnapshot(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	id := ctx.URLParamUint64("id")
	err := c.DocumentService.RemoveSnapshot(tenantId, uint(id))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
	return
}

// UpdateDocument
// @Tags	接口文档
// @summary	更新文档版本信息
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization			header	string									true	"Authentication header"
// @Param 	currProjectId			query	int										true	"当前项目ID"
// @Param 	UpdateDocumentVersionReq body 	serverDomain.UpdateDocumentVersionReq 	true 	"更新文档版本信息的请求参数"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/document/update_version	[post]
func (c *DocumentCtrl) UpdateDocument(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	var req domain.UpdateDocumentVersionReq
	if err := ctx.ReadJSON(&req); err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	err := c.DocumentService.UpdateDocument(tenantId, req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
	return
}

// GetShareLink
// @Tags	接口文档
// @summary	生成分享接口文档的链接
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization		header	string							true	"Authentication header"
// @Param 	currProjectId		query	int								true	"当前项目ID"
// @Param 	DocumentShareReq 	body 	serverDomain.DocumentShareReq 	true 	"生成分享接口文档的链接的请求参数"
// @success	200	{object}	_domain.Response{data=object{code=string}}
// @Router	/api/v1/document/share	[post]
func (c *DocumentCtrl) GetShareLink(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	var req domain.DocumentShareReq
	if err := ctx.ReadJSON(&req); err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	code, err := c.DocumentService.GenerateShareLink(tenantId, req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}
	res := iris.Map{"code": code}
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg, Data: res})
	return
}

// GetContentsByShareLink
// @Tags	接口文档
// @summary	查看分享的文档
// @accept	application/json
// @Produce	application/json
// @Param 	code 	query 	string 	true 	"查看接口文档的链接"
// @success	200	{object}	_domain.Response{data=serverDomain.DocumentRep}
// @Router	/api/v1/document/get_share_content	[get]
func (c *DocumentCtrl) GetContentsByShareLink(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	link := ctx.URLParam("code")
	if link == "" {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: "code can't be empty"})
		return
	}

	res, err := c.DocumentService.ContentByShare(tenantId, link)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: res})

	return
}

// GetDocumentDetail
// @Tags	接口文档
// @summary	查看文档详情
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization		header	string	true	"Authentication header"
// @Param 	currProjectId		query	int		true	"当前项目ID"
// @Param 	documentId 			query 	int 	true 	"文档ID"
// @Param 	endpointId 			query 	int 	true 	"endpointId"
// @Param 	interfaceId 		query 	int 	true 	"interfaceId"
// @success	200	{object}	_domain.Response{data=object{interface=model.EndpointInterface,servers=[]model.ServeServer}}
// @Router	/api/v1/document/share_detail	[get]
func (c *DocumentCtrl) GetDocumentDetail(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	documentId, _ := ctx.URLParamInt("documentId")
	endpointId, _ := ctx.URLParamInt("endpointId")
	interfaceId, _ := ctx.URLParamInt("interfaceId")
	if interfaceId == 0 {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: "interfaceId can't be empty"})
		return
	}
	if documentId != 0 && endpointId == 0 {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: "endpointId can't be empty"})
		return
	}

	res, err := c.DocumentService.GetDocumentDetail(tenantId, uint(documentId), uint(endpointId), uint(interfaceId))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: res})

	return
}

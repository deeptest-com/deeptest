package handler

import (
	"github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/service"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	commonUtils "github.com/aaronchen2k/deeptest/pkg/lib/comm"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/jinzhu/copier"
	"github.com/kataras/iris/v12"
	"github.com/snowlyg/multi"
	encoder "github.com/zwgblue/yaml-encoder"
	"go.uber.org/zap"
)

type EndpointCtrl struct {
	BaseCtrl
	EndpointService       *service.EndpointService       `inject:""`
	ServeService          *service.ServeService          `inject:""`
	ThirdPartySyncService *service.ThirdPartySyncService `inject:""`
}

// Index
// @Tags	设计器
// @summary	设计器列表
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization		header	string							true	"Authentication header"
// @Param 	currProjectId		query	int								true	"当前项目ID"
// @Param 	EndpointReqPaginate body 	serverDomain.EndpointReqPaginate true 	"设计器列表的请求体"
// @success	200	{object}	_domain.Response{data=object{result=[]model.Endpoint}}
// @Router	/api/v1/endpoint/index	[post]
func (c *EndpointCtrl) Index(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	var req serverDomain.EndpointReqPaginate
	if err := ctx.ReadJSON(&req); err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}
	res, _ := c.EndpointService.Paginate(tenantId, req)
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: res})

	return
}

// Save
// @Tags	设计器
// @summary	保存设计器
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization		header	string							true	"Authentication header"
// @Param 	currProjectId		query	int								true	"当前项目ID"
// @Param 	EndpointReq body 	serverDomain.EndpointReq true 	"保存设计器的请求参数"
// @success	200	{object}	_domain.Response{data=int}
// @Router	/api/v1/endpoint/save	[post]
func (c *EndpointCtrl) Save(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	var req serverDomain.EndpointReq
	err := ctx.ReadJSON(&req)

	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	userName := multi.GetUsername(ctx)
	if req.ID == 0 {
		req.CreateUser = userName
	} else {
		req.UpdateUser = userName
	}

	endpoint := c.requestParser(req)

	/*
		if endpoint.CategoryId == 0 {
			endpoint.CategoryId = 0
		}
	*/

	if res, err := c.EndpointService.Save(tenantId, endpoint); err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
	} else {
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: res})
	}

	return
}

// Detail
// @Tags	设计器
// @summary	设计器详情
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	id 				query 	int 	true 	"设计器id"
// @Param 	version 		query 	string false 	"设计器版本"
// @success	200	{object}	_domain.Response{data=model.Endpoint}
// @Router	/api/v1/endpoint/detail	[get]
func (c *EndpointCtrl) Detail(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	id := ctx.URLParamUint64("id")
	version := ctx.URLParamDefault("version", c.EndpointService.GetLatestVersion(tenantId, uint(id)))
	if id != 0 {
		res := c.EndpointService.GetById(tenantId, uint(id), version)
		userId := multi.GetUserId(ctx)
		res.IsFavorite = c.EndpointService.IsFavorite(tenantId, uint(id), userId)
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: res})
	} else {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
	}
}

// Delete
// @Tags	设计器
// @summary	删除设计器
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	id 				query 	int 	true 	"设计器id"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/endpoint/delete	[delete]
func (c *EndpointCtrl) Delete(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	id := ctx.URLParamUint64("id")

	err := c.EndpointService.DeleteById(tenantId, uint(id))

	if err == nil {
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
	} else {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
	}
}

// BatchDelete
// @Tags	设计器
// @summary	批量删除设计器
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	req 			query 	[]int 	true 	"设计器id"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/endpoint/batchDelete	[delete]
func (c *EndpointCtrl) BatchDelete(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	var req []uint
	if err := ctx.ReadJSON(&req); err == nil {
		c.EndpointService.BatchDelete(tenantId, req)
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
	} else {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
	}
}

// 构造参数构造auth，BasicAuth,BearerToken,OAuth20,ApiKey
func (c *EndpointCtrl) requestParser(req serverDomain.EndpointReq) (endpoint model.Endpoint) {

	for key, item := range req.Interfaces {
		req.Interfaces[key].Body = item.RequestBody.Examples
		if item.RequestBody.MediaType == "" {
			req.Interfaces[key].RequestBody.MediaType = "application/json"
		}
		req.Interfaces[key].BodyType = item.RequestBody.MediaType
		req.Interfaces[key].Name = req.Title
		/*
					if req.Interfaces[key].RequestBody.Examples == "" {
						var examples []map[string]string
			//			example := c.ServeService.Schema2Example(req.ServeId, item.RequestBody.SchemaItem.Content)
			//			examples = append(examples, map[string]string{"name": "default", "content": commonUtils.JsonEncode(example)})
			//			req.Interfaces[key].RequestBody.Examples = commonUtils.JsonEncode(examples)
					}
		*/

	}

	if req.CategoryId == 0 {
		req.CategoryId = -1
	}

	if req.Status == 0 {
		req.Status = 1
	}

	copier.CopyWithOption(&endpoint, &req, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	return
}

// Expire
// @Tags	设计器
// @summary	禁用设计器
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	id 				query 	int 	true 	"设计器id"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/endpoint/expire	[put]
func (c *EndpointCtrl) Expire(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	id := ctx.URLParamUint64("id")
	err := c.EndpointService.DisableById(tenantId, uint(id))
	if err == nil {
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
	} else {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
	}
}

// Publish
// @Tags	设计器
// @summary	发布设计器
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	id 				query 	int 	true 	"设计器id"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/endpoint/publish	[put]
func (c *EndpointCtrl) Publish(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	id := ctx.URLParamUint64("id")
	err := c.EndpointService.Publish(tenantId, uint(id))
	if err == nil {
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
	} else {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
	}
}

// Develop
// @Tags	设计器
// @summary	开发设计器
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	id 				query 	int 	true 	"设计器id"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/endpoint/develop	[put]
func (c *EndpointCtrl) Develop(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	id := ctx.URLParamUint64("id")
	err := c.EndpointService.Develop(tenantId, uint(id))
	if err == nil {
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
	} else {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
	}
}

// Copy
// @Tags	设计器
// @summary	复制设计器
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	id 				query 	int 	true 	"设计器id"
// @Param 	version 		query 	string 	false 	"设计器版本"
// @success	200	{object}	_domain.Response{data=int}
// @Router	/api/v1/endpoint/copy	[get]
func (c *EndpointCtrl) Copy(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	id := ctx.URLParamUint64("id")
	version := ctx.URLParamDefault("version", c.EndpointService.GetLatestVersion(tenantId, uint(id)))

	userId := multi.GetUserId(ctx)
	userName := multi.GetUsername(ctx)
	res, err := c.EndpointService.Copy(tenantId, uint(id), 0, userId, userName, version)
	if err == nil {
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: res, Msg: _domain.NoErr.Msg})
	} else {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
	}
}

// Yaml
// @Tags	设计器
// @summary	设计器信息转yaml
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization	header	string						true	"Authentication header"
// @Param 	currProjectId	query	int							true	"当前项目ID"
// @Param 	EndpointReq 	body 	serverDomain.EndpointReq 	true 	"设计器信息转yaml的请求参数"
// @success	200	{object}	_domain.Response{data=string}
// @Router	/api/v1/endpoint/yaml	[post]
func (c *EndpointCtrl) Yaml(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	var req serverDomain.EndpointReq
	if err := ctx.ReadJSON(&req); err == nil {
		endpoint := c.requestParser(req)
		c.EndpointService.SchemasConv(tenantId, &endpoint, nil)
		res := c.EndpointService.Yaml(tenantId, endpoint)
		var ret interface{}
		commonUtils.JsonDecode(commonUtils.JsonEncode(res), &ret)
		content, _ := encoder.NewEncoder(ret).Encode()
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: string(content)})
	} else {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
	}
	return
}

// UpdateStatus
// @Tags	设计器
// @summary	更新设计器状态
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	id 				query 	int	true 	"设计器id"
// @Param 	status 			query 	int	true 	"设计器状态"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/endpoint/updateStatus	[put]
func (c *EndpointCtrl) UpdateStatus(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	id := ctx.URLParamUint64("id")
	status := ctx.URLParamUint64("status")
	err := c.EndpointService.UpdateStatus(tenantId, uint(id), int64(status))
	if err == nil {
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
	} else {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
	}
}

// AddVersion
// @Tags	设计器
// @summary	添加设计器版本
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization		header	string							true	"Authentication header"
// @Param 	currProjectId		query	int								true	"当前项目ID"
// @Param 	EndpointVersionReq 	body 	serverDomain.EndpointVersionReq	true 	"添加设计器版本的请求参数"
// @success	200	{object}	_domain.Response{data=string}
// @Router	/api/v1/endpoint/version/add	[post]
func (c *EndpointCtrl) AddVersion(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	var req serverDomain.EndpointVersionReq
	if err := ctx.ReadJSON(&req); err == nil {
		var version model.EndpointVersion
		copier.CopyWithOption(&version, &req, copier.Option{IgnoreEmpty: true, DeepCopy: true})
		err = c.EndpointService.AddVersion(tenantId, &version)
		if err == nil {
			ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: req.Version})
		} else {
			ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		}
	} else {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
	}
}

// ListVersions
// @Tags	设计器
// @summary	设计器版本列表
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	id 				query 	int		true 	"设计器id"
// @success	200	{object}	_domain.Response{data=[]model.EndpointVersion}
// @Router	/api/v1/endpoint/version/list	[get]
func (c *EndpointCtrl) ListVersions(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	id := ctx.URLParamUint64("id")
	res, err := c.EndpointService.GetVersionsByEndpointId(tenantId, uint(id))
	if err == nil {
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: res})
	} else {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
	}
}

// BatchUpdateField
// @Tags	设计器
// @summary	批量更新字段内容
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization	header	string						true	"Authentication header"
// @Param 	currProjectId	query	int							true	"当前项目ID"
// @Param 	BatchUpdateReq 	body 	serverDomain.BatchUpdateReq	true 	"批量更新字段内容的请求参数"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/endpoint/batchUpdateField	[post]
func (c *EndpointCtrl) BatchUpdateField(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	var req serverDomain.BatchUpdateReq
	if err := ctx.ReadJSON(&req); err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	if err := c.EndpointService.BatchUpdateByField(tenantId, req); err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

// UpdateTag
// @Tags	设计器
// @summary	更新标签
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization	header	string						true	"Authentication header"
// @Param 	currProjectId	query	int							true	"当前项目ID"
// @Param 	EndpointTagReq 	body 	serverDomain.EndpointTagReq	true 	"更新标签的请求参数"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/endpoint/updateTag	[put]
func (c *EndpointCtrl) UpdateTag(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	var req serverDomain.EndpointTagReq
	if err := ctx.ReadJSON(&req); err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	projectId, _ := ctx.URLParamInt("currProjectId")
	if err := c.EndpointService.UpdateTags(tenantId, req, uint(projectId)); err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
	return
}

// UpdateAdvancedMockDisabled
// @Tags	设计器
// @summary	启用或者禁用接口所有期望
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization			header	string	true	"Authentication header"
// @Param 	currProjectId			query	int		true	"当前项目ID"
// @Param 	id 						body 	int		true 	"endpoint_id"
// @Param 	advancedMockDisabled 	body 	bool	true 	"接口的mock禁用状态 true:禁用 false:启用"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/endpoint/updateMockStatus	[post]
func (c *EndpointCtrl) UpdateAdvancedMockDisabled(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	var req model.Endpoint
	if err := ctx.ReadJSON(&req); err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	if err := c.EndpointService.UpdateAdvancedMockDisabled(tenantId, req.ID, req.AdvancedMockDisabled); err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

func (c *EndpointCtrl) SyncFromThirdParty(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	var req _domain.ReqId
	if err := ctx.ReadParams(&req); err != nil {
		logUtils.Errorf("参数解析失败", zap.String("错误:", err.Error()))
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	if err := c.EndpointService.SyncFromThirdParty(tenantId, req.Id); err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

/*
func (c *EndpointCtrl) Index() {
	c.EndpointService.GetVersionsByEndpointId(1)
}
*/

func (c *EndpointCtrl) GetDiff(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	endpointId, err := ctx.URLParamInt("endpointId")
	if err != nil {
		logUtils.Errorf("参数解析失败", zap.String("错误:", err.Error()))
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	if data, err := c.EndpointService.GetDiff(tenantId, uint(endpointId)); err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
	} else {
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data})
	}
}

func (c *EndpointCtrl) SaveDiff(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	var req serverDomain.EndpointDiffReq
	if err := ctx.ReadJSON(&req); err != nil {
		logUtils.Errorf("参数解析失败", zap.String("错误:", err.Error()))
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}
	userName := multi.GetUsername(ctx)
	if err := c.EndpointService.SaveDiff(tenantId, req.EndpointId, req.IsChanged, userName); err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

// UpdateName
// @Tags	设计器
// @summary	更新设计器名称
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	id 				query 	int	true 	"设计器id"
// @Param 	name 			query 	string	true 	"设计器状态"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/endpoint/updateName	[put]
func (c *EndpointCtrl) UpdateName(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	var req serverDomain.UpdateNameReq
	if err := ctx.ReadJSON(&req); err != nil {
		logUtils.Errorf("参数解析失败", zap.String("错误:", err.Error()))
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}
	err := c.EndpointService.UpdateName(tenantId, req.Id, req.Name)
	if err == nil {
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
	} else {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
	}
}

func (c *EndpointCtrl) ListFunctionsByThirdPartyClass(ctx iris.Context) {
	var req serverDomain.ImportThirdPartyEndpointReq
	if err := ctx.ReadJSON(&req); err != nil {
		logUtils.Errorf("参数解析失败", zap.String("错误:", err.Error()))
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	data, err := c.ThirdPartySyncService.ListFunctionsByClass(req.FilePath, req.ClassCode)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ErrThirdPartyFunctions.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg, Data: data})
}

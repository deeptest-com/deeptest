package handler

import (
	serverDomain "github.com/deeptest-com/deeptest/cmd/server/v1/domain"
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	"github.com/deeptest-com/deeptest/internal/pkg/domain"
	serverConsts "github.com/deeptest-com/deeptest/internal/server/consts"
	service "github.com/deeptest-com/deeptest/internal/server/modules/service"
	_domain "github.com/deeptest-com/deeptest/pkg/domain"
	"github.com/kataras/iris/v12"
	"github.com/snowlyg/multi"
)

type DiagnoseInterfaceCtrl struct {
	DiagnoseInterfaceService *service.DiagnoseInterfaceService `inject:""`
	ExtractorService         *service.ExtractorService         `inject:""`
	CheckpointService        *service.CheckpointService        `inject:""`

	DebugInterfaceService *service.DebugInterfaceService `inject:""`

	BaseCtrl
}

// Load
// @Tags	快捷调试
// @summary	获取测试接口列表
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	projectId		query	int		true	"项目ID"
// @success	200	{object}	_domain.Response{data=[]serverDomain.DiagnoseInterface}
// @Router	/api/v1/diagnoseInterfaces	[get]
func (c *DiagnoseInterfaceCtrl) Load(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	projectId, _ := ctx.URLParamInt("projectId")

	data, err := c.DiagnoseInterfaceService.Load(tenantId, projectId)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
}

// Get
// @Tags	快捷调试
// @summary	获取测试接口详情
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	id				path	int		true	"调试接口的id"
// @success	200	{object}	_domain.Response{data=model.DiagnoseInterface}
// @Router	/api/v1/diagnoseInterfaces/{id}	[get]
func (c *DiagnoseInterfaceCtrl) Get(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	id, _ := ctx.Params().GetInt("id")

	data, err := c.DiagnoseInterfaceService.Get(tenantId, id)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
}

// Save
// @Tags	快捷调试
// @summary	新建测试接口
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	DiagnoseInterfaceSaveReq				body	serverDomain.DiagnoseInterfaceSaveReq		true	"新建测试接口的请求体"
// @success	200	{object}	_domain.Response{data=[]serverDomain.DiagnoseInterface}
// @Router	/api/v1/diagnoseInterfaces	[post]
func (c *DiagnoseInterfaceCtrl) Save(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	req := serverDomain.DiagnoseInterfaceSaveReq{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	req.CreatedBy = multi.GetUserId(ctx)
	po, err := c.DiagnoseInterfaceService.Save(tenantId, req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	//data, err := c.DiagnoseInterfaceService.Load(int(po.ProjectId), int(po.ServeId))
	//if err != nil {
	//	ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
	//	return
	//}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: po})
}

// SaveDebugData
// @Tags	快捷调试
// @summary	保存测试调试接口
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	DebugData				body	domain.DebugData		true	"保存测试调试接口的请求体"
// @success	200	{object}	_domain.Response{data=domain.DebugData}
// @Router	/api/v1/diagnoseInterfaces/saveDebugData	[post]
func (c *DiagnoseInterfaceCtrl) SaveDebugData(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	req := domain.DebugData{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	_, err = c.DebugInterfaceService.CreateOrUpdate(tenantId, req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	loadReq := domain.DebugInfo{
		DiagnoseInterfaceId: req.DiagnoseInterfaceId,
		DebugInterfaceId:    req.DebugInterfaceId,
		UsedBy:              consts.DiagnoseDebug,
	}

	data, err := c.DebugInterfaceService.Load(tenantId, loadReq)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data})
}

// Update
// @Tags	快捷调试
// @summary	更新测试接口
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	DiagnoseInterfaceSaveReq				body	serverDomain.DiagnoseInterfaceSaveReq		true	"更新测试接口的请求体"
// @success	200	{object}	_domain.Response{data=[]serverDomain.DiagnoseInterface}
// @Router	/api/v1/diagnoseInterfaces	[put]
func (c *DiagnoseInterfaceCtrl) Update(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	req := serverDomain.DiagnoseInterfaceSaveReq{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	req.UpdatedBy = multi.GetUserId(ctx)
	po, err := c.DiagnoseInterfaceService.Save(tenantId, req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	//data, err := c.DiagnoseInterfaceService.Load(int(po.ProjectId), int(po.ServeId))
	//if err != nil {
	//	ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
	//	return
	//}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: po})
}

// Delete
// @Tags	快捷调试
// @summary	删除测试接口
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	id				path	int		true	"调试接口的id"
// @Param 	type			query	string	true	"type"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/diagnoseInterfaces/{id}	[delete]
func (c *DiagnoseInterfaceCtrl) Delete(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	id, _ := ctx.Params().GetInt("id")
	typ := ctx.URLParam("type")

	err := c.DiagnoseInterfaceService.Remove(tenantId, id, serverConsts.DiagnoseInterfaceType(typ))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code})
}

// Move 移动
// @Tags	快捷调试
// @summary	移动节点
// @accept 	application/json
// @Produce application/json
// @Param	Authorization				header	string									true	"Authentication header"
// @Param 	currProjectId				query	int										true	"当前项目ID"
// @Param 	DiagnoseInterfaceMoveReq	body	serverDomain.DiagnoseInterfaceMoveReq	true	"移动节点的请求体"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/diagnoseInterfaces/move	[post]
func (c *DiagnoseInterfaceCtrl) Move(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	projectId, _ := ctx.URLParamInt("currProjectId")

	var req serverDomain.DiagnoseInterfaceMoveReq
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	_, err = c.DiagnoseInterfaceService.Move(tenantId, uint(req.DragKey), uint(req.DropKey), req.DropPos, uint(projectId))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

// ImportInterfaces 导入接口
// @Tags	快捷调试
// @summary	导入接口
// @accept 	application/json
// @Produce application/json
// @Param	Authorization				header	string									true	"Authentication header"
// @Param 	currProjectId				query	int										true	"当前项目ID"
// @Param 	DiagnoseInterfaceImportReq	body	serverDomain.DiagnoseInterfaceImportReq	true	"导入接口的请求体"
// @success	200	{object}	_domain.Response{data=model.DiagnoseInterface}
// @Router	/api/v1/diagnoseInterfaces/importInterfaces	[post]
func (c *DiagnoseInterfaceCtrl) ImportInterfaces(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	req := serverDomain.DiagnoseInterfaceImportReq{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	req.CreateBy = multi.GetUserId(ctx)
	newNode, bizErr := c.DiagnoseInterfaceService.ImportInterfaces(tenantId, req)
	if bizErr != nil {
		ctx.JSON(_domain.Response{
			Code: _domain.SystemErr.Code,
		})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: newNode})
}

// ImportCurl 导入cURL命令
// @Tags	快捷调试
// @summary	导入cURL命令
// @accept 	application/json
// @Produce application/json
// @Param	Authorization				header	string								true	"Authentication header"
// @Param 	currProjectId				query	int									true	"当前项目ID"
// @Param 	DiagnoseCurlImportReq		body	serverDomain.DiagnoseCurlImportReq	true	"导入cURL命令的请求体"
// @success	200	{object}	_domain.Response{data=model.DiagnoseInterface}
// @Router	/api/v1/diagnoseInterfaces/importCurl	[post]
func (c *DiagnoseInterfaceCtrl) ImportCurl(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	req := serverDomain.DiagnoseCurlImportReq{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	req.CreateBy = multi.GetUserId(ctx)
	newNode, bizErr := c.DiagnoseInterfaceService.ImportCurl(tenantId, req)
	if bizErr != nil {
		ctx.JSON(_domain.Response{
			Code: _domain.SystemErr.Code,
			Msg:  bizErr.Error(),
		})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: newNode})
}

// Index
// @Tags	请求录制
// @summary	录制的请求转调试接口
// @accept 	application/json
// @Produce application/json
// @Param	Authorization			header	string					true	"Authentication header"
// @Param 	RecordReq		body	serverDomain.RecordReq	true	"录制的请求列表"
// @success	200	{object}	_domain.Response{}}
// @Router	/api/v1/records/importRecordData	[post]
func (c *DiagnoseInterfaceCtrl) ImportRecordData(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)

	var req serverDomain.RecordReq
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	req.UserId = multi.GetUserId(ctx)

	err = c.DiagnoseInterfaceService.ImportRecordData(tenantId, req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code})
}

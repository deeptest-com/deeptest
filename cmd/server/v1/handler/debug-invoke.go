package handler

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	service "github.com/aaronchen2k/deeptest/internal/server/modules/service"
	"github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/kataras/iris/v12"
)

type DebugInvokeCtrl struct {
	DebugInvokeService *service.DebugInvokeService `inject:""`
	ExtractorService   *service.ExtractorService   `inject:""`
	CheckpointService  *service.CheckpointService  `inject:""`
	BaseCtrl
}

// SubmitResult
// @Tags	接口调试
// @summary	Agent提交接口执行结果
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization				header	string										true	"Authentication header"
// @Param 	currProjectId				query	int											true	"当前项目ID"
// @Param 	SubmitDebugResultRequest	body	domain.SubmitDebugResultRequest				true	"Agent提交接口执行结果的请求体"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/debugs/invoke/submitResult	[post]
func (c *DebugInvokeCtrl) SubmitResult(ctx iris.Context) {
	req := domain.SubmitDebugResultRequest{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}
	var invoke model.DebugInvoke
	invoke, err = c.DebugInvokeService.SubmitResult(req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: invoke.ID})
}

// List
// @Tags	接口调试
// @summary	调试记录列表
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization		header	string	true	"Authentication header"
// @Param 	currProjectId		query	int		true	"当前项目ID"
// @Param 	debugInterfaceId	query	int		true	"debugInterfaceId"
// @Param 	endpointInterfaceId	query	int		true	"endpointInterfaceId"
// @success	200	{object}	_domain.Response{data=[]model.DebugInvoke}
// @Router	/api/v1/debugs/invoke	[get]
func (c *DebugInvokeCtrl) List(ctx iris.Context) {
	debugInterfaceId, err := ctx.URLParamInt("debugInterfaceId")
	endpointInterfaceId, err := ctx.URLParamInt("endpointInterfaceId")
	if debugInterfaceId <= 0 && endpointInterfaceId <= 0 {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	if debugInterfaceId < 0 {
		debugInterfaceId = 0
	}
	if endpointInterfaceId < 0 {
		endpointInterfaceId = 0
	}

	data, err := c.DebugInvokeService.ListByInterface(uint(debugInterfaceId), uint(endpointInterfaceId))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data})
}

// GetLastResp
// @Tags	接口调试
// @summary	获取最后调试记录响应
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization		header	string			true	"Authentication header"
// @Param 	currProjectId		query	int				true	"当前项目ID"
// @Param 	debugInterfaceId	query	int				true	"debugInterfaceId"
// @Param 	endpointInterfaceId	query	int				true	"endpointInterfaceId"
// @success	200	{object}	_domain.Response{data=domain.DebugResponse}
// @Router	/api/v1/debugs/invoke/getLastResp	[get]
func (c *DebugInvokeCtrl) GetLastResp(ctx iris.Context) {
	debugInterfaceId, err := ctx.URLParamInt("debugInterfaceId")
	endpointInterfaceId, err := ctx.URLParamInt("endpointInterfaceId")
	if debugInterfaceId <= 0 && endpointInterfaceId <= 0 {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	if debugInterfaceId < 0 {
		debugInterfaceId = 0
	}
	if endpointInterfaceId < 0 {
		endpointInterfaceId = 0
	}

	reqAndResp, err := c.DebugInvokeService.GetLastResp(uint(debugInterfaceId), uint(endpointInterfaceId))

	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
		return
	}
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: reqAndResp})
}

// GetResult 获取调用结果细节
func (c *DebugInvokeCtrl) GetResult(ctx iris.Context) {
	invokeId, err := ctx.URLParamInt("invokeId")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	result, err := c.DebugInvokeService.GetResult(invokeId)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: result})
}

// GetLog 获取调用日志
func (c *DebugInvokeCtrl) GetLog(ctx iris.Context) {
	invokeId, err := ctx.URLParamInt("invokeId")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	result, err := c.DebugInvokeService.GetLog(invokeId)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: result})
}

// GetAsInterface 详情
// @Tags	接口调试
// @summary	调试记录详情
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization	header	string			true	"Authentication header"
// @Param 	currProjectId	query	int				true	"当前项目ID"
// @Param 	id				path	int				true	"id"
// @success	200	{object}	_domain.Response{data=object{debugData=domain.DebugData,resp=domain.DebugResponse}}
// @Router	/api/v1/debugs/invoke/{id}	[get]
func (c *DebugInvokeCtrl) GetAsInterface(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	debugData, resultReq, resultResp, err := c.DebugInvokeService.GetAsInterface(id)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: iris.Map{
		"debugData": debugData,
		"req":       resultReq,
		"resp":      resultResp,
	}})
}

// Delete 删除
// @Tags	接口调试
// @summary	删除调试记录
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization	header	string			true	"Authentication header"
// @Param 	currProjectId	query	int				true	"当前项目ID"
// @Param 	id				path	int				true	"id"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/debugs/invoke/{id}	[delete]
func (c *DebugInvokeCtrl) Delete(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	err = c.DebugInvokeService.Delete(uint(id))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

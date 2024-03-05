package handler

import (
	"github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	agentExec "github.com/aaronchen2k/deeptest/internal/agent/exec"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/service"
	"github.com/aaronchen2k/deeptest/pkg/domain"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/kataras/iris/v12"
	"github.com/snowlyg/multi"
)

type ScenarioNodeCtrl struct {
	ScenarioNodeService *service.ScenarioNodeService `inject:""`
	ScenarioService     *service.ScenarioService     `inject:""`

	PerformanceTestPlanService *service.PerformanceTestPlanService `inject:""`

	BaseCtrl
}

// LoadTree
// @Tags	场景模块
// @summary	场景树状数据
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	scenarioId		query	int		true	"场景ID"
// @success	200	{object}	_domain.Response{data=agentExec.Processor}
// @Router	/api/v1/scenarios/load	[get]
func (c *ScenarioNodeCtrl) LoadTree(ctx iris.Context) {
	designFor := ctx.URLParam("designFor")
	scenarioOrPlanId, _ := ctx.URLParamInt("scenarioId")

	var scenarioId uint

	if designFor == consts.DesignForFunctionalTest.String() {
		scenarioId = uint(scenarioOrPlanId)
	} else if designFor == consts.DesignForPerformanceTest.String() {
		scenarioId, _ = c.PerformanceTestPlanService.GetScenarioId(scenarioOrPlanId)
	}

	scenario, err := c.ScenarioService.GetById(scenarioId)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	data, err := c.ScenarioNodeService.GetTree(scenario, false)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data})
}

// AddInterfacesFromDefine 添加接口
// @Tags	场景模块/编排节点
// @summary	添加定义接口
// @accept 	application/json
// @Produce application/json
// @Param	Authorization				header	string									true	"Authentication header"
// @Param 	currProjectId				query	int										true	"当前项目ID"
// @Param 	ScenarioAddInterfacesReq	body	serverDomain.ScenarioAddInterfacesReq	true	"添加定义接口的请求参数"
// @success	200	{object}	_domain.Response{data=model.Processor}
// @Router	/api/v1/scenarios/nodes/addInterfacesFromDefine	[post]
func (c *ScenarioNodeCtrl) AddInterfacesFromDefine(ctx iris.Context) {
	req := serverDomain.ScenarioAddInterfacesReq{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	req.CreateBy = multi.GetUserId(ctx)
	nodePo, bizErr := c.ScenarioNodeService.AddInterfacesFromDefine(req)
	if bizErr != nil {
		ctx.JSON(_domain.Response{
			Code: _domain.SystemErr.Code,
		})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: nodePo})
}

// AddInterfacesFromDiagnose 添加接口
// @Tags	场景模块/编排节点
// @summary	添加调试接口
// @accept 	application/json
// @Produce application/json
// @Param	Authorization						header	string											true	"Authentication header"
// @Param 	currProjectId						query	int												true	"当前项目ID"
// @Param 	ScenarioAddInterfacesFromTreeReq	body	serverDomain.ScenarioAddInterfacesFromTreeReq	true	"添加调试接口的请求参数"
// @success	200	{object}	_domain.Response{data=model.Processor}
// @Router	/api/v1/scenarios/nodes/addInterfacesFromTest	[post]
func (c *ScenarioNodeCtrl) AddInterfacesFromDiagnose(ctx iris.Context) {
	req := serverDomain.ScenarioAddInterfacesFromTreeReq{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	req.CreateBy = multi.GetUserId(ctx)
	nodePo, bizErr := c.ScenarioNodeService.AddInterfacesFromDiagnose(req)
	if bizErr != nil {
		ctx.JSON(_domain.Response{
			Code: _domain.SystemErr.Code,
		})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: nodePo})
}

// AddInterfacesFromCase 添加接口用例
// @Tags	场景模块/编排节点
// @summary	添加接口用例
// @accept 	application/json
// @Produce application/json
// @Param	Authorization						header	string										true	"Authentication header"
// @Param 	currProjectId						query	int											true	"当前项目ID"
// @Param 	ScenarioAddCasesFromTreeReq			body	serverDomain.ScenarioAddCasesFromTreeReq	true	"添加调接口用例的请求参数"
// @success	200	{object}	_domain.Response{data=model.Processor}
// @Router	/api/v1/scenarios/nodes/addInterfacesFromCase	[post]
func (c *ScenarioNodeCtrl) AddInterfacesFromCase(ctx iris.Context) {
	req := serverDomain.ScenarioAddCasesFromTreeReq{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	req.CreateBy = multi.GetUserId(ctx)
	nodePo, bizErr := c.ScenarioNodeService.AddInterfacesFromCase(req)
	if bizErr != nil {
		ctx.JSON(_domain.Response{
			Code: _domain.SystemErr.Code,
		})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: nodePo})
}

// AddProcessor 添加
// @Tags	场景模块/编排节点
// @summary	新建处理器
// @accept 	application/json
// @Produce application/json
// @Param	Authorization			header	string								true	"Authentication header"
// @Param 	currProjectId			query	int									true	"当前项目ID"
// @Param 	ScenarioAddScenarioReq	body	serverDomain.ScenarioAddScenarioReq	true	"新建处理器的请求参数"
// @success	200	{object}	_domain.Response{data=model.Processor}
// @Router	/api/v1/scenarios/nodes/addProcessor	[post]
func (c *ScenarioNodeCtrl) AddProcessor(ctx iris.Context) {
	projectId, err := ctx.URLParamInt("currProjectId")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	req := serverDomain.ScenarioAddScenarioReq{}
	err = ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	req.ProjectId = uint(projectId)
	req.CreateBy = multi.GetUserId(ctx)

	nodePo, bizErr := c.ScenarioNodeService.AddProcessor(req, "add")
	if bizErr != nil {
		ctx.JSON(_domain.Response{
			Code: _domain.SystemErr.Code,
		})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: nodePo})
}

// UpdateName 更新
// @Tags	场景模块/编排节点
// @summary	更新节点名称
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string							true	"Authentication header"
// @Param 	currProjectId	query	int								true	"当前项目ID"
// @Param 	id				path	int								true	"节点ID"
// @Param 	ScenarioNodeReq	body	serverDomain.ScenarioNodeReq	true	"更新节点名称的请求参数"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/scenarios/nodes/{id}/updateName	[put]
func (c *ScenarioNodeCtrl) UpdateName(ctx iris.Context) {
	var req serverDomain.ScenarioNodeReq
	err := ctx.ReadJSON(&req)
	if err != nil {
		logUtils.Errorf("参数验证失败", err.Error())
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	err = c.ScenarioNodeService.UpdateName(req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

// Delete 删除
// @Tags	场景模块/编排节点
// @summary	删除节点
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string							true	"Authentication header"
// @Param 	currProjectId	query	int								true	"当前项目ID"
// @Param 	id				path	int								true	"节点ID"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/scenarios/nodes/{id}	[delete]
func (c *ScenarioNodeCtrl) Delete(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	err = c.ScenarioNodeService.Delete(uint(id))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

// DisableOrNot 禁用
// @Tags	场景模块/编排节点
// @summary	禁用节点
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string							true	"Authentication header"
// @Param 	currProjectId	query	int								true	"当前项目ID"
// @Param 	id				path	int								true	"节点ID"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/scenarios/nodes/{id}/disableOrNot	[post]
func (c *ScenarioNodeCtrl) DisableOrNot(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	err = c.ScenarioNodeService.DisableOrNot(uint(id))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

// Move 移动
// @Tags	场景模块/编排节点
// @summary	移动节点
// @accept 	application/json
// @Produce application/json
// @Param	Authorization			header	string								true	"Authentication header"
// @Param 	currProjectId			query	int									true	"当前项目ID"
// @Param 	ScenarioNodeMoveReq		body	serverDomain.ScenarioNodeMoveReq	true	"移动节点的请求参数"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/scenarios/nodes/move	[put]
func (c *ScenarioNodeCtrl) Move(ctx iris.Context) {
	projectId, _ := ctx.URLParamInt("currProjectId")

	var req serverDomain.ScenarioNodeMoveReq
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	_, err = c.ScenarioNodeService.Move(uint(req.DragKey), uint(req.DropKey), req.DropPos, uint(projectId))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

// ImportCurl 导入cURL命令
// @Tags	场景模块/编排节点
// @summary	导入cURL命令
// @accept 	application/json
// @Produce application/json
// @Param	Authorization				header	string								true	"Authentication header"
// @Param 	DiagnoseCurlImportReq		body	serverDomain.ScenarioCurlImportReq	true	"导入cURL命令的请求体"
// @success	200	{object}	_domain.Response{data=model.Processor}
// @Router	/api/v1/scenarios/nodes/importCurl	[post]
func (c *ScenarioNodeCtrl) ImportCurl(ctx iris.Context) {
	req := serverDomain.ScenarioCurlImportReq{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	req.CreateBy = multi.GetUserId(ctx)
	newNode, bizErr := c.ScenarioNodeService.ImportCurl(req)
	if bizErr != nil {
		ctx.JSON(_domain.Response{
			Code: _domain.SystemErr.Code,
			Msg:  bizErr.Error(),
		})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: newNode})
}

func (c *ScenarioNodeCtrl) CopyProcessor(ctx iris.Context) {
	req := agentExec.Processor{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	var rootId uint
	createBy := multi.GetUserId(ctx)
	mod := "siblings"
	if req.EntityType == consts.ProcessorLogicElse {
		mod = "child"
	}

	bizErr := c.ScenarioNodeService.CopyProcessor(&req, createBy, mod, &rootId)
	if bizErr != nil {
		ctx.JSON(_domain.Response{
			Code: _domain.SystemErr.Code,
			Msg:  bizErr.Error(),
		})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: rootId})
}

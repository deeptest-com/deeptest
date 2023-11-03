package handler

import (
	"github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/core/web/validate"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/service"
	"github.com/aaronchen2k/deeptest/pkg/domain"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/snowlyg/multi"
	"strings"

	"github.com/kataras/iris/v12"
	"go.uber.org/zap"
)

type PlanCtrl struct {
	PlanService *service.PlanService `inject:""`
	BaseCtrl
}

// List
// @Tags	测试计划
// @summary	计划列表
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string							true	"Authentication header"
// @Param 	currProjectId	query	int								true	"当前项目ID"
// @Param 	PlanReqPaginate	query	serverDomain.PlanReqPaginate	true	"计划列表的请求参数"
// @success	200	{object}	_domain.Response{data=_domain.PageData{result=[]model.Plan}}
// @Router	/api/v1/plans	[get]
func (c *PlanCtrl) List(ctx iris.Context) {
	projectId, err := ctx.URLParamInt("currProjectId")
	if projectId == 0 {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	var req serverDomain.PlanReqPaginate
	err = ctx.ReadQuery(&req)
	if err != nil {
		errs := validate.ValidRequest(err)
		if len(errs) > 0 {
			logUtils.Errorf("参数验证失败", zap.String("错误", strings.Join(errs, ";")))
			ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: strings.Join(errs, ";")})
			return
		}
	}
	req.ConvertParams()
	req.Field = "updated_at"
	req.Order = "desc"
	data, err := c.PlanService.Paginate(req, projectId)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
}

// Get
// @Tags	测试计划
// @summary	计划详情
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	id				path	int		true	"计划ID"
// @Param 	detail			query	bool	true	"是否需要详情"
// @success	200	{object}	_domain.Response{data=model.Plan}
// @Router	/api/v1/plans/{id}	[get]
func (c *PlanCtrl) Get(ctx iris.Context) {
	var req _domain.ReqId
	if err := ctx.ReadParams(&req); err != nil {
		logUtils.Errorf("参数解析失败", zap.String("错误:", err.Error()))
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	detail, _ := ctx.URLParamBool("detail")

	plan, err := c.PlanService.GetById(req.Id, detail)

	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
		return
	}
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: plan, Msg: _domain.NoErr.Msg})
}

// Create
// @Tags	测试计划
// @summary	新建计划
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string		true	"Authentication header"
// @Param 	currProjectId	query	int			true	"当前项目ID"
// @Param 	Plan			body	model.Plan	true	"新建计划的请求参数"
// @success	200	{object}	_domain.Response{data=model.Plan}
// @Router	/api/v1/plans	[post]
func (c *PlanCtrl) Create(ctx iris.Context) {
	projectId, err := ctx.URLParamInt("currProjectId")
	if projectId == 0 {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: "projectId"})
		return
	}

	req := model.Plan{}
	err = ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	req.CreateUserId = multi.GetUserId(ctx)
	req.ProjectId = uint(projectId)
	req.Status = consts.Draft
	po, bizErr := c.PlanService.Create(req)
	if bizErr != nil {
		ctx.JSON(_domain.Response{Code: bizErr.Code, Data: nil})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: po, Msg: _domain.NoErr.Msg})
}

// Update
// @Tags	测试计划
// @summary	更新计划
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string		true	"Authentication header"
// @Param 	currProjectId	query	int			true	"当前项目ID"
// @Param 	Plan			body	model.Plan	true	"更新计划的请求参数"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/plans	[put]
func (c *PlanCtrl) Update(ctx iris.Context) {
	var req model.Plan
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	userId := multi.GetUserId(ctx)
	req.UpdateUserId = userId

	err = c.PlanService.Update(req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

// Delete
// @Tags	测试计划
// @summary	删除计划
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	id				path	int		true	"计划ID"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/plans/{id}	[delete]
func (c *PlanCtrl) Delete(ctx iris.Context) {
	var req _domain.ReqId
	err := ctx.ReadParams(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	err = c.PlanService.DeleteById(req.Id)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

// AddScenarios
// @Tags	测试计划
// @summary	添加场景
// @accept 	application/json
// @Produce application/json
// @Param	Authorization		header	string								true	"Authentication header"
// @Param 	currProjectId		query	int									true	"当前项目ID"
// @Param 	id					path	int									true	"计划ID"
// @Param 	PlanAddScenariosReq	body	serverDomain.PlanAddScenariosReq	true	"添加场景的请求参数"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/plans/{id}/addScenarios	[post]
func (c *PlanCtrl) AddScenarios(ctx iris.Context) {
	planId, _ := ctx.Params().GetInt("id")

	req := serverDomain.PlanAddScenariosReq{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	err = c.PlanService.AddScenarios(uint(planId), req.ScenarioIds)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code})
}

// RemoveScenario
// @Tags	测试计划
// @summary	移除场景
// @accept 	application/json
// @Produce application/json
// @Param	Authorization		header	string	true	"Authentication header"
// @Param 	currProjectId		query	int		true	"当前项目ID"
// @Param 	id					path	int		true	"计划ID"
// @Param 	scenarioId			query	int		true	"场景ID"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/plans/{id}/removeScenario	[post]
func (c *PlanCtrl) RemoveScenario(ctx iris.Context) {
	planId, _ := ctx.Params().GetInt("id")

	scenarioId, err := ctx.URLParamInt("scenarioId")

	err = c.PlanService.RemoveScenario(planId, scenarioId)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code})
}

// RemoveScenarios
// @Tags	测试计划
// @summary	批量移除场景
// @accept 	application/json
// @Produce application/json
// @Param	Authorization		header	string								true	"Authentication header"
// @Param 	currProjectId		query	int									true	"当前项目ID"
// @Param 	id					path	int									true	"计划ID"
// @Param 	PlanAddScenariosReq	body	serverDomain.PlanAddScenariosReq	true	"批量移除场景的请求参数"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/plans/{id}/removeScenarios	[post]
func (c *PlanCtrl) RemoveScenarios(ctx iris.Context) {
	planId, _ := ctx.Params().GetInt("id")

	req := serverDomain.PlanAddScenariosReq{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	err = c.PlanService.RemoveScenarios(planId, req.ScenarioIds)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code})
}

// StatusDropDownOptions
// @Tags	测试计划
// @summary	计划状态下拉选项
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @success	200	{object}	_domain.Response{data=map[consts.TestStatus]string}
// @Router	/api/v1/plans/statusDropDownOptions	[get]
func (c *PlanCtrl) StatusDropDownOptions(ctx iris.Context) {
	data := c.PlanService.StatusDropDownOptions()
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
}

// TestStageDropDownOptions
// @Tags	测试计划
// @summary	计划测试阶段下拉选项
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @success	200	{object}	_domain.Response{data=map[consts.TestStage]string}
// @Router	/api/v1/plans/testStageDropDownOptions	[get]
func (c *PlanCtrl) TestStageDropDownOptions(ctx iris.Context) {
	data := c.PlanService.TestStageDropDownOptions()
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
}

// Clone
// @Tags	测试计划
// @summary	克隆计划
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	id				path	int		true	"计划ID"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/plans/{id}/clone	[post]
func (c *PlanCtrl) Clone(ctx iris.Context) {
	userId := multi.GetUserId(ctx)

	var req _domain.ReqId
	if err := ctx.ReadParams(&req); err != nil {
		logUtils.Errorf("参数解析失败", zap.String("错误:", err.Error()))
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	plan, err := c.PlanService.Clone(req.Id, userId)

	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
		return
	}
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: plan, Msg: _domain.NoErr.Msg})
}

// PlanScenariosList
// @Tags	测试计划
// @summary	计划中的场景列表
// @accept 	application/json
// @Produce application/json
// @Param	Authorization				header	string									true	"Authentication header"
// @Param 	currProjectId				query	int										true	"当前项目ID"
// @Param 	planId						query	int										true	"计划ID"
// @Param 	PlanScenariosReqPaginate	query	serverDomain.PlanScenariosReqPaginate	true	"计划中的场景列表的请求参数"
// @success	200	{object}	_domain.Response{data=_domain.PageData{result=[]model.ScenarioDetail}}
// @Router	/api/v1/planScenariosList	[get]
func (c *PlanCtrl) PlanScenariosList(ctx iris.Context) {
	planId, err := ctx.URLParamInt("planId")
	if planId == 0 {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	var req serverDomain.PlanScenariosReqPaginate
	err = ctx.ReadQuery(&req)
	if err != nil {
		errs := validate.ValidRequest(err)
		if len(errs) > 0 {
			logUtils.Errorf("参数验证失败", zap.String("错误", strings.Join(errs, ";")))
			ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: strings.Join(errs, ";")})
			return
		}
	}
	req.ConvertParams()

	data, err := c.PlanService.PlanScenariosPaginate(req, uint(planId))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
}

// NotRelationScenarioList
// @Tags	测试计划
// @summary	计划中未绑定的场景列表
// @accept 	application/json
// @Produce application/json
// @Param	Authorization					header	string										true	"Authentication header"
// @Param 	currProjectId					query	int											true	"当前项目ID"
// @Param 	NotRelationScenarioReqPaginate	query	serverDomain.NotRelationScenarioReqPaginate	true	"计划中未绑定的场景列表的请求参数"
// @success	200	{object}	_domain.Response{data=_domain.PageData{result=[]model.Scenario}}
// @Router	/api/v1/notRelationScenarioList	[get]
func (c *PlanCtrl) NotRelationScenarioList(ctx iris.Context) {
	projectId, err := ctx.URLParamInt("currProjectId")
	if projectId == 0 {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	var req serverDomain.NotRelationScenarioReqPaginate
	err = ctx.ReadQuery(&req)
	if err != nil {
		errs := validate.ValidRequest(err)
		if len(errs) > 0 {
			logUtils.Errorf("参数验证失败", zap.String("错误", strings.Join(errs, ";")))
			ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: strings.Join(errs, ";")})
			return
		}
	}
	req.ConvertParams()

	data, err := c.PlanService.NotRelationScenarioList(req, projectId)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
}

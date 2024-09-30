package handler

import (
	"github.com/deeptest-com/deeptest/internal/agent/domain"
	"github.com/deeptest-com/deeptest/internal/server/modules/service"
	"github.com/deeptest-com/deeptest/pkg/domain"
	"github.com/kataras/iris/v12"
	"github.com/snowlyg/multi"
)

type PlanExecCtrl struct {
	PlanExecService *service.PlanExecService `inject:""`

	BaseCtrl
}

// LoadExecData
// @Tags	测试计划/执行计划
// @summary	加载执行计划
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	id				query	int		true	"计划ID"
// @Param 	environmentId	query	int		true	"环境ID"
// @success	200	{object}	_domain.Response{data=agentExec.PlanExecObj}
// @Router	/api/v1/plans/exec/loadExecPlan	[get]
func (c *PlanExecCtrl) LoadExecData(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	id, err := ctx.URLParamInt("id")
	environmentId, err := ctx.URLParamInt("environmentId")

	data, err := c.PlanExecService.LoadExecData(tenantId, id, environmentId)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
}

// LoadExecResult
// @Tags	测试计划/执行计划
// @summary	加载执行结果
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	planId			query	int		true	"计划ID"
// @success	200	{object}	_domain.Response{data=domain.Report}
// @Router	/api/v1/plans/exec/loadExecResult	[get]
func (c *PlanExecCtrl) LoadExecResult(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	scenarioId, err := ctx.URLParamInt("planId")

	data, err := c.PlanExecService.LoadExecResult(tenantId, scenarioId)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
}

// SubmitResult
// @Tags	测试计划/执行计划
// @summary	提交测试结果
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string						true	"Authentication header"
// @Param 	currProjectId	query	int							true	"当前项目ID"
// @Param 	id				path	int							true	"计划ID"
// @Param 	PlanExecResult	body	agentDomain.PlanExecResult	true	"提交计划测试结果的请求参数"
// @success	200	{object}	_domain.Response{data=model.PlanReport}
// @Router	/api/v1/plans/exec/submitResult/{id}	[post]
func (c *PlanExecCtrl) SubmitResult(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	planId, err := ctx.Params().GetInt("id")

	result := agentDomain.PlanExecResult{}
	err = ctx.ReadJSON(&result)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	userId := multi.GetUserId(ctx)
	report, err := c.PlanExecService.SaveReport(tenantId, planId, userId, result)

	// report.Logs = nil // otherwise will cause a json parse err on agent size
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: report})
}

// GetPlanReportNormalData
// @Tags	测试计划/执行计划
// @summary	获取计划执行中的静态内容
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	id				query	int		true	"计划ID"
// @Param 	environmentId	query	int		true	"环境ID"
// @success	200	{object}	_domain.Response{data=agentDomain.Report}
// @Router	/api/v1/plans/exec/getPlanReportNormalData	[get]
func (c *PlanExecCtrl) GetPlanReportNormalData(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	planId, err := ctx.URLParamInt("id")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	environmentId, err := ctx.URLParamInt("environmentId")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	data, err := c.PlanExecService.GetPlanReportNormalData(tenantId, uint(planId), uint(environmentId))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data})
}

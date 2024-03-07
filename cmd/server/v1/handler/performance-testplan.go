package handler

import (
	"github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/core/web/validate"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/service"
	"github.com/aaronchen2k/deeptest/pkg/domain"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/kataras/iris/v12"
	"github.com/snowlyg/multi"
	"go.uber.org/zap"
	"strings"
)

type PerformanceTestPlanCtrl struct {
	PerformanceTestPlanService *service.PerformanceTestPlanService `inject:""`
	BaseCtrl
}

// List
// @Tags	性能测试模块
// @summary	性能测试计划列表
// @accept 	application/json
// @Produce application/json
// @Param	Authorization		header	string								true	"Authentication header"
// @Param 	currProjectId		query	int									true	"当前项目ID"
// @Param 	PerformanceTestPlanReqPaginate	query	serverDomain.PerformanceTestPlanReqPaginate	true	"获取计划列表的请求对象"
// @success	200	{object}	_domain.Response{data=_domain.PageData{result=[]model.PerformanceTestPlan}}
// @Router	/api/v1/performanceTestPlans	[get]
func (c *PerformanceTestPlanCtrl) List(ctx iris.Context) {
	projectId, err := ctx.URLParamInt("currProjectId")
	if projectId == 0 {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	var req serverDomain.PerformanceTestPlanReqPaginate

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

	data, err := c.PerformanceTestPlanService.Paginate(req, projectId)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
}

// Get
// @Tags	性能测试模块
// @summary	性能测试计划详情
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	id				path	int		true	"计划ID"
// @success	200	{object}	_domain.Response{data=model.PerformanceTestPlan}
// @Router	/api/v1/performanceTestPlans/{id}	[get]
func (c *PerformanceTestPlanCtrl) Get(ctx iris.Context) {
	var req _domain.ReqId
	if err := ctx.ReadParams(&req); err != nil {
		logUtils.Errorf("参数解析失败", zap.String("错误:", err.Error()))
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}
	performanceTestPlan, err := c.PerformanceTestPlanService.GetById(req.Id)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
		return
	}
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: performanceTestPlan, Msg: _domain.NoErr.Msg})
}

func (c *PerformanceTestPlanCtrl) ListRunner(ctx iris.Context) {
	performanceScenarioId, err := ctx.URLParamInt("performanceScenarioId")
	if performanceScenarioId == 0 {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	runners, err := c.PerformanceTestPlanService.ListRunner(performanceScenarioId)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
		return
	}
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: runners, Msg: _domain.NoErr.Msg})
}

// Create
// @Tags	性能测试模块
// @summary	新建性能测试计划
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	PerformanceTestPlan				body	model.PerformanceTestPlan		true	"新建性能测试计划的请求参数"
// @success	200	{object}	_domain.Response{data=model.PerformanceTestPlan}
// @Router	/api/v1/performanceTestPlans	[post]
func (c *PerformanceTestPlanCtrl) Create(ctx iris.Context) {
	projectId, err := ctx.URLParamInt("currProjectId")
	if projectId == 0 {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	req := model.PerformanceTestPlan{}
	err = ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	req.ProjectId = uint(projectId)
	req.CreateUserId = multi.GetUserId(ctx)
	req.CreateUserName = multi.GetUsername(ctx)
	req.Status = consts.Draft

	po, err := c.PerformanceTestPlanService.Create(req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: po, Msg: _domain.NoErr.Msg})
}

// Update
// @Tags	性能测试模块
// @summary	更新性能测试计划
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	PerformanceTestPlan				body	model.PerformanceTestPlan		true	"更新性能测试计划的请求参数"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/performanceTestPlans	[put]
func (c *PerformanceTestPlanCtrl) Update(ctx iris.Context) {
	var req model.PerformanceTestPlan
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}
	req.UpdateUserName = multi.GetUsername(ctx)
	req.UpdateUserId = multi.GetUserId(ctx)
	err = c.PerformanceTestPlanService.Update(req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

// Delete
// @Tags	性能测试模块
// @summary	删除性能测试计划
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	id				path	int		true	"计划ID"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/performanceTestPlans/{id}	[delete]
func (c *PerformanceTestPlanCtrl) Delete(ctx iris.Context) {
	var req _domain.ReqId
	err := ctx.ReadParams(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	err = c.PerformanceTestPlanService.DeleteById(req.Id)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

// UpdateStatus
// @Tags	性能测试模块
// @summary	更新计划状态
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	id				path	int		true	"性能测试ID"
// @Param 	status			query	string	false	"性能测试状态"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/performanceTestPlans/{id}/updateStatus	[put]
func (c *PerformanceTestPlanCtrl) UpdateStatus(ctx iris.Context) {
	id, _ := ctx.Params().GetInt("id")
	status := ctx.URLParamDefault("status", "")
	if status == "" {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
	}

	updateUserId := multi.GetUserId(ctx)
	updateUserName := multi.GetUsername(ctx)
	err := c.PerformanceTestPlanService.UpdateStatus(uint(id), consts.TestStatus(status), updateUserId, updateUserName)
	if err == nil {
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
	} else {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
	}
}

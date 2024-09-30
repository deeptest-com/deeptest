package handler

import (
	"github.com/deeptest-com/deeptest/cmd/server/v1/domain"
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	"github.com/deeptest-com/deeptest/internal/server/core/web/validate"
	"github.com/deeptest-com/deeptest/internal/server/modules/model"
	"github.com/deeptest-com/deeptest/internal/server/modules/service"
	"github.com/deeptest-com/deeptest/pkg/domain"
	logUtils "github.com/deeptest-com/deeptest/pkg/lib/log"
	"github.com/kataras/iris/v12"
	"github.com/snowlyg/multi"
	"go.uber.org/zap"
	"strings"
)

type ScenarioCtrl struct {
	ScenarioService *service.ScenarioService `inject:""`
	BaseCtrl
}

// ListByProject
// @Tags	场景模块
// @summary	当前项目下的所有场景
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @success	200	{object}	_domain.Response{data=[]model.Scenario}
// @Router	/api/v1/scenarios/listByProject	[get]
func (c *ScenarioCtrl) ListByProject(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	projectId, err := ctx.URLParamInt("currProjectId")
	if projectId == 0 {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	res, err := c.ScenarioService.ListByProject(tenantId, projectId)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: res})

	return
}

// List
// @Tags	场景模块
// @summary	场景列表
// @accept 	application/json
// @Produce application/json
// @Param	Authorization		header	string								true	"Authentication header"
// @Param 	currProjectId		query	int									true	"当前项目ID"
// @Param 	ScenarioReqPaginate	query	serverDomain.ScenarioReqPaginate	true	"获取场景列表的请求参数"
// @success	200	{object}	_domain.Response{data=_domain.PageData{result=[]model.Scenario}}
// @Router	/api/v1/scenarios	[get]
func (c *ScenarioCtrl) List(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	projectId, err := ctx.URLParamInt("currProjectId")
	if projectId == 0 {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	var req serverDomain.ScenarioReqPaginate

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

	data, err := c.ScenarioService.Paginate(tenantId, req, projectId)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
}

// Get
// @Tags	场景模块
// @summary	场景详情
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	id				path	int		true	"场景ID"
// @success	200	{object}	_domain.Response{data=model.Scenario}
// @Router	/api/v1/scenarios/{id}	[get]
func (c *ScenarioCtrl) Get(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	var req _domain.ReqId
	if err := ctx.ReadParams(&req); err != nil {
		logUtils.Errorf("参数解析失败", zap.String("错误:", err.Error()))
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}
	scenario, err := c.ScenarioService.GetById(tenantId, req.Id)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
		return
	}
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: scenario, Msg: _domain.NoErr.Msg})
}

// Create
// @Tags	场景模块
// @summary	新建场景
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	Scenario				body	model.Scenario		true	"新建场景的请求参数"
// @success	200	{object}	_domain.Response{data=model.Scenario}
// @Router	/api/v1/scenarios	[post]
func (c *ScenarioCtrl) Create(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	projectId, err := ctx.URLParamInt("currProjectId")
	if projectId == 0 {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	req := model.Scenario{}
	err = ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	req.ProjectId = uint(projectId)
	req.CreateUserName = multi.GetUsername(ctx)
	req.CreateUserId = multi.GetUserId(ctx)
	req.Status = consts.Draft

	po, err := c.ScenarioService.Create(tenantId, req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: po, Msg: _domain.NoErr.Msg})
}

// Update
// @Tags	场景模块
// @summary	更新场景
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	Scenario				body	model.Scenario		true	"更新场景的请求参数"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/scenarios	[put]
func (c *ScenarioCtrl) Update(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	var req model.Scenario
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}
	req.UpdateUserName = multi.GetUsername(ctx)
	req.UpdateUserId = multi.GetUserId(ctx)
	err = c.ScenarioService.Update(tenantId, req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

// Delete
// @Tags	场景模块
// @summary	删除场景
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	id				path	int		true	"场景ID"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/scenarios/{id}	[delete]
func (c *ScenarioCtrl) Delete(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	var req _domain.ReqId
	err := ctx.ReadParams(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	err = c.ScenarioService.DeleteById(tenantId, req.Id)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

// AddPlans
// @Tags	场景模块
// @summary	关联计划
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	id				path	int		true	"场景ID"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/scenarios/{id}/addPlans	[post]
func (c *ScenarioCtrl) AddPlans(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	scenarioId, _ := ctx.Params().GetInt("id")

	planIds := make([]int, 0)
	err := ctx.ReadJSON(&planIds)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: "ids"})
		return
	}

	err = c.ScenarioService.AddPlans(tenantId, scenarioId, planIds)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code})
}

// Plans
// @Tags	场景模块
// @summary	关联计划列表
// @accept 	application/json
// @Produce application/json
// @Param	Authorization			header	string									true	"Authentication header"
// @Param 	currProjectId			query	int										true	"当前项目ID"
// @Param 	id						path	int										true	"场景ID"
// @Param 	ScenarioPlanReqPaginate	body	serverDomain.ScenarioPlanReqPaginate	true	"查询关联计划列表的参数"
// @success	200	{object}	_domain.Response{data=_domain.PageData{result=[]model.Plan}}
// @Router	/api/v1/scenarios/{id}/plans	[post]
func (c *ScenarioCtrl) Plans(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	projectId, err := ctx.URLParamInt("currProjectId")
	if projectId == 0 {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	scenarioId, _ := ctx.Params().GetInt("id")

	var req serverDomain.ScenarioPlanReqPaginate
	err = ctx.ReadJSON(&req)
	if err != nil {
		errs := validate.ValidRequest(err)
		if len(errs) > 0 {
			logUtils.Errorf("参数验证失败", zap.String("错误", strings.Join(errs, ";")))
			ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: strings.Join(errs, ";")})
			return
		}
	}

	req.ProjectId = uint(projectId)
	data, err := c.ScenarioService.PlanPaginate(tenantId, req, scenarioId)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})

}

// UpdateStatus
// @Tags	场景模块
// @summary	更新计划状态
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	id				path	int		true	"场景ID"
// @Param 	status			query	string	false	"场景状态"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/scenarios/{id}/updateStatus	[put]
func (c *ScenarioCtrl) UpdateStatus(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	id, _ := ctx.Params().GetInt("id")
	status := ctx.URLParamDefault("status", "")
	if status == "" {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
	}

	updateUserId := multi.GetUserId(ctx)
	updateUserName := multi.GetUsername(ctx)
	err := c.ScenarioService.UpdateStatus(tenantId, uint(id), consts.TestStatus(status), updateUserId, updateUserName)
	if err == nil {
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
	} else {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
	}
}

// UpdatePriority
// @Tags	场景模块
// @summary	更新优先级
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	id				path	int		true	"场景ID"
// @Param 	priority		query	string	false	"场景优先级"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/scenarios/{id}/updatePriority	[put]
func (c *ScenarioCtrl) UpdatePriority(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	id, _ := ctx.Params().GetInt("id")
	priority := ctx.URLParamDefault("priority", "")
	if priority == "" {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
	}

	updateUserId := multi.GetUserId(ctx)
	updateUserName := multi.GetUsername(ctx)
	err := c.ScenarioService.UpdatePriority(tenantId, uint(id), priority, updateUserId, updateUserName)
	if err == nil {
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
	} else {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
	}
}

// RemovePlans
// @Tags	场景模块
// @summary	取消计划关联
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	id				path	int		true	"场景ID"
// @Param 	""				body	[]int	true	"计划ID"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/scenarios/{id}/removePlans	[post]
func (c *ScenarioCtrl) RemovePlans(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	scenarioId, _ := ctx.Params().GetInt("id")

	planIds := make([]int, 0)
	err := ctx.ReadJSON(&planIds)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: "ids"})
		return
	}

	err = c.ScenarioService.RemovePlans(tenantId, scenarioId, planIds)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code})
}

package handler

import (
	"github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/core/web/validate"
	"github.com/aaronchen2k/deeptest/internal/server/modules/service"
	"github.com/aaronchen2k/deeptest/pkg/domain"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"strings"

	"github.com/kataras/iris/v12"
	"go.uber.org/zap"
)

type ScenarioReportCtrl struct {
	ScenarioReportService *service.ScenarioReportService `inject:""`
	BaseCtrl
}

// List
// @Tags	场景模块/场景报告
// @summary	结果列表
// @accept 	application/json
// @Produce application/json
// @Param	Authorization		header	string							true	"Authentication header"
// @Param 	currProjectId		query	int								true	"当前项目ID"
// @Param 	ReportReqPaginate	body	serverDomain.ReportReqPaginate	true	"获取结果列表的请求参数"
// @success	200	{object}	_domain.Response{data=_domain.PageData{result=[]model.ScenarioReport}}
// @Router	/api/v1/scenarios/reports	[post]
func (c *ScenarioReportCtrl) List(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	/*
		projectId, err := ctx.URLParamInt("currProjectId")
		if projectId == 0 {
			ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
			return
		}
	*/

	var req serverDomain.ReportReqPaginate
	if err := ctx.ReadJSON(&req); err != nil {
		errs := validate.ValidRequest(err)
		if len(errs) > 0 {
			logUtils.Errorf("参数验证失败", zap.String("错误", strings.Join(errs, ";")))
			ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: strings.Join(errs, ";")})
			return
		}
	}

	data, err := c.ScenarioReportService.Paginate(tenantId, req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
}

// Get
// @Tags	场景模块/场景报告
// @summary	结果详情
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	id				path	int		true	"报告ID"
// @success	200	{object}	_domain.Response{data=model.ScenarioReport}
// @Router	/api/v1/scenarios/reports/{id}	[get]
func (c *ScenarioReportCtrl) Get(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	var req _domain.ReqId
	if err := ctx.ReadParams(&req); err != nil {
		logUtils.Errorf("参数解析失败", zap.String("错误:", err.Error()))
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}
	report, err := c.ScenarioReportService.GetById(tenantId, req.Id)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
		return
	}
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: report, Msg: _domain.NoErr.Msg})
}

// Delete
// @Tags	场景模块/场景报告
// @summary	删除报告
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	id				path	int		true	"报告ID"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/scenarios/reports/{id}	[delete]
func (c *ScenarioReportCtrl) Delete(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	var req _domain.ReqId
	err := ctx.ReadParams(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	err = c.ScenarioReportService.DeleteById(tenantId, req.Id)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

// Create
// @Tags	场景模块/场景报告
// @summary	创建报告
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	id				path	int		true	"报告ID"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/scenarios/reports/{id}	[put]
func (c *ScenarioReportCtrl) Create(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	var req _domain.ReqId
	err := ctx.ReadParams(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	err = c.ScenarioReportService.CreatePlanReport(tenantId, req.Id)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

func (c *ScenarioReportCtrl) ReferBug(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	var req serverDomain.ReferBugReq
	if err := ctx.ReadJSON(&req); err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	err := c.ScenarioReportService.ReferBug(tenantId, req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

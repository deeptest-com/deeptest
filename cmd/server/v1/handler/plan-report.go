package handler

import (
	"github.com/deeptest-com/deeptest/cmd/server/v1/domain"
	integrationService "github.com/deeptest-com/deeptest/integration/service"
	"github.com/deeptest-com/deeptest/internal/server/core/web/validate"
	"github.com/deeptest-com/deeptest/internal/server/modules/service"
	"github.com/deeptest-com/deeptest/pkg/domain"
	logUtils "github.com/deeptest-com/deeptest/pkg/lib/log"
	"github.com/kataras/iris/v12"
	"go.uber.org/zap"
	"strings"
)

type PlanReportCtrl struct {
	ReportService            *service.PlanReportService        `inject:""`
	IntegrationReportService *integrationService.ReportService `inject:""`
	BaseCtrl
}

// List
// @Tags	测试报告
// @summary	结果列表
// @accept 	application/json
// @Produce application/json
// @Param	Authorization			header	string								true	"Authentication header"
// @Param 	currProjectId			query	int									true	"当前项目ID"
// @Param 	PlanReportReqPaginate	query	serverDomain.PlanReportReqPaginate	true	"测试报告列表的请求参数"
// @success	200	{object}	_domain.Response{data=_domain.PageData{result=[]model.PlanReportDetail}}
// @Router	/api/v1/plans/reports	[get]
func (c *PlanReportCtrl) List(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	projectId, err := ctx.URLParamInt("currProjectId")
	if projectId == 0 {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	var req serverDomain.PlanReportReqPaginate
	if err := ctx.ReadQuery(&req); err != nil {
		errs := validate.ValidRequest(err)
		if len(errs) > 0 {
			logUtils.Errorf("参数验证失败", zap.String("错误", strings.Join(errs, ";")))
			ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: strings.Join(errs, ";")})
			return
		}
	}
	req.ConvertParams()

	data, err := c.ReportService.Paginate(tenantId, req, projectId)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
}

// Get
// @Tags	测试报告
// @summary	结果详情
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	id				path	int		true	"测试报告ID"
// @success	200	{object}	_domain.Response{data=model.PlanReportDetail}
// @Router	/api/v1/plans/reports/{id}	[get]
func (c *PlanReportCtrl) Get(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	var req _domain.ReqId
	if err := ctx.ReadParams(&req); err != nil {
		logUtils.Errorf("参数解析失败", zap.String("错误:", err.Error()))
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}
	report, err := c.ReportService.GetById(tenantId, req.Id)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
		return
	}
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: report, Msg: _domain.NoErr.Msg})
}

// Delete
// @Tags	测试报告
// @summary	删除计划报告
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	id				path	int		true	"测试报告ID"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/plans/reports/{id}	[delete]
func (c *PlanReportCtrl) Delete(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	var req _domain.ReqId
	err := ctx.ReadParams(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	err = c.ReportService.DeleteById(tenantId, req.Id)

	if ctx.Method() == "DELETE" {
		//同步第三方
		go c.IntegrationReportService.DeleteReport(tenantId, req.Id)
	}

	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

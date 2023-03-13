package handler

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/service"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/kataras/iris/v12"
	"go.uber.org/zap"
)

type SummaryCtrl struct {
	SummaryDetailsService            *service.SummaryDetailsService            `inject:""`
	SummaryBugsService               *service.SummaryBugsService               `inject:""`
	SummaryProjectUserRankingService *service.SummaryProjectUserRankingService `inject:""`
	BaseCtrl
}

func (c *SummaryCtrl) Card(ctx iris.Context) {

	//now := time.Now()
	//req := v1.SummaryBugsReq{}
	//req.BugId = "2"
	//req.ProjectId = 1
	//req.BugClassify = "dfs"
	//req.BugCreateDate = now.Format("2006/01/02 15:04:05")
	//req.BugSeverity = "blocker"
	//req.BugState = "open"
	//req.ProductName = "ProductName"
	//req.ProductId = "Productid"
	//
	//var projectId v1.ReqProjectId
	//if err := ctx.ReadParams(&projectId); err != nil {
	//	logUtils.Errorf("参数解析失败", zap.String("错误:", err.Error()))
	//	ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Data: nil, Msg: _domain.ParamErr.Msg})
	//	return
	//}
	//
	//if projectId.ProjectId == 0 {
	//	data, err := c.SummaryBugsService.FindGroupByBugSeverity()
	//	if err != nil {
	//		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: err.Error()})
	//		return
	//	} else {
	//		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
	//		return
	//	}
	//} else if projectId.ProjectId == 1 {
	//	data, err := c.SummaryBugsService.CountByProjectId(req)
	//	if err != nil {
	//		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: err.Error()})
	//		return
	//	} else {
	//		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
	//		return
	//	}
	//} else if projectId.ProjectId == 2 {
	//	err := c.SummaryBugsService.Create(req)
	//	if err != nil {
	//		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: err.Error()})
	//		return
	//	} else {
	//		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: nil, Msg: _domain.NoErr.Msg})
	//		return
	//	}
	//} else if projectId.ProjectId == 3 {
	//	data, err := c.SummaryBugsService.Count()
	//	if err != nil {
	//		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: err.Error()})
	//		return
	//	} else {
	//		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
	//		return
	//	}
	//} else if projectId.ProjectId == 4 {
	//	data, err := c.SummaryBugsService.FindByProjectIdGroupByBugSeverity(req)
	//	if err != nil {
	//		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: err.Error()})
	//		return
	//	} else {
	//		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
	//		return
	//	}
	//}

	//projectId, err := ctx.URLParamInt("currProjectId")
	//if projectId == 0 {
	//	ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: "projectId"})
	//	return
	//}
	//
	//var req v1.ReportReqPaginate
	//if err := ctx.ReadQuery(&req); err != nil {
	//	errs := validate.ValidRequest(err)
	//	if len(errs) > 0 {
	//		logUtils.Errorf("参数验证失败", zap.String("错误", strings.Join(errs, ";")))
	//		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: strings.Join(errs, ";")})
	//		return
	//	}
	//}
	//req.ConvertParams()

	//data, err := c.ReportService.Paginate(req, projectId)
	//if err != nil {
	//	ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: err.Error()})
	//	return
	//}

	//ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
}

func (c *SummaryCtrl) Bugs(ctx iris.Context) {

	var projectId int64
	projectId, err := ctx.Params().GetInt64("projectId")

	if err != nil {
		logUtils.Errorf("参数解析失败", zap.String("错误:", err.Error()))
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Data: nil, Msg: _domain.ParamErr.Msg})
		return
	}

	data, err := c.SummaryBugsService.Bugs(projectId)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: err.Error()})
		return
	} else {
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
		return
	}
}

func (c *SummaryCtrl) Details(ctx iris.Context) {
	//var req _domain.ReqId
	//if err := ctx.ReadParams(&req); err != nil {
	//	logUtils.Errorf("参数解析失败", zap.String("错误:", err.Error()))
	//	ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Data: nil, Msg: _domain.ParamErr.Msg})
	//	return
	//}
	//report, err := c.ReportService.GetById(req.Id)
	//if err != nil {
	//	ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: _domain.SystemErr.Msg})
	//	return
	//}
	//ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: report, Msg: _domain.NoErr.Msg})
}

func (c *SummaryCtrl) ProjectUserRanking(ctx iris.Context) {
	//var req _domain.ReqId
	//err := ctx.ReadParams(&req)
	//if err != nil {
	//	ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: err.Error()})
	//	return
	//}

	//err = c.ReportService.DeleteById(req.Id)
	//if err != nil {
	//	ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: err.Error()})
	//	return
	//}
	//
	//ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: nil, Msg: _domain.NoErr.Msg})
}

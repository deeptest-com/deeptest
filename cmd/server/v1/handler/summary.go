package handler

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/service"
	"github.com/kataras/iris/v12"
)

type SummaryCtrl struct {
	SummaryDetailsService            *service.SummaryDetailsService            `inject:""`
	SummaryBugsService               *service.SummaryBugsService               `inject:""`
	SummaryProjectUserRankingService *service.SummaryProjectUserRankingService `inject:""`
	BaseCtrl
}

func (c *SummaryCtrl) Card(ctx iris.Context) {
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

	//var projectId v1.ReqProjectId
	//if err := ctx.ReadParams(&projectId); err != nil {
	//	logUtils.Errorf("参数解析失败", zap.String("错误:", err.Error()))
	//	ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Data: nil, Msg: _domain.ParamErr.Msg})
	//	return
	//}
	//
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
	//
	//data, err := c.ReportService.Paginate(req, projectId)
	//if err != nil {
	//	ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: err.Error()})
	//	return
	//}
	//
	//ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
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

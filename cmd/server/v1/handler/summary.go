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

	var projectId int64
	projectId, err := ctx.Params().GetInt64("projectId")

	if err != nil {
		logUtils.Errorf("参数解析失败", zap.String("错误:", err.Error()))
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Data: nil, Msg: _domain.ParamErr.Msg})
		return
	}

	data, err := c.SummaryDetailsService.Card(projectId)

	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: err.Error()})
		return
	} else {
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
		return
	}
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

	var userId int64
	userId, err := ctx.Params().GetInt64("userId")

	if err != nil {
		logUtils.Errorf("参数解析失败", zap.String("错误:", err.Error()))
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Data: nil, Msg: _domain.ParamErr.Msg})
		return
	}

	data, err := c.SummaryDetailsService.Details(userId)

	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: err.Error()})
		return
	} else {
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
		return
	}
}

func (c *SummaryCtrl) ProjectUserRanking(ctx iris.Context) {

	var projectId int64
	projectId, err := ctx.Params().GetInt64("projectId")

	if err != nil {
		logUtils.Errorf("参数解析失败", zap.String("错误:", err.Error()))
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Data: nil, Msg: _domain.ParamErr.Msg})
		return
	}

	data, err := c.SummaryProjectUserRankingService.ProjectUserRanking(projectId)

	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: err.Error()})
		return
	} else {
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
		return
	}
}

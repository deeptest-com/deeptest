package handler

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/service"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/kataras/iris/v12"
	"github.com/snowlyg/multi"
	"go.uber.org/zap"
)

type SummaryCtrl struct {
	SummaryService *service.SummaryService `inject:""`
	BaseCtrl
}

func (c *SummaryCtrl) Summary() {
	c.SummaryService.Collection()
	return
}

func (c *SummaryCtrl) Card(ctx iris.Context) {
	projectId, err := ctx.Params().GetInt64("projectId")

	if err != nil {
		logUtils.Errorf("参数解析失败", zap.String("错误:", err.Error()))
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	data, err := c.SummaryService.Card(projectId)

	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	} else {
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
		return
	}
}

func (c *SummaryCtrl) Bugs(ctx iris.Context) {
	projectId, err := ctx.Params().GetInt64("projectId")

	if err != nil {
		logUtils.Errorf("参数解析失败", zap.String("错误:", err.Error()))
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	data, err := c.SummaryService.Bugs(projectId)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	} else {
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
		return
	}
}

func (c *SummaryCtrl) Details(ctx iris.Context) {

	userId := multi.GetUserId(ctx)

	if userId == 0 {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: "userId为0，或获取userId失败"})
		return
	}

	data, err := c.SummaryService.Details(int64(userId))

	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	} else {
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
		return
	}
}

func (c *SummaryCtrl) ProjectUserRanking(ctx iris.Context) {

	cycle, err := ctx.Params().GetInt64("cycle")
	projectId, err := ctx.Params().GetInt64("projectId")

	if err != nil {
		logUtils.Errorf("参数解析失败", zap.String("错误:", err.Error()))
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	data, err := c.SummaryService.ProjectUserRanking(projectId, cycle)

	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	} else {
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
		return
	}
}

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

func (c *SummaryCtrl) Summary(ctx iris.Context) {
	//SAAS
	tenantId := c.getTenantId(ctx)
	c.SummaryService.Collection(tenantId)
	return
}

// Card
// @Tags	汇总
// @summary	汇总卡片位信息
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	projectId		path	int		true	"项目ID"
// @success	200	{object}	_domain.Response{data=serverDomain.ResSummaryCard}
// @Router	/api/v1/summary/card/{projectId}	[get]
func (c *SummaryCtrl) Card(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	projectId, err := ctx.Params().GetInt64("projectId")

	if err != nil {
		logUtils.Errorf("参数解析失败", zap.String("错误:", err.Error()))
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	data, err := c.SummaryService.Card(tenantId, projectId)

	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	} else {
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
		return
	}
}

// Bugs
// @Tags	汇总
// @summary	汇总bug信息
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	projectId		path	int		true	"项目ID"
// @success	200	{object}	_domain.Response{data=serverDomain.ResSummaryBugs}
// @Router	/api/v1/summary/bugs/{projectId}	[get]
func (c *SummaryCtrl) Bugs(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	projectId, err := ctx.Params().GetInt64("projectId")

	if err != nil {
		logUtils.Errorf("参数解析失败", zap.String("错误:", err.Error()))
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	data, err := c.SummaryService.Bugs(tenantId, projectId)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	} else {
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
		return
	}
}

// Details
// @Tags	汇总
// @summary	汇总项目详情
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @success	200	{object}	_domain.Response{data=serverDomain.ResSummaryDetail}
// @Router	/api/v1/summary/details	[get]
func (c *SummaryCtrl) Details(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	userId := multi.GetUserId(ctx)

	if userId == 0 {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: "userId为0，或获取userId失败"})
		return
	}

	data, err := c.SummaryService.Details(tenantId, int64(userId))

	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	} else {
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
		return
	}
}

// ProjectUserRanking
// @Tags	汇总
// @summary	汇总项目用户排行数据
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	cycle			path	int		true	"cycle"
// @Param 	projectId		path	int		true	"项目ID"
// @success	200	{object}	_domain.Response{data=serverDomain.ResRankingList}
// @Router	/api/v1/summary/projectUserRanking/{cycle}/{projectId}	[get]
func (c *SummaryCtrl) ProjectUserRanking(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	cycle, err := ctx.Params().GetInt64("cycle")
	projectId, err := ctx.Params().GetInt64("projectId")

	if err != nil {
		logUtils.Errorf("参数解析失败", zap.String("错误:", err.Error()))
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	data, err := c.SummaryService.ProjectUserRanking(tenantId, cycle, projectId)

	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	} else {
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
		return
	}
}

// Collection
// @Tags	汇总
// @summary	汇总数据
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	store			path	string	true	"string"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/summary/collection/{store}	[get]
func (c *SummaryCtrl) Collection(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	name := ctx.Params().GetString("store")
	switch name {
	case "details":
		c.SummaryService.CollectionDetails(tenantId)
	case "ranking":
		c.SummaryService.CollectionRanking(tenantId)
	case "bugs":
		c.SummaryService.CollectionBugs(tenantId)

	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: "", Msg: _domain.NoErr.Msg})
	return
}

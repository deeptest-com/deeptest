package handler

import (
	serverDomain "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/service"
	"github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/kataras/iris/v12"
)

type ExtractorCtrl struct {
	ExtractorService *service.ExtractorService `inject:""`
	BaseCtrl
}

// Get 		详情
// @Tags	提取器
// @summary	提取器详情
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	id				path	int		true	"提取器ID"
// @success	200	{object}	_domain.Response{data=model.DebugConditionExtractor}
// @Router	/api/v1/extractors/{id}	[get]
func (c *ExtractorCtrl) Get(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	extractor, err := c.ExtractorService.Get(tenantId, uint(id))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
		return
	}
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: extractor})
}

// QuickCreate 添加
// @Tags	提取器
// @summary	新建提取器
// @accept 	application/json
// @Produce application/json
// @Param	Authorization						header	string											true	"Authentication header"
// @Param 	currProjectId						query	int												true	"当前项目ID"
// @Param 	ExtractorConditionQuickCreateReq	body	serverDomain.ExtractorConditionQuickCreateReq	true	"新建提取器的请求参数"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/extractors	[post]
func (c *ExtractorCtrl) QuickCreate(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	var req serverDomain.ExtractorConditionQuickCreateReq
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	err = c.ExtractorService.QuickCreate(tenantId, req, consts.InterfaceDebug)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

// Update 更新
// @Tags	提取器
// @summary	更新提取器
// @accept 	application/json
// @Produce application/json
// @Param	Authorization			header	string							true	"Authentication header"
// @Param 	currProjectId			query	int								true	"当前项目ID"
// @Param 	DebugConditionExtractor	body	model.DebugConditionExtractor	true	"更新提取器的请求参数"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/extractors	[put]
func (c *ExtractorCtrl) Update(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	var extractor model.DebugConditionExtractor
	err := ctx.ReadJSON(&extractor)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	err = c.ExtractorService.Update(tenantId, &extractor)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

// ListExtractorVariableForCheckpoint
// @Tags	提取器
// @summary	提取器变量列表
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string				true	"Authentication header"
// @Param 	currProjectId	query	int					true	"当前项目ID"
// @Param 	DebugInfo		body	domain.DebugInfo	true	"提取器变量列表的请求参数"
// @success	200	{object}	_domain.Response{data=[]domain.Variable}
// @Router	/api/v1/extractors/listExtractorVariableForCheckpoint	[post]
func (c *ExtractorCtrl) ListExtractorVariableForCheckpoint(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	req := domain.DebugInfo{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	data, err := c.ExtractorService.ListExtractorVariableByInterface(tenantId, req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data})
}

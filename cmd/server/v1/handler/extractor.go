package handler

import (
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

// List
// @Tags	提取器
// @summary	提取器列表
// @accept 	application/json
// @Produce application/json
// @Param	Authorization		header	string	true	"Authentication header"
// @Param 	currProjectId		query	int		true	"当前项目ID"
// @Param 	debugInterfaceId	query	int		true	"debugInterfaceId"
// @Param 	endpointInterfaceId	query	int		true	"endpointInterfaceId"
// @Param 	usedBy				query	string	true	"usedBy"
// @success	200	{object}	_domain.Response{data=_domain.PageData{result=[]model.DebugInterfaceExtractor}}
// @Router	/api/v1/extractors	[get]
func (c *ExtractorCtrl) List(ctx iris.Context) {
	debugInterfaceId, err := ctx.URLParamInt("debugInterfaceId")
	endpointInterfaceId, err := ctx.URLParamInt("endpointInterfaceId")
	usedBy := ctx.URLParam("usedBy")

	if debugInterfaceId <= 0 && endpointInterfaceId <= 0 {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	if debugInterfaceId < 0 {
		debugInterfaceId = 0
	}
	if endpointInterfaceId < 0 {
		endpointInterfaceId = 0
	}

	data, err := c.ExtractorService.List(uint(debugInterfaceId), uint(endpointInterfaceId), consts.UsedBy(usedBy))

	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ret := []interface{}{}
	for _, item := range data {
		ret = append(ret, item)
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: ret})
}

// Get 详情
// @Tags	提取器
// @summary	提取器详情
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	id				path	int		true	"提取器ID"
// @success	200	{object}	_domain.Response{data=model.DebugInterfaceExtractor}
// @Router	/api/v1/extractors/{id}	[get]
func (c *ExtractorCtrl) Get(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	extractor, err := c.ExtractorService.Get(uint(id))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
		return
	}
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: extractor})
}

// Create 添加
// @Tags	提取器
// @summary	新建提取器
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	DebugInterfaceExtractor				body	model.DebugInterfaceExtractor		true	"新建提取器的请求参数"
// @success	200	{object}	_domain.Response{data=model.DebugInterfaceExtractor}
// @Router	/api/v1/extractors	[post]
func (c *ExtractorCtrl) Create(ctx iris.Context) {
	extractor := model.DebugInterfaceExtractor{}
	err := ctx.ReadJSON(&extractor)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	bizErr := c.ExtractorService.Create(&extractor)
	if bizErr.Code > 0 {
		ctx.JSON(_domain.Response{
			Code: bizErr.Code,
		})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: extractor, Msg: _domain.NoErr.Msg})
}

// CreateOrUpdateResult 新建或更新结果
// @Tags	提取器
// @summary	新建或更新提取器
// @accept 	application/json
// @Produce application/json
// @Param	Authorization			header	string							true	"Authentication header"
// @Param 	currProjectId			query	int								true	"当前项目ID"
// @Param 	DebugInterfaceExtractor	body	model.DebugInterfaceExtractor	true	"新建或更新提取器的请求参数"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/extractors/createOrUpdateResult	[post]
func (c *ExtractorCtrl) CreateOrUpdateResult(ctx iris.Context) {
	var extractor model.DebugInterfaceExtractor
	err := ctx.ReadJSON(&extractor)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	err = c.ExtractorService.CreateOrUpdateResult(&extractor, consts.InterfaceDebug)
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
// @Param 	DebugInterfaceExtractor	body	model.DebugInterfaceExtractor	true	"更新提取器的请求参数"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/extractors	[put]
func (c *ExtractorCtrl) Update(ctx iris.Context) {
	var extractor model.DebugInterfaceExtractor
	err := ctx.ReadJSON(&extractor)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	err = c.ExtractorService.Update(&extractor)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

// Delete 删除
// @Tags	提取器
// @summary	删除提取器
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	id				path	int		true	"提取器ID"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/extractors/{id}	[delete]
func (c *ExtractorCtrl) Delete(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	err = c.ExtractorService.Delete(uint(id))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

// ListExtractorVariableForCheckpoint
// @Tags	提取器
// @summary	提取器变量列表
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string			true	"Authentication header"
// @Param 	currProjectId	query	int				true	"当前项目ID"
// @Param 	DebugReq		body	domain.DebugReq	true	"提取器变量列表的请求参数"
// @success	200	{object}	_domain.Response{data=[]domain.Variable}
// @Router	/api/v1/extractors/listExtractorVariableForCheckpoint	[post]
func (c *ExtractorCtrl) ListExtractorVariableForCheckpoint(ctx iris.Context) {
	req := domain.DebugReq{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	data, err := c.ExtractorService.ListExtractorVariableByInterface(req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data})
}

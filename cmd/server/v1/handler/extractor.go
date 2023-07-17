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
func (c *ExtractorCtrl) List(ctx iris.Context) {
	debugInterfaceId, err := ctx.URLParamInt("debugInterfaceId")
	endpointInterfaceId, err := ctx.URLParamInt("endpointInterfaceId")
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

	data, err := c.ExtractorService.List(uint(debugInterfaceId), uint(endpointInterfaceId))

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

// CreateOrUpdateResult 新建或更新结果
func (c *ExtractorCtrl) CreateOrUpdateResult(ctx iris.Context) {
	var extractor model.DebugConditionExtractor
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
func (c *ExtractorCtrl) Update(ctx iris.Context) {
	var extractor model.DebugConditionExtractor
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

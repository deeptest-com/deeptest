package controller

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/service"
	"github.com/aaronchen2k/deeptest/pkg/domain"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"

	"github.com/kataras/iris/v12"
)

type ScenarioProcessorCtrl struct {
	ScenarioProcessorService *service.ScenarioProcessorService `inject:""`
	BaseCtrl
}

// Get 详情
func (c *ScenarioProcessorCtrl) Get(ctx iris.Context) {
	processorId, err := ctx.Params().GetInt("id")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Data: nil, Msg: _domain.ParamErr.Msg})
		return
	}

	processor, err := c.ScenarioProcessorService.Get(processorId)

	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: _domain.SystemErr.Msg})
		return
	}
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: processor})
}

// UpdateName 更新
func (c *ScenarioProcessorCtrl) UpdateName(ctx iris.Context) {
	var req model.ProcessorEntity
	err := ctx.ReadJSON(&req)
	if err != nil {
		logUtils.Errorf("参数验证失败", err.Error())
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Data: nil, Msg: err.Error()})
		return
	}

	err = c.ScenarioProcessorService.UpdateName(req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: nil, Msg: _domain.NoErr.Msg})
}

// Save 保存
func (c *ScenarioProcessorCtrl) Save(ctx iris.Context) {
	processorCategory := ctx.Params().Get("category")

	var err error

	if processorCategory == consts.ProcessorGroup.ToString() {
		var entity model.ProcessorGroup
		err = ctx.ReadJSON(&entity)
		err = c.ScenarioProcessorService.SaveGroup(entity)

	} else if processorCategory == consts.ProcessorLogic.ToString() {
		var entity model.ProcessorLogic
		err = ctx.ReadJSON(&entity)
		err = c.ScenarioProcessorService.SaveLogic(entity)

	} else if processorCategory == consts.ProcessorTimer.ToString() {
		var entity model.ProcessorTimer
		err = ctx.ReadJSON(&entity)
		err = c.ScenarioProcessorService.SaveTimer(entity)

	}

	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: nil, Msg: _domain.NoErr.Msg})
}

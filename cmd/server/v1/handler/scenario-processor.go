package handler

import (
	domain "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	agentExec "github.com/aaronchen2k/deeptest/internal/agent/exec"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/service"
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
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	processorEntity, err := c.ScenarioProcessorService.GetEntity(processorId)

	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
		return
	}
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: processorEntity})
}

// UpdateName 更新
func (c *ScenarioProcessorCtrl) UpdateName(ctx iris.Context) {
	var req agentExec.ProcessorEntityBase
	err := ctx.ReadJSON(&req)
	if err != nil {
		logUtils.Errorf("参数验证失败", err.Error())
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	err = c.ScenarioProcessorService.UpdateName(req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

// SaveProcessorInfo 更新
func (c *ScenarioProcessorCtrl) SaveProcessorInfo(ctx iris.Context) {
	var req domain.ScenarioProcessorInfo
	err := ctx.ReadJSON(&req)
	if err != nil {
		logUtils.Errorf("参数验证失败", err.Error())
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	err = c.ScenarioProcessorService.SaveProcessorInfo(req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

// Save 保存
func (c *ScenarioProcessorCtrl) Save(ctx iris.Context) {
	processorCategoryString := ctx.Params().Get("category")
	processorCategory := consts.ProcessorCategory(processorCategoryString)

	var err error
	var po interface{}

	if processorCategory == consts.ProcessorGroup {
		var entity model.ProcessorGroup
		err = ctx.ReadJSON(&entity)
		err = c.ScenarioProcessorService.SaveGroup(&entity)
		po = entity

	} else if processorCategory == consts.ProcessorLogic {
		var entity model.ProcessorLogic
		err = ctx.ReadJSON(&entity)
		err = c.ScenarioProcessorService.SaveLogic(&entity)
		po = entity

	} else if processorCategory == consts.ProcessorLoop {
		var entity model.ProcessorLoop
		err = ctx.ReadJSON(&entity)
		err = c.ScenarioProcessorService.SaveLoop(&entity)
		po = entity

	} else if processorCategory == consts.ProcessorTimer {
		var entity model.ProcessorTimer
		err = ctx.ReadJSON(&entity)
		err = c.ScenarioProcessorService.SaveTimer(&entity)
		po = entity

	} else if processorCategory == consts.ProcessorPrint {
		var entity model.ProcessorPrint
		err = ctx.ReadJSON(&entity)
		err = c.ScenarioProcessorService.SavePrint(&entity)
		po = entity

	} else if processorCategory == consts.ProcessorVariable {
		var entity model.ProcessorVariable
		err = ctx.ReadJSON(&entity)
		err = c.ScenarioProcessorService.SaveVariable(&entity)
		po = entity

	} else if processorCategory == consts.ProcessorCookie {
		var entity model.ProcessorCookie
		err = ctx.ReadJSON(&entity)
		err = c.ScenarioProcessorService.SaveCookie(&entity)
		po = entity

	} else if processorCategory == consts.ProcessorAssertion {
		var entity model.ProcessorAssertion
		err = ctx.ReadJSON(&entity)
		err = c.ScenarioProcessorService.SaveAssertion(&entity)
		po = entity

	} else if processorCategory == consts.ProcessorExtractor {
		var entity model.ProcessorExtractor
		err = ctx.ReadJSON(&entity)
		err = c.ScenarioProcessorService.SaveExtractor(&entity)
		po = entity

	} else if processorCategory == consts.ProcessorData {
		var entity model.ProcessorData
		err = ctx.ReadJSON(&entity)
		err = c.ScenarioProcessorService.SaveData(&entity)
		po = entity

	}

	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: po, Msg: _domain.NoErr.Msg})
}

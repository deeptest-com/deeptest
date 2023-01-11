package handler

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/service"
	"github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/kataras/iris/v12"
)

type CheckpointCtrl struct {
	CheckpointService *service.CheckpointService `inject:""`
	BaseCtrl
}

// List
func (c *CheckpointCtrl) List(ctx iris.Context) {
	interfaceId, err := ctx.URLParamInt("interfaceId")
	usedBy := ctx.URLParam("usedBy")
	if interfaceId == 0 || usedBy == "" {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: "interfaceId OR usedBy"})
		return
	}

	data, err := c.CheckpointService.List(interfaceId, consts.UsedBy(usedBy))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data})
}

// Get 详情
func (c *CheckpointCtrl) Get(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Data: nil, Msg: _domain.ParamErr.Msg})
		return
	}

	checkpoint, err := c.CheckpointService.Get(uint(id))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: _domain.SystemErr.Msg})
		return
	}
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: checkpoint})
}

// Create 添加
func (c *CheckpointCtrl) Create(ctx iris.Context) {
	checkpoint := model.InterfaceCheckpoint{}
	err := ctx.ReadJSON(&checkpoint)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	err = c.CheckpointService.Create(&checkpoint)
	if err != nil {
		ctx.JSON(_domain.Response{
			Code: _domain.SystemErr.Code,
			Data: nil,
		})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: checkpoint, Msg: _domain.NoErr.Msg})
}

// Update 更新
func (c *CheckpointCtrl) Update(ctx iris.Context) {
	var checkpoint model.InterfaceCheckpoint
	err := ctx.ReadJSON(&checkpoint)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	err = c.CheckpointService.Update(&checkpoint)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: err.Error()})
		return
	}
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: nil, Msg: _domain.NoErr.Msg})
}

// Delete 删除
func (c *CheckpointCtrl) Delete(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	err = c.CheckpointService.Delete(uint(id))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

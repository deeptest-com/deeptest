package handler

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/service"
	"github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/kataras/iris/v12"
)

type EnvironmentCtrl struct {
	EnvironmentService *service.EnvironmentService `inject:""`
	BaseCtrl
}

// List
func (c *EnvironmentCtrl) List(ctx iris.Context) {
	data, err := c.EnvironmentService.List()
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data})
}

// Get 详情
func (c *EnvironmentCtrl) Get(ctx iris.Context) {
	projectId, err := ctx.URLParamInt("currProjectId")
	id, err := ctx.Params().GetInt("id")

	if id <= 0 && projectId <= 0 {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Data: nil, Msg: _domain.ParamErr.Msg})
		return
	}

	env, err := c.EnvironmentService.Get(uint(id), uint(projectId))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: _domain.SystemErr.Msg})
		return
	}
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: env})
}

// Change 修改
func (c *EnvironmentCtrl) Change(ctx iris.Context) {
	projectId, err := ctx.URLParamInt("currProjectId")
	id, err := ctx.URLParamInt("id")

	if projectId == 0 || id == 0 {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: ""})
		return
	}

	err = c.EnvironmentService.Change(id, projectId)
	if err != nil {
		ctx.JSON(_domain.Response{
			Code: _domain.SystemErr.Code,
			Data: nil,
		})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

// Create 添加
func (c *EnvironmentCtrl) Create(ctx iris.Context) {
	projectId, err := ctx.URLParamInt("currProjectId")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	env := model.Environment{}
	err = ctx.ReadJSON(&env)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	err = c.EnvironmentService.Create(&env, uint(projectId))
	if err != nil {
		ctx.JSON(_domain.Response{
			Code: _domain.SystemErr.Code,
			Data: nil,
		})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: env, Msg: _domain.NoErr.Msg})
}

// Update 更新
func (c *EnvironmentCtrl) Update(ctx iris.Context) {
	var env model.Environment
	err := ctx.ReadJSON(&env)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	err = c.EnvironmentService.Update(&env)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: err.Error()})
		return
	}
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: nil, Msg: _domain.NoErr.Msg})
}

// Copy 添加
func (c *EnvironmentCtrl) Copy(ctx iris.Context) {
	id, err := ctx.URLParamInt("id")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	err = c.EnvironmentService.Copy(id)
	if err != nil {
		ctx.JSON(_domain.Response{
			Code: _domain.SystemErr.Code,
			Data: nil,
		})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

// Delete 删除
func (c *EnvironmentCtrl) Delete(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	err = c.EnvironmentService.Delete(uint(id))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

// CreateVar 添加
func (c *EnvironmentCtrl) CreateVar(ctx iris.Context) {
	envVar := model.EnvironmentVar{}
	err := ctx.ReadJSON(&envVar)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	err = c.EnvironmentService.CreateVar(&envVar)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ErrNameExist.Code, Data: nil})
		return
	}

	env, err := c.EnvironmentService.Get(envVar.EnvironmentId, 0)

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: env, Msg: _domain.NoErr.Msg})
}

// UpdateVar 更新
func (c *EnvironmentCtrl) UpdateVar(ctx iris.Context) {
	var envVar model.EnvironmentVar
	err := ctx.ReadJSON(&envVar)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	err = c.EnvironmentService.UpdateVar(&envVar)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: err.Error()})
		return
	}

	env, err := c.EnvironmentService.Get(envVar.EnvironmentId, 0)
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: env, Msg: _domain.NoErr.Msg})
}

// DeleteVar 删除
func (c *EnvironmentCtrl) DeleteVar(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	err = c.EnvironmentService.DeleteVar(uint(id))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	envVar, err := c.EnvironmentService.GetVar(uint(id))
	env, err := c.EnvironmentService.Get(envVar.EnvironmentId, 0)
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: env, Msg: _domain.NoErr.Msg})
}

// ClearVar 清除
func (c *EnvironmentCtrl) ClearVar(ctx iris.Context) {
	environmentId, err := ctx.URLParamInt("environmentId")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	err = c.EnvironmentService.ClearAllVar(uint(environmentId))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	env, err := c.EnvironmentService.Get(uint(environmentId), 0)
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: env, Msg: _domain.NoErr.Msg})
}

// DeleteShareVar 删除
func (c *EnvironmentCtrl) DeleteShareVar(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	err = c.EnvironmentService.DisableShareVar(uint(id))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	envVar, err := c.EnvironmentService.GetVar(uint(id))
	env, err := c.EnvironmentService.Get(envVar.EnvironmentId, 0)
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: env, Msg: _domain.NoErr.Msg})
}

// ClearShareVar 清除
func (c *EnvironmentCtrl) ClearShareVar(ctx iris.Context) {
	interfaceId, err := ctx.URLParamInt("interfaceId")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	err = c.EnvironmentService.DisableAllShareVar(uint(interfaceId))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	env, err := c.EnvironmentService.Get(uint(interfaceId), 0)
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: env, Msg: _domain.NoErr.Msg})
}

package handler

import (
	"github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
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
	projectId, err := ctx.URLParamInt("currProjectId")
	data, err := c.EnvironmentService.List(projectId)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data})
}

// Get 详情
func (c *EnvironmentCtrl) Get(ctx iris.Context) {
	projectId, err := ctx.URLParamInt("currProjectId")
	id, err := ctx.Params().GetInt("id")

	if id <= 0 && projectId <= 0 {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	env, err := c.EnvironmentService.Get(uint(id), uint(projectId))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
		return
	}
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: env})
}

// Change 修改
func (c *EnvironmentCtrl) Change(ctx iris.Context) {
	projectId, err := ctx.URLParamInt("currProjectId")
	id, err := ctx.URLParamInt("id")

	if projectId == 0 || id == 0 {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	err = c.EnvironmentService.Change(id, projectId)
	if err != nil {
		ctx.JSON(_domain.Response{
			Code: _domain.SystemErr.Code,
		})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

// Create 添加
func (c *EnvironmentCtrl) Create(ctx iris.Context) {
	projectId, err := ctx.URLParamInt("currProjectId")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	env := model.Environment{}
	err = ctx.ReadJSON(&env)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	err = c.EnvironmentService.Create(&env, uint(projectId))
	if err != nil {
		ctx.JSON(_domain.Response{
			Code: _domain.SystemErr.Code,
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
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	err = c.EnvironmentService.Update(&env)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

// Copy 添加
func (c *EnvironmentCtrl) Copy(ctx iris.Context) {
	id, err := ctx.URLParamInt("id")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	err = c.EnvironmentService.Copy(id)
	if err != nil {
		ctx.JSON(_domain.Response{
			Code: _domain.SystemErr.Code,
		})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

// Delete 删除
func (c *EnvironmentCtrl) Delete(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
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
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
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
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	err = c.EnvironmentService.UpdateVar(&envVar)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	env, err := c.EnvironmentService.Get(envVar.EnvironmentId, 0)
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: env, Msg: _domain.NoErr.Msg})
}

// DeleteVar 删除
func (c *EnvironmentCtrl) DeleteVar(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
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
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
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
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
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

func (c *EnvironmentCtrl) Save(ctx iris.Context) {
	var req serverDomain.EnvironmentReq
	if err := ctx.ReadJSON(&req); err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	if id, err := c.EnvironmentService.Save(req); err == nil {
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: id, Msg: _domain.NoErr.Msg})
	} else {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
	}
}

func (c *EnvironmentCtrl) Clone(ctx iris.Context) {
	id := ctx.URLParamIntDefault("id", 0)
	if id != 0 {
		if env, err := c.EnvironmentService.Clone(uint(id)); err != nil {
			ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		} else {
			ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: env.ID, Msg: _domain.NoErr.Msg})
		}
	}
}

func (c *EnvironmentCtrl) DeleteEnvironment(ctx iris.Context) {
	id := ctx.URLParamIntDefault("id", 0)
	if id != 0 {
		if err := c.EnvironmentService.DeleteEnvironment(uint(id)); err != nil {
			ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
			return
		}
	}
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

func (c *EnvironmentCtrl) ListAll(ctx iris.Context) {
	projectId := ctx.URLParamIntDefault("projectId", 0)
	if res, err := c.EnvironmentService.ListAll(uint(projectId)); err == nil {
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: res, Msg: _domain.NoErr.Msg})
	} else {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
	}
}

func (c *EnvironmentCtrl) SaveGlobal(ctx iris.Context) {
	var req []serverDomain.EnvironmentVariable
	projectId := ctx.URLParamIntDefault("currProjectId", 0)

	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	err = c.EnvironmentService.SaveGlobal(uint(projectId), req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

func (c *EnvironmentCtrl) ListGlobal(ctx iris.Context) {
	projectId := ctx.URLParamIntDefault("projectId", 0)
	res, err := c.EnvironmentService.ListGlobal(uint(projectId))
	if err == nil {
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: res, Msg: _domain.NoErr.Msg})
	} else {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
	}
}

func (c *EnvironmentCtrl) SaveParams(ctx iris.Context) {
	var req serverDomain.EnvironmentParamsReq
	if err := ctx.ReadJSON(&req); err == nil {
		if err = c.EnvironmentService.SaveParams(req); err == nil {
			ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
		} else {
			ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		}
	} else {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
	}
}

func (c *EnvironmentCtrl) ListParams(ctx iris.Context) {
	projectId := ctx.URLParamIntDefault("projectId", 0)
	res, err := c.EnvironmentService.ListParams(uint(projectId))
	if err == nil {
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: res, Msg: _domain.NoErr.Msg})
	} else {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
	}
}

func (c *EnvironmentCtrl) Order(ctx iris.Context) {
	var req serverDomain.EnvironmentIdsReq
	if err := ctx.ReadJSON(&req); err == nil {
		if err = c.EnvironmentService.SaveOrder(req); err == nil {
			ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
		} else {
			ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		}
	} else {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
	}
}

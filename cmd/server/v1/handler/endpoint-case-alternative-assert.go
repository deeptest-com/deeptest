package handler

import (
	"github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/service"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/kataras/iris/v12"
)

type EndpointCaseAlternativeAssertCtrl struct {
	EndpointCaseAlternativeAssertService *service.EndpointCaseAlternativeAssertService `inject:""`
	EndpointCaseAlternativeService       *service.EndpointCaseAlternativeService       `inject:""`
}

// List
func (c *EndpointCaseAlternativeAssertCtrl) List(ctx iris.Context) {
	alternativeCaseId, err := ctx.URLParamInt("alternativeCaseId")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	root, err := c.EndpointCaseAlternativeAssertService.List(uint(alternativeCaseId))

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: root})
}

// Save
func (c *EndpointCaseAlternativeAssertCtrl) Save(ctx iris.Context) {
	var req model.EndpointCaseAlternativeAssert
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	err = c.EndpointCaseAlternativeAssertService.Save(&req)

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: req})
}

// Delete 删除
func (c *EndpointCaseAlternativeAssertCtrl) Delete(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	err = c.EndpointCaseAlternativeAssertService.Delete(uint(id))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

// Disable 禁用
func (c *EndpointCaseAlternativeAssertCtrl) Disable(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	err = c.EndpointCaseAlternativeAssertService.Disable(uint(id))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

// Move 移动
func (c *EndpointCaseAlternativeAssertCtrl) Move(ctx iris.Context) {
	var req serverDomain.ConditionMoveReq
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	err = c.EndpointCaseAlternativeAssertService.Move(req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

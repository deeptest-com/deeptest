package handler

import (
	"github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/service"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/kataras/iris/v12"
	"github.com/snowlyg/multi"
)

type EndpointCaseAlternativeCtrl struct {
	EndpointCaseAlternativeService *service.EndpointCaseAlternativeService `inject:""`
	DebugInterfaceService          *service.DebugInterfaceService          `inject:""`
}

// LoadAlternative
func (c *EndpointCaseAlternativeCtrl) LoadAlternative(ctx iris.Context) {
	baseId, err := ctx.URLParamInt("baseId")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	root, err := c.EndpointCaseAlternativeService.LoadAlternative(uint(baseId))

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: root})
}

// LoadAlternativeSaved
func (c *EndpointCaseAlternativeCtrl) LoadAlternativeSaved(ctx iris.Context) {
	baseId, err := ctx.URLParamInt("baseId")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	ret, err := c.EndpointCaseAlternativeService.LoadAlternativeSaved(uint(baseId))

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: ret})
}

// SaveAlternativeCase
func (c *EndpointCaseAlternativeCtrl) SaveAlternativeCase(ctx iris.Context) {
	var req serverDomain.EndpointCaseAlternativeSaveReq
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	req.CreateUserId = multi.GetUserId(ctx)
	req.CreateUserName = multi.GetUsername(ctx)

	ret, err := c.EndpointCaseAlternativeService.SaveAlternativeCase(req)

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: ret})
}

// GenerateCases
//func (c *EndpointCaseAlternativeCtrl) GenerateCases(ctx iris.Context) {
//	req := serverDomain.EndpointCaseAlternativeGenerateReq{}
//	err := ctx.ReadJSON(&req)
//	if err != nil {
//		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
//		return
//	}
//
//	req.CreateUserName = multi.GetUsername(ctx)
//	req.CreateUserId = multi.GetUserId(ctx)
//
//	err = c.EndpointCaseAlternativeService.GenerateFromSpec(req)
//
//	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code})
//}

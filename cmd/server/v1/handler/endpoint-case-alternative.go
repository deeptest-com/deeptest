package handler

import (
	"github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	agentExec "github.com/aaronchen2k/deeptest/internal/agent/exec"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
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
	endpointId, err := ctx.URLParamInt("endpointId")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}
	method := ctx.URLParam("method")

	root, err := c.EndpointCaseAlternativeService.LoadAlternative(uint(endpointId), consts.HttpMethod(method))

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

// CreateBenchmark
func (c *EndpointCaseAlternativeCtrl) CreateBenchmark(ctx iris.Context) {
	var req serverDomain.EndpointCaseBenchmarkCreateReq
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	req.CreateUserId = multi.GetUserId(ctx)
	req.CreateUserName = multi.GetUsername(ctx)

	ret, err := c.EndpointCaseAlternativeService.CreateBenchmarkCase(req)

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: ret})
}

// SaveAlternative
func (c *EndpointCaseAlternativeCtrl) SaveAlternative(ctx iris.Context) {
	var req serverDomain.EndpointCaseAlternativeSaveReq
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	req.CreateUserId = multi.GetUserId(ctx)
	req.CreateUserName = multi.GetUsername(ctx)

	ret, err := c.EndpointCaseAlternativeService.SaveAlternative(req)

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: ret})
}

// LoadCaseForExec
func (c *EndpointCaseAlternativeCtrl) LoadCaseForExec(ctx iris.Context) {
	var req agentExec.CasesExecObj
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	req.UserId = multi.GetUserId(ctx)

	ret, err := c.EndpointCaseAlternativeService.LoadCaseForExec(req)

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: ret})
}

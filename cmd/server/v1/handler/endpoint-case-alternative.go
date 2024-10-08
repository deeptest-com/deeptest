package handler

import (
	"github.com/deeptest-com/deeptest/cmd/server/v1/domain"
	agentExec "github.com/deeptest-com/deeptest/internal/agent/exec"
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	"github.com/deeptest-com/deeptest/internal/server/modules/service"
	_domain "github.com/deeptest-com/deeptest/pkg/domain"
	"github.com/kataras/iris/v12"
	"github.com/snowlyg/multi"
)

type EndpointCaseAlternativeCtrl struct {
	BaseCtrl
	EndpointCaseAlternativeService *service.EndpointCaseAlternativeService `inject:""`
	DebugInterfaceService          *service.DebugInterfaceService          `inject:""`
}

// LoadAlternative
func (c *EndpointCaseAlternativeCtrl) LoadAlternative(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	endpointId, err := ctx.URLParamInt("endpointId")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}
	method := ctx.URLParam("method")

	root, err := c.EndpointCaseAlternativeService.LoadAlternative(tenantId, uint(endpointId), consts.HttpMethod(method))

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: root})
}

// LoadFactor
func (c *EndpointCaseAlternativeCtrl) LoadFactor(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	caseId, err := ctx.URLParamInt("caseId")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	ret, err := c.EndpointCaseAlternativeService.LoadFactor(tenantId, uint(caseId))

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: ret})
}

// CreateBenchmark
func (c *EndpointCaseAlternativeCtrl) CreateBenchmark(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	var req serverDomain.EndpointCaseBenchmarkCreateReq
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	req.CreateUserId = multi.GetUserId(ctx)
	req.CreateUserName = multi.GetUsername(ctx)

	ret, err := c.EndpointCaseAlternativeService.CreateBenchmarkCase(tenantId, req)

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: ret})
}

// SaveFactor
func (c *EndpointCaseAlternativeCtrl) SaveFactor(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	var req serverDomain.EndpointCaseFactorSaveReq
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	err = c.EndpointCaseAlternativeService.SaveFactor(tenantId, req)

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code})
}

// SaveCase
func (c *EndpointCaseAlternativeCtrl) SaveCase(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	var req serverDomain.EndpointCaseAlternativeSaveReq
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	req.CreateUserId = multi.GetUserId(ctx)
	req.CreateUserName = multi.GetUsername(ctx)

	ret, err := c.EndpointCaseAlternativeService.SaveCase(tenantId, req)

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: ret})
}

// LoadCasesForExec
func (c *EndpointCaseAlternativeCtrl) LoadCasesForExec(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	var req agentExec.CasesExecReq

	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	req.UserId = multi.GetUserId(ctx)

	ret, err := c.EndpointCaseAlternativeService.LoadCasesForExec(tenantId, req)

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: ret})
}

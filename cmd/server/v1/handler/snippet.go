package handler

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	scriptHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/script"
	service "github.com/aaronchen2k/deeptest/internal/server/modules/service"
	"github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/kataras/iris/v12"
	"github.com/snowlyg/multi"
	"time"
)

type SnippetCtrl struct {
	SnippetService        *service.SnippetService        `inject:""`
	DebugInterfaceService *service.DebugInterfaceService `inject:""`
	BaseCtrl
}

// Get 详情
// @Tags	脚本
// @summary	获取详情
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	name			query	string	true	"脚本名"
// @success	200	{object}	_domain.Response{data=model.Snippet}
// @Router	/api/v1/snippets	[get]
func (c *SnippetCtrl) Get(ctx iris.Context) {
	name := ctx.URLParam("name")

	snippet, err := c.SnippetService.Get(scriptHelper.ScriptType(name))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: snippet})
}

func (c *SnippetCtrl) ListJslibNames(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	projectId, err := ctx.URLParamInt("currProjectId")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	snippets, err := c.SnippetService.ListJslibNames(tenantId, projectId)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: snippets})
}

func (c *SnippetCtrl) GetJslibs(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	projectId, err := ctx.URLParamInt("currProjectId")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	snippets, err := c.SnippetService.GetJslibs(tenantId, projectId)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: snippets})
}

func (c *SnippetCtrl) GetJslibsForAgent(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	projectId, err := ctx.URLParamInt("projectId")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	agentLoadedLibs := map[uint]time.Time{}
	err = ctx.ReadJSON(&agentLoadedLibs)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	snippets, err := c.SnippetService.GetJslibsForAgent(tenantId, agentLoadedLibs, projectId)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: snippets})
}

func (c *SnippetCtrl) ListVar(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	req := domain.DebugInfo{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	req.UserId = multi.GetUserId(ctx)
	data := c.SnippetService.ListVar(tenantId, req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
}

func (c *SnippetCtrl) ListMock(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	list := c.SnippetService.ListMock(tenantId)
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: list, Msg: _domain.NoErr.Msg})
}

func (c *SnippetCtrl) ListSysFunc(ctx iris.Context) {
	list := c.SnippetService.ListSysFunc()
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: list, Msg: _domain.NoErr.Msg})
}

func (c *SnippetCtrl) ListCustomFunc(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	projectId, err := ctx.URLParamInt("currProjectId")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}
	list := c.SnippetService.ListCustomFunc(tenantId, uint(projectId))
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: list, Msg: _domain.NoErr.Msg})
}

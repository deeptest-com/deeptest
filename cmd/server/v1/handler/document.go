package handler

import (
	domain "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/service"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/kataras/iris/v12"
)

type DocumentCtrl struct {
	DocumentService *service.DocumentService `inject:""`
}

func (c *DocumentCtrl) Index(ctx iris.Context) {
	var req domain.DocumentReq
	if err := ctx.ReadJSON(&req); err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	res, _ := c.DocumentService.Content(req)
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: res})

	return
}

func (c *DocumentCtrl) DocumentVersionList(ctx iris.Context) {
	projectId, _ := ctx.URLParamInt("currProjectId")

	var req domain.DocumentVersionListReq
	if err := ctx.ReadJSON(&req); err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	res, err := c.DocumentService.GetDocumentVersionList(uint(projectId), req.NeedLatest)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg, Data: res})
	return
}

func (c *DocumentCtrl) Publish(ctx iris.Context) {
	projectId, _ := ctx.URLParamInt("currProjectId")

	var req domain.DocumentVersionReq
	if err := ctx.ReadJSON(&req); err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	err := c.DocumentService.Publish(req, uint(projectId))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
	return
}

func (c *DocumentCtrl) DeleteSnapshot(ctx iris.Context) {
	id := ctx.URLParamUint64("id")
	err := c.DocumentService.RemoveSnapshot(uint(id))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
	return
}

func (c *DocumentCtrl) UpdateDocument(ctx iris.Context) {
	var req domain.UpdateDocumentVersionReq
	if err := ctx.ReadJSON(&req); err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	err := c.DocumentService.UpdateDocument(req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
	return
}

func (c *DocumentCtrl) GetShareLink(ctx iris.Context) {
	var req domain.DocumentShareReq
	if err := ctx.ReadJSON(&req); err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	code, err := c.DocumentService.GenerateShareLink(req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}
	res := iris.Map{"code": code}
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg, Data: res})
	return
}

func (c *DocumentCtrl) GetContentsByShareLink(ctx iris.Context) {
	link := ctx.URLParam("code")
	if link == "" {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: "code can't be empty"})
		return
	}

	res, err := c.DocumentService.ContentByShare(link)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: res})

	return
}

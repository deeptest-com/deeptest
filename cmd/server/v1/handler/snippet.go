package handler

import (
	service "github.com/aaronchen2k/deeptest/internal/server/modules/service"
	"github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/kataras/iris/v12"
)

type SnippetCtrl struct {
	SnippetService *service.SnippetService `inject:""`

	BaseCtrl
}

// Get 详情
func (c *SnippetCtrl) Get(ctx iris.Context) {
	name := ctx.URLParam("name")

	snippet, err := c.SnippetService.Get(name)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: _domain.SystemErr.Msg})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: snippet})
}

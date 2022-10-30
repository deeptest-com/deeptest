package handler

import (
	"github.com/aaronchen2k/deeptest/internal/agent/service"
	"github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/kataras/iris/v12"
)

type SpecCtrl struct {
	SpecService *service.SpecService `inject:""`
}

// Load 解析定义文件
func (c *SpecCtrl) Load(ctx iris.Context) {
	pth := ctx.URLParam("path")
	typ := ctx.URLParam("type")

	content, err := c.SpecService.Load(pth, typ)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: err.Error()})
		return
	}

	ret := iris.Map{"content": content}
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: ret, Msg: _domain.NoErr.Msg})
}

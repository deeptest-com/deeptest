package handler

import (
	"github.com/aaronchen2k/deeptest/cmd/agent/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/agent/service"
	"github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/kataras/iris/v12"
)

type SpecCtrl struct {
	SpecService *service.SpecService `inject:""`
}

// Load 解析定义文件
func (c *SpecCtrl) Load(ctx iris.Context) {
	req := domain.SubmitSpecReq{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: err.Error()})
		return
	}

	doc3, info, err := c.SpecService.SubmitSpec(req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: err.Error()})
		return
	}

	ret := iris.Map{"doc": doc3, "info": info}
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: ret, Msg: _domain.NoErr.Msg})
}

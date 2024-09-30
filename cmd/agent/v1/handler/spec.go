package handler

import (
	agentDomain "github.com/deeptest-com/deeptest/cmd/agent/v1/domain"
	agentService "github.com/deeptest-com/deeptest/internal/agent/service"
	_domain "github.com/deeptest-com/deeptest/pkg/domain"
	"github.com/kataras/iris/v12"
)

type SpecCtrl struct {
}

// Parse 解析定义文件
func (c *SpecCtrl) Parse(ctx iris.Context) {
	req := agentDomain.ParseSpecReq{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	err = agentService.ParseSpec(req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

package handler

import (
	"github.com/aaronchen2k/deeptest/cmd/agent/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/agent/service"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/kataras/iris/v12"
)

type ExecInterfaceCtrl struct {
}

// Call
func (c *ExecInterfaceCtrl) Call(ctx iris.Context) {
	req := agentDomain.InterfaceCall{}
	req.TenantId = c.getTenantId(ctx)
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	resultReq, resultResp, err := service.RunInterface(req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: iris.Map{
		"req":  resultReq,
		"resp": resultResp,
	}})
}

func (c *ExecInterfaceCtrl) getTenantId(ctx iris.Context) consts.TenantId {
	return "123"
}

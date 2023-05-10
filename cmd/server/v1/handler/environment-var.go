package handler

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/service"
	"github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/kataras/iris/v12"
)

type EnvironmentVarCtrl struct {
	EnvironmentService *service.EnvironmentService `inject:""`
	BaseCtrl
}

// List
func (c *EnvironmentVarCtrl) List(ctx iris.Context) {
	serverId, err := ctx.URLParamInt("serverId")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	data, _ := c.EnvironmentService.GetVarsByServer(uint(serverId))

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data})
}

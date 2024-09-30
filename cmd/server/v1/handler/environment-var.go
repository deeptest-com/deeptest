package handler

import (
	"github.com/deeptest-com/deeptest/internal/server/modules/service"
	"github.com/deeptest-com/deeptest/pkg/domain"
	"github.com/kataras/iris/v12"
)

type EnvironmentVarCtrl struct {
	EnvironmentService *service.EnvironmentService `inject:""`
	BaseCtrl
}

// List
// @Tags	环境管理/全局变量
// @summary	列出环境变量
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	serverId		query	int		true	"serverId"
// @success	200	{object}	_domain.Response{data=[]domain.GlobalVar}
// @Router	/api/v1/environments/envVars	[get]
func (c *EnvironmentVarCtrl) List(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	serverId, err := ctx.URLParamInt("serverId")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	data, _ := c.EnvironmentService.GetVarsByServer(tenantId, uint(serverId))

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data})
}

// ListByEnvId
// @Tags	环境管理/全局变量
// @summary	根据环境列出环境变量
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	envId			query	int		true	"环境ID"
// @success	200	{object}	_domain.Response{data=[]domain.GlobalVar}
// @Router	/api/v1/environments/varsByEnv	[get]
func (c *EnvironmentVarCtrl) ListByEnvId(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	envId, err := ctx.URLParamInt("envId")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code})
		return
	}

	data, _ := c.EnvironmentService.GetVarsByEnv(tenantId, uint(envId))

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data})
}

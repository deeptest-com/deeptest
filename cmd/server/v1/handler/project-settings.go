package handler

import (
	"github.com/deeptest-com/deeptest/cmd/server/v1/domain"
	"github.com/deeptest-com/deeptest/internal/pkg/config"
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	"github.com/deeptest-com/deeptest/internal/pkg/core/cron"
	"github.com/deeptest-com/deeptest/internal/server/modules/service"
	_domain "github.com/deeptest-com/deeptest/pkg/domain"
	"github.com/deeptest-com/deeptest/saas/common"
	"github.com/deeptest-com/deeptest/saas/tenant"
	"github.com/kataras/iris/v12"
)

type ProjectSettingsCtrl struct {
	BaseCtrl
	Cron                   *cron.ServerCron                `inject:""`
	ProjectSettingsService *service.ProjectSettingsService `inject:""`
	ThirdPartySyncService  *service.ThirdPartySyncService  `inject:""`
}

// SaveSwaggerSync
// @Tags	自动同步
// @summary	保存同步信息
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string						true	"Authentication header"
// @Param 	currProjectId	query	int							true	"当前项目ID"
// @Param 	SwaggerSyncReq	body	serverDomain.SwaggerSyncReq	true	"保存同步信息的请求参数"
// @success	200	{object}	_domain.Response{data=model.SwaggerSync}
// @Router	/api/v1/serves/saveSwaggerSync	[post]
func (c *ProjectSettingsCtrl) SaveSwaggerSync(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	var req serverDomain.SwaggerSyncReq
	if err := ctx.ReadJSON(&req); err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	//
	if req.Switch == consts.SwitchOFF {

	}

	projectId, _ := ctx.URLParamInt("currProjectId")
	if req.ProjectId == 0 {
		req.ProjectId = uint(projectId)
	}
	res, err := c.ProjectSettingsService.SaveSwaggerSync(tenantId, req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg, Data: res})
}

// SwaggerSyncDetail
// @Tags	自动同步
// @summary	获取同步信息
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @success	200	{object}	_domain.Response{data=model.SwaggerSync}
// @Router	/api/v1/serves/swaggerSyncDetail	[get]
//func (c *ProjectSettingsCtrl) SwaggerSyncDetail(ctx iris.Context) {
//	tenantId := c.getTenantId(ctx)
//	projectId := ctx.URLParamUint64("currProjectId")
//	res, err := c.ProjectSettingsService.SwaggerSyncDetail(tenantId, uint(projectId))
//	if err != nil {
//		res.CategoryId = -1
//		res.SyncType = consts.FullCover
//		res.Cron = "23 * * * *"
//	}
//	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg, Data: res})
//}

func (c *ProjectSettingsCtrl) InitSwaggerCron() {
	//SAAS
	if config.CONFIG.Saas.Switch {
		tenants := tenant.NewTenant().GetInfos()
		for _, tenant := range tenants {
			go common.AsyncCatchErrRun(func() {
				c.initSwaggerCron(tenant.Id)
			})
			break
		}
	} else {
		//default
		c.initSwaggerCron("")
	}
}

// GetMock - Get Project Mock Settings
func (c *ProjectSettingsCtrl) GetMock(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	projectId := ctx.URLParamUint64("currProjectId")

	res, err := c.ProjectSettingsService.GetMock(tenantId, uint(projectId))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg, Data: res})
}

// SaveMock - Save Project Mock Settings
func (c *ProjectSettingsCtrl) SaveMock(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	var req serverDomain.MockReq
	if err := ctx.ReadJSON(&req); err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	projectId, _ := ctx.URLParamInt("currProjectId")
	if req.ProjectId == 0 {
		req.ProjectId = uint(projectId)
	}

	res, err := c.ProjectSettingsService.SaveMock(tenantId, req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg, Data: res})
}

func (c *ProjectSettingsCtrl) InitThirdPartySyncCron() {
	//SAAS
	if config.CONFIG.Saas.Switch {
		tenants := tenant.NewTenant().GetInfos()
		for _, tenant := range tenants {
			go common.AsyncCatchErrRun(func() {
				c.ThirdPartySyncService.AddThirdPartySyncCron(tenant.Id)
			})
		}
	} else {
		c.ThirdPartySyncService.AddThirdPartySyncCron("")
	}

}
func (c *ProjectSettingsCtrl) initSwaggerCron(tenantId consts.TenantId) {
	syncList, err := c.ProjectSettingsService.SwaggerSyncList(tenantId)
	if err != nil {
		return
	}
	for _, item := range syncList {
		c.ProjectSettingsService.AddSwaggerCron(tenantId, item)
	}
}

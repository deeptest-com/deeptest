package handler

import (
	"fmt"
	serverDomain "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/config"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/core/web/validate"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	"github.com/aaronchen2k/deeptest/internal/server/modules/service"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/aaronchen2k/deeptest/pkg/lib/cron/task"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/aaronchen2k/deeptest/saas/tenant"
	"github.com/kataras/iris/v12"
	"github.com/snowlyg/multi"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strings"
)

type ProjectCronCtrl struct {
	ProjectCronService    *service.ProjectCronService    `inject:""`
	LecangCronService     *service.LecangCronService     `inject:""`
	ThirdPartySyncService *service.ThirdPartySyncService `inject:""`
	ProjectCronRepo       *repo.ProjectCronRepo          `inject:""`
	DB                    *gorm.DB                       `inject:""`
	Proxy                 *task.Proxy                    `inject:""`
	BaseCtrl
}

// List
// @Tags	定时任务
// @summary	定时任务列表
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string							true	"Authentication header"
// @Param 	currProjectId	query	int								true	"当前项目ID"
// @Param 	ProjectCronReqPaginate	query	serverDomain.ProjectCronReqPaginate	true	"定时任务列表的请求参数"
// @success	200	{object}	_domain.Response{data=_domain.PageData{result=[]model.ProjectCron}}
// @Router	/api/v1/project/cron	[get]
func (c *ProjectCronCtrl) List(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	projectId, err := ctx.URLParamInt("currProjectId")
	if projectId == 0 {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	var req serverDomain.ProjectCronReqPaginate
	if err := ctx.ReadQuery(&req); err != nil {
		errs := validate.ValidRequest(err)
		if len(errs) > 0 {
			logUtils.Errorf("参数验证失败", zap.String("错误", strings.Join(errs, ";")))
			ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: strings.Join(errs, ";")})
			return
		}
	}
	req.ProjectId = uint(projectId)

	data, err := c.ProjectCronService.Paginate(tenantId, req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
}

// Get
// @Tags	定时任务
// @summary	定时任务详情
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	id				path	int		true	"任务ID"
// @success	200	{object}	_domain.Response{data=model.ProjectCronReq}
// @Router	/api/v1/project/cron/{id}	[get]
func (c *ProjectCronCtrl) Get(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)

	var req _domain.ReqId
	if err := ctx.ReadParams(&req); err != nil {
		logUtils.Errorf("参数解析失败", zap.String("错误:", err.Error()))
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	data, err := c.ProjectCronService.Get(tenantId, req.Id)

	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
		return
	}
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
}

// Save
// @Tags	定时任务
// @summary	保存定时任务
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string						true	"Authentication header"
// @Param 	currProjectId	query	int							true	"当前项目ID"
// @Param 	ProjectCronReq	body	model.ProjectCron	true	"保存定时任务的请求参数"
// @success	200	{object}	_domain.Response{data=int}
// @Router	/api/v1/project/cron	[post]
func (c *ProjectCronCtrl) Save(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	projectId, err := ctx.URLParamInt("currProjectId")
	if projectId == 0 {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: "projectId"})
		return
	}

	req := model.ProjectCron{}
	err = ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}
	req.ProjectId = uint(projectId)
	req.CreateUserId = multi.GetUserId(ctx)

	cron, err := c.ProjectCronService.Save(tenantId, req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
		return
	}

	c.RemoveCron(tenantId, cron)
	c.addCron(tenantId, cron)

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: cron.ID, Msg: _domain.NoErr.Msg})
}

// UpdateSwitchStatus
// @Tags	定时任务
// @summary	更新定时任务状态
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string						true	"Authentication header"
// @Param 	currProjectId	query	int							true	"当前项目ID"
// @Param 	id		body	number	true	"任务id"
// @Param 	switch	body	number	true	"开关状态"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/project/cron/updateStatus	[post]
func (c *ProjectCronCtrl) UpdateSwitchStatus(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)

	req := model.ProjectCron{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	err = c.ProjectCronService.UpdateSwitchStatus(tenantId, req.ID, req.Switch)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
		return
	}

	cron, err := c.ProjectCronService.Get(tenantId, req.ID)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	if req.Switch == consts.SwitchON {
		c.addCron(tenantId, cron)
	} else {
		c.RemoveCron(tenantId, cron)
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

// Delete
// @Tags	定时任务
// @summary	删除定时任务
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	id				path	int		true	"任务ID"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/project/cron/{id}	[delete]
func (c *ProjectCronCtrl) Delete(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)

	var req _domain.ReqId
	err := ctx.ReadParams(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	err = c.ProjectCronService.Delete(tenantId, req.Id)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	cron, err := c.ProjectCronService.Get(tenantId, req.Id)
	c.RemoveCron(tenantId, cron)

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

// Clone
// @Tags	定时任务
// @summary	克隆定时任务
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	id				path	int		true	"任务ID"
// @success	200	{object}	_domain.Response{data=int}
// @Router	/api/v1/project/cron/{id}/clone	[get]
func (c *ProjectCronCtrl) Clone(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	userId := multi.GetUserId(ctx)

	var req _domain.ReqId
	if err := ctx.ReadParams(&req); err != nil {
		logUtils.Errorf("参数解析失败", zap.String("错误:", err.Error()))
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	cron, err := c.ProjectCronService.Clone(tenantId, req.Id, userId)

	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
		return
	}

	c.addCron(tenantId, cron)

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: cron.ID, Msg: _domain.NoErr.Msg})
}

// EngineeringOptions
// @Tags	定时任务
// @summary	获取工程下拉选项
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	url				query	string	true	"环境URL"
// @success	200	{object}	_domain.Response{data=int}
// @Router	/api/v1/project/cron/engineeringOptions	[get]
func (c *ProjectCronCtrl) EngineeringOptions(ctx iris.Context) {
	url := ctx.URLParam("url")
	if url == "" {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: "环境URL不能为空"})
		return
	}

	data, err := c.ThirdPartySyncService.GetEngineeringOptions(url)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
}

// ServiceOptions
// @Tags	定时任务
// @summary	获取服务下拉选项
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	url				query	string	true	"环境URL"
// @Param 	engineering		query	string	true	"工程code"
// @success	200	{object}	_domain.Response{data=int}
// @Router	/api/v1/project/cron/serviceOptions	[get]
func (c *ProjectCronCtrl) ServiceOptions(ctx iris.Context) {
	url := ctx.URLParam("url")
	if url == "" {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: "环境URL不能为空"})
		return
	}

	engineering := ctx.URLParam("engineering")

	data, err := c.ThirdPartySyncService.GetServiceOptions(engineering, url)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
}

// AllServiceList
// @Tags	定时任务
// @summary	获取所有服务列表
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	url				query	string	true	"环境URL"
// @success	200	{object}	_domain.Response{data=int}
// @Router	/api/v1/project/cron/allServiceList	[get]
func (c *ProjectCronCtrl) AllServiceList(ctx iris.Context) {
	url := ctx.URLParam("url")
	if url == "" {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: "环境URL不能为空"})
		return
	}

	data, err := c.ThirdPartySyncService.GetAllServiceList(url)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
}

func (c *ProjectCronCtrl) InitProjectCron() {
	if config.CONFIG.Saas.Switch {
		tenants := tenant.NewTenant().GetInfos()
		for _, tenantIem := range tenants {
			c.addCronForSaas(tenantIem.Id)
		}
	} else {
		c.addCronForSaas("")
	}
}

func (c *ProjectCronCtrl) addCronForSaas(tenantId consts.TenantId) {
	cronList, _ := c.ProjectCronRepo.ListAllCron(tenantId)
	for _, item := range cronList {
		c.addCron(tenantId, item)
	}
}

func (c *ProjectCronCtrl) addCron(tenantId consts.TenantId, cron model.ProjectCron) {
	options := make(map[string]interface{})
	options["projectId"] = cron.ProjectId
	options["taskId"] = cron.ConfigId
	options["tenantId"] = tenantId

	c.Proxy.Init(tenantId, cron.Source, c.ProjectCronService.UpdateCronExecTimeById, fmt.Sprintf("%d", cron.ConfigId), cron.Cron)
	err := c.Proxy.Add(options)

	if err != nil {
		_ = c.ProjectCronService.UpdateExecErr(tenantId, cron.ID, err.Error())
		logUtils.Errorf("addCronErr:%+v, cron:%+v", err, cron)
	}
}

func (c *ProjectCronCtrl) RemoveCron(tenantId consts.TenantId, cron model.ProjectCron) {
	c.Proxy.Init(tenantId, cron.Source, nil, fmt.Sprintf("%d", cron.ConfigId), "")
	c.Proxy.Remove()
}

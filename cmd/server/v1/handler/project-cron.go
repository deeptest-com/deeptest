package handler

import (
	serverDomain "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/core/web/validate"
	"github.com/aaronchen2k/deeptest/internal/server/modules/service"
	third_party "github.com/aaronchen2k/deeptest/internal/server/modules/service/third-party"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/kataras/iris/v12"
	"github.com/snowlyg/multi"
	"go.uber.org/zap"
	"strings"
)

type ProjectCronCtrl struct {
	ProjectCronService    *service.ProjectCronService    `inject:""`
	LecangCronService     *third_party.LecangCronService `inject:""`
	ThirdPartySyncService *service.ThirdPartySyncService `inject:""`
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
	req.ConvertParams()
	req.ProjectId = uint(projectId)

	data, err := c.ProjectCronService.Paginate(req)
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
	var req _domain.ReqId
	if err := ctx.ReadParams(&req); err != nil {
		logUtils.Errorf("参数解析失败", zap.String("错误:", err.Error()))
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	data, err := c.ProjectCronService.Get(req.Id)

	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
		return
	}
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
}

// Save
// @Tags	定时任务
// @summary	新建定时任务
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string						true	"Authentication header"
// @Param 	currProjectId	query	int							true	"当前项目ID"
// @Param 	ProjectCronReq	body	serverDomain.ProjectCronReq	true	"新建定时任务的请求参数"
// @success	200	{object}	_domain.Response{data=int}
// @Router	/api/v1/project/cron	[post]
func (c *ProjectCronCtrl) Save(ctx iris.Context) {
	projectId, err := ctx.URLParamInt("currProjectId")
	if projectId == 0 {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: "projectId"})
		return
	}

	req := serverDomain.ProjectCronReq{}
	err = ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}
	req.ProjectId = uint(projectId)
	req.CreateUserId = multi.GetUserId(ctx)
	req.Switch = consts.SwitchON

	cronId, err := c.ProjectCronService.Save(req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: cronId, Msg: _domain.NoErr.Msg})
}

// Update
// @Tags	定时任务
// @summary	更新定时任务
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string						true	"Authentication header"
// @Param 	currProjectId	query	int							true	"当前项目ID"
// @Param 	ProjectCronReq	body	serverDomain.ProjectCronReq	true	"更新定时任务的请求参数"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/project/cron	[put]
func (c *ProjectCronCtrl) Update(ctx iris.Context) {
	req := serverDomain.ProjectCronReq{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}
	req.UpdateUserId = multi.GetUserId(ctx)

	err = c.ProjectCronService.Update(req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
		return
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
	var req _domain.ReqId
	err := ctx.ReadParams(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	err = c.ProjectCronService.Delete(req.Id)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

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
	userId := multi.GetUserId(ctx)

	var req _domain.ReqId
	if err := ctx.ReadParams(&req); err != nil {
		logUtils.Errorf("参数解析失败", zap.String("错误:", err.Error()))
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	cronId, err := c.ProjectCronService.Clone(req.Id, userId)

	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
		return
	}
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: cronId, Msg: _domain.NoErr.Msg})
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
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
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
// @Param 	engineeringCode	query	string	true	"工程code"
// @success	200	{object}	_domain.Response{data=int}
// @Router	/api/v1/project/cron/serviceOptions	[get]
func (c *ProjectCronCtrl) ServiceOptions(ctx iris.Context) {
	url := ctx.URLParam("url")
	if url == "" {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: "环境URL不能为空"})
		return
	}

	engineeringCode := ctx.URLParam("engineeringCode")

	data, err := c.ThirdPartySyncService.GetServiceOptions(engineeringCode, url)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
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
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
}

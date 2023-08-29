package handler

import (
	"github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/core/web/validate"
	"github.com/aaronchen2k/deeptest/internal/server/modules/service"
	"github.com/aaronchen2k/deeptest/pkg/domain"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/kataras/iris/v12"
	"github.com/snowlyg/multi"
	"go.uber.org/zap"
	"strings"
)

type ProjectCtrl struct {
	ProjectService                *service.ProjectService                `inject:""`
	ProjectRecentlyVisitedService *service.ProjectRecentlyVisitedService `inject:""`
	BaseCtrl
}

// List
// @Tags	项目管理
// @summary	项目列表
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string							true	"Authentication header"
// @Param 	currProjectId	query	int								true	"当前项目ID"
// @Param 	ProjectReq		query	serverDomain.ProjectReqPaginate	true	"项目列表"
// @success	200	{object}	_domain.Response{data=_domain.PageData{result=[]model.Project}}
// @Router	/api/v1/projects	[get]
func (c *ProjectCtrl) List(ctx iris.Context) {
	userId := multi.GetUserId(ctx)

	var req serverDomain.ProjectReqPaginate
	if err := ctx.ReadQuery(&req); err != nil {
		errs := validate.ValidRequest(err)
		if len(errs) > 0 {
			logUtils.Errorf("参数验证失败", zap.String("错误", strings.Join(errs, ";")))
			ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: strings.Join(errs, ";")})
			return
		}
	}
	req.ConvertParams()

	data, err := c.ProjectService.Paginate(req, userId)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
}

// Get
// @Tags	项目管理
// @summary	项目详情
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param	id				path	int		true	"项目ID"
// @success	200	{object}	_domain.Response{data=model.Project}
// @Router	/api/v1/projects/{id}	[get]
func (c *ProjectCtrl) Get(ctx iris.Context) {
	var req _domain.ReqId
	if err := ctx.ReadParams(&req); err != nil {
		logUtils.Errorf("参数解析失败", zap.String("错误:", err.Error()))
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}
	project, err := c.ProjectService.Get(req.Id)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
		return
	}
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: project, Msg: _domain.NoErr.Msg})
}

// Create
// @Tags	项目管理
// @summary	创建项目
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization	header	string					true	"Authentication header"
// @Param 	currProjectId	query	int						true	"当前项目ID"
// @Param 	ProjectReq 		body 	serverDomain.ProjectReq true 	"Create project Request Object"
// @success	200	{object}	_domain.Response{data=object{id=int}}
// @Router	/api/v1/projects	[post]
func (c *ProjectCtrl) Create(ctx iris.Context) {
	userId := multi.GetUserId(ctx)

	req := serverDomain.ProjectReq{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	id, bizErr := c.ProjectService.Create(req, userId)
	if bizErr.Code != 0 {
		ctx.JSON(_domain.Response{Code: bizErr.Code, Data: nil, Msg: bizErr.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: iris.Map{"id": id}, Msg: _domain.NoErr.Msg})
}

// Update
// @Tags	项目管理
// @summary	更新项目
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization	header	string					true	"Authentication header"
// @Param 	currProjectId	query	int						true	"当前项目ID"
// @Param 	ProjectReq 		body 	serverDomain.ProjectReq true 	"update project Request Object"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/projects	[put]
func (c *ProjectCtrl) Update(ctx iris.Context) {

	var req serverDomain.ProjectReq
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	err = c.ProjectService.Update(req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

// Delete
// @Tags	项目管理
// @summary	删除项目
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization	header	string			true	"Authentication header"
// @Param 	currProjectId	query	int				true	"当前项目ID"
// @Param	id				path	int				true	"项目ID"
// @Param 	_domain.ReqId	query	_domain.ReqId	true	"请求参数"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/projects/{id}	[delete]
func (c *ProjectCtrl) Delete(ctx iris.Context) {
	var req _domain.ReqId
	err := ctx.ReadParams(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	err = c.ProjectService.DeleteById(req.Id)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

// GetByUser
// @Tags	项目管理
// @summary	获取用户参与的项目
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization	header	string			true	"Authentication header"
// @Param 	currProjectId	query	int				true	"当前项目ID"
// @success	200	{object}	_domain.Response{data=object{projects=[]model.ProjectMemberRole, currProject=model.Project, recentProjects=[]model.Project}}
// @Router	/api/v1/projects/getByUser	[get]
func (c *ProjectCtrl) GetByUser(ctx iris.Context) {
	userId := multi.GetUserId(ctx)

	projects, currProject, recentProjects, err := c.ProjectService.GetByUser(userId)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ret := iris.Map{"projects": projects, "currProject": currProject, "recentProjects": recentProjects}
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: ret, Msg: _domain.NoErr.Msg})
}

// ChangeProject
// @Tags	项目管理
// @summary	切换用户默认项目
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	id 				body 	int 	true 	"项目ID"
// @success	200	{object}	_domain.Response{data=object{id=int}}
// @Router	/api/v1/projects/changeProject	[post]
func (c *ProjectCtrl) ChangeProject(ctx iris.Context) {
	userId := multi.GetUserId(ctx)

	req := serverDomain.ProjectReq{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	err = c.ProjectService.ChangeProject(req.Id, userId)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	_, _ = c.ProjectRecentlyVisitedService.Create(userId, req.Id)

	projects, currProject, recentProjects, err := c.ProjectService.GetByUser(userId)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ret := iris.Map{"projects": projects, "currProject": currProject, "recentProjects": recentProjects}
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: ret, Msg: _domain.NoErr.Msg})
}

// Members
// @Tags	项目管理
// @summary	获取项目成员
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization		header	string							true	"Authentication header"
// @Param 	currProjectId		query	int								true	"当前项目ID"
// @Param 	ProjectReqPaginate	query	serverDomain.ProjectReqPaginate	true	"获取项目成员的分页参数"
// @success	200	{object}	_domain.Response{data=_domain.PageData{result=[]serverDomain.MemberResp}}
// @Router	/api/v1/projects/members	[get]
func (c *ProjectCtrl) Members(ctx iris.Context) {
	projectId, err := ctx.URLParamInt("currProjectId")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	var req serverDomain.ProjectReqPaginate
	if err := ctx.ReadQuery(&req); err != nil {
		errs := validate.ValidRequest(err)
		if len(errs) > 0 {
			ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: strings.Join(errs, ";")})
			return
		}
	}
	req.ConvertParams()

	data, err := c.ProjectService.Members(req, projectId)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
}

// RemoveMember
// @Tags	项目管理
// @summary	删除项目成员
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization				header	string								true	"Authentication header"
// @Param 	currProjectId				query	int									true	"当前项目ID"
// @Param 	ProjectMemberRemoveReq 		body 	serverDomain.ProjectMemberRemoveReq true 	"删除项目成员的请求体"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/projects/removeMember	[post]
func (c *ProjectCtrl) RemoveMember(ctx iris.Context) {
	req := serverDomain.ProjectMemberRemoveReq{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	err = c.ProjectService.RemoveMember(req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

// ChangeUserRole
// @Tags	项目管理
// @summary	更新项目成员的角色
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization				header	string								true	"Authentication header"
// @Param 	currProjectId				query	int									true	"当前项目ID"
// @Param 	UpdateProjectMemberReq 		body 	serverDomain.UpdateProjectMemberReq true 	"更新项目成员角色的请求体"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/projects/changeUserRole	[post]
func (c *ProjectCtrl) ChangeUserRole(ctx iris.Context) {
	req := serverDomain.UpdateProjectMemberReq{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	err = c.ProjectService.UpdateMemberRole(req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

// Apply
// @Tags	项目管理
// @summary	申请项目成员
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization	header	string							true	"Authentication header"
// @Param 	currProjectId	query	int								true	"当前项目ID"
// @Param 	ApplyProjectReq body 	serverDomain.ApplyProjectReq	true 	"申请项目成员的请求体"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/projects/apply	[post]
func (c *ProjectCtrl) Apply(ctx iris.Context) {
	req := serverDomain.ApplyProjectReq{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}
	req.ApplyUserId = multi.GetUserId(ctx)
	err = c.ProjectService.Apply(req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

// Audit
// @Tags	项目管理
// @summary	审批操作
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization	header	string							true	"Authentication header"
// @Param 	currProjectId	query	int								true	"当前项目ID"
// @Param 	AuditProjectReq body 	serverDomain.AuditProjectReq 	true 	"审批操作的请求体"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/projects/audit	[post]
func (c *ProjectCtrl) Audit(ctx iris.Context) {
	req := serverDomain.AuditProjectReq{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	userId := multi.GetUserId(ctx)
	err = c.ProjectService.Audit(req.Id, userId, req.Status)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

// AuditList
// @Tags	项目管理
// @summary	申请加入审批列表
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization			header	string								true	"Authentication header"
// @Param 	currProjectId			query	int									true	"当前项目ID"
// @Param 	AuditProjectPaginate	body 	serverDomain.AuditProjectPaginate 	true 	"申请加入审批列表的请求体"
// @success	200	{object}	_domain.Response{data=_domain.PageData{result=[]model.ProjectMemberAudit}}
// @Router	/api/v1/projects/auditList	[post]
func (c *ProjectCtrl) AuditList(ctx iris.Context) {

	req := serverDomain.AuditProjectPaginate{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}
	req.AuditUserId = multi.GetUserId(ctx)
	req.ApplyUserId = multi.GetUserId(ctx)
	res, err := c.ProjectService.AuditList(req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg, Data: res})
}

// AuditUsers
// @Tags	项目管理
// @summary	申请加入项目的审批人
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	projectId		query	int		true	"要申请的项目ID"
// @success	200	{object}	_domain.Response{data=[]model.SysUser}
// @Router	/api/v1/projects/auditUsers	[get]
func (c *ProjectCtrl) AuditUsers(ctx iris.Context) {
	projectId, err := ctx.URLParamInt("projectId")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}
	if projectId == 0 {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: "projectId can't be empty"})
		return
	}
	res, err := c.ProjectService.AuditUsers(uint(projectId))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg, Data: res})
}

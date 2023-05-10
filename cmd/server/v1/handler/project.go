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
		ctx.JSON(_domain.Response{Code: bizErr.Code, Data: nil})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: iris.Map{"id": id}, Msg: _domain.NoErr.Msg})
}

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

func (c *ProjectCtrl) Members(ctx iris.Context) {
	projectId, err := ctx.URLParamInt("id")
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

func (c *ProjectCtrl) Audit(ctx iris.Context) {
	req := serverDomain.AuditProjectReq{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	userId := multi.GetUserId(ctx)
	err = c.ProjectService.Audit(req.ProjectId, userId, req.Status)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

func (c *ProjectCtrl) AuditList(ctx iris.Context) {

	req := serverDomain.AuditProjectPaginate{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}
	req.AuditUserId = multi.GetUserId(ctx)
	res, err := c.ProjectService.AuditList(req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg, Data: res})
}

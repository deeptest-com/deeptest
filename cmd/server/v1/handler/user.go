package handler

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/core/web/validate"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	"github.com/aaronchen2k/deeptest/internal/server/modules/service"
	"github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/aaronchen2k/deeptest/pkg/lib/log"
	"strings"

	"github.com/kataras/iris/v12"
	"github.com/snowlyg/multi"
	"go.uber.org/zap"
)

type UserCtrl struct {
	UserService *service.UserService `inject:""`
	UserRepo    *repo.UserRepo       `inject:""`
}

func (c *UserCtrl) ListAll(ctx iris.Context) {
	var req v1.UserReqPaginate

	if err := ctx.ReadQuery(&req); err != nil {
		errs := validate.ValidRequest(err)
		if len(errs) > 0 {
			_logUtils.Errorf("参数验证失败", zap.String("错误", strings.Join(errs, ";")))
			ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: strings.Join(errs, ";")})
			return
		}
	}

	data, err := c.UserRepo.Paginate(req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
}

// GetUser 详情
func (c *UserCtrl) GetUser(ctx iris.Context) {
	var req _domain.ReqId
	if err := ctx.ReadParams(&req); err != nil {
		_logUtils.Errorf("参数解析失败", zap.String("错误:", err.Error()))
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Data: nil, Msg: _domain.ParamErr.Msg})
		return
	}
	user, err := c.UserRepo.FindById(req.Id)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: _domain.SystemErr.Msg})
		return
	}
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: user, Msg: _domain.NoErr.Msg})
}

// Profile 个人信息
func (c *UserCtrl) Profile(ctx iris.Context) {
	user, err := c.UserRepo.FindById(multi.GetUserId(ctx))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: _domain.SystemErr.Msg})
		return
	}
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: user, Msg: _domain.NoErr.Msg})
}

// Message 消息
func (c *UserCtrl) Message(ctx iris.Context) {
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: nil, Msg: _domain.NoErr.Msg})
}

// CreateUser 添加
func (c *UserCtrl) CreateUser(ctx iris.Context) {
	req := v1.UserReq{}
	if err := ctx.ReadJSON(&req); err != nil {
		errs := validate.ValidRequest(err)
		if len(errs) > 0 {
			_logUtils.Errorf("参数验证失败", zap.String("错误", strings.Join(errs, ";")))
			ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: strings.Join(errs, ";")})
			return
		}
	}
	id, err := c.UserRepo.Create(req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: iris.Map{"id": id}, Msg: _domain.NoErr.Msg})
}

// UpdateUser 更新
func (c *UserCtrl) UpdateUser(ctx iris.Context) {
	var reqId _domain.ReqId
	if err := ctx.ReadParams(&reqId); err != nil {
		_logUtils.Errorf("参数解析失败", zap.String("错误:", err.Error()))
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Data: nil, Msg: _domain.ParamErr.Msg})
		return
	}

	var req v1.UserReq
	if err := ctx.ReadJSON(&req); err != nil {
		errs := validate.ValidRequest(err)
		if len(errs) > 0 {
			_logUtils.Errorf("参数验证失败", zap.String("错误", strings.Join(errs, ";")))
			ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: strings.Join(errs, ";")})
			return
		}
	}

	err := c.UserRepo.Update(reqId.Id, req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: err.Error()})
		return
	}
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: nil, Msg: _domain.NoErr.Msg})
}

// DeleteUser 删除
func (c *UserCtrl) DeleteUser(ctx iris.Context) {
	var req _domain.ReqId
	if err := ctx.ReadParams(&req); err != nil {
		_logUtils.Errorf("参数解析失败", zap.String("错误:", err.Error()))
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Data: nil, Msg: _domain.ParamErr.Msg})
		return
	}
	err := c.UserRepo.DeleteById(req.Id)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: nil, Msg: _domain.NoErr.Msg})
}

// Logout 退出
func (c *UserCtrl) Logout(ctx iris.Context) {
	token := multi.GetVerifiedToken(ctx)
	if token == nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: "授权凭证为空"})
		return
	}
	err := c.UserRepo.DelToken(string(token))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: err.Error()})
		return
	}
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: nil, Msg: _domain.NoErr.Msg})
}

// Clear 清空 token
func (c *UserCtrl) Clear(ctx iris.Context) {
	token := multi.GetVerifiedToken(ctx)
	if token == nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: "授权凭证为空"})
		return
	}
	if err := c.UserRepo.CleanToken(multi.AdminAuthority, string(token)); err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: err.Error()})
		return
	}
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: nil, Msg: _domain.NoErr.Msg})
}

// ChangeAvatar 修改头像
func (c *UserCtrl) ChangeAvatar(ctx iris.Context) {
	avatar := &model.Avatar{}
	if err := ctx.ReadJSON(avatar); err != nil {
		errs := validate.ValidRequest(err)
		if len(errs) > 0 {
			_logUtils.Errorf("参数验证失败", zap.String("错误", strings.Join(errs, ";")))
			ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: strings.Join(errs, ";")})
			return
		}
	}
	err := c.UserRepo.UpdateAvatar(multi.GetUserId(ctx), avatar.Avatar)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: err.Error()})
		return
	}
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: nil, Msg: _domain.NoErr.Msg})
}

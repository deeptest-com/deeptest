package controller

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	logUtils "github.com/aaronchen2k/deeptest/internal/pkg/lib/log"
	"github.com/aaronchen2k/deeptest/internal/server/core/web/validate"
	serverDomain "github.com/aaronchen2k/deeptest/internal/server/modules/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/service"
	"strings"

	"github.com/kataras/iris/v12"
	"go.uber.org/zap"
)

type TestStepCtrl struct {
	TestStepService *service.TestStepService `inject:""`
	FileService     *service.FileService     `inject:""`
	BaseCtrl
}

func NewTestStepCtrl() *TestStepCtrl {
	return &TestStepCtrl{}
}

// Get 详情
func (c *TestStepCtrl) Get(ctx iris.Context) {
	var reqId _domain.ReqId
	if err := ctx.ReadParams(&reqId); err != nil {
		logUtils.Errorf("参数解析失败", zap.String("错误:", err.Error()))
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Data: nil, Msg: _domain.ParamErr.Msg})
		return
	}

	step, err := c.TestStepService.FindById(reqId.Id)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: _domain.SystemErr.Msg})
		return
	}
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: step, Msg: _domain.NoErr.Msg})
}

// Create 添加
func (c *TestStepCtrl) Create(ctx iris.Context) {
	f, fh, err := ctx.FormFile("file")
	if err != nil {
		logUtils.Errorf("文件上传失败", zap.String("ctx.FormFile(\"file\")", err.Error()))
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: err.Error()})
		return
	}
	defer f.Close()

	imagePath, err := c.FileService.UploadFile(ctx, fh)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: err.Error()})
		return
	}

	req := serverDomain.TestStepReq{}
	if err := ctx.ReadForm(&req); err != nil {
		errs := validate.ValidRequest(err)
		if len(errs) > 0 {
			logUtils.Errorf("参数验证失败", zap.String("错误", strings.Join(errs, ";")))
			ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: strings.Join(errs, ";")})
			return
		}
	}

	req.Image = imagePath
	id, err := c.TestStepService.Create(req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: c.ErrCode(err), Data: nil})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: iris.Map{"id": id}, Msg: _domain.NoErr.Msg})
}

// Update 更新
func (c *TestStepCtrl) Update(ctx iris.Context) {
	var reqId _domain.ReqId
	if err := ctx.ReadParams(&reqId); err != nil {
		logUtils.Errorf("参数解析失败", zap.String("错误:", err.Error()))
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Data: nil, Msg: _domain.ParamErr.Msg})
		return
	}

	var req serverDomain.TestStepReq
	if err := ctx.ReadJSON(&req); err != nil {
		errs := validate.ValidRequest(err)
		if len(errs) > 0 {
			logUtils.Errorf("参数验证失败", zap.String("错误", strings.Join(errs, ";")))
			ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: strings.Join(errs, ";")})
			return
		}
	}

	err := c.TestStepService.Update(reqId.Id, req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: err.Error()})
		return
	}
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: nil, Msg: _domain.NoErr.Msg})
}

// Delete 删除
func (c *TestStepCtrl) Delete(ctx iris.Context) {
	var req _domain.ReqId
	if err := ctx.ReadParams(&req); err != nil {
		logUtils.Errorf("参数解析失败", zap.String("错误:", err.Error()))
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Data: nil, Msg: _domain.ParamErr.Msg})
		return
	}
	err := c.TestStepService.DeleteById(req.Id)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: nil, Msg: _domain.NoErr.Msg})
}

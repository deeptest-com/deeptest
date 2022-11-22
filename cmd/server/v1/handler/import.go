package handler

import (
	domain "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	commService "github.com/aaronchen2k/deeptest/internal/pkg/service"
	service "github.com/aaronchen2k/deeptest/internal/server/modules/service"
	"github.com/aaronchen2k/deeptest/pkg/domain"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/kataras/iris/v12"
)

type ImportCtrl struct {
	ImportService *service.ImportService `inject:""`

	FileService *commService.FileService `inject:""`

	BaseCtrl
}

func (c *ImportCtrl) ImportSpec(ctx iris.Context) {
	targetId, err := ctx.URLParamInt("targetId")
	if targetId == 0 {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: "targetId"})
		return
	}

	req := domain.InterfaceImportReq{}
	err = ctx.ReadJSON(&req)
	if err != nil {
		logUtils.Errorf("参数验证失败", err.Error())
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Data: nil, Msg: err.Error()})
		return
	}

	c.ImportService.Import(req, targetId)

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: nil, Msg: _domain.NoErr.Msg})

	return
}

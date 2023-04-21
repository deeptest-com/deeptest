package handler

import (
	"fmt"
	domain "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	commService "github.com/aaronchen2k/deeptest/internal/pkg/service"
	service "github.com/aaronchen2k/deeptest/internal/server/modules/service"
	"github.com/aaronchen2k/deeptest/pkg/domain"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/kataras/iris/v12"
)

type ImportCtrl struct {
	ImportService *service.ImportService   `inject:""`
	YapiService   *service.YapiService     `inject:""`
	FileService   *commService.FileService `inject:""`

	BaseCtrl
}

func (c *ImportCtrl) ImportSpec(ctx iris.Context) {
	targetId, err := ctx.URLParamInt("targetId")
	if targetId == 0 {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	var req openapi3.T
	err = ctx.ReadJSON(&req)
	if err != nil {
		logUtils.Errorf("参数验证失败", err.Error())
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	c.ImportService.Import(req, targetId)

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})

	return
}

func (c *ImportCtrl) ImportYapi(ctx iris.Context) {
	projectId, err := ctx.URLParamInt("currProjectId")
	fmt.Println("projectId", projectId)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	req := domain.InterfaceYapiReq{}
	err = ctx.ReadJSON(&req)
	if err != nil {
		logUtils.Errorf("参数验证失败", err.Error())
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}
	fmt.Println("InterfaceYapiReq", req)

	req.ProjectId = projectId
	c.YapiService.ImportYapiProject(req)

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})

	return
}

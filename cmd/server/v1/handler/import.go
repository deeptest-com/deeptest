package handler

import (
	"fmt"
	domain "github.com/deeptest-com/deeptest/cmd/server/v1/domain"
	commService "github.com/deeptest-com/deeptest/internal/pkg/service"
	service "github.com/deeptest-com/deeptest/internal/server/modules/service"
	"github.com/deeptest-com/deeptest/pkg/domain"
	logUtils "github.com/deeptest-com/deeptest/pkg/lib/log"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/kataras/iris/v12"
)

type ImportCtrl struct {
	ImportService *service.ImportService   `inject:""`
	YapiService   *service.YapiService     `inject:""`
	FileService   *commService.FileService `inject:""`

	BaseCtrl
}

// ImportSpec
// @Tags	导入模块
// @summary	导入OpenApi文件
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	targetId 		query 	int 	true 	"targetId"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/import/importSpec	[post]
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

// ImportYapi
// @Tags	导入模块
// @summary	导入yapi项目接口
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization		header	string							true	"Authentication header"
// @Param 	currProjectId		query	int								true	"当前项目ID"
// @Param 	InterfaceYapiReq 	body 	serverDomain.InterfaceYapiReq 	true 	"导入yapi项目接口的请求参数"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/import/importYapi	[post]
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

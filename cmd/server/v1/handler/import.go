package handler

import (
	domain "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	commService "github.com/aaronchen2k/deeptest/internal/pkg/service"
	service "github.com/aaronchen2k/deeptest/internal/server/modules/service"
	"github.com/aaronchen2k/deeptest/pkg/domain"
	_fileUtils "github.com/aaronchen2k/deeptest/pkg/lib/file"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/kataras/iris/v12"
)

type ImportCtrl struct {
	ImportService *service.ImportService `inject:""`

	FileService *commService.FileService `inject:""`

	BaseCtrl
}

func (c *ImportCtrl) ImportSpecFromContent(ctx iris.Context) {
	targetId, _ := ctx.URLParamInt("targetId")
	typ := ctx.URLParam("type")

	req := domain.InterfaceImportReq{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		logUtils.Errorf("参数验证失败", err.Error())
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Data: nil, Msg: err.Error()})
		return
	}

	fileContent := []byte(req.Content)
	var content []byte

	if typ == "postman" {
		content = fileContent // already be converted to openapi3 format
	} else if typ == "openapi2" {
		content, _ = c.ImportService.OpenApi2To3(content)
	} else if typ == "openapi3" {
		content = fileContent
	}

	c.ImportService.Import(content, targetId)

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: nil, Msg: _domain.NoErr.Msg})

	return
}
func (c *ImportCtrl) ImportSpecFromForm(ctx iris.Context) {
	targetId, _ := ctx.URLParamInt("targetId")
	typ := ctx.URLParam("type")

	f, fh, err := ctx.FormFile("file")
	if err != nil {
		logUtils.Errorf("获取上传文件失败， %s。", err.Error())
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: err.Error()})
		return
	}
	defer f.Close()

	pth, err := c.FileService.UploadFile(ctx, fh, "spec")

	fileContent := _fileUtils.ReadFileBuf(pth)
	var content []byte

	if typ == "postman" {
		content, _ = c.ImportService.PostmanToOpenApi3(pth)
	} else if typ == "openapi2" {
		content, _ = c.ImportService.OpenApi2To3(content)
	} else if typ == "openapi" {
		content = fileContent
	}

	c.ImportService.Import(content, targetId)

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: nil, Msg: _domain.NoErr.Msg})

	return
}

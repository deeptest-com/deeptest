package controller

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	logUtils "github.com/aaronchen2k/deeptest/internal/pkg/lib/log"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/service"
	"github.com/kataras/iris/v12"
	"go.uber.org/zap"
)

type FileCtrl struct {
	FileService *service.FileService `inject:""`
}

// Upload 上传文件
func (c *FileCtrl) Upload(ctx iris.Context) {
	f, fh, err := ctx.FormFile("file")
	if err != nil {
		logUtils.Errorf("文件上传失败", zap.String("ctx.FormFile(\"file\")", err.Error()))
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: err.Error()})
		return
	}
	defer f.Close()

	pth, err := c.FileService.UploadFile(ctx, fh)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: err.Error()})
		return
	}
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: pth, Msg: _domain.NoErr.Msg})
}

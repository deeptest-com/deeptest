package service

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	dateUtils "github.com/aaronchen2k/deeptest/pkg/lib/date"
	_fileUtils "github.com/aaronchen2k/deeptest/pkg/lib/file"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/kataras/iris/v12"
	"github.com/snowlyg/helper/dir"
	"mime/multipart"
	"path/filepath"
	"time"
)

type ProcessorDataService struct {
}

// Upload 上传文件
func (s *ProcessorDataService) Upload(ctx iris.Context, fh *multipart.FileHeader) (ret v1.ProcessorDataUploadResp, err error) {
	filename, err := _fileUtils.GetUploadFileName(fh.Filename)
	if err != nil {
		logUtils.Errorf("获取文件名失败，错误%s", err.Error())
		return
	}

	targetDir := filepath.Join(consts.DirUpload, dateUtils.DateStr(time.Now()))
	targetDir = filepath.Join(consts.WorkDir, targetDir)

	err = dir.InsureDir(targetDir)
	if err != nil {
		logUtils.Errorf("文件上传失败，错误%s", err.Error())
		return
	}

	_, err = ctx.SaveFormFile(fh, filepath.Join(targetDir, filename))
	if err != nil {
		logUtils.Errorf("文件上传失败，错误%s", "保存文件到本地")
		return
	}

	ret.Path = filepath.Join(targetDir, filename)
	ret.Data = ""

	return
}

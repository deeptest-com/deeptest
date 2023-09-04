package commService

import (
	"errors"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	dateUtils "github.com/aaronchen2k/deeptest/pkg/lib/date"
	_fileUtils "github.com/aaronchen2k/deeptest/pkg/lib/file"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	stringUtils "github.com/aaronchen2k/deeptest/pkg/lib/string"
	"mime/multipart"
	"path/filepath"
	"strings"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/snowlyg/helper/dir"
	"github.com/snowlyg/helper/str"
)

var (
	ErrEmpty = errors.New("请上传正确的文件")
)

type FileService struct {
}

// UploadFile 上传文件
func (s *FileService) UploadFile(ctx iris.Context, fh *multipart.FileHeader) (ret string, err error) {
	filename, err := _fileUtils.GetUploadFileName(fh.Filename)
	if err != nil {
		logUtils.Errorf("获取文件名失败，错误%s", err.Error())
		return
	}

	targetDir := filepath.Join(consts.DirUpload, dateUtils.DateStr(time.Now()))
	absDir := filepath.Join(dir.GetCurrentAbPath(), targetDir)

	err = dir.InsureDir(targetDir)
	if err != nil {
		logUtils.Errorf("文件上传失败，错误%s", err.Error())
		return
	}

	pth := filepath.Join(absDir, filename)
	_, err = ctx.SaveFormFile(fh, pth)
	if err != nil {
		logUtils.Errorf("文件上传失败，错误%s", "保存文件到本地")
		return
	}

	ret = filepath.Join(targetDir, filename)

	return
}

// GetFileName 获取文件名称
func GetFileName(name string) (string, error) {
	fns := strings.Split(strings.TrimPrefix(name, "./"), ".")
	if len(fns) != 2 {
		logUtils.Errorf("文件名错误 %s", name)
		return "", ErrEmpty
	}

	base := fns[0]
	ext := fns[1]
	return str.Join(base, "-", stringUtils.Uuid(), ".", ext), nil
}

package service

import (
	"errors"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	dateUtils "github.com/aaronchen2k/deeptest/pkg/lib/date"
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
	Sname *repo.FileNameRepo `inject:""`
}

func NewFileService() *FileService {
	return &FileService{}
}

// UploadFile 上传文件
func (s *FileService) UploadFile(ctx iris.Context, fh *multipart.FileHeader, folder string) (pth string, err error) {
	filename, err := GetFileName(fh.Filename)
	if err != nil {
		logUtils.Errorf("获取文件名失败，错误%s", err.Error())
		return
	}

	relaDir := filepath.Join("static", "upload", folder, dateUtils.DateStr(time.Now()))
	absDir := filepath.Join(dir.GetCurrentAbPath(), relaDir)

	err = dir.InsureDir(absDir)
	if err != nil {
		logUtils.Errorf("文件上传失败，错误%s", err.Error())
		return
	}

	_, err = ctx.SaveFormFile(fh, filepath.Join(absDir, filename))
	if err != nil {
		logUtils.Errorf("文件上传失败，错误%s", "保存文件到本地")
		return
	}

	pth = filepath.Join(relaDir, filename)

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
func (s *FileService) UploadFileSaveName(ctx iris.Context, fh *multipart.FileHeader, folder string) (res *model.FileName, err error) {
	filename, err := GetFileName(fh.Filename)

	if err != nil {
		logUtils.Errorf("获取文件名失败，错误%s", err.Error())
		return
	}
	//name, err := s.Sname.GetPathName("ceshi.csv")
	//if err != nil {
	//	return nil, err
	//}
	//fmt.Println("1111111111111111111111111111111111", name.PathName)

	relaDir := filepath.Join("static", "upload", folder, dateUtils.DateStr(time.Now())) //转换后的路径static\upload\files\2023-01-03
	absDir := filepath.Join(dir.GetCurrentAbPath(), relaDir)

	err = dir.InsureDir(absDir)
	if err != nil {
		logUtils.Errorf("文件上传失败，错误%s", err.Error())
		return
	}

	_, err = ctx.SaveFormFile(fh, filepath.Join(absDir, filename))
	if err != nil {
		logUtils.Errorf("文件上传失败，错误%s", "保存文件到本地")
		return
	}

	pth := filepath.Join(relaDir, filename)
	k := strings.Replace(pth, "\\", "/", -1)

	res = model.NewFileName(fh.Filename, k)

	_, err = s.Sname.SaveName(res)
	if err != nil {
		return
	}

	return
}

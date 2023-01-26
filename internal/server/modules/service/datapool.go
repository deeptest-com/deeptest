package service

import (
	"errors"
	"fmt"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	dateUtils "github.com/aaronchen2k/deeptest/pkg/lib/date"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	stringUtils "github.com/aaronchen2k/deeptest/pkg/lib/string"
	"github.com/kataras/iris/v12"
	"github.com/snowlyg/helper/dir"
	"github.com/snowlyg/helper/str"
	"mime/multipart"
	"path/filepath"
	"strings"
	"time"
)

type DatapoolService struct {
	DatapoolRepo *repo.DatapoolRepo `inject:""`
}

func NewDatapoolService() *DatapoolService {
	return &DatapoolService{}
}

func (s *DatapoolService) List(projectId uint) (ret []v1.DatapoolReq, err error) {
	ret, err = s.DatapoolRepo.List(projectId)

	return
}

func (s *DatapoolService) Get(id uint) (model.Datapool, error) {
	return s.DatapoolRepo.Get(id)
}

func (s *DatapoolService) Save(req *model.Datapool) (err error) {
	return s.DatapoolRepo.Save(req)
}

func (s *DatapoolService) SaveData(req v1.DatapoolReq) (err error) {
	return s.DatapoolRepo.SaveData(req)
}

func (s *DatapoolService) Delete(id uint) (err error) {
	return s.DatapoolRepo.Delete(id)
}

// Upload 上传文件
func (s *DatapoolService) Upload(ctx iris.Context, fh *multipart.FileHeader, datapoolId int) (ret v1.DatapoolUploadResp, err error) {
	filename, err := GetFileName(fh.Filename)
	if err != nil {
		logUtils.Errorf("获取文件名失败，错误%s", err.Error())
		return
	}

	targetDir := filepath.Join(consts.DirData, dateUtils.DateStr(time.Now()))
	targetDir = filepath.Join(dir.GetCurrentAbPath(), targetDir)

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

// GetFileName 获取文件名称
func GetFileName(name string) (ret string, err error) {
	fns := strings.Split(strings.TrimPrefix(name, "./"), ".")
	if len(fns) != 2 {
		msg := fmt.Sprintf("文件名错误 %s", name)

		logUtils.Info(msg)
		err = errors.New(msg)

		return
	}

	base := fns[0]
	ext := fns[1]

	ret = str.Join(base, "-", stringUtils.Uuid(), ".", ext)

	return
}

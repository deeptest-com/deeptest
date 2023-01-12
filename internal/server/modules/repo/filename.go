package repo

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"gorm.io/gorm"
)

type FileNameRepo struct {
	DB *gorm.DB `inject:""`
}

func NewFileNameRepo() *FileNameRepo {
	return &FileNameRepo{}
}

func (r *FileNameRepo) SaveName(req *model.FileName) (name string, err error) {
	err = r.DB.Create(&req).Error
	name = req.PathName
	return
}

func (r *FileNameRepo) GetPathName(filename string) (f *model.FileName, err error) {
	//var f model.FileName

	err = r.DB.Where("file_name = ?", filename).Order("created_at DESC").Limit(1).Find(&f).Error
	//pathname = f.PathName

	//First(&file.CreatedAt).Error

	return

}

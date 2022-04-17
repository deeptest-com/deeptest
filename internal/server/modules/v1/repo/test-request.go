package repo

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	"gorm.io/gorm"
)

type TestRequestRepo struct {
	DB *gorm.DB `inject:""`
}

func NewRequestRepo(db *gorm.DB) *TestRequestRepo {
	return &TestRequestRepo{DB: db}
}

func (r *TestRequestRepo) List(interfaceId int) (pos []model.TestRequest, err error) {
	err = r.DB.
		Select("id", "name").
		Where("interface_id=?", interfaceId).
		Where("NOT deleted").
		Order("created_at DESC").
		Find(&pos).Error
	return
}

func (r *TestRequestRepo) Get(id uint) (field model.TestRequest, err error) {
	err = r.DB.
		Where("id=?", id).
		Where("NOT deleted").
		First(&field).Error
	return
}

func (r *TestRequestRepo) Save(request *model.TestRequest) (err error) {
	err = r.DB.Save(request).Error
	return
}

func (r *TestRequestRepo) Delete(id uint) (err error) {
	err = r.DB.Model(&model.TestRequest{}).
		Where("id=?", id).
		Update("deleted", true).
		Error

	return
}

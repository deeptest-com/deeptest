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

func (r *TestRequestRepo) ListByProject(projectId int) (pos []*model.TestRequest, err error) {
	err = r.DB.
		Where("project_id=?", projectId).
		Where("NOT deleted").
		Order("parent_id ASC, ordr ASC").
		Find(&pos).Error
	return
}

func (r *TestRequestRepo) Get(fieldId uint) (field model.TestRequest, err error) {
	err = r.DB.
		Where("id=?", fieldId).
		Where("NOT deleted").
		First(&field).Error
	return
}

func (r *TestRequestRepo) Save(field *model.TestRequest) (err error) {
	err = r.DB.Save(field).Error
	return
}

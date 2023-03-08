package repo

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"go/types"
	"gorm.io/gorm"
)

type SummaryDetailsRepo struct {
	DB *gorm.DB `inject:""`
}

func NewSummaryDetailsRepo() *SummaryDetailsRepo {
	return &SummaryDetailsRepo{}
}

func (r *SummaryDetailsRepo) FindByProjectId(projectId uint) (summaryDetails []model.SummaryDetails, err error) {
	err = r.DB.Model(&model.SummaryDetails{}).Where("project_id = ?", projectId).Find(&summaryDetails).Error
	return
}

func (r *SummaryDetailsRepo) Find() (summaryDetails []model.SummaryDetails, err error) {
	err = r.DB.Model(&model.SummaryDetails{}).Find(&summaryDetails).Error
	return
}

func (r *SummaryDetailsRepo) UpdateByProjectId(projectId uint, key string, value types.Object) (err error) {
	err = r.DB.Update(key, value).Where("project_id= ? ", projectId).Error
	return
}

func (r *SummaryDetailsRepo) Create(summaryDetails model.SummaryDetails) (err error) {
	err = r.DB.Model(&model.SummaryDetails{}).Create(summaryDetails).Error
	return
}

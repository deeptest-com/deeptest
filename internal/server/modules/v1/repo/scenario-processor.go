package repo

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	"gorm.io/gorm"
)

type ScenarioProcessorRepo struct {
	DB *gorm.DB `inject:""`
}

func (r *ScenarioProcessorRepo) UpdateName(id int, name string) (err error) {
	err = r.DB.Model(&model.TestProcessor{}).
		Where("id = ?", id).
		Update("name", name).Error

	return
}

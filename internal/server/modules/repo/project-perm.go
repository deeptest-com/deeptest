package repo

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"gorm.io/gorm"
)

type ProjectPerm struct {
	DB *gorm.DB `inject:""`
}

func NewProjectPerm() *ProjectPerm {
	return &ProjectPerm{}
}

func (p *ProjectPerm) GetRecordByNameAndAction(name, action string) (projectPerm model.ProjectPerm, err error) {
	err = p.DB.Model(&model.ProjectPerm{}).
		Where("name = ?", name).
		Where("act = ?", action).
		First(&projectPerm).Error
	return
}

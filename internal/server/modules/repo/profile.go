package repo

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"gorm.io/gorm"
)

type ProfileRepo struct {
	DB       *gorm.DB  `inject:""`
	RoleRepo *RoleRepo `inject:""`
}

func NewProfileRepo() *ProfileRepo {
	return &ProfileRepo{}
}

func (r *ProfileRepo) FindByUserId(userId uint) (profile model.SysUserProfile, err error) {
	db := r.DB.Model(&model.SysUserProfile{}).Where("user_id = ?", userId)
	err = db.First(&profile).Error
	return
}

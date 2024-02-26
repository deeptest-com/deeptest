package repo

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"gorm.io/gorm"
)

type ProfileRepo struct {
	*BaseRepo `inject:""`
	DB        *gorm.DB  `inject:""`
	RoleRepo  *RoleRepo `inject:""`
}

func (r *ProfileRepo) FindByUserId(tenantId consts.TenantId, userId uint) (profile model.SysUserProfile, err error) {
	db := r.GetDB(tenantId).Model(&model.SysUserProfile{}).Where("user_id = ?", userId)
	err = db.First(&profile).Error
	return
}

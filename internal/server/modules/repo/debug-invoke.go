package repo

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"gorm.io/gorm"
)

type DebugInvokeRepo struct {
	*BaseRepo `inject:""`
	DB        *gorm.DB `inject:""`
}

func (r *DebugInvokeRepo) Get(id uint) (po model.DebugInvoke, err error) {
	err = r.DB.
		Where("id=?", id).
		Where("NOT deleted").
		First(&po).Error
	return
}

func (r *DebugInvokeRepo) Save(invocation *model.DebugInvoke) (err error) {
	err = r.DB.Save(invocation).Error
	return
}

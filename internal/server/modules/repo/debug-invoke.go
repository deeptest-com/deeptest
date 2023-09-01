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

func (r *DebugInvokeRepo) ChangeProcessorOwner(oldProcessId, newProcessId, debugInterfaceId, endpointInterfaceId uint) (err error) {
	values := map[string]interface{}{
		"scenario_processor_id": newProcessId,
		"debug_interface_id":    debugInterfaceId,
		"endpoint_interface_id": endpointInterfaceId,
	}

	err = r.DB.Model(&model.DebugInvoke{}).
		Where("scenario_processor_id = ? AND NOT deleted", oldProcessId).
		Updates(values).Error

	return
}

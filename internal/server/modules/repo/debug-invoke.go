package repo

import (
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	"github.com/deeptest-com/deeptest/internal/server/modules/model"
	"gorm.io/gorm"
)

type DebugInvokeRepo struct {
	*BaseRepo `inject:""`
	DB        *gorm.DB `inject:""`
}

func (r *DebugInvokeRepo) Get(tenantId consts.TenantId, id uint) (po model.DebugInvoke, err error) {
	err = r.GetDB(tenantId).
		Where("id=?", id).
		Where("NOT deleted").
		First(&po).Error
	return
}

func (r *DebugInvokeRepo) Save(tenantId consts.TenantId, invocation *model.DebugInvoke) (err error) {
	err = r.GetDB(tenantId).Save(invocation).Error
	return
}

func (r *DebugInvokeRepo) ChangeProcessorOwner(tenantId consts.TenantId, oldProcessId, newProcessId, debugInterfaceId, endpointInterfaceId uint) (err error) {
	values := map[string]interface{}{
		"scenario_processor_id": newProcessId,
		"debug_interface_id":    debugInterfaceId,
		"endpoint_interface_id": endpointInterfaceId,
	}

	err = r.GetDB(tenantId).Model(&model.DebugInvoke{}).
		Where("scenario_processor_id = ? AND NOT deleted", oldProcessId).
		Updates(values).Error

	return
}

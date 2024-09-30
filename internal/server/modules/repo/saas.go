package repo

import (
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	"github.com/deeptest-com/deeptest/internal/server/modules/model"
)

type SaasRepo struct {
	*BaseRepo `inject:""`
}

func (r *SaasRepo) GetUserList(tenantId consts.TenantId) (data []model.SysUser, err error) {
	db := r.GetDB(tenantId)
	err = db.Model(&model.SysUser{}).Find(&data).Error
	return
}

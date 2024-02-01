package repo

import "github.com/aaronchen2k/deeptest/internal/server/modules/model"

type SaasRepo struct {
	*BaseRepo `inject:""`
}

func (r *SaasRepo) GetUserList(tenantId string) (data []model.SysUser, err error) {
	db := r.GetDB(tenantId)
	err = db.Model(&model.SysUser{}).Find(&data).Error
	return
}

package repo

import (
	"fmt"
	v1 "github.com/deeptest-com/deeptest/cmd/server/v1/domain"
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	"github.com/deeptest-com/deeptest/internal/server/modules/model"
	"gorm.io/gorm"
)

type SysAgentRepo struct {
	*BaseRepo `inject:""`
	DB        *gorm.DB `inject:""`
}

func (r *SysAgentRepo) List(tenantId consts.TenantId, keywords string) (pos []model.SysAgent, err error) {
	db := r.GetDB(tenantId).Model(&model.SysAgent{}).
		Where("NOT deleted")

	if keywords != "" {
		db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", keywords))
	}

	err = db.Find(&pos).Error

	return
}

func (r *SysAgentRepo) Get(tenantId consts.TenantId, id uint) (po model.SysAgent, err error) {
	err = r.GetDB(tenantId).
		Where("id = ?", id).
		First(&po).Error

	return
}

func (r *SysAgentRepo) Save(tenantId consts.TenantId, po *model.SysAgent) (err error) {
	err = r.GetDB(tenantId).Model(po).
		Save(&po).Error

	return
}

func (r *SysAgentRepo) UpdateName(tenantId consts.TenantId, to v1.AgentReq) (err error) {
	err = r.GetDB(tenantId).Model(&model.SysAgent{}).
		Where("id = ?", to.Id).
		Updates(map[string]interface{}{"name": to.Name, "update_user": to.UpdateUser}).Error

	return
}

func (r *SysAgentRepo) Delete(tenantId consts.TenantId, id uint) (err error) {
	err = r.GetDB(tenantId).Model(&model.SysAgent{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{"deleted": true}).Error

	return
}

func (r *SysAgentRepo) Disable(tenantId consts.TenantId, id uint) (err error) {
	err = r.GetDB(tenantId).Model(&model.SysAgent{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{"disabled": gorm.Expr("NOT disabled")}).Error

	return
}

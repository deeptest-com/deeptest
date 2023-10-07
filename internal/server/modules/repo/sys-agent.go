package repo

import (
	"fmt"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"gorm.io/gorm"
)

type SysAgentRepo struct {
	DB *gorm.DB `inject:""`
}

func (r *SysAgentRepo) List(keywords string) (pos []model.SysAgent, err error) {
	db := r.DB.Model(&model.SysAgent{}).
		Where("NOT deleted")

	if keywords != "" {
		db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", keywords))
	}

	err = db.Find(&pos).Error

	return
}

func (r *SysAgentRepo) Get(id uint) (po model.SysAgent, err error) {
	err = r.DB.
		Where("id = ?", id).
		First(&po).Error

	return
}

func (r *SysAgentRepo) Save(po *model.SysAgent) (err error) {
	err = r.DB.Model(po).
		Save(&po).Error

	return
}

func (r *SysAgentRepo) UpdateName(to v1.AgentReq) (err error) {
	err = r.DB.Model(&model.SysAgent{}).
		Where("id = ?", to.Id).
		Updates(map[string]interface{}{"name": to.Name, "update_user": to.UpdateUser}).Error

	return
}

func (r *SysAgentRepo) Delete(id uint) (err error) {
	err = r.DB.Model(&model.SysAgent{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{"deleted": true}).Error

	return
}

func (r *SysAgentRepo) Disable(id uint) (err error) {
	err = r.DB.Model(&model.SysAgent{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{"disabled": gorm.Expr("NOT disabled")}).Error

	return
}

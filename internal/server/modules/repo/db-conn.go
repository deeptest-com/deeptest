package repo

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	model "github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"gorm.io/gorm"
)

type DatabaseConnRepo struct {
	DB *gorm.DB `inject:""`
}

func (r *DatabaseConnRepo) List(envId uint) (po []model.DatabaseConn, err error) {
	err = r.DB.Model(&model.DatabaseConn{}).
		Where("environment_id = ? AND NOT deleted", envId).
		Find(&po).Error

	return
}

func (r *DatabaseConnRepo) Get(id uint) (po model.DatabaseConn, err error) {
	err = r.DB.Model(&model.DatabaseConn{}).
		Where("id = ?", id).First(&po).Error

	return
}

func (r *DatabaseConnRepo) Save(po *model.DatabaseConn) (err error) {
	err = r.DB.Save(po).Error

	return
}

func (r *DatabaseConnRepo) Delete(id uint) (err error) {
	err = r.DB.Model(&model.DatabaseConn{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{"deleted": true}).Error

	return
}

func (r *DatabaseConnRepo) Disable(id uint) (err error) {
	err = r.DB.Model(&model.DatabaseConn{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{"disabled": gorm.Expr("NOT disabled")}).Error

	return
}

func (r *DatabaseConnRepo) UpdateName(req v1.DbConnReq) (err error) {
	err = r.DB.Model(&model.DatabaseConn{}).
		Where("id = ?", req.Id).
		Updates(map[string]interface{}{"name": req.Name, "update_user": req.UpdateUser}).Error

	return
}

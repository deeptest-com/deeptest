package repo

import (
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

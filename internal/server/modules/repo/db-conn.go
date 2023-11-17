package repo

import (
	"errors"
	"fmt"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	model "github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"gorm.io/gorm"
)

type DatabaseConnRepo struct {
	DB *gorm.DB `inject:""`
}

func (r *DatabaseConnRepo) List(keywords string, projectId int, ignoreDisabled bool) (pos []model.DatabaseConn, err error) {
	db := r.DB.Model(&model.DatabaseConn{}).
		Where("project_id = ? AND NOT deleted", projectId)

	if ignoreDisabled {
		db.Where("NOT disabled")
	}

	if keywords != "" {
		db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", keywords))
	}

	err = db.Find(&pos).Error

	return
}

func (r *DatabaseConnRepo) Get(id uint) (po model.DatabaseConn, err error) {
	err = r.DB.Model(&model.DatabaseConn{}).
		Where("id = ?", id).First(&po).Error

	return
}

func (r *DatabaseConnRepo) GetByName(id, projectId uint, name string) (po model.DatabaseConn, err error) {
	err = r.DB.Model(&model.DatabaseConn{}).
		Where("id != ? AND project_id = ? AND name = ? and not deleted", id, projectId, name).First(&po).Error

	return
}

func (r *DatabaseConnRepo) Save(po *model.DatabaseConn) (err error) {
	exist, _ := r.GetByName(po.ID, po.ProjectId, po.Name)
	if exist.ID > 0 {
		err = errors.New("名称不能和已存在的记录相同")
		return
	}

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
	exist, _ := r.GetByName(req.Id, req.ProjectId, req.Name)
	if exist.ID > 0 {
		err = errors.New("名称不能和已存在的记录相同")
		return
	}

	err = r.DB.Model(&model.DatabaseConn{}).
		Where("id = ?", req.Id).
		Updates(map[string]interface{}{"name": req.Name, "update_user": req.UpdateUser}).Error

	return
}

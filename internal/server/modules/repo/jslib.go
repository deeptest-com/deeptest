package repo

import (
	"errors"
	"fmt"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	model "github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"gorm.io/gorm"
)

type JslibRepo struct {
	*BaseRepo `inject:""`
	DB        *gorm.DB `inject:""`
}

func (r *JslibRepo) List(tenantId consts.TenantId, keywords string, projectId int, ignoreDisabled bool) (pos []model.SysJslib, err error) {
	db := r.GetDB(tenantId).Model(&model.SysJslib{}).
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

func (r *JslibRepo) Get(tenantId consts.TenantId, id uint) (po model.SysJslib, err error) {
	err = r.GetDB(tenantId).Model(&model.SysJslib{}).
		Where("id = ?", id).First(&po).Error

	return
}
func (r *JslibRepo) GetByName(tenantId consts.TenantId, id, projectId uint, name string) (po model.SysJslib, err error) {
	err = r.GetDB(tenantId).Model(&model.SysJslib{}).
		Where("id != ? AND project_id = ? AND name = ? and not deleted", id, projectId, name).First(&po).Error

	return
}

func (r *JslibRepo) Save(tenantId consts.TenantId, po *model.SysJslib) (err error) {
	exist, _ := r.GetByName(tenantId, po.ID, po.ProjectId, po.Name)
	if exist.ID > 0 {
		err = errors.New("名称不能和已存在的记录相同")
		return
	}

	err = r.GetDB(tenantId).Save(po).Error

	return
}

func (r *JslibRepo) UpdateName(tenantId consts.TenantId, to v1.JslibReq) (err error) {
	exist, _ := r.GetByName(tenantId, to.Id, to.ProjectId, to.Name)
	if exist.ID > 0 {
		err = errors.New("名称不能和已存在的记录相同")
		return
	}

	err = r.GetDB(tenantId).Model(&model.SysJslib{}).
		Where("id = ?", to.Id).
		Updates(map[string]interface{}{"name": to.Name, "update_user": to.UpdateUser}).Error

	return
}

func (r *JslibRepo) Delete(tenantId consts.TenantId, id uint) (err error) {
	err = r.GetDB(tenantId).Model(&model.SysJslib{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{"deleted": true}).Error

	return
}

func (r *JslibRepo) Disable(tenantId consts.TenantId, id uint) (err error) {
	err = r.GetDB(tenantId).Model(&model.SysJslib{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{"disabled": gorm.Expr("NOT disabled")}).Error

	return
}

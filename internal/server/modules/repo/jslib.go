package repo

import (
	"errors"
	"fmt"
	v1 "github.com/deeptest-com/deeptest/cmd/server/v1/domain"
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	model "github.com/deeptest-com/deeptest/internal/server/modules/model"
	"gorm.io/gorm"
)

type JslibRepo struct {
	*BaseRepo `inject:""`
	DB        *gorm.DB `inject:""`
}

func (r *JslibRepo) List(tenantId consts.TenantId, keywords string, projectId int, ignoreDisabled bool) (pos []model.Jslib, err error) {
	db := r.GetDB(tenantId).Model(&model.Jslib{}).
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

func (r *JslibRepo) Get(tenantId consts.TenantId, id uint) (po model.Jslib, err error) {
	err = r.GetDB(tenantId).Model(&model.Jslib{}).
		Where("id = ?", id).First(&po).Error

	return
}
func (r *JslibRepo) GetByName(tenantId consts.TenantId, id, projectId uint, name string) (po model.Jslib, err error) {
	err = r.GetDB(tenantId).Model(&model.Jslib{}).
		Where("id != ? AND project_id = ? AND name = ? and not deleted", id, projectId, name).First(&po).Error

	return
}

func (r *JslibRepo) Save(tenantId consts.TenantId, po *model.Jslib) (err error) {
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

	err = r.GetDB(tenantId).Model(&model.Jslib{}).
		Where("id = ?", to.Id).
		Updates(map[string]interface{}{"name": to.Name, "update_user": to.UpdateUser}).Error

	return
}

func (r *JslibRepo) Delete(tenantId consts.TenantId, id uint) (err error) {
	err = r.GetDB(tenantId).Model(&model.Jslib{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{"deleted": true}).Error

	return
}

func (r *JslibRepo) Disable(tenantId consts.TenantId, id uint) (err error) {
	err = r.GetDB(tenantId).Model(&model.Jslib{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{"disabled": gorm.Expr("NOT disabled")}).Error

	return
}

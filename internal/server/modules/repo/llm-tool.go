package repo

import (
	"errors"
	"fmt"
	v1 "github.com/deeptest-com/deeptest/cmd/server/v1/domain"
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	model "github.com/deeptest-com/deeptest/internal/server/modules/model"
	"gorm.io/gorm"
)

type LlmToolRepo struct {
	*BaseRepo `inject:""`
	DB        *gorm.DB `inject:""`
}

func (r *LlmToolRepo) List(tenantId consts.TenantId, keywords string, projectId int, ignoreDisabled bool) (pos []model.LlmTool, err error) {
	db := r.GetDB(tenantId).Model(&model.LlmTool{}).
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

func (r *LlmToolRepo) Get(tenantId consts.TenantId, id uint) (po model.LlmTool, err error) {
	err = r.GetDB(tenantId).Model(&model.LlmTool{}).
		Where("id = ?", id).First(&po).Error

	return
}

func (r *LlmToolRepo) GetByName(tenantId consts.TenantId, id, projectId uint, name string) (po model.LlmTool, err error) {
	err = r.GetDB(tenantId).Model(&model.LlmTool{}).
		Where("id != ? AND project_id = ? AND name = ? and not deleted", id, projectId, name).First(&po).Error

	return
}

func (r *LlmToolRepo) Save(tenantId consts.TenantId, po *model.LlmTool) (err error) {
	exist, _ := r.GetByName(tenantId, po.ID, po.ProjectId, po.Name)
	if exist.ID > 0 {
		err = errors.New("名称不能和已存在的记录相同")
		return
	}

	err = r.GetDB(tenantId).Save(po).Error
	if err != nil {
		return
	}

	count := r.Count(tenantId, po.ProjectId)
	if count == 1 {
		err = r.SetDefault(tenantId, po.ID)
	}

	return
}

func (r *LlmToolRepo) Delete(tenantId consts.TenantId, id uint) (err error) {
	err = r.GetDB(tenantId).Model(&model.LlmTool{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{"deleted": true}).Error

	return
}

func (r *LlmToolRepo) Disable(tenantId consts.TenantId, id uint) (err error) {
	err = r.GetDB(tenantId).Model(&model.LlmTool{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{"disabled": gorm.Expr("NOT disabled")}).Error

	return
}

func (r *LlmToolRepo) SetDefault(tenantId consts.TenantId, id uint) (err error) {
	err = r.GetDB(tenantId).Model(&model.LlmTool{}).
		Where("NOT deleted").
		Updates(map[string]interface{}{"is_default": false}).Error

	if err != nil {
		return
	}

	err = r.GetDB(tenantId).Model(&model.LlmTool{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{"is_default": true}).Error

	return
}

func (r *LlmToolRepo) UpdateName(tenantId consts.TenantId, req v1.DbConnReq) (err error) {
	exist, _ := r.GetByName(tenantId, req.Id, req.ProjectId, req.Name)
	if exist.ID > 0 {
		err = errors.New("名称不能和已存在的记录相同")
		return
	}

	err = r.GetDB(tenantId).Model(&model.LlmTool{}).
		Where("id = ?", req.Id).
		Updates(map[string]interface{}{"name": req.Name, "update_user": req.UpdateUser}).Error

	return
}

// Paginate
func (r *LlmToolRepo) Count(tenantId consts.TenantId, projectId uint) (count int64) {
	r.GetDB(tenantId).Model(&model.LlmTool{}).
		Where("project_id = ? AND NOT deleted", projectId).
		Count(&count)

	return
}

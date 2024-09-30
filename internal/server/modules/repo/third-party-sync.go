package repo

import (
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	"github.com/deeptest-com/deeptest/internal/server/modules/model"
	"gorm.io/gorm"
	"time"
)

type ThirdPartySyncRepo struct {
	*BaseRepo `inject:""`
	DB        *gorm.DB `inject:""`
}

func (r *ThirdPartySyncRepo) AllData(tenantId consts.TenantId) (data []model.ThirdPartySync, err error) {
	err = r.GetDB(tenantId).Where("switch = ?", consts.SwitchON).Find(&data).Error

	return
}

func (r *ThirdPartySyncRepo) UpdateExecTimeById(tenantId consts.TenantId, id uint) (err error) {
	return r.GetDB(tenantId).Model(&model.ThirdPartySync{}).Where("id=?", id).Update("exec_time", time.Now()).Error
}

func (r *ThirdPartySyncRepo) GetByProjectAndServe(tenantId consts.TenantId, projectId, serveId uint) (data model.ThirdPartySync, err error) {
	err = r.GetDB(tenantId).Where("project_id = ? and serve_id = ?", projectId, serveId).First(&data).Error

	return
}

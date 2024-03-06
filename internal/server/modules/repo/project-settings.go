package repo

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	mockData "github.com/aaronchen2k/deeptest/internal/pkg/helper/openapi-mock/openapi/generator/data"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"time"
)

type ProjectSettingsRepo struct {
	*BaseRepo       `inject:""`
	CategoryRepo    *CategoryRepo    `inject:""`
	EnvironmentRepo *EnvironmentRepo `inject:""`
}

func (r *ProjectSettingsRepo) SaveSwaggerSync(tenantId consts.TenantId, sync *model.SwaggerSync) (err error) {
	return r.Save(tenantId, sync.ID, sync)
}

func (r *ProjectSettingsRepo) GetSwaggerSync(tenantId consts.TenantId, projectId uint) (sync model.SwaggerSync, err error) {
	err = r.GetDB(tenantId).First(&sync, "project_id=?", projectId).Error
	return
}

func (r *ProjectSettingsRepo) GetSwaggerSyncById(tenantId consts.TenantId, id uint) (sync model.SwaggerSync, err error) {
	err = r.GetDB(tenantId).First(&sync, "id=?", id).Error
	return
}

func (r *ProjectSettingsRepo) GetSwaggerSyncList(tenantId consts.TenantId) (res []model.SwaggerSync, err error) {
	err = r.GetDB(tenantId).Find(&res).Error
	return
}

func (r *ProjectSettingsRepo) UpdateSwaggerSyncExecTimeById(tenantId consts.TenantId, id uint) (err error) {
	return r.GetDB(tenantId).Model(&model.SwaggerSync{}).Where("id=?", id).Update("exec_time", time.Now()).Error
}

func (r *ProjectSettingsRepo) GetMock(tenantId consts.TenantId, projectId uint) (po model.ProjectMockSetting, err error) {
	r.GetDB(tenantId).First(&po, "project_id=?", projectId)

	if po.ID == 0 {
		po = model.ProjectMockSetting{
			Priority:  consts.MockPrioritySmart,
			ProjectId: projectId,
		}
	}

	if po.Priority == consts.MockPriorityExample {
		po.UseExamples = mockData.IfPresent
	} else {
		po.UseExamples = mockData.No
	}

	return
}
func (r *ProjectSettingsRepo) SaveMock(tenantId consts.TenantId, po *model.ProjectMockSetting) (err error) {
	err = r.Save(tenantId, po.ID, po)
	return
}

func (r *ProjectSettingsRepo) DeleteSwaggerSyncById(tenantId consts.TenantId, id uint) error {
	return r.GetDB(tenantId).Model(&model.SwaggerSync{}).
		Where("id = ?", id).
		Update("deleted", true).Error
}

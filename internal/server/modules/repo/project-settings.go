package repo

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"time"
)

type ProjectSettingsRepo struct {
	*BaseRepo       `inject:""`
	CategoryRepo    *CategoryRepo    `inject:""`
	EnvironmentRepo *EnvironmentRepo `inject:""`
}

func (r *ProjectSettingsRepo) SaveSwaggerSync(sync *model.SwaggerSync) (err error) {
	return r.Save(sync.ID, sync)
}

func (r *ProjectSettingsRepo) GetSwaggerSync(projectId uint) (sync model.SwaggerSync, err error) {
	err = r.DB.First(&sync, "project_id=?", projectId).Error
	return
}

func (r *ProjectSettingsRepo) GetSwaggerSyncById(id uint) (sync model.SwaggerSync, err error) {
	err = r.DB.First(&sync, "id=?", id).Error
	return
}

func (r *ProjectSettingsRepo) GetSwaggerSyncList() (res []model.SwaggerSync, err error) {
	err = r.DB.Find(&res).Error
	return
}

func (r *ProjectSettingsRepo) UpdateSwaggerSyncExecTimeById(id uint) (err error) {
	return r.DB.Model(&model.SwaggerSync{}).Where("id=?", id).Update("exec_time", time.Now()).Error
}

func (r *ProjectSettingsRepo) GetMock(projectId uint) (po model.ProjectMockSetting, err error) {
	r.DB.First(&po, "project_id=?", projectId)

	if po.ID == 0 {
		po = model.ProjectMockSetting{
			Priority:  consts.MockPrioritySmart,
			ProjectId: projectId,
		}
	}

	return
}
func (r *ProjectSettingsRepo) SaveMock(po *model.ProjectMockSetting) (err error) {
	err = r.Save(po.ID, po)
	return
}

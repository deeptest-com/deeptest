package repo

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"gorm.io/gorm"
	"time"
)

type ThirdPartySyncRepo struct {
	*BaseRepo `inject:""`
	DB        *gorm.DB `inject:""`
}

func (r *ThirdPartySyncRepo) AllData() (data []model.ThirdPartySync, err error) {
	err = r.DB.Where("switch = ?", consts.SwitchON).Find(&data).Error

	return
}

func (r *ThirdPartySyncRepo) UpdateExecTimeById(id uint) (err error) {
	return r.DB.Model(&model.ThirdPartySync{}).Where("id=?", id).Update("exec_time", time.Now()).Error
}

func (r *ThirdPartySyncRepo) GetByProjectAndServe(projectId, serveId uint) (data model.ThirdPartySync, err error) {
	err = r.DB.Where("project_id = ? and serve_id = ?", projectId, serveId).First(&data).Error

	return
}

package repo

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"gorm.io/gorm"
	"strconv"
)

type PerformanceRunnerRepo struct {
	DB               *gorm.DB `inject:""`
	*BaseRepo        `inject:""`
	ProjectRepo      *ProjectRepo      `inject:""`
	ScenarioRepo     *ScenarioRepo     `inject:""`
	ScenarioNodeRepo *ScenarioNodeRepo `inject:""`
}

func (r *PerformanceRunnerRepo) List(scenarioId int) (pos []model.PerformanceRunner, err error) {
	err = r.DB.Model(&model.PerformanceRunner{}).
		Select(fmt.Sprintf("%s.*, a.name, a.url web_address", model.PerformanceRunner{}.TableName())).
		Joins(fmt.Sprintf("LEFT JOIN %s a ON %s.agent_id=a.id",
			model.SysAgent{}.TableName(), model.PerformanceRunner{}.TableName())).
		Where("scenario_id = ?", scenarioId).
		Find(&pos).Error

	return
}

func (r *PerformanceRunnerRepo) Get(id uint) (performanceTestPlan model.PerformanceRunner, err error) {
	err = r.DB.Where("id = ?", id).
		First(&performanceTestPlan).
		Error

	return performanceTestPlan, nil
}

func (r *PerformanceRunnerRepo) Save(po *model.PerformanceRunner) (err error) {
	err = r.DB.Save(po).Error
	if err != nil {
		return
	}

	err = r.UpdateSerialNumber(po.ID, po.ProjectId)
	if err != nil {
		return
	}

	return
}

func (r *PerformanceRunnerRepo) DeleteById(id uint) (err error) {
	err = r.DB.Model(&model.PerformanceRunner{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{"deleted": true}).Error

	return
}

func (r *PerformanceRunnerRepo) UpdateSerialNumber(id, projectId uint) (err error) {
	var project model.Project
	project, err = r.ProjectRepo.Get(projectId)
	if err != nil {
		return
	}

	err = r.DB.Model(&model.PerformanceRunner{}).Where("id=?", id).
		Update("serial_number", project.ShortName+"-RN-"+strconv.Itoa(int(id))).Error
	return
}

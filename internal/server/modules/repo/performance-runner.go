package repo

import (
	agentDomain "github.com/aaronchen2k/deeptest/cmd/agent/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"gorm.io/gorm"
	"sort"
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
	sql := "SELECT r.id, r.agent_id, r.weight, r.name, r.web_address, r.serial_number, r.scenario_id, r.project_id, " +
		"a.name, a.url web_address " +
		"FROM biz_performance_runner r " +
		"LEFT JOIN sys_agent a ON r.agent_id = a.id"

	err = r.DB.Raw(sql).
		Scan(&pos).Error

	return
}

func (r *PerformanceRunnerRepo) Get(id uint) (performanceTestPlan model.PerformanceRunner, err error) {
	err = r.DB.Where("id = ?", id).
		First(&performanceTestPlan).
		Error

	return performanceTestPlan, nil
}

func (r *PerformanceRunnerRepo) Select(req agentDomain.PerformanceRunnerSelectionReq) (err error) {
	err = r.DB.
		Where("scenario_id = ?", req.ScenarioId).
		Delete(&model.PerformanceRunner{}).Error
	if err != nil {
		return
	}

	ids := req.Ids
	sort.Ints(ids)

	for _, id := range ids {
		po := model.PerformanceRunner{
			Weight:     1,
			AgentId:    uint(id),
			ScenarioId: uint(req.ScenarioId),
			ProjectId:  uint(req.ProjectId),
		}

		err = r.DB.Save(&po).Error
		if err != nil {
			return
		}

		err = r.UpdateSerialNumber(po.ID, po.ProjectId)
		if err != nil {
			return
		}
	}

	return
}

func (r *PerformanceRunnerRepo) DeleteById(id uint) (err error) {
	err = r.DB.
		Where("id = ?", id).
		Delete(&model.PerformanceRunner{}).Error

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

func (r *PerformanceRunnerRepo) ListExistOnes(ids string) (ret []int) {
	r.DB.Model(model.PerformanceRunner{}).
		Select("id").
		Where("id IN (?)", ids).
		Scan(&ret)

	return
}

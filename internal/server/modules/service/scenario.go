package service

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	repo2 "github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	"github.com/aaronchen2k/deeptest/pkg/domain"
)

type ScenarioService struct {
	ScenarioRepo     *repo2.ScenarioRepo     `inject:""`
	ScenarioNodeRepo *repo2.ScenarioNodeRepo `inject:""`
}

func (s *ScenarioService) ListByProject(serveId int) (pos []model.Scenario, err error) {
	pos, err = s.ScenarioRepo.ListByProject(serveId)
	return
}

func (s *ScenarioService) Paginate(req v1.ScenarioReqPaginate, projectId int) (ret _domain.PageData, err error) {
	ret, err = s.ScenarioRepo.Paginate(req, projectId)

	if err != nil {
		return
	}

	return
}

func (s *ScenarioService) GetById(id uint) (model.Scenario, error) {
	return s.ScenarioRepo.Get(id)
}

func (s *ScenarioService) Create(req model.Scenario) (po model.Scenario, err error) {
	po, err = s.ScenarioRepo.Create(req)

	s.ScenarioNodeRepo.CreateDefault(po.ID, req.ProjectId, req.CreateUserId)

	return
}

func (s *ScenarioService) Update(req model.Scenario) error {
	return s.ScenarioRepo.Update(req)
}

func (s *ScenarioService) DeleteById(id uint) error {
	return s.ScenarioRepo.DeleteById(id)
}

func (s *ScenarioService) AddPlans(scenarioId int, planIds []int) (err error) {
	err = s.ScenarioRepo.AddPlans(uint(scenarioId), planIds)
	return
}

func (s *ScenarioService) RemovePlans(scenarioId int, planIds []int) (err error) {
	if len(planIds) == 0 {
		return
	}
	err = s.ScenarioRepo.RemovePlans(uint(scenarioId), planIds)
	return
}

func (s *ScenarioService) PlanPaginate(req v1.ScenarioPlanReqPaginate, scenarioId int) (ret _domain.PageData, err error) {
	ret, err = s.ScenarioRepo.PlanList(req, scenarioId)
	return
}

func (s *ScenarioService) UpdateStatus(id uint, status consts.TestStatus, updateUserId uint, updateUserName string) (err error) {
	err = s.ScenarioRepo.UpdateStatus(id, status, updateUserId, updateUserName)
	return
}

func (s *ScenarioService) UpdatePriority(id uint, priority string, updateUserId uint, updateUserName string) (err error) {
	err = s.ScenarioRepo.UpdatePriority(id, priority, updateUserId, updateUserName)
	return
}

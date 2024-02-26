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
	UserRepo         *repo2.UserRepo         `inject:""`
}

func (s *ScenarioService) ListByProject(tenantId consts.TenantId, serveId int) (pos []model.Scenario, err error) {
	pos, err = s.ScenarioRepo.ListByProject(tenantId, serveId)
	return
}

func (s *ScenarioService) Paginate(tenantId consts.TenantId, req v1.ScenarioReqPaginate, projectId int) (ret _domain.PageData, err error) {
	ret, err = s.ScenarioRepo.Paginate(tenantId, req, projectId)

	if err != nil {
		return
	}

	return
}

func (s *ScenarioService) GetById(tenantId consts.TenantId, id uint) (scenario model.Scenario, err error) {
	scenario, err = s.ScenarioRepo.Get(tenantId, id)
	if err != nil {
		return
	}

	user, _ := s.UserRepo.GetByUserId(tenantId, scenario.CreateUserId)
	scenario.CreatorName = user.Name
	return
}

func (s *ScenarioService) Create(tenantId consts.TenantId, req model.Scenario) (po model.Scenario, err error) {
	po, err = s.ScenarioRepo.Create(tenantId, req)

	s.ScenarioNodeRepo.CreateDefault(tenantId, po.ID, req.ProjectId, req.CreateUserId)

	return
}

func (s *ScenarioService) Update(tenantId consts.TenantId, req model.Scenario) error {
	return s.ScenarioRepo.Update(tenantId, req)
}

func (s *ScenarioService) DeleteById(tenantId consts.TenantId, id uint) error {
	return s.ScenarioRepo.DeleteById(tenantId, id)
}

func (s *ScenarioService) AddPlans(tenantId consts.TenantId, scenarioId int, planIds []int) (err error) {
	err = s.ScenarioRepo.AddPlans(tenantId, uint(scenarioId), planIds)
	return
}

func (s *ScenarioService) RemovePlans(tenantId consts.TenantId, scenarioId int, planIds []int) (err error) {
	if len(planIds) == 0 {
		return
	}
	err = s.ScenarioRepo.RemovePlans(tenantId, uint(scenarioId), planIds)
	return
}

func (s *ScenarioService) PlanPaginate(tenantId consts.TenantId, req v1.ScenarioPlanReqPaginate, scenarioId int) (ret _domain.PageData, err error) {
	ret, err = s.ScenarioRepo.PlanList(tenantId, req, scenarioId)
	return
}

func (s *ScenarioService) UpdateStatus(tenantId consts.TenantId, id uint, status consts.TestStatus, updateUserId uint, updateUserName string) (err error) {
	err = s.ScenarioRepo.UpdateStatus(tenantId, id, status, updateUserId, updateUserName)
	return
}

func (s *ScenarioService) UpdatePriority(tenantId consts.TenantId, id uint, priority string, updateUserId uint, updateUserName string) (err error) {
	err = s.ScenarioRepo.UpdatePriority(tenantId, id, priority, updateUserId, updateUserName)
	return
}

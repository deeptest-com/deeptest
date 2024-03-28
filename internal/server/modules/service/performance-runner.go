package service

import (
	agentDomain "github.com/aaronchen2k/deeptest/cmd/agent/v1/domain"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	repo "github.com/aaronchen2k/deeptest/internal/server/modules/repo"
)

type PerformanceRunnerService struct {
	PerformanceRunnerRepo *repo.PerformanceRunnerRepo `inject:""`
	UserRepo              *repo.UserRepo              `inject:""`
}

func (s *PerformanceRunnerService) List(scenarioId int) (pos []model.PerformanceRunner, err error) {
	pos, err = s.PerformanceRunnerRepo.List(uint(scenarioId))

	return
}

func (s *PerformanceRunnerService) GetById(id uint) (performanceTestPlan model.PerformanceRunner, err error) {
	performanceTestPlan, err = s.PerformanceRunnerRepo.Get(id)
	if err != nil {
		return
	}

	return
}

func (s *PerformanceRunnerService) Select(req agentDomain.PerformanceRunnerSelectionReq) (err error) {
	err = s.PerformanceRunnerRepo.Select(req)

	return
}

func (s *PerformanceRunnerService) DeleteById(id uint) error {
	return s.PerformanceRunnerRepo.DeleteById(id)
}

func (s *PerformanceRunnerService) UpdateIsConductor(req v1.PerformanceTestUpdateIsConductorReq) (err error) {
	return s.PerformanceRunnerRepo.UpdateIsConductor(req)
}

func (s *PerformanceRunnerService) UpdateWeight(req v1.PerformanceTestUpdateWeightReq) (err error) {
	return s.PerformanceRunnerRepo.UpdateWeight(req)
}

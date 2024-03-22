package service

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	repo "github.com/aaronchen2k/deeptest/internal/server/modules/repo"
)

type PerformanceRunnerService struct {
	PerformanceRunnerRepo *repo.PerformanceRunnerRepo `inject:""`
	UserRepo              *repo.UserRepo              `inject:""`
}

func (s *PerformanceRunnerService) List(scenarioId int) (pos []model.PerformanceRunner, err error) {
	pos, err = s.PerformanceRunnerRepo.List(scenarioId)

	return
}

func (s *PerformanceRunnerService) GetById(id uint) (performanceTestPlan model.PerformanceRunner, err error) {
	performanceTestPlan, err = s.PerformanceRunnerRepo.Get(id)
	if err != nil {
		return
	}

	return
}

func (s *PerformanceRunnerService) Save(req *model.PerformanceRunner) (err error) {
	err = s.PerformanceRunnerRepo.Save(req)

	return
}

func (s *PerformanceRunnerService) DeleteById(id uint) error {
	return s.PerformanceRunnerRepo.DeleteById(id)
}

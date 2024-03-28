package service

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	repo "github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	"github.com/aaronchen2k/deeptest/pkg/domain"
)

type PerformanceTestPlanService struct {
	PerformanceTestPlanRepo *repo.PerformanceTestPlanRepo `inject:""`
	PerformanceRunnerRepo   *repo.PerformanceRunnerRepo   `inject:""`
	UserRepo                *repo.UserRepo                `inject:""`
}

func (s *PerformanceTestPlanService) Paginate(req v1.PerformanceTestPlanReqPaginate, projectId int) (ret _domain.PageData, err error) {
	ret, err = s.PerformanceTestPlanRepo.Paginate(req, projectId)

	if err != nil {
		return
	}

	return
}

func (s *PerformanceTestPlanService) GetById(id uint) (performanceTestPlan model.PerformanceTestPlan, err error) {
	performanceTestPlan, err = s.PerformanceTestPlanRepo.Get(id)
	if err != nil {
		return
	}

	user, _ := s.UserRepo.GetByUserId(performanceTestPlan.CreateUserId)
	performanceTestPlan.CreatorName = user.Name
	return
}

func (s *PerformanceTestPlanService) Create(req model.PerformanceTestPlan) (po model.PerformanceTestPlan, err error) {
	po, err = s.PerformanceTestPlanRepo.Create(req)

	return
}

func (s *PerformanceTestPlanService) Update(req model.PerformanceTestPlan) error {
	return s.PerformanceTestPlanRepo.Update(req)
}

func (s *PerformanceTestPlanService) DeleteById(id uint) error {
	return s.PerformanceTestPlanRepo.DeleteById(id)
}

func (s *PerformanceTestPlanService) UpdateStatus(id uint, status consts.TestStatus, updateUserId uint, updateUserName string) (err error) {
	err = s.PerformanceTestPlanRepo.UpdateStatus(id, status, updateUserId, updateUserName)
	return
}

func (s *PerformanceTestPlanService) GetScenarioId(planId int) (scenarioId uint, err error) {
	po, err := s.PerformanceTestPlanRepo.Get(uint(planId))

	scenarioId = po.ScenarioId

	return
}

func (s *PerformanceTestPlanService) ListRunner(performanceScenarioId int) (runners []model.PerformanceRunner, err error) {
	runners, err = s.PerformanceTestPlanRepo.ListRunner(uint(performanceScenarioId))

	return
}

func (s *PerformanceTestPlanService) GetConductor(planId int) (ret model.PerformanceRunner, err error) {
	plan, err := s.PerformanceTestPlanRepo.Get(uint(planId))

	runners, err := s.PerformanceRunnerRepo.List(plan.ScenarioId)

	for _, runner := range runners {
		if runner.IsConductor {
			ret = runner
			break
		}
	}

	if !ret.IsConductor && len(runners) > 0 {
		ret = runners[0]
	}

	return
}

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

	//s.PerformanceTestPlanNodeRepo.CreateDefault(po.ID, req.ProjectId, req.CreateUserId)

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

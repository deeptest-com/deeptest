package service

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	repo "github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	"github.com/aaronchen2k/deeptest/pkg/domain"
)

type PlanService struct {
	PlanRepo *repo.PlanRepo `inject:""`
}

func NewPlanService() *PlanService {
	return &PlanService{}
}

func (s *PlanService) Paginate(req v1.PlanReqPaginate, projectId int) (ret _domain.PageData, err error) {
	ret, err = s.PlanRepo.Paginate(req, projectId)

	if err != nil {
		return
	}

	return
}

func (s *PlanService) GetById(id uint) (model.Plan, error) {
	return s.PlanRepo.Get(id)
}

func (s *PlanService) Create(req model.Plan) (po model.Plan, bizErr *_domain.BizErr) {
	po, bizErr = s.PlanRepo.Create(req)

	return
}

func (s *PlanService) Update(req model.Plan) error {
	return s.PlanRepo.Update(req)
}

func (s *PlanService) DeleteById(id uint) error {
	return s.PlanRepo.DeleteById(id)
}

func (s *PlanService) AddScenarios(planId int, scenarioIds []int) (err error) {
	err = s.PlanRepo.AddScenarios(planId, scenarioIds)
	return
}

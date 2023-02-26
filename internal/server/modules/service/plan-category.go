package service

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	repo "github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/jinzhu/copier"
)

type PlanCategoryService struct {
	PlanCategoryRepo *repo.PlanCategoryRepo `inject:""`
}

func (s *PlanCategoryService) GetTree(projectId int) (root *v1.PlanCategory, err error) {
	root, err = s.PlanCategoryRepo.GetTree(uint(projectId))

	return
}

func (s *PlanCategoryService) Get(scenarioId int) (root model.PlanCategory, err error) {
	root, err = s.PlanCategoryRepo.Get(uint(scenarioId))

	return
}

func (s *PlanCategoryService) Create(req v1.PlanCategoryCreateReq) (ret model.PlanCategory, err *_domain.BizErr) {
	target, _ := s.PlanCategoryRepo.Get(uint(req.TargetId))
	if target.ID == 0 {
		return
	}

	ret = model.PlanCategory{
		Name:      req.Name,
		ParentId:  req.TargetId,
		ProjectId: req.ProjectId,
	}

	if req.Mode == "child" {
		ret.ParentId = target.ID
	} else if req.Mode == "brother" {
		ret.ParentId = target.ParentId
	}

	ret.Ordr = s.PlanCategoryRepo.GetMaxOrder(ret.ParentId)

	s.PlanCategoryRepo.Save(&ret)

	if req.Mode == "parent" { // move interface to new folder
		target.ParentId = ret.ID
		s.PlanCategoryRepo.Save(&target)
	}

	return
}

func (s *PlanCategoryService) Update(req v1.PlanCategoryReq) (err error) {
	err = s.PlanCategoryRepo.Update(req)
	return
}

func (s *PlanCategoryService) UpdateName(req v1.PlanCategoryReq) (err error) {
	err = s.PlanCategoryRepo.UpdateName(req.Id, req.Name)
	return
}

func (s *PlanCategoryService) Delete(id uint) (err error) {
	err = s.PlanCategoryRepo.Delete(id)
	return
}

func (s *PlanCategoryService) ListToByPlan(id uint) (ret []*model.PlanCategory, err error) {
	pos, _ := s.PlanCategoryRepo.ListByProject(id)

	for _, po := range pos {
		to := model.PlanCategory{}
		copier.CopyWithOption(&to, po, copier.Option{DeepCopy: true})

		ret = append(ret, &to)
	}

	return
}

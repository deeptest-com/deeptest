package service

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	serverConsts "github.com/aaronchen2k/deeptest/internal/server/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	repo "github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
)

type CategoryService struct {
	CategoryRepo *repo.CategoryRepo `inject:""`
}

func (s *CategoryService) GetTree(typ serverConsts.CategoryDiscriminator, projectId, serveId int) (root *v1.Category, err error) {
	root, err = s.CategoryRepo.GetTree(typ, uint(projectId), uint(serveId))

	return
}

func (s *CategoryService) Get(scenarioId int) (root model.Category, err error) {
	root, err = s.CategoryRepo.Get(uint(scenarioId))

	return
}

func (s *CategoryService) Create(req v1.CategoryCreateReq) (ret model.Category, err *_domain.BizErr) {
	target, _ := s.CategoryRepo.Get(req.TargetId)
	if target.ID == 0 {
		return
	}

	ret = model.Category{
		Name:      req.Name,
		ParentId:  req.TargetId,
		ProjectId: req.ProjectId,
	}

	if req.Mode == "child" {
		ret.ParentId = target.ID
	} else if req.Mode == "brother" {
		ret.ParentId = target.ParentId
	}

	ret.Ordr = s.CategoryRepo.GetMaxOrder(ret.ParentId, req.Type, req.ProjectId)

	s.CategoryRepo.Save(&ret)

	if req.Mode == "parent" { // move interface to new folder
		target.ParentId = ret.ID
		s.CategoryRepo.Save(&target)
	}

	return
}

func (s *CategoryService) Update(req v1.CategoryReq) (err error) {
	err = s.CategoryRepo.Update(req)
	return
}

func (s *CategoryService) UpdateName(req v1.CategoryReq) (err error) {
	err = s.CategoryRepo.UpdateName(req.Id, req.Name)
	return
}

func (s *CategoryService) Delete(id uint) (err error) {
	err = s.deleteNodeAndChildren(id)
	return
}

func (s *CategoryService) Move(srcId, targetId uint, pos serverConsts.DropPos, typ serverConsts.CategoryDiscriminator, projectId uint) (
	srcScenarioNode model.Category, err error) {
	srcScenarioNode, err = s.CategoryRepo.Get(srcId)

	srcScenarioNode.ParentId, srcScenarioNode.Ordr = s.CategoryRepo.UpdateOrder(pos, targetId, typ, projectId)
	err = s.CategoryRepo.UpdateOrdAndParent(srcScenarioNode)

	return
}

func (s *CategoryService) deleteNodeAndChildren(nodeId uint) (err error) {
	err = s.CategoryRepo.Delete(nodeId)
	if err == nil {
		children, _ := s.CategoryRepo.GetChildren(nodeId)
		for _, child := range children {
			s.deleteNodeAndChildren(child.ID)
		}
	}

	return
}

package service

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	serverConsts "github.com/aaronchen2k/deeptest/internal/server/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	repo "github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/jinzhu/copier"
)

type ScenarioCategoryService struct {
	ScenarioCategoryRepo *repo.ScenarioCategoryRepo `inject:""`
}

func (s *ScenarioCategoryService) GetTree(projectId, serveId int) (root *v1.ScenarioCategory, err error) {
	root, err = s.ScenarioCategoryRepo.GetTree(uint(projectId), uint(serveId))

	return
}

func (s *ScenarioCategoryService) Get(scenarioId int) (root model.ScenarioCategory, err error) {
	root, err = s.ScenarioCategoryRepo.Get(uint(scenarioId))

	return
}

func (s *ScenarioCategoryService) Create(req v1.ScenarioCategoryCreateReq) (ret model.ScenarioCategory, err *_domain.BizErr) {
	target, _ := s.ScenarioCategoryRepo.Get(uint(req.TargetId))
	if target.ID == 0 {
		return
	}

	ret = model.ScenarioCategory{
		Name:      req.Name,
		ParentId:  req.TargetId,
		ProjectId: req.ProjectId,
	}

	if req.Mode == "child" {
		ret.ParentId = target.ID
	} else if req.Mode == "brother" {
		ret.ParentId = target.ParentId
	}

	ret.Ordr = s.ScenarioCategoryRepo.GetMaxOrder(ret.ParentId)

	s.ScenarioCategoryRepo.Save(&ret)

	if req.Mode == "parent" { // move interface to new folder
		target.ParentId = ret.ID
		s.ScenarioCategoryRepo.Save(&target)
	}

	return
}

func (s *ScenarioCategoryService) Update(req v1.ScenarioCategoryReq) (err error) {
	err = s.ScenarioCategoryRepo.Update(req)
	return
}

func (s *ScenarioCategoryService) UpdateName(req v1.ScenarioCategoryReq) (err error) {
	err = s.ScenarioCategoryRepo.UpdateName(req.Id, req.Name)
	return
}

func (s *ScenarioCategoryService) Delete(id uint) (err error) {
	err = s.deleteScenarioNodeAndChildren(id)
	return
}

func (s *ScenarioCategoryService) Move(srcId, targetId uint, pos serverConsts.DropPos, projectId uint) (
	srcScenarioNode model.ScenarioCategory, err error) {
	srcScenarioNode, err = s.ScenarioCategoryRepo.Get(srcId)

	srcScenarioNode.ParentId, srcScenarioNode.Ordr = s.ScenarioCategoryRepo.UpdateOrder(pos, targetId)
	err = s.ScenarioCategoryRepo.UpdateOrdAndParent(srcScenarioNode)

	return
}

func (s *ScenarioCategoryService) deleteScenarioNodeAndChildren(nodeId uint) (err error) {
	err = s.ScenarioCategoryRepo.Delete(nodeId)
	if err == nil {
		children, _ := s.ScenarioCategoryRepo.GetChildren(nodeId)
		for _, child := range children {
			s.deleteScenarioNodeAndChildren(child.ID)
		}
	}

	return
}

func (s *ScenarioCategoryService) ListToByScenario(id, serveId uint) (ret []*model.ScenarioCategory, err error) {
	pos, _ := s.ScenarioCategoryRepo.ListByProject(id, serveId)

	for _, po := range pos {
		to := model.ScenarioCategory{}
		copier.CopyWithOption(&to, po, copier.Option{DeepCopy: true})

		ret = append(ret, &to)
	}

	return
}

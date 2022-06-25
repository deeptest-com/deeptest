package service

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	serverDomain "github.com/aaronchen2k/deeptest/internal/server/modules/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/repo"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
)

type ScenarioNodeService struct {
	ScenarioNodeRepo *repo.ScenarioNodeRepo `inject:""`
	ScenarioRepo     *repo.ScenarioRepo     `inject:""`
}

func (s *ScenarioNodeService) GetTree(scenarioId int) (root *model.TestProcessor, err error) {
	root, err = s.ScenarioNodeRepo.GetTree(scenarioId)

	return
}

func (s *ScenarioNodeService) AddInterfaces(req serverDomain.ScenarioAddInterfacesReq) (err *_domain.BizErr) {
	targetProcessor, _ := s.ScenarioNodeRepo.Get(req.TargetId)

	for _, interfaceNode := range req.SelectedNodes {
		s.createDirOrInterface(interfaceNode, targetProcessor)
	}

	return
}

func (s *ScenarioNodeService) AddProcessor(req serverDomain.ScenarioAddScenarioReq) (ret model.TestProcessor, err *_domain.BizErr) {
	targetProcessor, _ := s.ScenarioNodeRepo.Get(uint(req.TargetProcessorId))
	if targetProcessor.ID == 0 {
		return
	}

	ret = model.TestProcessor{
		Name:           req.Name,
		EntityCategory: req.ProcessorCategory,
		EntityType:     req.ProcessorType,

		ScenarioId: targetProcessor.ScenarioId,
	}

	if req.Mode == "child" {
		ret.ParentId = targetProcessor.ID

	} else if req.Mode == "parent" && req.TargetProcessorCategory == consts.ProcessorInterface {
		ret.ParentId = targetProcessor.ParentId
	}

	s.ScenarioNodeRepo.Save(&ret)

	if req.Mode == "parent" {
		targetProcessor.ParentId = ret.ID
		s.ScenarioNodeRepo.Save(&targetProcessor)
	}

	return
}

func (s *ScenarioNodeService) createDirOrInterface(interfaceNode serverDomain.InterfaceSimple, parentProcessor model.TestProcessor) (
	err *_domain.BizErr) {

	if !interfaceNode.IsDir {
		processor := model.TestProcessor{
			Name:           interfaceNode.Name,
			ScenarioId:     parentProcessor.ScenarioId,
			EntityCategory: consts.ProcessorInterface,
			InterfaceId:    uint(interfaceNode.Id),
			ParentId:       parentProcessor.ID,
		}
		s.ScenarioNodeRepo.Save(&processor)

	} else {
		processor := model.TestProcessor{
			Name:           interfaceNode.Name,
			ScenarioId:     parentProcessor.ScenarioId,
			EntityCategory: consts.ProcessorGroup,
			ParentId:       parentProcessor.ID,
		}
		s.ScenarioNodeRepo.Save(&processor)

		for _, child := range interfaceNode.Children {
			s.createDirOrInterface(child, processor)
		}
	}

	return
}

func (s *ScenarioNodeService) Get(id int) (po model.TestProcessor, err error) {
	po, _ = s.ScenarioNodeRepo.Get(uint(id))
	return
}

//func (s *ScenarioNodeService) Save(interf *model.ScenarioNode) (err error) {
//	err = s.ScenarioNodeRepo.Save(interf)
//
//	return
//}
//func (s *ScenarioNodeService) Create(req serverDomain.ScenarioNodeReq) (interf *model.ScenarioNode, err error) {
//	interf = &model.ScenarioNode{Name: req.Name, ProjectId: uint(req.ProjectId),
//		IsDir: req.Type == serverConsts.Dir}
//
//	var dropPos serverConsts.DropPos
//	if req.Mode == serverConsts.Child {
//		dropPos = serverConsts.Inner
//	} else {
//		dropPos = serverConsts.After
//	}
//
//	interf.ParentId, interf.Ordr = s.ScenarioNodeRepo.UpdateOrder(dropPos, uint(req.Target))
//	err = s.ScenarioNodeRepo.Save(interf)
//
//	return
//}
//func (s *ScenarioNodeService) Update(id int, req serverDomain.ScenarioNodeReq) (err error) {
//
//	return
//}
func (s *ScenarioNodeService) UpdateName(req serverDomain.ScenarioNodeReq) (err error) {
	err = s.ScenarioNodeRepo.UpdateName(req.Id, req.Name)
	return
}

//func (s *ScenarioNodeService) Delete(projectId, id uint) (err error) {
//	err = s.deleteScenarioNodeAndChildren(projectId, id)
//
//	return
//}
//
//func (s *ScenarioNodeService) Move(srcId, targetId uint, pos serverConsts.DropPos, projectId uint) (
//	srcScenarioNode model.ScenarioNode, err error) {
//	srcScenarioNode, err = s.ScenarioNodeRepo.Get(srcId)
//
//	srcScenarioNode.ParentId, srcScenarioNode.Ordr = s.ScenarioNodeRepo.UpdateOrder(pos, targetId)
//	err = s.ScenarioNodeRepo.UpdateOrdAndParent(srcScenarioNode)
//
//	return
//}
//
//func (s *ScenarioNodeService) deleteScenarioNodeAndChildren(projectId, interfId uint) (err error) {
//	err = s.ScenarioNodeRepo.Delete(interfId)
//	if err == nil {
//		children, _ := s.ScenarioNodeRepo.GetChildren(projectId, interfId)
//		for _, child := range children {
//			s.deleteScenarioNodeAndChildren(child.ProjectId, child.ID)
//		}
//	}
//
//	return
//}

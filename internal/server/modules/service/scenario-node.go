package service

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/agent/exec"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	serverConsts "github.com/aaronchen2k/deeptest/internal/server/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	repo2 "github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/jinzhu/copier"
)

type ScenarioNodeService struct {
	ScenarioNodeRepo      *repo2.ScenarioNodeRepo      `inject:""`
	ScenarioProcessorRepo *repo2.ScenarioProcessorRepo `inject:""`
	ScenarioRepo          *repo2.ScenarioRepo          `inject:""`
}

func (s *ScenarioNodeService) GetTree(scenarioId int) (root *agentExec.Processor, err error) {
	root, err = s.ScenarioNodeRepo.GetTree(uint(scenarioId), false)

	return
}

func (s *ScenarioNodeService) AddInterfaces(req v1.ScenarioAddInterfacesReq) (ret model.Processor, err *_domain.BizErr) {
	targetProcessor, _ := s.ScenarioProcessorRepo.Get(req.TargetId)

	for _, interfaceNode := range req.SelectedNodes {
		ret, _ = s.createDirOrInterface(interfaceNode, targetProcessor)
	}

	return
}

func (s *ScenarioNodeService) AddProcessor(req v1.ScenarioAddScenarioReq) (ret model.Processor, err *_domain.BizErr) {
	targetProcessor, _ := s.ScenarioProcessorRepo.Get(uint(req.TargetProcessorId))
	if targetProcessor.ID == 0 {
		return
	}

	ret = model.Processor{
		Name:           req.Name,
		EntityCategory: req.ProcessorCategory,
		EntityType:     req.ProcessorType,

		ScenarioId: targetProcessor.ScenarioId,
		ProjectId:  req.ProjectId,
	}

	if req.Mode == "child" {
		ret.ParentId = targetProcessor.ID
	} else if req.Mode == "parent" && req.TargetProcessorCategory == consts.ProcessorInterface {
		ret.ParentId = targetProcessor.ParentId
	}
	ret.Ordr = s.ScenarioNodeRepo.GetMaxOrder(ret.ParentId)

	s.ScenarioNodeRepo.Save(&ret)

	if req.Mode == "parent" {
		targetProcessor.ParentId = ret.ID
		s.ScenarioNodeRepo.Save(&targetProcessor)
	}

	if ret.EntityType == consts.ProcessorLogicElse { // create default entity
		entity := model.ProcessorLogic{
			ProcessorEntityBase: agentExec.ProcessorEntityBase{
				ProcessorID:       ret.ID,
				ProcessorCategory: ret.EntityCategory,
				ProcessorType:     ret.EntityType,
			},
		}
		s.ScenarioProcessorRepo.SaveLogic(&entity)
	}

	return
}

func (s *ScenarioNodeService) createDirOrInterface(interfaceNode v1.InterfaceSimple, parentProcessor model.Processor) (
	ret model.Processor, err *_domain.BizErr) {

	if interfaceNode.ParentId == 0 {
		for _, child := range interfaceNode.Children {
			s.createDirOrInterface(child, parentProcessor)
		}

	} else if interfaceNode.IsDir {
		processor := model.Processor{
			Name:           interfaceNode.Name,
			ScenarioId:     parentProcessor.ScenarioId,
			EntityCategory: consts.ProcessorGroup,
			EntityType:     consts.ProcessorGroupDefault,
			ParentId:       parentProcessor.ID,
			ProjectId:      parentProcessor.ProjectId,
		}
		processor.Ordr = s.ScenarioNodeRepo.GetMaxOrder(processor.ParentId)
		s.ScenarioNodeRepo.Save(&processor)

		for _, child := range interfaceNode.Children {
			s.createDirOrInterface(child, processor)
		}

	} else {
		processor := model.Processor{
			Name:           interfaceNode.Name,
			ScenarioId:     parentProcessor.ScenarioId,
			EntityCategory: consts.ProcessorInterface,
			EntityType:     consts.ProcessorInterfaceDefault,
			InterfaceId:    uint(interfaceNode.Id),
			ParentId:       parentProcessor.ID,
			ProjectId:      parentProcessor.ProjectId,
		}
		processor.Ordr = s.ScenarioNodeRepo.GetMaxOrder(processor.ParentId)
		s.ScenarioNodeRepo.Save(&processor)

		ret = processor
	}

	return
}

func (s *ScenarioNodeService) UpdateName(req v1.ScenarioNodeReq) (err error) {
	err = s.ScenarioNodeRepo.UpdateName(req.Id, req.Name)
	return
}

func (s *ScenarioNodeService) Delete(id uint) (err error) {
	err = s.deleteScenarioNodeAndChildren(id)

	return
}

func (s *ScenarioNodeService) Move(srcId, targetId uint, pos serverConsts.DropPos, projectId uint) (
	srcScenarioNode model.Processor, err error) {
	srcScenarioNode, err = s.ScenarioNodeRepo.Get(srcId)

	srcScenarioNode.ParentId, srcScenarioNode.Ordr = s.ScenarioNodeRepo.UpdateOrder(pos, targetId)
	err = s.ScenarioNodeRepo.UpdateOrdAndParent(srcScenarioNode)

	return
}

func (s *ScenarioNodeService) deleteScenarioNodeAndChildren(nodeId uint) (err error) {
	err = s.ScenarioNodeRepo.Delete(nodeId)
	if err == nil {
		children, _ := s.ScenarioNodeRepo.GetChildren(nodeId)
		for _, child := range children {
			s.deleteScenarioNodeAndChildren(child.ID)
		}
	}

	return
}

func (s *ScenarioNodeService) ListToByScenario(id uint) (ret []*agentExec.Processor, err error) {
	pos, _ := s.ScenarioNodeRepo.ListByScenario(id)

	for _, po := range pos {
		to := agentExec.Processor{}
		copier.CopyWithOption(&to, po, copier.Option{DeepCopy: true})

		ret = append(ret, &to)
	}

	return
}

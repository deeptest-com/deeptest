package service

import (
	"encoding/json"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/agent/exec"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	serverConsts "github.com/aaronchen2k/deeptest/internal/server/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	repo "github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/jinzhu/copier"
	"github.com/kataras/iris/v12"
)

type ScenarioNodeService struct {
	ScenarioNodeRepo         *repo.ScenarioNodeRepo      `inject:""`
	ScenarioProcessorRepo    *repo.ScenarioProcessorRepo `inject:""`
	ScenarioProcessorService *ScenarioProcessorService   `inject:""`
	ScenarioRepo             *repo.ScenarioRepo          `inject:""`

	ProcessorInterfaceRepo *repo.ProcessorInterfaceRepo `inject:""`
	EndpointInterfaceRepo  *repo.EndpointInterfaceRepo  `inject:""`
}

func (s *ScenarioNodeService) GetTree(scenario model.Scenario, withDetail bool) (root *agentExec.Processor, err error) {
	pos, err := s.ScenarioNodeRepo.ListByScenario(scenario.ID)
	if err != nil {
		return
	}

	tos := s.ToTos(pos, withDetail)

	root = tos[0]
	root.Name = scenario.Name
	root.Slots = iris.Map{"icon": "icon"}

	s.ScenarioNodeRepo.MakeTree(tos[1:], root)

	root.Session = agentExec.Session{}

	root.ScenarioId = scenario.ID

	return
}

func (s *ScenarioNodeService) ToTos(pos []*model.Processor, withDetail bool) (tos []*agentExec.Processor) {
	for _, po := range pos {
		to := agentExec.Processor{
			ProcessorBase: agentExec.ProcessorBase{
				IsLeaf:     s.ScenarioNodeRepo.IsLeaf(*po),
				Session:    agentExec.Session{},
				ScenarioId: po.ScenarioId,
			},
		}
		copier.CopyWithOption(&to, po, copier.Option{DeepCopy: true})

		if withDetail {
			entity, _ := s.ScenarioProcessorService.GetEntityTo(&to)
			to.EntityRaw, _ = json.Marshal(entity)
		}

		// just to avoid json marshal error for IProcessorEntity
		to.Entity = agentExec.ProcessorGroup{}

		tos = append(tos, &to)
	}

	return
}

func (s *ScenarioNodeService) AddInterfaces(req v1.ScenarioAddInterfacesReq) (ret model.Processor, err error) {
	targetProcessor, _ := s.ScenarioProcessorRepo.Get(req.TargetId)

	if s.ScenarioNodeRepo.IsLeaf(targetProcessor) {
		targetProcessor, _ = s.ScenarioProcessorRepo.Get(targetProcessor.ParentId)
	}

	for _, interfaceId := range req.InterfaceIds {
		ret, _ = s.addInterface(interfaceId, req.CreateBy, targetProcessor)
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
		CreatedBy:  req.CreateBy,
	}

	if req.Mode == "child" {
		ret.ParentId = targetProcessor.ID
	} else if req.Mode == "brother" {
		ret.ParentId = targetProcessor.ParentId
	} else if req.Mode == "parent" && req.TargetProcessorCategory == consts.ProcessorInterface {
		ret.ParentId = targetProcessor.ParentId
	}

	ret.Ordr = s.ScenarioNodeRepo.GetMaxOrder(ret.ParentId)

	s.ScenarioNodeRepo.Save(&ret)

	if req.Mode == "parent" { // move interface to new folder
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

func (s *ScenarioNodeService) addInterface(endpointInterfaceId int, createBy uint, parentProcessor model.Processor) (
	ret model.Processor, err error) {

	endpointInterface, err := s.EndpointInterfaceRepo.Get(uint(endpointInterfaceId))
	if err != nil {
		return
	}

	processor := model.Processor{
		Name: endpointInterface.Name,

		EntityCategory: consts.ProcessorInterface,
		EntityType:     consts.ProcessorInterfaceDefault,

		EndpointInterfaceId: endpointInterface.ID,
		EntityId:            0, // set to 0 for interface processor node
		ParentId:            parentProcessor.ID,
		ScenarioId:          parentProcessor.ScenarioId,
		ProjectId:           parentProcessor.ProjectId,
		CreatedBy:           createBy,
	}
	processor.Ordr = s.ScenarioNodeRepo.GetMaxOrder(processor.ParentId)
	s.ScenarioNodeRepo.Save(&processor)

	//interfaceProcessor := model.ProcessorInterface{}
	// interfaceProcessor, err = s.ScenarioProcessorService.CloneInterface(uint(endpointInterfaceId), processor)
	//if err != nil {
	//	return
	//}
	//s.ScenarioProcessorRepo.UpdateEntityId(processor.ID, interfaceProcessor.ID)

	ret = processor

	return
}

//func (s *ScenarioNodeService) createDirOrInterface(interfaceId int, parentProcessor model.Processor) (
//	ret model.Processor, err error) {
//
//	if interfaceNode.ParentId == 0 {
//		for _, child := range interfaceNode.Children {
//			s.createDirOrInterface(child, parentProcessor)
//		}
//
//	} else if !interfaceNode.IsLeaf {
//		processor := model.Processor{
//			Name:           interfaceNode.Name,
//			ScenarioId:     parentProcessor.ScenarioId,
//			EntityCategory: consts.ProcessorGroup,
//			EntityType:     consts.ProcessorGroupDefault,
//			ParentId:       parentProcessor.ID,
//			ProjectId:      parentProcessor.ProjectId,
//		}
//		processor.Ordr = s.ScenarioNodeRepo.GetMaxOrder(processor.ParentId)
//		s.ScenarioNodeRepo.Save(&processor)
//
//		for _, child := range interfaceNode.Children {
//			s.createDirOrInterface(child, processor)
//		}
//
//	} else {
//		processor := model.Processor{
//			Name:           interfaceNode.Name,
//			ScenarioId:     parentProcessor.ScenarioId,
//			EntityCategory: consts.ProcessorInterface,
//			EntityType:     consts.ProcessorInterfaceDefault,
//			//EntityId:       interfaceProcessor.ID,
//			EndpointInterfaceId: uint(interfaceNode.Id),
//			ParentId:    parentProcessor.ID,
//			ProjectId:   parentProcessor.ProjectId,
//		}
//		processor.Ordr = s.ScenarioNodeRepo.GetMaxOrder(processor.ParentId)
//		s.ScenarioNodeRepo.Save(&processor)
//
//		interfaceProcessor := model.ProcessorInterface{}
//		interfaceProcessor, err = s.ScenarioProcessorService.CloneInterface(uint(interfaceNode.Id), processor)
//		if err != nil {
//			return
//		}
//
//		s.ScenarioProcessorRepo.UpdateEntityId(processor.ID, interfaceProcessor.ID)
//
//		ret = processor
//	}
//
//	return
//}

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

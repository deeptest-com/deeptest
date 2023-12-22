package service

import (
	domain "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	agentExec "github.com/aaronchen2k/deeptest/internal/agent/exec"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	"github.com/jinzhu/copier"
)

type ScenarioProcessorService struct {
	ScenarioProcessorRepo *repo.ScenarioProcessorRepo `inject:""`
	EndpointInterfaceRepo *repo.EndpointInterfaceRepo `inject:""`

	ConditionRepo *repo.ConditionRepo `inject:""`

	DebugInterfaceRepo *repo.DebugInterfaceRepo `inject:""`
	ServeServerRepo    *repo.ServeServerRepo    `inject:""`
	DatapoolRepo       *repo.DatapoolRepo       `inject:""`

	ExtractorService         *ExtractorService         `inject:""`
	CheckpointService        *CheckpointService        `inject:""`
	DebugInterfaceService    *DebugInterfaceService    `inject:""`
	ScenarioInterfaceService *ScenarioInterfaceService `inject:""`
}

func (s *ScenarioProcessorService) GetEntity(id int) (ret interface{}, err error) {
	ret, err = s.ScenarioProcessorRepo.GetEntity(uint(id))
	return
}

func (s *ScenarioProcessorService) UpdateName(req agentExec.ProcessorEntityBase) (err error) {
	err = s.ScenarioProcessorRepo.UpdateName(req.ProcessorID, req.Name)
	return
}

func (s *ScenarioProcessorService) SaveBasicInfo(req domain.ScenarioProcessorInfo) (err error) {
	err = s.ScenarioProcessorRepo.SaveBasicInfo(req)
	return
}

func (s *ScenarioProcessorService) SaveGroup(req *model.ProcessorGroup) (err error) {
	err = s.ScenarioProcessorRepo.SaveGroup(req)
	s.ScenarioProcessorRepo.UpdateName(req.ProcessorID, req.Name)
	return
}

func (s *ScenarioProcessorService) SaveTimer(req *model.ProcessorTimer) (err error) {
	err = s.ScenarioProcessorRepo.SaveTimer(req)
	s.ScenarioProcessorRepo.UpdateName(req.ProcessorID, req.Name)
	return
}

func (s *ScenarioProcessorService) SavePrint(req *model.ProcessorPrint) (err error) {
	err = s.ScenarioProcessorRepo.SavePrint(req)
	s.ScenarioProcessorRepo.UpdateName(req.ProcessorID, req.Name)
	return
}

func (s *ScenarioProcessorService) SaveLogic(req *model.ProcessorLogic) (err error) {
	err = s.ScenarioProcessorRepo.SaveLogic(req)
	s.ScenarioProcessorRepo.UpdateName(req.ProcessorID, req.Name)
	return
}

func (s *ScenarioProcessorService) SaveLoop(req *model.ProcessorLoop) (err error) {
	/*
		if req.ProcessorType == consts.ProcessorLoopTime {
			req.Name = fmt.Sprintf("迭代%d次", req.Times)
		}
	*/

	err = s.ScenarioProcessorRepo.UpdateName(req.ProcessorID, req.Name)

	err = s.ScenarioProcessorRepo.SaveLoop(req)

	return
}

func (s *ScenarioProcessorService) SaveVariable(req *model.ProcessorVariable) (err error) {
	err = s.ScenarioProcessorRepo.SaveVariable(req)
	s.ScenarioProcessorRepo.UpdateName(req.ProcessorID, req.Name)
	return
}

func (s *ScenarioProcessorService) SaveCookie(req *model.ProcessorCookie) (err error) {
	err = s.ScenarioProcessorRepo.SaveCookie(req)
	s.ScenarioProcessorRepo.UpdateName(req.ProcessorID, req.Name)
	return
}

func (s *ScenarioProcessorService) SaveAssertion(req *model.ProcessorAssertion) (err error) {
	err = s.ScenarioProcessorRepo.SaveAssertion(req)
	s.ScenarioProcessorRepo.UpdateName(req.ProcessorID, req.Name)
	return
}

/*
func (s *ScenarioProcessorService) SaveExtractor(req *model.ProcessorExtractor) (err error) {
	err = s.ScenarioProcessorRepo.SaveExtractor(req)
	s.ScenarioProcessorRepo.UpdateName(req.ProcessorID, req.Name)
	return
}
*/

func (s *ScenarioProcessorService) SaveData(req *model.ProcessorData) (err error) {
	err = s.ScenarioProcessorRepo.SaveData(req)
	s.ScenarioProcessorRepo.UpdateName(req.ProcessorID, req.Name)
	return
}

func (s *ScenarioProcessorService) SaveCustomCode(req *model.ProcessorCustomCode) (err error) {
	err = s.ScenarioProcessorRepo.SaveCustomCode(req)
	s.ScenarioProcessorRepo.UpdateName(req.ProcessorID, req.Name)
	return
}

func (s *ScenarioProcessorService) SaveInterface(req *model.ProcessorComm) (err error) {
	err = s.ScenarioProcessorRepo.UpdateName(req.ProcessorID, req.Name)
	return
}

func (s *ScenarioProcessorService) GetEntityTo(processorTo *agentExec.Processor) (ret agentExec.IProcessorEntity, err error) {
	processor, _ := s.ScenarioProcessorRepo.Get(processorTo.ID)

	switch processor.EntityCategory {
	case consts.ProcessorInterface:
		debugData, _ := s.DebugInterfaceService.GetDetail(processor.EntityId)

		interfaceEntity := agentExec.ProcessorInterface{}
		copier.CopyWithOption(&interfaceEntity, debugData, copier.Option{DeepCopy: true})

		//server, _ := s.ServeServerRepo.Get(debugData.ServerId)
		//interfaceEntity.BaseUrl = server.Url
		interfaceEntity.BaseUrl = ""

		interfaceEntity.ProcessorID = processor.ID
		interfaceEntity.ParentID = processor.ParentId
		interfaceEntity.ProcessorCategory = consts.ProcessorInterface
		interfaceEntity.ProcessorType = consts.ProcessorInterfaceDefault
		interfaceEntity.ProcessorInterfaceSrc = debugData.ProcessorInterfaceSrc

		interfaceEntity.PreConditions, _ = s.ConditionRepo.ListTo(processor.EntityId, processor.EndpointInterfaceId, consts.ScenarioDebug, "false", consts.ConditionSrcPre)
		interfaceEntity.PostConditions, _ = s.ConditionRepo.ListTo(processor.EntityId, processor.EndpointInterfaceId, consts.ScenarioDebug, "false", consts.ConditionSrcPost)

		ret = &interfaceEntity

	case consts.ProcessorRoot:
		commEntityPo, _ := s.ScenarioProcessorRepo.GetRoot(processor)

		ret = agentExec.ProcessorRoot{}
		copier.CopyWithOption(&ret, commEntityPo, copier.Option{DeepCopy: true})

	case consts.ProcessorGroup:
		entityPo, _ := s.ScenarioProcessorRepo.GetGroup(processor)
		ret = agentExec.ProcessorGroup{}
		copier.CopyWithOption(&ret, entityPo, copier.Option{DeepCopy: true})

	case consts.ProcessorLogic:
		entityPo, _ := s.ScenarioProcessorRepo.GetLogic(processor)
		ret = agentExec.ProcessorLogic{}
		copier.CopyWithOption(&ret, entityPo, copier.Option{DeepCopy: true})

	case consts.ProcessorLoop:
		entityPo, _ := s.ScenarioProcessorRepo.GetLoop(processor)
		ret = agentExec.ProcessorLoop{}
		copier.CopyWithOption(&ret, entityPo, copier.Option{DeepCopy: true})

	case consts.ProcessorVariable:
		entityPo, _ := s.ScenarioProcessorRepo.GetVariable(processor)
		ret = agentExec.ProcessorVariable{}
		copier.CopyWithOption(&ret, entityPo, copier.Option{DeepCopy: true})

	case consts.ProcessorTimer:
		entityPo, _ := s.ScenarioProcessorRepo.GetTimer(processor)
		ret = agentExec.ProcessorTimer{}
		copier.CopyWithOption(&ret, entityPo, copier.Option{DeepCopy: true})

	case consts.ProcessorPrint:
		entityPo, _ := s.ScenarioProcessorRepo.GetPrint(processor)
		ret = agentExec.ProcessorPrint{}
		copier.CopyWithOption(&ret, entityPo, copier.Option{DeepCopy: true})

	case consts.ProcessorCookie:
		entityPo, _ := s.ScenarioProcessorRepo.GetCookie(processor)
		ret = agentExec.ProcessorCookie{}
		copier.CopyWithOption(&ret, entityPo, copier.Option{DeepCopy: true})

	case consts.ProcessorAssertion:
		entityPo, _ := s.ScenarioProcessorRepo.GetAssertion(processor)
		ret = agentExec.ProcessorAssertion{}
		copier.CopyWithOption(&ret, entityPo, copier.Option{DeepCopy: true})

	case consts.ProcessorData:
		entityPo, _ := s.ScenarioProcessorRepo.GetData(processor)
		processorData := agentExec.ProcessorData{}
		copier.CopyWithOption(&processorData, entityPo, copier.Option{DeepCopy: true})

		datapool, _ := s.DatapoolRepo.Get(entityPo.DatapoolId)
		processorData.DatapoolName = datapool.Name

		ret = processorData

	case consts.ProcessorCustomCode:
		entityPo, _ := s.ScenarioProcessorRepo.GetCustomCode(processor)
		ret = agentExec.ProcessorCustomCode{}
		copier.CopyWithOption(&ret, entityPo, copier.Option{DeepCopy: true})

	default:
	}

	return
}

//func (s *ScenarioProcessorService) CloneInterface(interfaceId uint, processor modelRef.Processor) (ret modelRef.ProcessorInterface, err error) {
//	interf, err := s.EndpointInterfaceRepo.GetDetail(interfaceId)
//	if err != nil {
//		return
//	}
//
//	copier.CopyWithOption(&ret, interf, copier.Option{DeepCopy: true})
//
//	ret.ProcessorId = processor.ID
//	ret.ScenarioId = processor.ScenarioId
//	ret.ID = 0
//	ret.CreatedAt = nil
//
//	err = s.ScenarioInterfaceRepo.SaveInterface(&ret)
//
//	s.CopyExtractors(interfaceId, ret.ID, processor)
//	s.CopyCheckpoints(interfaceId, ret.ID, processor)
//
//	return
//}
//
//func (s *ScenarioProcessorService) CopyExtractors(interfaceId, processorInterfaceId uint, processor modelRef.Processor) {
//	pos, _ := s.ExtractorService.Index(interfaceId, consts.InterfaceDebug)
//
//	for _, po := range pos {
//		extractor := modelRef.DebugInterfaceExtractor{}
//
//		copier.CopyWithOption(&extractor, po, copier.Option{DeepCopy: true})
//		extractor.ID = 0
//		extractor.UsedBy = consts.ScenarioDebug
//		extractor.EndpointInterfaceId = processorInterfaceId
//		extractor.ScenarioId = processor.ScenarioId
//
//		s.ExtractorRepo.SaveDebugData(&extractor)
//	}
//
//	return
//}
//
//func (s *ScenarioProcessorService) CopyCheckpoints(interfaceId, processorInterfaceId uint, processor modelRef.Processor) {
//	pos, _ := s.CheckpointService.Index(interfaceId, consts.InterfaceDebug)
//
//	for _, po := range pos {
//		checkpoint := modelRef.InterfaceCheckpoint{}
//
//		copier.CopyWithOption(&checkpoint, po, copier.Option{DeepCopy: true})
//		checkpoint.ID = 0
//		checkpoint.UsedBy = consts.ScenarioDebug
//		checkpoint.EndpointInterfaceId = processorInterfaceId
//		checkpoint.ScenarioId = processor.ScenarioId
//
//		s.CheckpointRepo.SaveDebugData(&checkpoint)
//	}
//
//	return
//}

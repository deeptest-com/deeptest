package service

import (
	domain "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	agentExec "github.com/aaronchen2k/deeptest/internal/agent/exec"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	v1 "github.com/aaronchen2k/deeptest/internal/pkg/domain"
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
	EnvironmentService       *EnvironmentService       `inject:""`
	DebugSceneService        *DebugSceneService        `inject:""`
}

func (s *ScenarioProcessorService) GetEntity(tenantId consts.TenantId, id int) (ret interface{}, err error) {
	ret, err = s.ScenarioProcessorRepo.GetEntity(tenantId, uint(id))
	return
}

func (s *ScenarioProcessorService) UpdateName(tenantId consts.TenantId, req agentExec.ProcessorEntityBase) (err error) {
	err = s.ScenarioProcessorRepo.UpdateName(tenantId, req.ProcessorID, req.Name)
	return
}

func (s *ScenarioProcessorService) SaveBasicInfo(tenantId consts.TenantId, req domain.ScenarioProcessorInfo) (err error) {
	err = s.ScenarioProcessorRepo.SaveBasicInfo(tenantId, req)
	return
}

func (s *ScenarioProcessorService) SaveGroup(tenantId consts.TenantId, req *model.ProcessorGroup) (err error) {
	err = s.ScenarioProcessorRepo.SaveGroup(tenantId, req)
	s.ScenarioProcessorRepo.UpdateName(tenantId, req.ProcessorID, req.Name)
	return
}

func (s *ScenarioProcessorService) SaveTimer(tenantId consts.TenantId, req *model.ProcessorTimer) (err error) {
	err = s.ScenarioProcessorRepo.SaveTimer(tenantId, req)
	s.ScenarioProcessorRepo.UpdateName(tenantId, req.ProcessorID, req.Name)
	return
}

func (s *ScenarioProcessorService) SavePrint(tenantId consts.TenantId, req *model.ProcessorPrint) (err error) {
	err = s.ScenarioProcessorRepo.SavePrint(tenantId, req)
	s.ScenarioProcessorRepo.UpdateName(tenantId, req.ProcessorID, req.Name)
	return
}

func (s *ScenarioProcessorService) SaveLogic(tenantId consts.TenantId, req *model.ProcessorLogic) (err error) {
	err = s.ScenarioProcessorRepo.SaveLogic(tenantId, req)
	s.ScenarioProcessorRepo.UpdateName(tenantId, req.ProcessorID, req.Name)
	return
}

func (s *ScenarioProcessorService) SaveLoop(tenantId consts.TenantId, req *model.ProcessorLoop) (err error) {
	/*
		if req.ProcessorType == consts.ProcessorLoopTime {
			req.Name = fmt.Sprintf("迭代%d次", req.Times)
		}
	*/

	err = s.ScenarioProcessorRepo.UpdateName(tenantId, req.ProcessorID, req.Name)

	err = s.ScenarioProcessorRepo.SaveLoop(tenantId, req)

	return
}

func (s *ScenarioProcessorService) SaveVariable(tenantId consts.TenantId, req *model.ProcessorVariable) (err error) {
	err = s.ScenarioProcessorRepo.SaveVariable(tenantId, req)
	s.ScenarioProcessorRepo.UpdateName(tenantId, req.ProcessorID, req.Name)
	return
}

func (s *ScenarioProcessorService) SaveCookie(tenantId consts.TenantId, req *model.ProcessorCookie) (err error) {
	err = s.ScenarioProcessorRepo.SaveCookie(tenantId, req)
	s.ScenarioProcessorRepo.UpdateName(tenantId, req.ProcessorID, req.Name)
	return
}

func (s *ScenarioProcessorService) SaveAssertion(tenantId consts.TenantId, req *model.ProcessorAssertion) (err error) {
	err = s.ScenarioProcessorRepo.SaveAssertion(tenantId, req)
	s.ScenarioProcessorRepo.UpdateName(tenantId, req.ProcessorID, req.Name)
	return
}

/*
func (s *ScenarioProcessorService) SaveExtractor(req *model.ProcessorExtractor) (err error) {
	err = s.ScenarioProcessorRepo.SaveExtractor(req)
	s.ScenarioProcessorRepo.UpdateName(req.ProcessorID, req.Name)
	return
}
*/

func (s *ScenarioProcessorService) SaveData(tenantId consts.TenantId, req *model.ProcessorData) (err error) {
	err = s.ScenarioProcessorRepo.SaveData(tenantId, req)
	s.ScenarioProcessorRepo.UpdateName(tenantId, req.ProcessorID, req.Name)
	return
}

func (s *ScenarioProcessorService) SaveCustomCode(tenantId consts.TenantId, req *model.ProcessorCustomCode) (err error) {
	err = s.ScenarioProcessorRepo.SaveCustomCode(tenantId, req)
	s.ScenarioProcessorRepo.UpdateName(tenantId, req.ProcessorID, req.Name)
	return
}

func (s *ScenarioProcessorService) SaveInterface(tenantId consts.TenantId, req *model.ProcessorComm) (err error) {
	err = s.ScenarioProcessorRepo.UpdateName(tenantId, req.ProcessorID, req.Name)
	return
}

func (s *ScenarioProcessorService) GetEntityTo(tenantId consts.TenantId, processorTo *agentExec.Processor) (ret agentExec.IProcessorEntity, err error) {
	processor, _ := s.ScenarioProcessorRepo.Get(tenantId, processorTo.ID)

	switch processor.EntityCategory {
	case consts.ProcessorInterface:
		debugData, _ := s.DebugInterfaceService.GetDetail(tenantId, processor.EntityId)

		//合并全局参数到debugdata 到 GlobalParams，在场景执行中全局参数使用的是 debugData.GlobalParams，所以要提取合并
		globalParams, _ := s.EnvironmentService.GetGlobalParams(tenantId, debugData.ProjectId)
		globalParams = s.DebugInterfaceService.MergeGlobalParams(globalParams, debugData.GlobalParams)
		endpointInterfaceGlobalParams, _ := s.EndpointInterfaceRepo.GetGlobalParams(tenantId, debugData.EndpointInterfaceId, debugData.ProjectId)
		debugData.GlobalParams = s.MergeGlobalParams(endpointInterfaceGlobalParams, globalParams)

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

		interfaceEntity.PreConditions, _ = s.ConditionRepo.ListTo(tenantId, processor.EntityId, processor.EndpointInterfaceId, consts.ScenarioDebug, "false", consts.ConditionSrcPre)
		interfaceEntity.PostConditions, _ = s.ConditionRepo.ListTo(tenantId, processor.EntityId, processor.EndpointInterfaceId, consts.ScenarioDebug, "false", consts.ConditionSrcPost)

		ret = &interfaceEntity

	case consts.ProcessorRoot:
		commEntityPo, _ := s.ScenarioProcessorRepo.GetRoot(tenantId, processor)

		ret = agentExec.ProcessorRoot{}
		copier.CopyWithOption(&ret, commEntityPo, copier.Option{DeepCopy: true})

	case consts.ProcessorGroup:
		entityPo, _ := s.ScenarioProcessorRepo.GetGroup(tenantId, processor)
		ret = agentExec.ProcessorGroup{}
		copier.CopyWithOption(&ret, entityPo, copier.Option{DeepCopy: true})

	case consts.ProcessorLogic:
		entityPo, _ := s.ScenarioProcessorRepo.GetLogic(tenantId, processor)
		ret = agentExec.ProcessorLogic{}
		copier.CopyWithOption(&ret, entityPo, copier.Option{DeepCopy: true})

	case consts.ProcessorLoop:
		entityPo, _ := s.ScenarioProcessorRepo.GetLoop(tenantId, processor)
		ret = agentExec.ProcessorLoop{}
		copier.CopyWithOption(&ret, entityPo, copier.Option{DeepCopy: true})

	case consts.ProcessorVariable:
		entityPo, _ := s.ScenarioProcessorRepo.GetVariable(tenantId, processor)
		ret = agentExec.ProcessorVariable{}
		copier.CopyWithOption(&ret, entityPo, copier.Option{DeepCopy: true})

	case consts.ProcessorTimer:
		entityPo, _ := s.ScenarioProcessorRepo.GetTimer(tenantId, processor)
		ret = agentExec.ProcessorTimer{}
		copier.CopyWithOption(&ret, entityPo, copier.Option{DeepCopy: true})

	case consts.ProcessorPrint:
		entityPo, _ := s.ScenarioProcessorRepo.GetPrint(tenantId, processor)
		ret = agentExec.ProcessorPrint{}
		copier.CopyWithOption(&ret, entityPo, copier.Option{DeepCopy: true})

	case consts.ProcessorCookie:
		entityPo, _ := s.ScenarioProcessorRepo.GetCookie(tenantId, processor)
		ret = agentExec.ProcessorCookie{}
		copier.CopyWithOption(&ret, entityPo, copier.Option{DeepCopy: true})

	case consts.ProcessorAssertion:
		entityPo, _ := s.ScenarioProcessorRepo.GetAssertion(tenantId, processor)
		ret = agentExec.ProcessorAssertion{}
		copier.CopyWithOption(&ret, entityPo, copier.Option{DeepCopy: true})

	case consts.ProcessorData:
		entityPo, _ := s.ScenarioProcessorRepo.GetData(tenantId, processor)
		processorData := agentExec.ProcessorData{}
		copier.CopyWithOption(&processorData, entityPo, copier.Option{DeepCopy: true})

		datapool, _ := s.DatapoolRepo.Get(tenantId, entityPo.DatapoolId)
		processorData.DatapoolName = datapool.Name

		ret = processorData

	case consts.ProcessorCustomCode:
		entityPo, _ := s.ScenarioProcessorRepo.GetCustomCode(tenantId, processor)
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

func (s *ScenarioProcessorService) MergeGlobalParams(endpointInterfaceGlobalParams []model.EndpointInterfaceGlobalParam, globalParams []v1.GlobalParam) (ret []model.DebugInterfaceGlobalParam) {

	for _, item := range globalParams {
		b := true
		for _, param := range endpointInterfaceGlobalParams {
			if param.Name == item.Name && param.In == item.In && param.Disabled {
				b = false
				break
			}
		}

		if b {
			param := model.DebugInterfaceGlobalParam{GlobalParam: v1.GlobalParam{Name: item.Name, DefaultValue: item.DefaultValue, In: item.In, Type: item.Type}}
			ret = append(ret, param)
		}

	}

	return
}

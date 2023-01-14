package service

import (
	agentExec "github.com/aaronchen2k/deeptest/internal/agent/exec"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	"github.com/jinzhu/copier"
)

type ScenarioProcessorService struct {
	ScenarioProcessorRepo *repo.ScenarioProcessorRepo  `inject:""`
	ScenarioInterfaceRepo *repo.ProcessorInterfaceRepo `inject:""`
	InterfaceRepo         *repo.InterfaceRepo          `inject:""`
	ExtractorRepo         *repo.ExtractorRepo          `inject:""`
	CheckpointRepo        *repo.CheckpointRepo         `inject:""`

	ExtractorService  *ExtractorService  `inject:""`
	CheckpointService *CheckpointService `inject:""`
}

func (s *ScenarioProcessorService) GetEntity(id int) (ret interface{}, err error) {
	ret, err = s.ScenarioProcessorRepo.GetEntity(uint(id))
	return
}

func (s *ScenarioProcessorService) UpdateName(req agentExec.ProcessorEntityBase) (err error) {
	err = s.ScenarioProcessorRepo.UpdateName(req.ProcessorID, req.Name)
	return
}

func (s *ScenarioProcessorService) SaveGroup(req *model.ProcessorGroup) (err error) {
	err = s.ScenarioProcessorRepo.SaveGroup(req)
	return
}

func (s *ScenarioProcessorService) SaveTimer(req *model.ProcessorTimer) (err error) {
	err = s.ScenarioProcessorRepo.SaveTimer(req)
	return
}

func (s *ScenarioProcessorService) SavePrint(req *model.ProcessorPrint) (err error) {
	err = s.ScenarioProcessorRepo.SavePrint(req)
	return
}

func (s *ScenarioProcessorService) SaveLogic(req *model.ProcessorLogic) (err error) {
	err = s.ScenarioProcessorRepo.SaveLogic(req)
	return
}

func (s *ScenarioProcessorService) SaveLoop(req *model.ProcessorLoop) (err error) {
	err = s.ScenarioProcessorRepo.SaveLoop(req)
	return
}

func (s *ScenarioProcessorService) SaveVariable(req *model.ProcessorVariable) (err error) {
	err = s.ScenarioProcessorRepo.SaveVariable(req)
	return
}

func (s *ScenarioProcessorService) SaveCookie(req *model.ProcessorCookie) (err error) {
	err = s.ScenarioProcessorRepo.SaveCookie(req)
	return
}

func (s *ScenarioProcessorService) SaveAssertion(req *model.ProcessorAssertion) (err error) {
	err = s.ScenarioProcessorRepo.SaveAssertion(req)
	return
}

func (s *ScenarioProcessorService) SaveExtractor(req *model.ProcessorExtractor) (err error) {
	err = s.ScenarioProcessorRepo.SaveExtractor(req)
	return
}

func (s *ScenarioProcessorService) SaveData(req *model.ProcessorData) (err error) {
	err = s.ScenarioProcessorRepo.SaveData(req)
	return
}

func (s *ScenarioProcessorService) CloneInterface(interfaceId uint, processor model.Processor) (ret model.ProcessorInterface, err error) {
	interf, err := s.InterfaceRepo.GetDetail(interfaceId)
	if err != nil {
		return
	}

	copier.CopyWithOption(&ret, interf, copier.Option{DeepCopy: true})

	ret.ProcessorId = processor.ID
	ret.ScenarioId = processor.ScenarioId
	ret.ID = 0
	ret.CreatedAt = nil

	err = s.ScenarioInterfaceRepo.SaveInterface(&ret)

	s.CopyExtractors(interfaceId, ret.ID, processor)
	s.CopyCheckpoints(interfaceId, ret.ID, processor)

	return
}

func (s *ScenarioProcessorService) CopyExtractors(interfaceId, processorInterfaceId uint, processor model.Processor) {
	pos, _ := s.ExtractorService.List(interfaceId, consts.Interface)

	for _, po := range pos {
		extractor := model.InterfaceExtractor{}

		copier.CopyWithOption(&extractor, po, copier.Option{DeepCopy: true})
		extractor.ID = 0
		extractor.UsedBy = consts.Scenario
		extractor.InterfaceId = processorInterfaceId
		extractor.ProcessorId = processor.ID
		extractor.ScenarioId = processor.ScenarioId

		s.ExtractorRepo.Save(&extractor)
	}

	return
}

func (s *ScenarioProcessorService) CopyCheckpoints(interfaceId, processorInterfaceId uint, processor model.Processor) {
	pos, _ := s.CheckpointService.List(interfaceId, consts.Interface)

	for _, po := range pos {
		checkpoint := model.InterfaceCheckpoint{}

		copier.CopyWithOption(&checkpoint, po, copier.Option{DeepCopy: true})
		checkpoint.ID = 0
		checkpoint.UsedBy = consts.Scenario
		checkpoint.InterfaceId = processorInterfaceId
		checkpoint.ScenarioId = processor.ScenarioId

		s.CheckpointRepo.Save(&checkpoint)
	}

	return
}

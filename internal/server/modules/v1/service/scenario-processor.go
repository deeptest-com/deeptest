package service

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/repo"
)

type ScenarioProcessorService struct {
	ScenarioProcessorRepo *repo.ScenarioProcessorRepo `inject:""`
}

func (s *ScenarioProcessorService) Get(processorId int) (ret interface{}, err error) {
	processor, _ := s.ScenarioProcessorRepo.Get(uint(processorId))

	if processor.EntityCategory == consts.ProcessorInterface {
		ret, _ = s.ScenarioProcessorRepo.GetInterface(uint(processorId), processor)

	} else if processor.EntityCategory == consts.ProcessorGroup {
		ret, _ = s.ScenarioProcessorRepo.GetGroup(uint(processorId), processor)

	} else if processor.EntityCategory == consts.ProcessorLogic {
		ret, _ = s.ScenarioProcessorRepo.GetLogic(uint(processorId), processor)

	} else if processor.EntityCategory == consts.ProcessorVariable {
		ret, _ = s.ScenarioProcessorRepo.GetVariable(uint(processorId), processor)

	} else if processor.EntityCategory == consts.ProcessorTimer {
		ret, _ = s.ScenarioProcessorRepo.GetTimer(uint(processorId), processor)

	} else if processor.EntityCategory == consts.ProcessorCookie {
		ret, _ = s.ScenarioProcessorRepo.GetCookie(uint(processorId), processor)

	}

	return
}

func (s *ScenarioProcessorService) UpdateName(req model.ProcessorEntity) (err error) {
	err = s.ScenarioProcessorRepo.UpdateName(req.ProcessorId, req.Name)
	return
}

func (s *ScenarioProcessorService) SaveGroup(req model.ProcessorGroup) (err error) {
	err = s.ScenarioProcessorRepo.SaveGroup(req)
	return
}

func (s *ScenarioProcessorService) SaveTimer(req model.ProcessorTimer) (err error) {
	err = s.ScenarioProcessorRepo.SaveTimer(req)
	return
}

func (s *ScenarioProcessorService) SaveLogic(req model.ProcessorLogic) (err error) {
	err = s.ScenarioProcessorRepo.SaveLogic(req)
	return
}

func (s *ScenarioProcessorService) SaveVariable(req model.ProcessorVariable) (err error) {
	err = s.ScenarioProcessorRepo.SaveVariable(req)
	return
}

func (s *ScenarioProcessorService) SaveCookie(req model.ProcessorCookie) (err error) {
	err = s.ScenarioProcessorRepo.SaveCookie(req)
	return
}

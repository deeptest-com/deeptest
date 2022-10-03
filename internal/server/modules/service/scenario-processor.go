package service

import (
	agentDomain "github.com/aaronchen2k/deeptest/internal/agent/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
)

type ScenarioProcessorService struct {
	ScenarioProcessorRepo *repo.ScenarioProcessorRepo `inject:""`
}

func (s *ScenarioProcessorService) Get(processorId int) (ret interface{}, err error) {
	processor, _ := s.ScenarioProcessorRepo.Get(uint(processorId))

	if processor.EntityCategory == consts.ProcessorInterface {
		ret, _ = s.ScenarioProcessorRepo.GetInterface(processor)

	} else if processor.EntityCategory == consts.ProcessorGroup {
		ret, _ = s.ScenarioProcessorRepo.GetGroup(processor)

	} else if processor.EntityCategory == consts.ProcessorLogic {
		ret, _ = s.ScenarioProcessorRepo.GetLogic(processor)

	} else if processor.EntityCategory == consts.ProcessorLoop {
		ret, _ = s.ScenarioProcessorRepo.GetLoop(processor)

	} else if processor.EntityCategory == consts.ProcessorVariable {
		ret, _ = s.ScenarioProcessorRepo.GetVariable(processor)

	} else if processor.EntityCategory == consts.ProcessorTimer {
		ret, _ = s.ScenarioProcessorRepo.GetTimer(processor)

	} else if processor.EntityCategory == consts.ProcessorPrint {
		ret, _ = s.ScenarioProcessorRepo.GetPrint(processor)

	} else if processor.EntityCategory == consts.ProcessorCookie {
		ret, _ = s.ScenarioProcessorRepo.GetCookie(processor)

	} else if processor.EntityCategory == consts.ProcessorAssertion {
		ret, _ = s.ScenarioProcessorRepo.GetAssertion(processor)

	} else if processor.EntityCategory == consts.ProcessorExtractor {
		ret, _ = s.ScenarioProcessorRepo.GetExtractor(processor)

	} else if processor.EntityCategory == consts.ProcessorData {
		ret, _ = s.ScenarioProcessorRepo.GetData(processor)

	}

	return
}

func (s *ScenarioProcessorService) UpdateName(req agentDomain.ProcessorEntity) (err error) {
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

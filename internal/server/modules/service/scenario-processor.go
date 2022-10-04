package service

import (
	agentDomain "github.com/aaronchen2k/deeptest/internal/agent/exec"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
)

type ScenarioProcessorService struct {
	ScenarioProcessorRepo *repo.ScenarioProcessorRepo `inject:""`
}

func (s *ScenarioProcessorService) GetEntity(id int) (ret interface{}, err error) {
	ret, err = s.ScenarioProcessorRepo.GetEntity(uint(id))
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

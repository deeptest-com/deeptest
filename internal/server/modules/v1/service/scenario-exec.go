package service

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/repo"
	_stringUtils "github.com/aaronchen2k/deeptest/pkg/lib/string"
	"time"
)

type ScenarioExecService struct {
	ScenarioProcessorRepo *repo.ScenarioProcessorRepo `inject:""`
	ScenarioRepo          *repo.ScenarioRepo          `inject:""`
	TestResultRepo        *repo.TestResultRepo        `inject:""`
}

func (s *ScenarioExecService) Load(scenarioId int) (result model.TestResult, err error) {
	scenario, err := s.ScenarioRepo.Get(uint(scenarioId))
	if err != nil {
		return
	}

	result.Name = scenario.Name

	return
}

func (s *ScenarioExecService) ExecScenario(scenarioId int) (err error) {
	scenario, err := s.ScenarioRepo.Get(uint(scenarioId))
	if err != nil {
		return
	}

	result, err := s.TestResultRepo.FindInProgressResult(uint(scenarioId))
	if result.ID > 0 {
		s.ResetResult(&result, scenario)
	} else {
		result, _ = s.CreateResult(scenario)
	}

	rootProcessor, err := s.ScenarioProcessorRepo.GetRootProcessor(scenario.ID)
	if err != nil {
		return
	}

	s.ExecRecursiveProcessor(rootProcessor)

	return
}

func (s *ScenarioExecService) ExecRecursiveProcessor(processor model.TestProcessor) (err error) {
	if s.isContainerProcessor(processor.EntityCategory) {
		if s.isExecutableContainerProcessor(processor.EntityCategory) {
			s.ExecContainerProcessor(processor)
		}

		children, _ := s.ScenarioProcessorRepo.GetChildrenProcessor(processor.ID, processor.ScenarioId)
		for _, child := range children {
			s.ExecRecursiveProcessor(child)
		}
	} else if processor.EntityCategory == consts.ProcessorInterface {
		s.ExecInterface(processor)
	} else {
		s.ExecActionProcessor(processor)
	}

	return
}

func (s *ScenarioExecService) ExecContainerProcessor(processor model.TestProcessor) (err error) {

	return
}

func (s *ScenarioExecService) ExecActionProcessor(processor model.TestProcessor) (err error) {

	return
}

func (s *ScenarioExecService) ExecInterface(interf model.TestProcessor) (err error) {

	return
}

func (s *ScenarioExecService) CreateResult(scenario model.TestScenario) (result model.TestResult, err error) {
	startTime := time.Now()
	result = model.TestResult{
		Name:           scenario.Name,
		StartTime:      &startTime,
		ProgressStatus: consts.InProgress,
		ScenarioId:     scenario.ID,
	}

	s.TestResultRepo.Create(&result)

	return
}

func (s *ScenarioExecService) ResetResult(result *model.TestResult, scenario model.TestScenario) (err error) {
	result.Name = scenario.Name

	startTime := time.Now()
	result.StartTime = &startTime

	s.TestResultRepo.ResetResult(*result)
	s.TestResultRepo.ClearLogs(result.ID)

	return
}

func (s *ScenarioExecService) isContainerProcessor(category consts.ProcessorCategory) bool {
	arr := []string{
		consts.ProcessorRoot.ToString(),
		//consts.ProcessorThreadGroup.ToString(),
		consts.ProcessorGroup.ToString(),
		consts.ProcessorLogic.ToString(),
		consts.ProcessorLoop.ToString(),
		consts.ProcessorData.ToString(),
	}
	return _stringUtils.FindInArr(category.ToString(), arr)
}

func (s *ScenarioExecService) isExecutableContainerProcessor(category consts.ProcessorCategory) bool {
	arr := []string{
		//consts.ProcessorThreadGroup.ToString(),
		consts.ProcessorLogic.ToString(),
		consts.ProcessorLoop.ToString(),
		consts.ProcessorData.ToString(),
	}
	return _stringUtils.FindInArr(category.ToString(), arr)
}

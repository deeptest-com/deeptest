package service

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/repo"
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

func (s *ScenarioExecService) Exec(scenarioId int) (err error) {
	scenario, err := s.ScenarioRepo.Get(uint(scenarioId))
	if err != nil {
		return
	}

	po, err := s.TestResultRepo.FindInProgressResult(uint(scenarioId))
	if po.ID > 0 {
		s.RefreshResult(po, scenario)
	} else {
		s.CreateResult(scenario)
	}

	return
}

func (s *ScenarioExecService) CreateResult(scenario model.TestScenario) (err error) {
	startTime := time.Now()
	result := model.TestResult{
		Name:           scenario.Name,
		StartTime:      &startTime,
		ProgressStatus: consts.InProgress,
		ScenarioId:     scenario.ID,
	}

	s.TestResultRepo.Create(&result)

	return
}

func (s *ScenarioExecService) RefreshResult(po model.TestResult, scenario model.TestScenario) (err error) {
	po.Name = scenario.Name

	startTime := time.Now()
	po.StartTime = &startTime

	s.TestResultRepo.RefreshResult(po)

	return
}

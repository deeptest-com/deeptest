package service

import (
	"github.com/aaronchen2k/deeptest/internal/agent/exec"
	execDomain "github.com/aaronchen2k/deeptest/internal/agent/exec/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	"sync"
)

var (
	breakMap sync.Map
)

type ScenarioExecService struct {
	ScenarioRepo     *repo.ScenarioRepo     `inject:""`
	ScenarioNodeRepo *repo.ScenarioNodeRepo `inject:""`
	TestReportRepo   *repo.ReportRepo       `inject:""`
	TestLogRepo      *repo.LogRepo          `inject:""`

	EnvironmentService *EnvironmentService `inject:""`
}

func (s *ScenarioExecService) Load(scenarioId int) (result domain.Report, err error) {
	scenario, err := s.ScenarioRepo.Get(uint(scenarioId))
	if err != nil {
		return
	}

	result.Name = scenario.Name

	return
}

func (s *ScenarioExecService) LoadExecData(scenarioId int) (execReq agentExec.ProcessorExecObj, err error) {
	rootProcessor, _ := s.ScenarioNodeRepo.GetTree(uint(scenarioId), true)
	execReq.Variables, _ = s.EnvironmentService.ListVariableForExec(uint(scenarioId))

	execReq.RootProcessor = rootProcessor

	return
}

func (s *ScenarioExecService) SaveReport(scenarioId int, rootResult execDomain.Result) (err error) {
	scenario, _ := s.ScenarioRepo.Get(uint(scenarioId))
	rootResult.Name = scenario.Name

	report := model.Report{
		Name:      scenario.Name,
		StartTime: rootResult.StartTime,
		EndTime:   rootResult.EndTime,
		Duration:  rootResult.EndTime.Unix() - rootResult.StartTime.Unix(),

		ProgressStatus: rootResult.ProgressStatus,
		ResultStatus:   rootResult.ResultStatus,

		ScenarioId: scenario.ID,
		ProjectId:  scenario.ProjectId,
	}

	s.countRequest(rootResult, &report)
	s.summarizeInterface(&report)

	s.TestReportRepo.Create(&report)
	s.TestLogRepo.CreateLogs(rootResult, &report)

	return
}

func (s ScenarioExecService) countRequest(result execDomain.Result, report *model.Report) {
	if result.ProcessorType == consts.ProcessorInterfaceDefault {
		s.countInterface(result.InterfaceId, result.ResultStatus, report)

		report.TotalRequestNum++

		switch result.ResultStatus {
		case consts.Pass:
			report.PassRequestNum++

		case consts.Fail:
			report.FailRequestNum++
			report.ResultStatus = consts.Fail

		default:
		}

	} else if result.ProcessorType == consts.ProcessorAssertionDefault {
		switch result.ResultStatus {
		case consts.Pass:
			report.PassAssertionNum++

		case consts.Fail:
			report.FailAssertionNum++
			report.ResultStatus = consts.Fail

		default:
		}
	}

	if result.Children == nil {
		return
	}

	for _, log := range result.Children {
		s.countRequest(*log, report)
	}
}

func (s ScenarioExecService) countInterface(interfaceId uint, status consts.ResultStatus, report *model.Report) {
	if report.InterfaceStatusMap == nil {
		report.InterfaceStatusMap = map[uint]map[consts.ResultStatus]int{}
	}

	if report.InterfaceStatusMap[interfaceId] == nil {
		report.InterfaceStatusMap[interfaceId] = map[consts.ResultStatus]int{}
		report.InterfaceStatusMap[interfaceId][consts.Pass] = 0
		report.InterfaceStatusMap[interfaceId][consts.Fail] = 0
	}

	switch status {
	case consts.Pass:
		report.InterfaceStatusMap[interfaceId][consts.Pass]++

	case consts.Fail:
		report.InterfaceStatusMap[interfaceId][consts.Fail]++

	default:
	}
}

func (s ScenarioExecService) summarizeInterface(report *model.Report) {
	for _, val := range report.InterfaceStatusMap {
		if val[consts.Fail] > 0 {
			report.FailInterfaceNum++
		} else {
			report.PassInterfaceNum++
		}

		report.TotalInterfaceNum++
	}
}

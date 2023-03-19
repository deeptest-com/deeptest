package service

import (
	"github.com/aaronchen2k/deeptest/internal/agent/exec"
	agentDomain "github.com/aaronchen2k/deeptest/internal/agent/exec/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
)

type PlanExecService struct {
	PlanRepo       *repo.PlanRepo       `inject:""`
	PlanReportRepo *repo.PlanReportRepo `inject:""`
	TestLogRepo    *repo.LogRepo        `inject:""`

	ScenarioExecService *ScenarioExecService `inject:""`

	EnvironmentService *EnvironmentService `inject:""`
	DatapoolService    *DatapoolService    `inject:""`
}

func (s *PlanExecService) LoadExecResult(planId int) (result domain.Report, err error) {
	plan, err := s.PlanRepo.Get(uint(planId))
	if err != nil {
		return
	}

	result.Name = plan.Name

	return
}

func (s *PlanExecService) LoadExecData(planId int) (ret agentExec.PlanExecObj, err error) {
	plan, err := s.PlanRepo.Get(uint(planId))
	if err != nil {
		return
	}

	scenarios, err := s.PlanRepo.ListScenario(plan.ID)
	for _, scenario := range scenarios {
		scenarioExecObj, _ := s.ScenarioExecService.LoadExecData(scenario.ID)
		ret.Scenarios = append(ret.Scenarios, scenarioExecObj)
	}

	ret.Name = plan.Name

	return
}

func (s *PlanExecService) SaveReport(planId int, result agentDomain.PlanExecResult) (
	report model.PlanReport, err error) {

	report.ProjectId = uint(planId)
	report.ProgressStatus = consts.End
	report.ResultStatus = consts.Pass

	for _, scenarioResult := range result.Scenarios {
		scenarioReport, _ := s.ScenarioExecService.SaveReport(scenarioResult.ID, *scenarioResult)
		s.CombineReport(scenarioReport, &report)
	}

	report.Duration = report.EndTime.Unix() - report.StartTime.Unix()

	return
}
func (s *PlanExecService) CombineReport(scenarioReport model.ScenarioReport, planReport *model.PlanReport) (
	report model.PlanReport, err error) {

	report.InterfaceStatusMap = map[uint]map[consts.ResultStatus]int{}

	if report.StartTime == nil || report.StartTime.Unix() > scenarioReport.StartTime.Unix() {
		report.StartTime = scenarioReport.StartTime
	}
	if report.EndTime == nil || report.EndTime.Unix() < scenarioReport.EndTime.Unix() {
		report.EndTime = scenarioReport.EndTime
	}

	if scenarioReport.ProgressStatus != consts.End {
		report.ProgressStatus = scenarioReport.ProgressStatus
	}
	if scenarioReport.ResultStatus != consts.Pass {
		report.ResultStatus = scenarioReport.ResultStatus
	}

	report.TotalRequestNum += scenarioReport.TotalRequestNum
	report.PassRequestNum += scenarioReport.PassRequestNum
	report.FailRequestNum += scenarioReport.FailRequestNum

	report.TotalAssertionNum += scenarioReport.TotalAssertionNum
	report.PassAssertionNum += scenarioReport.PassAssertionNum
	report.FailAssertionNum += scenarioReport.FailAssertionNum

	for keyId := range scenarioReport.InterfaceStatusMap {
		if report.InterfaceStatusMap[keyId] == nil {
			report.InterfaceStatusMap[keyId] = map[consts.ResultStatus]int{}
			report.InterfaceStatusMap[keyId][consts.Pass] = 0
			report.InterfaceStatusMap[keyId][consts.Fail] = 0
		}

		if _, ok := scenarioReport.InterfaceStatusMap[keyId][consts.Pass]; ok {
			report.InterfaceStatusMap[keyId][consts.Pass] += scenarioReport.InterfaceStatusMap[keyId][consts.Pass]
		}
		if _, ok := scenarioReport.InterfaceStatusMap[keyId][consts.Fail]; ok {
			report.InterfaceStatusMap[keyId][consts.Fail] += scenarioReport.InterfaceStatusMap[keyId][consts.Fail]
		}

	}

	return
}

func (s *PlanExecService) countRequest(result agentDomain.ScenarioExecResult, report *model.PlanReport) {
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

func (s *PlanExecService) countInterface(interfaceId uint, status consts.ResultStatus, report *model.PlanReport) {
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

func (s *PlanExecService) summarizeInterface(report *model.PlanReport) {
	for _, val := range report.InterfaceStatusMap {
		if val[consts.Fail] > 0 {
			report.FailInterfaceNum++
		} else {
			report.PassInterfaceNum++
		}

		report.TotalInterfaceNum++
	}
}

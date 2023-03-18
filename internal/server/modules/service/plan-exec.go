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
	//plan, _ := s.PlanRepo.Get(uint(planId))
	//rootResult.Name = plan.Name
	//
	//report := model.PlanReport{
	//	Name:      plan.Name,
	//	StartTime: rootResult.StartTime,
	//	EndTime:   rootResult.EndTime,
	//	Duration:  rootResult.EndTime.Unix() - rootResult.StartTime.Unix(),
	//
	//	ProgressStatus: rootResult.ProgressStatus,
	//	ResultStatus:   rootResult.ResultStatus,
	//
	//	PlanId:    plan.ID,
	//	ProjectId: plan.ProjectId,
	//}
	//
	//s.countRequest(rootResult, &report)
	//s.summarizeInterface(&report)
	//
	//s.PlanReportRepo.Create(&report)
	//s.TestLogRepo.CreateLogs(rootResult, &report)

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

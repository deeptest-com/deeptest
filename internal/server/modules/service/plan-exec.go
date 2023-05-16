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
	PlanRepo           *repo.PlanRepo           `inject:""`
	PlanReportRepo     *repo.PlanReportRepo     `inject:""`
	ScenarioReportRepo *repo.ScenarioReportRepo `inject:""`
	TestLogRepo        *repo.LogRepo            `inject:""`
	EnvironmentRepo    *repo.EnvironmentRepo    `inject:""`

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

func (s *PlanExecService) LoadExecData(planId, environmentId int) (ret agentExec.PlanExecObj, err error) {
	plan, err := s.PlanRepo.Get(uint(planId))
	if err != nil {
		return
	}

	scenarios, err := s.PlanRepo.ListScenario(plan.ID)
	for _, scenario := range scenarios {
		scenarioExecObj, _ := s.ScenarioExecService.LoadExecData(scenario.ID, uint(environmentId))
		ret.Scenarios = append(ret.Scenarios, scenarioExecObj)
	}

	ret.Name = plan.Name

	return
}

func (s *PlanExecService) SaveReport(planId int, userId uint, result agentDomain.PlanExecResult) (
	report model.PlanReport, err error) {
	plan, err := s.PlanRepo.Get(uint(planId))
	if err != nil {
		return
	}
	projectId := plan.ProjectId

	report.PlanId = uint(planId)
	report.ProjectId = projectId
	report.Name = plan.Name
	report.ExecEnvId = result.EnvId
	report.CreateUserId = userId
	report.ProgressStatus = consts.End
	report.ResultStatus = consts.Pass

	scenarioIds := make([]uint, 0)
	for _, scenarioResult := range result.Scenarios {
		scenarioReport, _ := s.ScenarioExecService.GenerateReport(int(scenarioResult.ScenarioId), userId, *scenarioResult)
		s.CombineReport(scenarioReport, &report)
		scenarioIds = append(scenarioIds, scenarioReport.ID)
	}

	report.Duration = report.EndTime.Unix() - report.StartTime.Unix()
	_ = s.PlanReportRepo.Create(&report)

	_ = s.ScenarioReportRepo.BatchUpdatePlanReportId(scenarioIds, report.ID)

	return
}
func (s *PlanExecService) CombineReport(scenarioReport model.ScenarioReport, planReport *model.PlanReport) (err error) {

	planReport.InterfaceStatusMap = map[uint]map[consts.ResultStatus]int{}

	if planReport.StartTime == nil || planReport.StartTime.Unix() > scenarioReport.StartTime.Unix() {
		planReport.StartTime = scenarioReport.StartTime
	}
	if planReport.EndTime == nil || planReport.EndTime.Unix() < scenarioReport.EndTime.Unix() {
		planReport.EndTime = scenarioReport.EndTime
	}

	if scenarioReport.ProgressStatus != consts.End {
		planReport.ProgressStatus = scenarioReport.ProgressStatus
	}
	if scenarioReport.ResultStatus != consts.Pass {
		planReport.ResultStatus = scenarioReport.ResultStatus
	}

	planReport.TotalRequestNum += scenarioReport.TotalRequestNum
	planReport.PassRequestNum += scenarioReport.PassRequestNum
	planReport.FailRequestNum += scenarioReport.FailRequestNum

	planReport.TotalAssertionNum += scenarioReport.TotalAssertionNum
	planReport.PassAssertionNum += scenarioReport.PassAssertionNum
	planReport.FailAssertionNum += scenarioReport.FailAssertionNum

	planReport.TotalInterfaceNum += scenarioReport.TotalInterfaceNum
	planReport.PassInterfaceNum += scenarioReport.PassInterfaceNum
	planReport.FailInterfaceNum += scenarioReport.FailInterfaceNum

	planReport.TotalScenarioNum += 1
	if scenarioReport.ResultStatus == consts.Pass {
		planReport.PassScenarioNum += 1
	} else if scenarioReport.ResultStatus == consts.Fail {
		planReport.FailScenarioNum += 1
	}

	for keyId := range scenarioReport.InterfaceStatusMap {
		if planReport.InterfaceStatusMap[keyId] == nil {
			planReport.InterfaceStatusMap[keyId] = map[consts.ResultStatus]int{}
			planReport.InterfaceStatusMap[keyId][consts.Pass] = 0
			planReport.InterfaceStatusMap[keyId][consts.Fail] = 0
		}

		if _, ok := scenarioReport.InterfaceStatusMap[keyId][consts.Pass]; ok {
			planReport.InterfaceStatusMap[keyId][consts.Pass] += scenarioReport.InterfaceStatusMap[keyId][consts.Pass]
		}
		if _, ok := scenarioReport.InterfaceStatusMap[keyId][consts.Fail]; ok {
			planReport.InterfaceStatusMap[keyId][consts.Fail] += scenarioReport.InterfaceStatusMap[keyId][consts.Fail]
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

func (s *PlanExecService) GetPlanReportNormalData(planId, environmentId uint) (ret agentDomain.PlanNormalData, err error) {
	plan, err := s.PlanRepo.Get(planId)
	if err != nil {
		return
	}

	environment, err := s.EnvironmentRepo.Get(environmentId)
	if err != nil {
		return
	}

	scenarioNum, err := s.PlanRepo.GetScenarioNumByPlan(planId)
	if err != nil {
		return
	}

	ret.PlanId = planId
	ret.PlanName = plan.Name
	ret.ExecEnv = environment.Name
	ret.TotalScenarioNum = int(scenarioNum)
	return

}

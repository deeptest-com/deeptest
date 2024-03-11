package service

import (
	"encoding/json"
	"github.com/aaronchen2k/deeptest/internal/agent/exec"
	agentDomain "github.com/aaronchen2k/deeptest/internal/agent/exec/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	"gorm.io/gorm"
)

type PlanExecService struct {
	PlanRepo           *repo.PlanRepo           `inject:""`
	PlanReportRepo     *repo.PlanReportRepo     `inject:""`
	ScenarioReportRepo *repo.ScenarioReportRepo `inject:""`
	TestLogRepo        *repo.LogRepo            `inject:""`
	EnvironmentRepo    *repo.EnvironmentRepo    `inject:""`
	ScenarioNodeRepo   *repo.ScenarioNodeRepo   `inject:""`

	ScenarioExecService *ScenarioExecService `inject:""`

	EnvironmentService *EnvironmentService `inject:""`
	DatapoolService    *DatapoolService    `inject:""`
}

func (s *PlanExecService) LoadExecResult(tenantId consts.TenantId, planId int) (result domain.Report, err error) {
	plan, err := s.PlanRepo.Get(tenantId, uint(planId))
	if err != nil {
		return
	}

	result.Name = plan.Name

	return
}

func (s *PlanExecService) LoadExecData(tenantId consts.TenantId, planId, environmentId int) (ret agentExec.PlanExecObj, err error) {
	_ = s.PlanRepo.UpdateCurrEnvId(tenantId, uint(planId), uint(environmentId))

	plan, err := s.PlanRepo.Get(tenantId, uint(planId))
	if err != nil {
		return
	}

	scenarios, err := s.PlanRepo.ListScenario(tenantId, plan.ID)
	for _, scenario := range scenarios {
		scenarioExecObj, _ := s.ScenarioExecService.LoadExecData(tenantId, scenario.ID, uint(environmentId))
		ret.Scenarios = append(ret.Scenarios, scenarioExecObj)
	}

	ret.Name = plan.Name

	return
}

func (s *PlanExecService) SaveReport(tenantId consts.TenantId, planId int, userId uint, result agentDomain.PlanExecResult) (
	report model.PlanReport, err error) {
	plan, err := s.PlanRepo.Get(tenantId, uint(planId))
	if err != nil {
		return
	}
	projectId := plan.ProjectId

	report.PlanId = uint(planId)
	report.ProjectId = projectId
	report.Name = plan.Name
	report.ExecEnvId = uint(result.EnvironmentId)
	report.CreateUserId = userId
	report.ProgressStatus = consts.End
	report.ResultStatus = consts.Pass

	stat, _ := json.Marshal(result.Stat)
	report.StatRaw = string(stat)

	scenarioReportIds := make([]uint, 0)
	for _, scenarioResult := range result.Scenarios {
		scenarioReport, _ := s.ScenarioExecService.GenerateReport(tenantId, int(scenarioResult.ScenarioId), userId, *scenarioResult)
		s.CombineReport(tenantId, scenarioReport, &report)
		scenarioReportIds = append(scenarioReportIds, scenarioResult.ScenarioReportId)
	}

	//report.Duration = report.EndTime.UnixMilli() - report.StartTime.UnixMilli()
	_ = s.PlanReportRepo.Create(tenantId, &report)

	_ = s.ScenarioReportRepo.BatchUpdatePlanReportId(tenantId, scenarioReportIds, report.ID)

	return
}
func (s *PlanExecService) CombineReport(tenantId consts.TenantId, scenarioReport model.ScenarioReport, planReport *model.PlanReport) (err error) {

	planReport.InterfaceStatusMap = map[uint]map[consts.ResultStatus]int{}

	if planReport.StartTime == nil || planReport.StartTime.UnixMilli() > scenarioReport.StartTime.UnixMilli() {
		planReport.StartTime = scenarioReport.StartTime
	}
	if planReport.EndTime == nil || planReport.EndTime.UnixMilli() < scenarioReport.EndTime.UnixMilli() {
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

	planReport.TotalProcessorNum += scenarioReport.TotalProcessorNum
	planReport.FinishProcessorNum += scenarioReport.FinishProcessorNum

	planReport.Duration += scenarioReport.Duration

	planReport.TotalScenarioNum += 1
	if scenarioReport.ResultStatus == consts.Fail {
		planReport.FailScenarioNum += 1
	} else {
		planReport.PassScenarioNum += 1
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
		s.countInterface(result.DebugInterfaceId, result.ResultStatus, report)

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

func (s *PlanExecService) GetPlanReportNormalData(tenantId consts.TenantId, planId, environmentId uint) (ret agentDomain.Report, err error) {
	plan, err := s.PlanRepo.Get(tenantId, planId)
	if err != nil {
		return
	}

	environment, err := s.EnvironmentRepo.Get(tenantId, environmentId)
	if err != nil {
		return
	}

	planScenarioRelation, err := s.PlanRepo.ListScenarioRelation(tenantId, planId)
	if err != nil && err != gorm.ErrRecordNotFound {
		return
	}

	if len(planScenarioRelation) > 0 {
		scenarioIds := make([]uint, 0)
		for _, v := range planScenarioRelation {
			scenarioIds = append(scenarioIds, v.ScenarioId)
		}

		interfaceNum, err := s.ScenarioNodeRepo.GetNumberByScenariosAndEntityCategory(tenantId, scenarioIds, "processor_interface")
		if err != nil {
			return ret, err
		}
		assertionNum, err := s.ScenarioNodeRepo.GetNumberByScenariosAndEntityCategory(tenantId, scenarioIds, "processor_assertion")
		if err != nil {
			return ret, err
		}
		processorNum, err := s.ScenarioNodeRepo.GetNumberByScenariosAndEntityCategory(tenantId, scenarioIds, "")
		if err != nil {
			return ret, err
		}
		ret.TotalInterfaceNum = int(interfaceNum)
		ret.TotalAssertionNum = int(assertionNum)
		ret.TotalProcessorNum = int(processorNum)
	}

	ret.PlanId = planId
	ret.PlanName = plan.Name
	ret.ExecEnv = environment.Name
	ret.TotalScenarioNum = len(planScenarioRelation)
	return

}

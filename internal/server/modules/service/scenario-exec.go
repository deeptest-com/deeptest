package service

import (
	"encoding/json"
	agentDomain2 "github.com/deeptest-com/deeptest/internal/agent/domain"
	agentExec "github.com/deeptest-com/deeptest/internal/agent/exec"
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	"github.com/deeptest-com/deeptest/internal/pkg/domain"
	"github.com/deeptest-com/deeptest/internal/server/modules/model"
	"github.com/deeptest-com/deeptest/internal/server/modules/repo"
	"sync"
)

var (
	breakMap sync.Map
)

type ScenarioExecService struct {
	ScenarioRepo          *repo.ScenarioRepo          `inject:""`
	ScenarioProcessorRepo *repo.ScenarioProcessorRepo `inject:""`
	ScenarioNodeRepo      *repo.ScenarioNodeRepo      `inject:""`
	ScenarioReportRepo    *repo.ScenarioReportRepo    `inject:""`
	DebugInterfaceRepo    *repo.DebugInterfaceRepo    `inject:""`
	TestLogRepo           *repo.LogRepo               `inject:""`
	EnvironmentRepo       *repo.EnvironmentRepo       `inject:""`

	EndpointInterfaceRepo *repo.EndpointInterfaceRepo `inject:""`
	EndpointRepo          *repo.EndpointRepo          `inject:""`
	ServeServerRepo       *repo.ServeServerRepo       `inject:""`

	DebugInvokeService    *DebugInvokeService    `inject:""`
	SceneService          *SceneService          `inject:""`
	EnvironmentService    *EnvironmentService    `inject:""`
	DatapoolService       *DatapoolService       `inject:""`
	ScenarioNodeService   *ScenarioNodeService   `inject:""`
	ScenarioReportService *ScenarioReportService `inject:""`

	ExecConditionService *ExecConditionService `inject:""`
}

func (s *ScenarioExecService) LoadExecResult(tenantId consts.TenantId, scenarioId int) (result domain.Report, err error) {
	scenario, err := s.ScenarioRepo.Get(tenantId, uint(scenarioId))
	if err != nil {
		return
	}

	result.Name = scenario.Name

	return
}

func (s *ScenarioExecService) LoadExecData(tenantId consts.TenantId, scenarioId, environmentId uint) (ret agentExec.ScenarioExecObj, err error) {
	scenario, err := s.ScenarioRepo.Get(tenantId, scenarioId)
	if err != nil {
		return
	}

	// get processor tree
	ret.ScenarioId = scenarioId
	ret.Name = scenario.Name
	ret.RootProcessor, _ = s.ScenarioNodeService.GetTree(tenantId, scenario, true)

	// get variables
	s.SceneService.LoadEnvVarMapByScenario(tenantId, &ret.ExecScene, scenarioId, environmentId)
	s.SceneService.LoadProjectSettings(tenantId, &ret.ExecScene, scenario.ProjectId)

	return
}

func (s *ScenarioExecService) SaveReport(tenantId consts.TenantId, scenarioId int, userId uint, rootResult agentDomain2.ScenarioExecResult) (report model.ScenarioReport, err error) {
	scenario, _ := s.ScenarioRepo.Get(tenantId, uint(scenarioId))
	rootResult.Name = scenario.Name

	report = model.ScenarioReport{
		Name:      scenario.Name,
		StartTime: rootResult.StartTime,
		EndTime:   rootResult.EndTime,
		Duration:  rootResult.EndTime.UnixMilli() - rootResult.StartTime.UnixMilli(),

		ProgressStatus: rootResult.ProgressStatus,
		ResultStatus:   rootResult.ResultStatus,

		ScenarioId:   scenario.ID,
		ProjectId:    scenario.ProjectId,
		CreateUserId: userId,
		ExecEnvId:    rootResult.EnvironmentId,
	}

	stat, _ := json.Marshal(rootResult.Stat)
	report.StatRaw = string(stat)

	// generate report
	s.countRequest(tenantId, rootResult, &report)
	s.summarizeInterface(&report)

	s.ScenarioReportRepo.Create(tenantId, &report)

	// deal with interface and custom code processor's conditions
	processorToInvokeIdMap := map[uint]uint{}
	for _, result := range rootResult.Children {
		err = s.dealwithResult(tenantId, result, &processorToInvokeIdMap)
	}

	// create logs
	s.TestLogRepo.CreateLogs(tenantId, rootResult, &report, processorToInvokeIdMap)

	logs, _ := s.ScenarioReportService.GetById(tenantId, report.ID)
	report.Logs = logs.Logs
	report.Priority = scenario.Priority

	return
}

func (s *ScenarioExecService) dealwithResult(tenantId consts.TenantId, result *agentDomain2.ScenarioExecResult, processorToInvokeIdMap *map[uint]uint) (
	err error) {

	processor, err := s.ScenarioProcessorRepo.Get(tenantId, result.ProcessorId)
	debugInterface, err := s.DebugInterfaceRepo.Get(tenantId, processor.EntityId)

	if result.ProcessorType == consts.ProcessorInterfaceDefault {
		request := domain.DebugData{}
		json.Unmarshal([]byte(result.ReqContent), &request)
		request.DebugInterfaceId = debugInterface.ID
		request.EndpointInterfaceId = debugInterface.EndpointInterfaceId
		request.CaseInterfaceId = debugInterface.CaseInterfaceId
		request.DiagnoseInterfaceId = debugInterface.DiagnoseInterfaceId
		request.ScenarioProcessorId = debugInterface.ScenarioProcessorId
		request.UsedBy = consts.ScenarioDebug
		request.ServeId = debugInterface.ServeId
		request.ServerId = debugInterface.ServerId
		request.ProjectId = debugInterface.ProjectId
		request.BaseUrl = debugInterface.BaseUrl

		response := domain.DebugResponse{}
		json.Unmarshal([]byte(result.RespContent), &response)

		req := domain.SubmitDebugResultRequest{
			Request:        request,
			Response:       response,
			PreConditions:  result.PreConditions,
			PostConditions: result.PostConditions,
		}
		invoke, _ := s.DebugInvokeService.SubmitResult(tenantId, req)
		(*processorToInvokeIdMap)[result.ProcessorId] = invoke.ID

	} else if len(result.Children) > 0 {
		for _, result := range result.Children {
			err = s.dealwithResult(tenantId, result, processorToInvokeIdMap)
		}
	}

	return
}

func (s *ScenarioExecService) GenerateReport(tenantId consts.TenantId, scenarioId int, userId uint, rootResult agentDomain2.ScenarioExecResult) (report model.ScenarioReport, err error) {
	scenario, _ := s.ScenarioRepo.Get(tenantId, uint(scenarioId))
	rootResult.Name = scenario.Name

	report = model.ScenarioReport{
		Name:      scenario.Name,
		StartTime: rootResult.StartTime,
		EndTime:   rootResult.EndTime,
		Duration:  rootResult.EndTime.Unix() - rootResult.StartTime.Unix(),

		ProgressStatus: rootResult.ProgressStatus,
		ResultStatus:   rootResult.ResultStatus,

		ScenarioId:   scenario.ID,
		ProjectId:    scenario.ProjectId,
		CreateUserId: userId,
	}

	s.countRequest(tenantId, rootResult, &report)
	s.summarizeInterface(&report)

	//s.ScenarioReportRepo.CreateExpression(&report)
	//s.TestLogRepo.CreateLogs(rootResult, &report)

	return
}

func (s *ScenarioExecService) countRequest(tenantId consts.TenantId, result agentDomain2.ScenarioExecResult, report *model.ScenarioReport) {
	report.TotalProcessorNum++
	report.FinishProcessorNum++
	if result.ProcessorType == consts.ProcessorInterfaceDefault {
		s.countInterface(result.DebugInterfaceId, result.ResultStatus, report)

		report.TotalRequestNum++
		report.Duration += result.Cost

		switch result.ResultStatus {
		case consts.Pass:
			report.PassRequestNum++

		case consts.Fail:
			report.FailRequestNum++
			report.ResultStatus = consts.Fail

		default:
		}

	} else if result.ProcessorType == consts.ProcessorAssertionDefault {
		report.TotalAssertionNum++
		switch result.ResultStatus {
		case consts.Pass:
			report.PassAssertionNum++

		case consts.Fail:
			report.FailAssertionNum++
			report.ResultStatus = consts.Fail

		default:
		}
	} else if result.ProcessorType == consts.ProcessorCustomCodeDefault {
		report.TotalAssertionNum += result.Stat.CheckpointPass + result.Stat.CheckpointFail

		report.PassAssertionNum += result.Stat.CheckpointPass
		report.FailAssertionNum += result.Stat.CheckpointFail

		if result.Stat.CheckpointFail > 0 {
			report.ResultStatus = consts.Fail
		}
	}

	if result.Children == nil {
		return
	}

	for _, log := range result.Children {
		s.countRequest(tenantId, *log, report)
	}
}

func (s *ScenarioExecService) countInterface(interfaceId uint, status consts.ResultStatus, report *model.ScenarioReport) {
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

func (s *ScenarioExecService) summarizeInterface(report *model.ScenarioReport) {
	for _, val := range report.InterfaceStatusMap {
		if val[consts.Fail] > 0 {
			report.FailInterfaceNum++
		} else {
			report.PassInterfaceNum++
		}

		report.TotalInterfaceNum++
	}
}

func (s *ScenarioExecService) GetScenarioNormalData(tenantId consts.TenantId, id, environmentId uint) (ret agentDomain2.Report, err error) {
	_ = s.ScenarioRepo.UpdateCurrEnvId(tenantId, id, environmentId)
	ret.ScenarioId = id

	environment, err := s.EnvironmentRepo.Get(tenantId, environmentId)
	if err != nil {
		return
	}
	ret.ExecEnv = environment.Name

	scenarioIds := []uint{id}
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

	return

}

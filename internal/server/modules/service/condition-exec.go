package service

import (
	"encoding/json"
	valueUtils "github.com/deeptest-com/deeptest/internal/agent/exec/utils/value"
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	"github.com/deeptest-com/deeptest/internal/pkg/domain"
	repo "github.com/deeptest-com/deeptest/internal/server/modules/repo"
)

type ExecConditionService struct {
	ExtractorRepo   *repo.ExtractorRepo   `inject:""`
	CheckpointRepo  *repo.CheckpointRepo  `inject:""`
	DatabaseOptRepo *repo.DatabaseOptRepo `inject:""`

	ScriptRepo         *repo.ScriptRepo         `inject:""`
	ResponseDefineRepo *repo.ResponseDefineRepo `inject:""`
	ShareVarService    *ShareVarService         `inject:""`
}

func (s *ExecConditionService) SavePreConditionResult(tenantId consts.TenantId, invokeId,
	debugInterfaceId, caseInterfaceId, endpointInterfaceId, serveId, processorId, scenarioId uint, usedBy consts.UsedBy,
	preConditions []domain.InterfaceExecCondition) (err error) {

	for _, condition := range preConditions {
		if condition.Type == consts.ConditionTypeScript {
			var scriptBase domain.ScriptBase
			json.Unmarshal(condition.Raw, &scriptBase)
			if scriptBase.Disabled {
				continue
			}

			s.dealwithScriptResult(tenantId, scriptBase, invokeId,
				debugInterfaceId, caseInterfaceId, endpointInterfaceId, serveId, processorId, scenarioId, usedBy)

		} else if condition.Type == consts.ConditionTypeDatabase {
			var databaseOptBase domain.DatabaseOptBase
			json.Unmarshal(condition.Raw, &databaseOptBase)
			if databaseOptBase.Disabled {
				continue
			}

			s.dealwithDbOptResult(tenantId, databaseOptBase, invokeId,
				debugInterfaceId, caseInterfaceId, endpointInterfaceId, serveId, processorId, scenarioId, usedBy)
		}
	}

	return
}

func (s *ExecConditionService) SavePostConditionResult(tenantId consts.TenantId, invokeId,
	debugInterfaceId, caseInterfaceId, endpointInterfaceId, serveId, processorId, scenarioId uint, usedBy consts.UsedBy,
	postConditions []domain.InterfaceExecCondition) (err error) {

	for _, condition := range postConditions {
		if condition.Type == consts.ConditionTypeScript {
			var scriptBase domain.ScriptBase
			json.Unmarshal(condition.Raw, &scriptBase)
			if scriptBase.Disabled {
				continue
			}

			s.dealwithScriptResult(tenantId, scriptBase, invokeId,
				debugInterfaceId, caseInterfaceId, endpointInterfaceId, serveId, processorId, scenarioId, usedBy)

		} else if condition.Type == consts.ConditionTypeDatabase {
			var databaseOptBase domain.DatabaseOptBase
			json.Unmarshal(condition.Raw, &databaseOptBase)
			if databaseOptBase.Disabled {
				continue
			}

			s.dealwithDbOptResult(tenantId, databaseOptBase, invokeId,
				debugInterfaceId, caseInterfaceId, endpointInterfaceId, serveId, processorId, scenarioId, usedBy)

		} else if condition.Type == consts.ConditionTypeExtractor {
			var extractorBase domain.ExtractorBase
			json.Unmarshal(condition.Raw, &extractorBase)
			if extractorBase.Disabled {
				continue
			}

			s.dealwithExtractorResult(tenantId, extractorBase, invokeId,
				debugInterfaceId, caseInterfaceId, endpointInterfaceId, serveId, processorId, scenarioId, usedBy)

		} else if condition.Type == consts.ConditionTypeCheckpoint {
			var checkpointBase domain.CheckpointBase
			json.Unmarshal(condition.Raw, &checkpointBase)
			if checkpointBase.Disabled {
				continue
			}

			s.dealwithCheckoutResult(tenantId, checkpointBase, invokeId)

		} else if condition.Type == consts.ConditionTypeResponseDefine {
			var responseDefineBase domain.ResponseDefineBase
			json.Unmarshal(condition.Raw, &responseDefineBase)
			if responseDefineBase.Disabled {
				continue
			}

			s.dealwithResponseDefineResult(tenantId, responseDefineBase, invokeId)
		}
	}

	return
}

func (s *ExecConditionService) dealwithScriptResult(tenantId consts.TenantId, scriptBase domain.ScriptBase, invokeId,
	debugInterfaceId, caseInterfaceId, endpointInterfaceId, serveId, processorId, scenarioId uint, usedBy consts.UsedBy) (err error) {
	scriptBase.InvokeId = invokeId

	s.ScriptRepo.UpdateResult(tenantId, scriptBase)
	s.ScriptRepo.CreateLog(tenantId, scriptBase)

	for _, settings := range scriptBase.VariableSettings {
		value := valueUtils.InterfaceToStr(settings.Value)

		s.ShareVarService.Save(tenantId, settings.Name, value, settings.ValueType,
			invokeId, debugInterfaceId, caseInterfaceId, endpointInterfaceId, serveId, processorId, scenarioId,
			consts.Public, usedBy)
	}

	return
}

func (s *ExecConditionService) dealwithDbOptResult(tenantId consts.TenantId, databaseOptBase domain.DatabaseOptBase, invokeId,
	debugInterfaceId, caseInterfaceId, endpointInterfaceId, serveId, processorId, scenarioId uint, usedBy consts.UsedBy) (err error) {

	databaseOptBase.InvokeId = invokeId

	s.DatabaseOptRepo.UpdateResult(tenantId, databaseOptBase)
	s.DatabaseOptRepo.CreateLog(tenantId, databaseOptBase)

	if databaseOptBase.ResultStatus == consts.Pass {
		s.ShareVarService.Save(tenantId, databaseOptBase.Variable, databaseOptBase.Result, databaseOptBase.ResultType,
			invokeId, debugInterfaceId, caseInterfaceId, endpointInterfaceId, serveId, processorId, scenarioId,
			databaseOptBase.Scope, usedBy)
	}

	return
}

func (s *ExecConditionService) dealwithExtractorResult(tenantId consts.TenantId, extractorBase domain.ExtractorBase, invokeId,
	debugInterfaceId, caseInterfaceId, endpointInterfaceId, serveId, processorId, scenarioId uint, usedBy consts.UsedBy) (err error) {

	extractorBase.InvokeId = invokeId

	s.ExtractorRepo.UpdateResult(tenantId, extractorBase)
	s.ExtractorRepo.CreateLog(tenantId, extractorBase)

	if extractorBase.ResultStatus == consts.Pass {
		s.ShareVarService.Save(tenantId, extractorBase.Variable, extractorBase.Result, extractorBase.ResultType,
			invokeId, debugInterfaceId, caseInterfaceId, endpointInterfaceId, serveId, processorId, scenarioId,
			extractorBase.Scope, usedBy)
	}

	return
}

func (s *ExecConditionService) dealwithCheckoutResult(tenantId consts.TenantId, checkpointBase domain.CheckpointBase, invokeId uint) (err error) {
	checkpointBase.InvokeId = invokeId

	s.CheckpointRepo.UpdateResult(tenantId, checkpointBase)
	s.CheckpointRepo.CreateLog(tenantId, checkpointBase)

	return
}

func (s *ExecConditionService) dealwithResponseDefineResult(tenantId consts.TenantId, responseDefineBase domain.ResponseDefineBase, invokeId uint) (err error) {

	responseDefineBase.InvokeId = invokeId

	s.ResponseDefineRepo.UpdateResult(tenantId, responseDefineBase)
	s.ResponseDefineRepo.CreateLog(tenantId, responseDefineBase)

	return
}

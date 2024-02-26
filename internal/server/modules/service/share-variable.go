package service

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	repo "github.com/aaronchen2k/deeptest/internal/server/modules/repo"
)

type ShareVarService struct {
	ShareVariableRepo *repo.ShareVariableRepo `inject:""`

	DiagnoseInterfaceRepo *repo.DiagnoseInterfaceRepo `inject:""`
	EndpointCaseRepo      *repo.EndpointCaseRepo      `inject:""`
	DebugInterfaceRepo    *repo.DebugInterfaceRepo    `inject:""`
	EndpointInterfaceRepo *repo.EndpointInterfaceRepo `inject:""`
	EndpointRepo          *repo.EndpointRepo          `inject:""`
	ServeServerRepo       *repo.ServeServerRepo       `inject:""`
	ScenarioProcessorRepo *repo.ScenarioProcessorRepo `inject:""`
}

func (s *ShareVarService) Save(tenantId consts.TenantId, name, value string, resultType consts.ExtractorResultType,
	invokeId, debugInterfaceId, caseInterfaceId, endpointInterfaceId, serveId, processorId, scenarioId uint,
	scope consts.ExtractorScope, usedBy consts.UsedBy) (err error) {

	po := model.ShareVariable{
		Name:                name,
		Value:               value,
		ValueType:           resultType,
		InvokeId:            invokeId,
		DebugInterfaceId:    debugInterfaceId,
		CaseInterfaceId:     caseInterfaceId,
		EndpointInterfaceId: endpointInterfaceId,
		ServeId:             serveId,
		ScenarioProcessorId: processorId,
		ScenarioId:          scenarioId,
		Scope:               scope,
		UsedBy:              usedBy,
	}

	if usedBy == consts.ScenarioDebug {
		po.ID, err = s.ShareVariableRepo.GetExistByScenarioDebug(tenantId, name, scenarioId)
	} else {
		po.ID, err = s.ShareVariableRepo.GetExistByInterfaceDebug(tenantId, name, serveId, usedBy)
	}

	err = s.ShareVariableRepo.Save(tenantId, &po)

	return
}

func (s *ShareVarService) List(tenantId consts.TenantId, debugInterfaceId, endpointInterfaceId, diagnoseInterfaceId, caseInterfaceId, scenarioProcessorId uint,
	usedBy consts.UsedBy) (
	shareVariables []domain.GlobalVar) {

	var serveId uint

	if diagnoseInterfaceId > 0 {
		diagnoseInterface, _ := s.DiagnoseInterfaceRepo.Get(tenantId, diagnoseInterfaceId)
		serveId = diagnoseInterface.ServeId

	} else if caseInterfaceId > 0 {
		caseInterface, _ := s.EndpointCaseRepo.Get(tenantId, caseInterfaceId)
		serveId = caseInterface.ServeId

	} else if endpointInterfaceId > 0 {
		endpointInterface, _ := s.EndpointInterfaceRepo.Get(tenantId, endpointInterfaceId)
		endpoint, _ := s.EndpointRepo.Get(tenantId, endpointInterface.EndpointId)
		serveId = endpoint.ServeId

	} else if debugInterfaceId > 0 {
		debugInterface, _ := s.DebugInterfaceRepo.Get(tenantId, debugInterfaceId)
		endpointInterface, _ := s.EndpointInterfaceRepo.Get(tenantId, debugInterface.EndpointInterfaceId)
		endpoint, _ := s.EndpointRepo.Get(tenantId, endpointInterface.EndpointId)
		serveId = endpoint.ServeId
	}

	shareVariables, _ = s.ListForDebug(tenantId, serveId, scenarioProcessorId, usedBy)

	return
}

func (s *ShareVarService) Delete(tenantId consts.TenantId, id int) (err error) {
	err = s.ShareVariableRepo.Delete(tenantId, id)

	return
}

func (s *ShareVarService) Clear(tenantId consts.TenantId, debugReq domain.DebugInfo) (err error) {
	if debugReq.ScenarioProcessorId > 0 {
		processor, _ := s.ScenarioProcessorRepo.Get(tenantId, debugReq.ScenarioProcessorId)
		err = s.ShareVariableRepo.DeleteAllByScenarioId(tenantId, processor.ScenarioId)

	} else if debugReq.DebugInterfaceId > 0 {
		debugData, _ := s.DebugInterfaceRepo.Get(tenantId, debugReq.DebugInterfaceId)
		err = s.ShareVariableRepo.DeleteAllByServeId(tenantId, debugData.ServeId)

	} else if debugReq.EndpointInterfaceId > 0 {
		endpointInterface, _ := s.EndpointInterfaceRepo.Get(tenantId, debugReq.EndpointInterfaceId)
		endpoint, _ := s.EndpointRepo.Get(tenantId, endpointInterface.EndpointId)
		err = s.ShareVariableRepo.DeleteAllByServeId(tenantId, endpoint.ServeId)

	}

	return
}

func (s *ShareVarService) ListForDebug(tenantId consts.TenantId, serveId, scenarioProcessorId uint, usedBy consts.UsedBy) (ret []domain.GlobalVar, err error) {
	var pos []model.ShareVariable

	if scenarioProcessorId > 0 {
		pos, err = s.ShareVariableRepo.ListForScenarioDebug(tenantId, scenarioProcessorId)
	} else {
		pos, err = s.ShareVariableRepo.ListForInterfaceDebug(tenantId, serveId, usedBy)
	}

	for _, po := range pos {
		ret = append(ret, domain.GlobalVar{
			VarId:      po.ID,
			Name:       po.Name,
			LocalValue: po.Value,
			ValueType:  po.ValueType,
		})
	}

	return
}

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
	DebugInterfaceRepo    *repo.DebugInterfaceRepo    `inject:""`
	EndpointInterfaceRepo *repo.EndpointInterfaceRepo `inject:""`
	EndpointRepo          *repo.EndpointRepo          `inject:""`
	ServeServerRepo       *repo.ServeServerRepo       `inject:""`
	ScenarioProcessorRepo *repo.ScenarioProcessorRepo `inject:""`
}

func (s *ShareVarService) Save(name, value string, debugInterfaceId, endpointInterfaceId, serveId, processorId, scenarioId uint,
	scope consts.ExtractorScope, usedBy consts.UsedBy) (err error) {

	po := model.ShareVariable{
		Name:                name,
		Value:               value,
		DebugInterfaceId:    debugInterfaceId,
		EndpointInterfaceId: endpointInterfaceId,
		ServeId:             serveId,
		ScenarioProcessorId: processorId,
		ScenarioId:          scenarioId,
		Scope:               scope,
		UsedBy:              usedBy,
	}

	if usedBy == consts.InterfaceDebug || usedBy == consts.TestDebug {
		po.ID, err = s.ShareVariableRepo.GetExistByInterfaceDebug(name, serveId, usedBy)
	} else if usedBy == consts.ScenarioDebug {
		po.ID, err = s.ShareVariableRepo.GetExistByScenarioDebug(name, scenarioId)
	}

	err = s.ShareVariableRepo.Save(&po)

	return
}

func (s *ShareVarService) List(debugInterfaceId, endpointInterfaceId, diagnoseInterfaceId, scenarioProcessorId uint,
	usedBy consts.UsedBy) (
	shareVariables []domain.GlobalVar) {

	var serveId uint

	if diagnoseInterfaceId > 0 {
		diagnoseInterface, _ := s.DiagnoseInterfaceRepo.Get(diagnoseInterfaceId)
		serveId = diagnoseInterface.ServeId

	} else if endpointInterfaceId > 0 {
		endpointInterface, _ := s.EndpointInterfaceRepo.Get(endpointInterfaceId)
		endpoint, _ := s.EndpointRepo.Get(endpointInterface.EndpointId)
		serveId = endpoint.ServeId

	} else if debugInterfaceId > 0 {
		debugInterface, _ := s.DebugInterfaceRepo.Get(debugInterfaceId)
		endpointInterface, _ := s.EndpointInterfaceRepo.Get(debugInterface.EndpointInterfaceId)
		endpoint, _ := s.EndpointRepo.Get(endpointInterface.EndpointId)
		serveId = endpoint.ServeId
	}

	shareVariables, _ = s.ListForDebug(serveId, scenarioProcessorId, usedBy)

	return
}

func (s *ShareVarService) Delete(id int) (err error) {
	err = s.ShareVariableRepo.Delete(id)

	return
}

func (s *ShareVarService) Clear(endpointOrProcessorId int, usedBy consts.UsedBy) (err error) {
	if usedBy == consts.InterfaceDebug {
		endpoint, _ := s.EndpointRepo.Get(uint(endpointOrProcessorId))
		err = s.ShareVariableRepo.DeleteAllByServeId(endpoint.ServeId)

	} else if usedBy == consts.ScenarioDebug {
		processor, _ := s.ScenarioProcessorRepo.Get(uint(endpointOrProcessorId))
		err = s.ShareVariableRepo.DeleteAllByScenarioId(processor.ScenarioId)

	}

	return
}

func (s *ShareVarService) ListForDebug(serveId, scenarioProcessorId uint, usedBy consts.UsedBy) (ret []domain.GlobalVar, err error) {
	var pos []model.ShareVariable

	if scenarioProcessorId > 0 {
		pos, err = s.ShareVariableRepo.ListByScenarioDebug(scenarioProcessorId)
	} else {
		pos, err = s.ShareVariableRepo.ListByInterfaceDebug(serveId, usedBy)
	}

	for _, po := range pos {
		ret = append(ret, domain.GlobalVar{
			VarId:      po.ID,
			Name:       po.Name,
			LocalValue: po.Value,
		})
	}

	return
}

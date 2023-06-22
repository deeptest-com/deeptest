package service

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	_httpUtils "github.com/aaronchen2k/deeptest/pkg/lib/http"
)

type DebugSceneService struct {
	EndpointInterfaceRepo *repo.EndpointInterfaceRepo `inject:""`
	EndpointRepo          *repo.EndpointRepo          `inject:""`
	ServeServerRepo       *repo.ServeServerRepo       `inject:""`
	ScenarioProcessorRepo *repo.ScenarioProcessorRepo `inject:""`
	EnvironmentRepo       *repo.EnvironmentRepo       `inject:""`

	ShareVarService *ShareVarService `inject:""`

	EnvironmentService *EnvironmentService `inject:""`
}

func (s *DebugSceneService) LoadScene(endpointInterfaceId, debugServerId, scenarioProcessorId uint, usedBy consts.UsedBy) (
	baseUrl string, shareVars []domain.GlobalVar, envVars []domain.GlobalVar,
	globalVars []domain.GlobalVar, globalParams []domain.GlobalParam) {

	if endpointInterfaceId > 0 {
		interf, _ := s.EndpointInterfaceRepo.Get(endpointInterfaceId)
		endpoint, _ := s.EndpointRepo.Get(interf.EndpointId)

		if debugServerId == 0 {
			debugServerId = endpoint.ServerId
		}
	}

	serveServer, _ := s.ServeServerRepo.Get(debugServerId)

	baseUrl = _httpUtils.AddSepIfNeeded(serveServer.Url)
	envId := serveServer.EnvironmentId
	environment, _ := s.EnvironmentRepo.Get(envId)

	shareVars, _ = s.ShareVarService.ListForDebug(debugServerId, scenarioProcessorId, usedBy)
	envVars, _ = s.EnvironmentService.GetVarsByEnv(envId)
	globalVars, _ = s.EnvironmentService.GetGlobalVars(environment.ProjectId)
	globalParams, _ = s.EnvironmentService.GetGlobalParams(environment.ProjectId)

	return
}

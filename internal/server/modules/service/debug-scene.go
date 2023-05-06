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

func (s *DebugSceneService) LoadScene(endpointInterfaceId, scenarioProcessorId uint, usedBy consts.UsedBy) (
	baseUrl string, shareVariables []domain.GlobalVar) {

	var serveId, serverId uint

	interf, _ := s.EndpointInterfaceRepo.Get(endpointInterfaceId)
	endpoint, _ := s.EndpointRepo.Get(interf.EndpointId)
	serveId = endpoint.ServeId
	serverId = endpoint.ServerId

	serveServer, _ := s.ServeServerRepo.Get(serverId)
	baseUrl = _httpUtils.AddSepIfNeeded(serveServer.Url)

	shareVariables, _ = s.ShareVarService.listForDebug(serveId, scenarioProcessorId, usedBy)

	return
}

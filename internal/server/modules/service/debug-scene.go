package service

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	_httpUtils "github.com/aaronchen2k/deeptest/pkg/lib/http"
)

type DebugSceneService struct {
	EndpointRepo    *repo.EndpointRepo    `inject:""`
	ServeServerRepo *repo.ServeServerRepo `inject:""`
	EnvironmentRepo *repo.EnvironmentRepo `inject:""`

	EnvironmentService *EnvironmentService `inject:""`
	VariableService    *VariableService    `inject:""`
}

func (s *DebugSceneService) LoadScene(endpointId, InterfaceId uint, usedBy consts.UsedBy) (
	baseUrl string, shareVariables []domain.ShareVars, envVars []domain.EnvVars,
	globalEnvVars []domain.GlobalEnvVars, globalParamVars []domain.GlobalParamVars) {

	endpoint, _ := s.EndpointRepo.Get(endpointId)
	projectId := endpoint.ProjectId
	serverId := endpoint.ServerId

	serveServer, _ := s.ServeServerRepo.Get(serverId)
	baseUrl = _httpUtils.AddSepIfNeeded(serveServer.Url)
	envId := serveServer.EnvironmentId

	shareVariables, _ = s.VariableService.GetVariablesByInterface(InterfaceId, usedBy)
	envVars, _ = s.EnvironmentService.GetVarsByEnv(envId)

	globalEnvVars, _ = s.EnvironmentService.GetGlobalVars(projectId)
	globalParamVars, _ = s.EnvironmentService.GetGlobalParams(projectId)

	// interf, _ := s.ProcessorInterfaceRepo.Get(req.InterfaceId)
	//req.Datapools, _ = s.DatapoolService.ListForExec(interf.ProjectId)

	return
}

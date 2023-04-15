package service

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	"path"
)

type DebugSceneService struct {
	EndpointRepo    *repo.EndpointRepo    `inject:""`
	ServeServerRepo *repo.ServeServerRepo `inject:""`
	EnvironmentRepo *repo.EnvironmentRepo `inject:""`

	EnvironmentService *EnvironmentService `inject:""`
	VariableService    *VariableService    `inject:""`
}

func (s *DebugSceneService) LoadScene(req *v1.DebugRequest) (err error) {

	endpointId := req.EndpointId
	InterfaceId := req.InterfaceId

	endpoint, _ := s.EndpointRepo.Get(endpointId)
	projectId := endpoint.ProjectId
	serverId := endpoint.ServerId

	server, _ := s.ServeServerRepo.Get(serverId)
	req.Url = path.Join(server.Url, req.Url)
	envId := server.EnvironmentId

	req.EnvVars, _ = s.EnvironmentService.GetVarsByEnv(envId)
	req.ShareVariables, _ = s.VariableService.GetVariablesByInterface(InterfaceId, req.UsedBy)

	req.GlobalEnvVars, _ = s.EnvironmentService.GetGlobalVars(projectId)
	req.GlobalParamVars, _ = s.EnvironmentService.GetGlobalParams(projectId)

	// interf, _ := s.ProcessorInterfaceRepo.Get(req.InterfaceId)
	//req.Datapools, _ = s.DatapoolService.ListForExec(interf.ProjectId)

	return
}

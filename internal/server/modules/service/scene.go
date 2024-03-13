package service

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
)

type SceneService struct {
	ScenarioNodeRepo   *repo.ScenarioNodeRepo   `inject:""`
	EnvironmentRepo    *repo.EnvironmentRepo    `inject:""`
	DebugInterfaceRepo *repo.DebugInterfaceRepo `inject:""`

	EndpointInterfaceRepo *repo.EndpointInterfaceRepo `inject:""`
	EndpointRepo          *repo.EndpointRepo          `inject:""`
	ServeServerRepo       *repo.ServeServerRepo       `inject:""`

	ShareVarService     *ShareVarService     `inject:""`
	EnvironmentService  *EnvironmentService  `inject:""`
	DatapoolService     *DatapoolService     `inject:""`
	ScenarioNodeService *ScenarioNodeService `inject:""`
}

func (s *SceneService) LoadEnvVarMapByScenario(tenantId consts.TenantId, scene *domain.ExecScene, scenarioId, environmentId uint) {
	scene.EnvToVariables = domain.EnvToVariables{}
	scene.DebugInterfaceToEnvMap = domain.InterfaceToEnvMap{}

	// load baseUrl for interface processors
	processors, _ := s.ScenarioNodeRepo.ListByScenario(tenantId, scenarioId)

	for _, processor := range processors {
		if processor.EntityType != consts.ProcessorInterfaceDefault {
			continue
		}

		var server = s.GetExecServer(tenantId, processor.EntityId, processor.EndpointInterfaceId, environmentId)

		scene.DebugInterfaceToEnvMap[processor.EntityId] = server.EnvironmentId

		scene.EnvToVariables[server.EnvironmentId] = append(scene.EnvToVariables[server.EnvironmentId], domain.GlobalVar{
			Name:        consts.KEY_BASE_URL,
			LocalValue:  server.Url,
			RemoteValue: server.Url,
		})
	}

	// load vars by env
	vars, _ := s.EnvironmentRepo.GetVars(tenantId, environmentId)
	for _, v := range vars {
		scene.EnvToVariables[environmentId] = append(scene.EnvToVariables[environmentId], domain.GlobalVar{
			Name:        v.Name,
			LocalValue:  v.LocalValue,
			RemoteValue: v.RemoteValue,
		})

	}

	return
}

func (s *SceneService) GetExecServer(tenantId consts.TenantId, debugInterfaceId, endpointInterfaceId, environmentId uint) (server model.ServeServer) {
	interf, _ := s.EndpointInterfaceRepo.Get(tenantId, endpointInterfaceId)

	if environmentId > 0 { // select a env to exec

		var serveId uint

		if debugInterfaceId > 0 {
			debugInterface, _ := s.DebugInterfaceRepo.Get(tenantId, debugInterfaceId)
			serveId = debugInterface.ServeId

		} else {
			endpoint, _ := s.EndpointRepo.Get(tenantId, interf.EndpointId)
			serveId = endpoint.ServeId

		}

		server, _ = s.ServeServerRepo.FindByServeAndExecEnv(tenantId, serveId, environmentId)

	} else { // TODO: may not used, since now environmentId always not empty
		var serverId uint
		if debugInterfaceId > 0 { // from debug interface
			debugInterface, _ := s.DebugInterfaceRepo.Get(tenantId, debugInterfaceId)
			serverId = debugInterface.ServerId

		} else { // from endpoint interface
			endpoint, _ := s.EndpointRepo.Get(tenantId, interf.EndpointId)
			serverId = endpoint.ServerId

		}

		server, _ = s.ServeServerRepo.Get(tenantId, serverId)
	}
	return
}

func (s *SceneService) LoadEnvVars(tenantId consts.TenantId, scene *domain.ExecScene, serverId, debugInterfaceId uint) (projectId uint, err error) {

	scene.EnvToVariables = domain.EnvToVariables{}
	scene.DebugInterfaceToEnvMap = domain.InterfaceToEnvMap{}

	serveServer, _ := s.ServeServerRepo.Get(tenantId, serverId)
	envId := serveServer.EnvironmentId

	scene.DebugInterfaceToEnvMap[debugInterfaceId] = envId

	vars, _ := s.EnvironmentRepo.GetVars(tenantId, envId)
	for _, v := range vars {
		scene.EnvToVariables[envId] = append(scene.EnvToVariables[envId], domain.GlobalVar{
			Name:        v.Name,
			LocalValue:  v.LocalValue,
			RemoteValue: v.RemoteValue,
		})
	}

	return
}

func (s *SceneService) LoadProjectSettings(tenantId consts.TenantId, scene *domain.ExecScene, projectId uint) {
	scene.GlobalParams, _ = s.EnvironmentService.GetGlobalParams(tenantId, projectId)
	scene.GlobalVars, _ = s.EnvironmentService.GetGlobalVars(tenantId, projectId)

	scene.Datapools, _ = s.DatapoolService.ListForExec(tenantId, projectId)
}

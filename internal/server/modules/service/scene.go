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

func (s *SceneService) LoadEnvVarMapByScenario(scene *domain.ExecScene, scenarioId, environmentId uint) {
	scene.EnvToVariables = domain.EnvToVariables{}
	scene.DebugInterfaceToEnvMap = domain.InterfaceToEnvMap{}

	processors, _ := s.ScenarioNodeRepo.ListByScenario(scenarioId)

	for _, processor := range processors {
		if processor.EntityType != consts.ProcessorInterfaceDefault {
			continue
		}

		var server = s.GetExecServer(processor.EntityId, processor.EndpointInterfaceId, environmentId)
		envId := server.EnvironmentId

		scene.DebugInterfaceToEnvMap[processor.EntityId] = envId

		scene.EnvToVariables[envId] = append(scene.EnvToVariables[envId], domain.GlobalVar{
			Name:        consts.KEY_BASE_URL,
			LocalValue:  server.Url,
			RemoteValue: server.Url,
		})

		vars, _ := s.EnvironmentRepo.GetVars(envId)
		for _, v := range vars {
			scene.EnvToVariables[envId] = append(scene.EnvToVariables[envId], domain.GlobalVar{
				Name:        v.Name,
				LocalValue:  v.LocalValue,
				RemoteValue: v.RemoteValue,
			})
		}
	}

	return
}

func (s *SceneService) GetExecServer(debugInterfaceId, endpointInterfaceId, environmentId uint) (server model.ServeServer) {
	interf, _ := s.EndpointInterfaceRepo.Get(endpointInterfaceId)

	if environmentId > 0 { // select a env to exec
		var serveId uint

		if debugInterfaceId > 0 {
			debugInterface, _ := s.DebugInterfaceRepo.Get(debugInterfaceId)
			serveId = debugInterface.ServeId

		} else {
			endpoint, _ := s.EndpointRepo.Get(interf.EndpointId)
			serveId = endpoint.ServeId

		}

		server, _ = s.ServeServerRepo.FindByServeAndExecEnv(serveId, environmentId)

	} else {
		var serverId uint
		if debugInterfaceId > 0 { // from debug interface
			debugInterface, _ := s.DebugInterfaceRepo.Get(debugInterfaceId)
			serverId = debugInterface.ServerId

		} else { // from endpoint interface
			endpoint, _ := s.EndpointRepo.Get(interf.EndpointId)
			serverId = endpoint.ServerId

		}

		server, _ = s.ServeServerRepo.Get(serverId)
	}
	return
}

func (s *SceneService) LoadEnvVars(scene *domain.ExecScene, serverId, debugInterfaceId uint) (projectId uint, err error) {

	scene.EnvToVariables = domain.EnvToVariables{}
	scene.DebugInterfaceToEnvMap = domain.InterfaceToEnvMap{}

	serveServer, _ := s.ServeServerRepo.Get(serverId)
	envId := serveServer.EnvironmentId

	scene.DebugInterfaceToEnvMap[debugInterfaceId] = envId

	vars, _ := s.EnvironmentRepo.GetVars(envId)
	for _, v := range vars {
		scene.EnvToVariables[envId] = append(scene.EnvToVariables[envId], domain.GlobalVar{
			Name:        v.Name,
			LocalValue:  v.LocalValue,
			RemoteValue: v.RemoteValue,
		})
	}

	return
}

func (s *SceneService) LoadProjectSettings(scene *domain.ExecScene, projectId uint) {
	scene.GlobalParams, _ = s.EnvironmentService.GetGlobalParams(projectId)
	scene.GlobalVars, _ = s.EnvironmentService.GetGlobalVars(projectId)

	scene.Datapools, _ = s.DatapoolService.ListForExec(projectId)
}

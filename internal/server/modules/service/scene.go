package service

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
)

type SceneService struct {
	ScenarioNodeRepo *repo.ScenarioNodeRepo `inject:""`
	EnvironmentRepo  *repo.EnvironmentRepo  `inject:""`

	EndpointInterfaceRepo *repo.EndpointInterfaceRepo `inject:""`
	EndpointRepo          *repo.EndpointRepo          `inject:""`
	ServeServerRepo       *repo.ServeServerRepo       `inject:""`

	ShareVarService     *ShareVarService     `inject:""`
	EnvironmentService  *EnvironmentService  `inject:""`
	DatapoolService     *DatapoolService     `inject:""`
	ScenarioNodeService *ScenarioNodeService `inject:""`
}

func (s *SceneService) LoadEnvVarMapByScenario(scene *domain.ExecScene, scenarioId uint) {
	scene.EnvToVariables = domain.EnvToVariables{}
	scene.InterfaceToEnvMap = domain.InterfaceToEnvMap{}

	processors, _ := s.ScenarioNodeRepo.ListByScenario(scenarioId)

	for _, processor := range processors {
		if processor.EntityType != consts.ProcessorInterfaceDefault {
			continue
		}

		interf, _ := s.EndpointInterfaceRepo.Get(processor.EndpointInterfaceId)
		endpoint, _ := s.EndpointRepo.Get(interf.EndpointId)
		serveServer, _ := s.ServeServerRepo.Get(endpoint.ServerId)
		envId := serveServer.EnvironmentId

		scene.InterfaceToEnvMap[processor.EndpointInterfaceId] = envId

		scene.EnvToVariables[envId] = append(scene.EnvToVariables[envId], domain.GlobalVar{
			Name:        consts.KEY_BASE_URL,
			LocalValue:  serveServer.Url,
			RemoteValue: serveServer.Url,
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

func (s *SceneService) LoadEnvVarMapByEndpointInterface(scene *domain.ExecScene, endpointInterfaceId, debugServerId uint) (projectId uint, err error) {
	scene.EnvToVariables = domain.EnvToVariables{}
	scene.InterfaceToEnvMap = domain.InterfaceToEnvMap{}

	interf, _ := s.EndpointInterfaceRepo.Get(endpointInterfaceId)
	endpoint, _ := s.EndpointRepo.Get(interf.EndpointId)

	if debugServerId == 0 {
		debugServerId = endpoint.ServerId
	}
	serveServer, _ := s.ServeServerRepo.Get(debugServerId)

	envId := serveServer.EnvironmentId
	projectId = endpoint.ProjectId

	scene.InterfaceToEnvMap[endpointInterfaceId] = envId

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

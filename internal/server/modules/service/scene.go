package service

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
)

type SceneService struct {
	ScenarioRepo       *repo.ScenarioRepo       `inject:""`
	ScenarioNodeRepo   *repo.ScenarioNodeRepo   `inject:""`
	ScenarioReportRepo *repo.ScenarioReportRepo `inject:""`
	TestLogRepo        *repo.LogRepo            `inject:""`
	EnvironmentRepo    *repo.EnvironmentRepo    `inject:""`

	EndpointInterfaceRepo *repo.EndpointInterfaceRepo `inject:""`
	EndpointRepo          *repo.EndpointRepo          `inject:""`
	ServeServerRepo       *repo.ServeServerRepo       `inject:""`

	ShareVarService     *ShareVarService     `inject:""`
	EnvironmentService  *EnvironmentService  `inject:""`
	DatapoolService     *DatapoolService     `inject:""`
	ScenarioNodeService *ScenarioNodeService `inject:""`
}

func (s *SceneService) LoadEnvVarMapByScenario(scenarioId uint) (
	envToVariablesMap domain.EnvToVariables, interfaceToEnvMap domain.InterfaceToEnvMap, err error) {

	envToVariablesMap = domain.EnvToVariables{}
	interfaceToEnvMap = domain.InterfaceToEnvMap{}

	processors, err := s.ScenarioNodeRepo.ListByScenario(scenarioId)

	for _, processor := range processors {
		if processor.EntityType != consts.ProcessorInterfaceDefault {
			continue
		}

		interf, _ := s.EndpointInterfaceRepo.Get(processor.EndpointInterfaceId)
		endpoint, _ := s.EndpointRepo.Get(interf.EndpointId)
		serveServer, _ := s.ServeServerRepo.Get(endpoint.ServerId)
		envId := serveServer.EnvironmentId

		interfaceToEnvMap[processor.EndpointInterfaceId] = envId

		envToVariablesMap[envId] = append(envToVariablesMap[envId], domain.GlobalVar{
			Name:        consts.KEY_BASE_URL,
			LocalValue:  serveServer.Url,
			RemoteValue: serveServer.Url,
		})

		vars, _ := s.EnvironmentRepo.GetVars(envId)
		for _, v := range vars {
			envToVariablesMap[envId] = append(envToVariablesMap[envId], domain.GlobalVar{
				Name:        v.Name,
				LocalValue:  v.LocalValue,
				RemoteValue: v.RemoteValue,
			})
		}
	}

	return
}

func (s *SceneService) LoadEnvVarMapByEndpointInterface(endpointInterfaceId uint) (
	envToVariablesMap domain.EnvToVariables, interfaceToEnvMap domain.InterfaceToEnvMap, projectId uint, err error) {

	envToVariablesMap = domain.EnvToVariables{}
	interfaceToEnvMap = domain.InterfaceToEnvMap{}

	interf, _ := s.EndpointInterfaceRepo.Get(endpointInterfaceId)
	endpoint, _ := s.EndpointRepo.Get(interf.EndpointId)
	serveServer, _ := s.ServeServerRepo.Get(endpoint.ServerId)
	envId := serveServer.EnvironmentId
	projectId = endpoint.ProjectId

	interfaceToEnvMap[endpointInterfaceId] = envId

	vars, _ := s.EnvironmentRepo.GetVars(envId)
	for _, v := range vars {
		envToVariablesMap[envId] = append(envToVariablesMap[envId], domain.GlobalVar{
			Name:        v.Name,
			LocalValue:  v.LocalValue,
			RemoteValue: v.RemoteValue,
		})
	}

	return
}

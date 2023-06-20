package service

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	repo "github.com/aaronchen2k/deeptest/internal/server/modules/repo"
)

type VariableService struct {
	DebugInterfaceRepo *repo.DebugInterfaceRepo `inject:""`

	EndpointInterfaceRepo *repo.EndpointInterfaceRepo `inject:""`
	EndpointRepo          *repo.EndpointRepo          `inject:""`

	ExtractorRepo   *repo.ExtractorRepo   `inject:""`
	EnvironmentRepo *repo.EnvironmentRepo `inject:""`

	EnvironmentService *EnvironmentService `inject:""`
	ShareVarService    *ShareVarService    `inject:""`
	DatapoolService    *DatapoolService    `inject:""`
}

func (s *VariableService) GetCombinedVarsForCheckpoint(debugInterfaceId, endpointInterfaceId, scenarioProcessorId uint) (
	ret map[string]interface{}, datapools domain.Datapools, err error) {

	debugEnv, _ := s.EnvironmentService.GetDebugEnvByEndpointInterface(debugInterfaceId, endpointInterfaceId)

	interf, _ := s.EndpointInterfaceRepo.Get(endpointInterfaceId)
	endpoint, _ := s.EndpointRepo.Get(interf.EndpointId)

	shareVariables, _ := s.ShareVarService.ListForDebug(endpoint.ServeId, scenarioProcessorId)
	envVars, _ := s.EnvironmentService.GetVarsByEnv(debugEnv.ID)
	globalVars, _ := s.EnvironmentService.GetGlobalVars(debugEnv.ProjectId)
	datapools, _ = s.DatapoolService.ListForExec(debugEnv.ProjectId)

	ret = CombineVariables(shareVariables, envVars, globalVars)

	return
}

func CombineVariables(shareVariables, envVars, globalVars []domain.GlobalVar) (
	ret map[string]interface{}) {
	ret = map[string]interface{}{}

	variableMap := map[string]interface{}{}

	for _, item := range globalVars {
		variableMap[item.Name] = item.LocalValue
	}
	for _, item := range envVars { // overwrite previous ones
		variableMap[item.Name] = item.LocalValue
	}
	for _, item := range shareVariables { // overwrite previous ones
		variableMap[item.Name] = item.LocalValue
	}

	// value is a  object
	for key, val := range variableMap {
		valMp, isMap := val.(map[string]interface{})

		if isMap {
			for propKey, v := range valMp {
				ret[fmt.Sprintf("%s.%s", key, propKey)] = v
			}

		} else {
			ret[fmt.Sprintf("%s", key)] = val

		}
	}

	return
}

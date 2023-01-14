package service

import (
	"fmt"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	agentExecDomain "github.com/aaronchen2k/deeptest/internal/agent/exec/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	repo "github.com/aaronchen2k/deeptest/internal/server/modules/repo"
)

type VariableService struct {
	InterfaceRepo   *repo.InterfaceRepo   `inject:""`
	ExtractorRepo   *repo.ExtractorRepo   `inject:""`
	EnvironmentRepo *repo.EnvironmentRepo `inject:""`
}

func (s *VariableService) GetVariablesByInterface(interfaceId uint, usedBy consts.UsedBy) (ret map[string]interface{}, err error) {
	interf, err := s.InterfaceRepo.Get(interfaceId)

	environmentVariables, _ := s.EnvironmentRepo.ListVariableByProject(interf.ProjectId)

	interfaceExtractorVariables, _ :=
		s.ExtractorRepo.ListValidExtractorVariableForInterface(interfaceId, interf.ProjectId, usedBy)

	ret = MergeVariables(environmentVariables, interfaceExtractorVariables, nil)

	return
}

func MergeVariables(environmentVariables []model.EnvironmentVar, interfaceExtractorVariables []v1.Variable,
	processorExecVariables []agentExecDomain.ExecVariable) (
	ret map[string]interface{}) {

	ret = map[string]interface{}{}

	variableMap := map[string]interface{}{}
	for _, item := range environmentVariables {
		variableMap[item.Name] = item.RightValue
	}
	for _, item := range interfaceExtractorVariables { // overwrite previous ones
		variableMap[item.Name] = item.Value
	}
	for _, item := range processorExecVariables { // overwrite previous ones
		variableMap[item.Name] = item.Value
	}

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

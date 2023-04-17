package service

import (
	"fmt"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	repo "github.com/aaronchen2k/deeptest/internal/server/modules/repo"
)

type VariableService struct {
	InterfaceRepo          *repo.InterfaceRepo          `inject:""`
	ProcessorInterfaceRepo *repo.ProcessorInterfaceRepo `inject:""`

	ExtractorRepo   *repo.ExtractorRepo   `inject:""`
	EnvironmentRepo *repo.EnvironmentRepo `inject:""`
}

func (s *VariableService) GetEnvVarsByInterface(interfaceId uint, usedBy consts.UsedBy) (ret map[string]interface{}, err error) {
	var projectId uint

	if usedBy == consts.InterfaceDebug {
		interf, _ := s.InterfaceRepo.Get(interfaceId)
		projectId = interf.ProjectId
	} else {
		interf, _ := s.ProcessorInterfaceRepo.Get(interfaceId)
		projectId = interf.ProjectId
	}

	environmentVariables, _ := s.EnvironmentRepo.ListVariableByProject(projectId)

	ret = CombineVariables(environmentVariables, nil)

	return
}

func (s *VariableService) GetShareVarsByInterface(interfaceId uint, usedBy consts.UsedBy) (ret []domain.ShareVars, err error) {
	//var projectId uint
	//
	//if usedBy == consts.InterfaceDebug {
	//	interf, _ := s.InterfaceRepo.Get(interfaceId)
	//	projectId = interf.ProjectId
	//} else {
	//	interf, _ := s.ProcessorInterfaceRepo.Get(interfaceId)
	//	projectId = interf.ProjectId
	//}
	//
	//interfaceExtractorVariables, _ :=
	//	s.ExtractorRepo.ListValidExtractorVarForInterface(interfaceId, projectId, usedBy)

	//ret = CombineVariables(nil, interfaceExtractorVariables)

	return
}

func CombineVariables(environmentVariables []model.EnvironmentVar, interfaceExtractorVariables []v1.Variable) (
	ret map[string]interface{}) {

	ret = map[string]interface{}{}

	variableMap := map[string]interface{}{}
	for _, item := range environmentVariables {
		variableMap[item.Name] = item.RightValue
	}
	for _, item := range interfaceExtractorVariables { // overwrite previous ones
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

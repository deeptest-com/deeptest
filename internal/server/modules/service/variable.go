package service

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/business"
	"github.com/aaronchen2k/deeptest/internal/server/modules/helper/request"
	repo2 "github.com/aaronchen2k/deeptest/internal/server/modules/repo"
)

type VariableService struct {
	InterfaceRepo   *repo2.InterfaceRepo   `inject:""`
	ExtractorRepo   *repo2.ExtractorRepo   `inject:""`
	EnvironmentRepo *repo2.EnvironmentRepo `inject:""`
	ExecContext     *business.ExecContext  `inject:""`
}

func (s *VariableService) GetVariablesByInterface(interfaceId uint) (ret map[string]interface{}, err error) {
	interf, err := s.InterfaceRepo.Get(interfaceId)

	environmentVariables, _ := s.EnvironmentRepo.ListVariableByProject(interf.ProjectId)
	interfaceExtractorVariables, _ := s.ExtractorRepo.ListValidExtractorVariable(interfaceId, interf.ProjectId)

	ret = requestHelper.MergeVariables(environmentVariables, interfaceExtractorVariables, nil)

	return
}

func (s *VariableService) GetVariablesByInterfaceAndProcessor(interfaceId, processorId, projectId uint) (ret map[string]interface{}, err error) {
	environmentVariables, _ := s.EnvironmentRepo.ListVariableByProject(projectId)
	execVariables := s.ExecContext.ListCachedVariable(processorId) // including interface extracted variables

	ret = requestHelper.MergeVariables(environmentVariables, nil, execVariables)

	return
}

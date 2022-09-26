package service

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/business"
	requestHelper "github.com/aaronchen2k/deeptest/internal/server/modules/v1/helper/request"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/repo"
)

type VariableService struct {
	InterfaceRepo   *repo.InterfaceRepo   `inject:""`
	ExtractorRepo   *repo.ExtractorRepo   `inject:""`
	EnvironmentRepo *repo.EnvironmentRepo `inject:""`
	ExecContext     *business.ExecContext `inject:""`
}

func (s *VariableService) GetVariablesByInterface(interfaceId uint) (ret [][]interface{}, err error) {
	interf, err := s.InterfaceRepo.Get(interfaceId)

	environmentVariables, _ := s.EnvironmentRepo.ListByInterface(interfaceId)
	interfaceExtractorVariables, _ := s.ExtractorRepo.ListExtractorVariableByProject(interf.ProjectId)

	ret = requestHelper.MergeVariables(environmentVariables, interfaceExtractorVariables, nil)

	return
}

func (s *VariableService) GetVariablesByInterfaceAndProcessor(interfaceId, processorId uint) (ret [][]interface{}, err error) {
	environmentVariables, _ := s.EnvironmentRepo.ListByInterface(interfaceId)
	execVariables := s.ExecContext.ListVariable(processorId)
	// interfaceExtractorVariables saved in execVariables

	ret = requestHelper.MergeVariables(environmentVariables, nil, execVariables)

	return
}

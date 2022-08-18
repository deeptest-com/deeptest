package business

import (
	serverDomain "github.com/aaronchen2k/deeptest/internal/server/modules/v1/domain"
	execHelper "github.com/aaronchen2k/deeptest/internal/server/modules/v1/helper/exec"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/repo"
)

type ExecRequestService struct {
	ExecContextService *ExecContextService    `inject:""`
	ScenarioNodeRepo   *repo.ScenarioNodeRepo `inject:""`
}

func NewExecRequestService() *ExecRequestService {
	return &ExecRequestService{}
}

func (s *ExecRequestService) ReplaceProcessorVariables(req *serverDomain.InvocationRequest, interfaceProcessor *model.TestProcessor) (
	err error) {

	variables := s.ExecContextService.ListVariable(interfaceProcessor.ID)

	execHelper.ReplaceVariablesForInvocation(req, variables)

	return
}

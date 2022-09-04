package business

import (
	serverDomain "github.com/aaronchen2k/deeptest/internal/server/modules/v1/domain"
	execHelper "github.com/aaronchen2k/deeptest/internal/server/modules/v1/helper/exec"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/repo"
)

type ExecRequest struct {
	ExecContextService *ExecContext           `inject:""`
	ScenarioNodeRepo   *repo.ScenarioNodeRepo `inject:""`
}

func NewExecRequestService() *ExecRequest {
	return &ExecRequest{}
}

func (s *ExecRequest) ReplaceProcessorVariables(req *serverDomain.InvocationRequest, interfaceProcessor *model.Processor) (
	err error) {

	variables := s.ExecContextService.ListVariable(interfaceProcessor.ID)
	execHelper.ReplaceExecVariablesForInvocation(req, variables)

	return
}

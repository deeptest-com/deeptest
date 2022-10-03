package runDomain

import (
	"github.com/aaronchen2k/deeptest/internal/agent/run"
)

type ProcessorGroup struct {
	BaseProcessorRun
}

func (s *ProcessorGroup) Run(r *run.SessionRunner) (ret *run.StageResult, err error) {
	// run

	for _, child := range s.Children {
		child.Run(r)
	}

	return
}

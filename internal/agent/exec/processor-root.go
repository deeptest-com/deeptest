package agentExec

import (
	"github.com/aaronchen2k/deeptest/internal/agent/exec/domain"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/utils/exec"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
)

type ProcessorRoot struct {
	ID uint `json:"id" yaml:"id"`
	ProcessorEntityBase
}

func (entity ProcessorRoot) Run(processor *Processor, session *Session) (log domain.Result, err error) {
	logUtils.Infof("root entity")

	processor.Result = domain.Result{
		ID:                entity.ProcessorID,
		Name:              entity.Name,
		ProcessorCategory: entity.ProcessorCategory,
		ProcessorType:     entity.ProcessorType,
	}

	exec.SendExecMsg(processor.Result, session.WsMsg)

	for _, child := range processor.Children {
		child.Run(session)
	}

	return
}

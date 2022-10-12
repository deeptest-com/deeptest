package agentExec

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/domain"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/utils/exec"
	"strings"
)

type ProcessorPrint struct {
	ID uint `json:"id" yaml:"id"`
	ProcessorEntityBase

	Expression string `json:"expression" yaml:"expression"`
}

func (entity ProcessorPrint) Run(processor *Processor, session *Session) (log domain.Result, err error) {
	processor.Result = domain.Result{
		ID:                entity.ProcessorID,
		Name:              entity.Name,
		ProcessorCategory: entity.ProcessorCategory,
		ProcessorType:     entity.ProcessorType,
		ParentId:          entity.ParentID,
	}

	value, err := EvaluateGovaluateExpressionByScope(entity.Expression, processor.ID)

	processor.Result.Summary = strings.ReplaceAll(fmt.Sprintf("%s为\"%v\"。",
		entity.Expression, value), "<nil>", "空")

	exec.SendExecMsg(processor.Result, session.WsMsg)

	return
}

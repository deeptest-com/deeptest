package agentExec

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/domain"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/utils/exec"
)

type ProcessorAssertion struct {
	ID uint `json:"id" yaml:"id"`
	ProcessorEntityBase

	Expression string `json:"expression" yaml:"expression"`
}

func (entity ProcessorAssertion) Run(processor *Processor, session *Session) (result domain.Result, err error) {
	processor.Result = domain.Result{
		ID:                entity.ProcessorID,
		Name:              entity.Name,
		ProcessorCategory: entity.ProcessorCategory,
		ProcessorType:     entity.ProcessorType,
		ParentId:          entity.ParentID,
	}

	ret, err := EvaluateGovaluateExpressionByScope(entity.Expression, entity.ID)

	pass, _ := ret.(bool)

	var status string
	processor.Result.ResultStatus, status = getResultStatus(pass)

	processor.Result.Summary = fmt.Sprintf("断言\"%s\"结果为\"%s\"。", entity.Expression, status)

	processor.Parent.Result.Children = append(processor.Parent.Result.Children, &processor.Result)
	exec.SendExecMsg(processor.Result, session.WsMsg)

	return
}

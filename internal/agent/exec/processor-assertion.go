package agentExec

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/domain"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/utils/exec"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"time"
)

type ProcessorAssertion struct {
	ID uint `json:"id" yaml:"id"`
	ProcessorEntityBase

	Expression string `json:"expression" yaml:"expression"`
}

func (entity ProcessorAssertion) Run(processor *Processor, session *Session) (err error) {
	logUtils.Infof("assertion entity")

	startTime := time.Now()
	processor.Result = &domain.Result{
		ID:                int(entity.ProcessorID),
		Name:              entity.Name,
		ProcessorCategory: entity.ProcessorCategory,
		ProcessorType:     entity.ProcessorType,
		StartTime:         &startTime,
		ParentId:          int(entity.ParentID),
	}

	ret, err := EvaluateGovaluateExpressionByScope(entity.Expression, entity.ID)

	pass, _ := ret.(bool)

	var status string
	processor.Result.ResultStatus, status = getResultStatus(pass)

	processor.Result.Summary = fmt.Sprintf("断言\"%s\"结果为\"%s\"。", entity.Expression, status)

	processor.AddResultToParent()
	exec.SendExecMsg(*processor.Result, session.WsMsg)

	endTime := time.Now()
	processor.Result.EndTime = &endTime

	return
}

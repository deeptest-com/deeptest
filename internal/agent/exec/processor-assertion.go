package agentExec

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/domain"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/utils/exec"
	commonUtils "github.com/aaronchen2k/deeptest/pkg/lib/comm"
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
	processor.Result = &agentDomain.ScenarioExecResult{
		ID:                int(entity.ProcessorID),
		Name:              entity.Name,
		ProcessorCategory: entity.ProcessorCategory,
		ProcessorType:     entity.ProcessorType,
		StartTime:         &startTime,
		ParentId:          int(entity.ParentID),
		ScenarioId:        processor.ScenarioId,
		ProcessorId:       processor.ID,
		LogId:             session.Step.GetId(),
		ParentLogId:       processor.Parent.Result.LogId,
	}

	ret, err := EvaluateGovaluateExpressionByScope(entity.Expression, processor.ID)

	pass, _ := ret.(bool)

	var status string
	processor.Result.ResultStatus, status = getResultStatus(pass)

	//processor.Result.Summary = fmt.Sprintf("断言\"%s\"结果为\"%s\"。", entity.Expression, status)
	processor.Result.Summary = fmt.Sprintf("结果为\"%s\"。", status)
	detail := map[string]interface{}{"结果": status, "表达式": entity.Expression}
	processor.Result.Detail = commonUtils.JsonEncode(detail)
	processor.AddResultToParent()
	execUtils.SendExecMsg(*processor.Result, session.WsMsg)

	endTime := time.Now()
	processor.Result.EndTime = &endTime

	return
}

package agentExec

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/domain"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/utils/exec"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"time"
)

type ProcessorPrint struct {
	ID uint `json:"id" yaml:"id"`
	ProcessorEntityBase

	RightValue string `json:"rightValue" yaml:"rightValue"`
}

func (entity ProcessorPrint) Run(processor *Processor, session *Session) (err error) {
	logUtils.Infof("print entity")

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

	value := ReplaceVariableValue(entity.RightValue)
	//processor.Result.Summary = strings.ReplaceAll(fmt.Sprintf("%s为\"%v\"。", entity.RightValue, value), "<nil>", "空")
	processor.Result.Summary = fmt.Sprintf("%s", entity.RightValue)
	processor.Result.Detail = map[string]interface{}{"结果": value}

	processor.AddResultToParent()
	execUtils.SendExecMsg(*processor.Result, session.WsMsg)

	endTime := time.Now()
	processor.Result.EndTime = &endTime

	return
}

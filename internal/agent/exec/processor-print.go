package agentExec

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/domain"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/utils/exec"
	commonUtils "github.com/aaronchen2k/deeptest/pkg/lib/comm"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	uuid "github.com/satori/go.uuid"
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
		LogId:             uuid.NewV4(),
		ParentLogId:       processor.Parent.Result.LogId,
	}

	value := ReplaceVariableValue(entity.RightValue)
	//processor.Result.Summary = strings.ReplaceAll(fmt.Sprintf("%s为\"%v\"。", entity.RightValue, value), "<nil>", "空")
	processor.Result.Summary = fmt.Sprintf("%s", entity.RightValue)
	detail := map[string]interface{}{"结果": value}
	processor.Result.Detail = commonUtils.JsonEncode(detail)

	processor.AddResultToParent()
	execUtils.SendExecMsg(*processor.Result, session.WsMsg)

	endTime := time.Now()
	processor.Result.EndTime = &endTime

	return
}

package agentExec

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/domain"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/utils/exec"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	commonUtils "github.com/aaronchen2k/deeptest/pkg/lib/comm"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	uuid "github.com/satori/go.uuid"
	"strings"
	"time"
)

type ProcessorPrint struct {
	ID uint `json:"id" yaml:"id"`
	ProcessorEntityBase

	RightValue string `json:"rightValue" yaml:"rightValue"`
}

func (entity ProcessorPrint) Run(processor *Processor, session *Session) (err error) {
	defer func() {
		if errX := recover(); errX != nil {
			processor.Error(session, errX)
		}
	}()
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
		Round:             processor.Round,
	}

	value := ReplaceVariableValue(entity.RightValue, session.ExecUuid)
	value = strings.TrimSpace(value)

	//processor.Result.Summary = strings.ReplaceAll(fmt.Sprintf("%s为\"%v\"。", entity.RightValue, value), "<nil>", "空")
	processor.Result.Summary = fmt.Sprintf("%s", entity.RightValue)
	detail := map[string]interface{}{"name": entity.Name, "result": value}
	processor.Result.Detail = commonUtils.JsonEncode(detail)

	processor.AddResultToParent()
	execUtils.SendExecMsg(*processor.Result, consts.Processor, session.WsMsg)

	endTime := time.Now()
	processor.Result.EndTime = &endTime

	return
}

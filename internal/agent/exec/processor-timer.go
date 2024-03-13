package agentExec

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/domain"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/utils/exec"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	commonUtils "github.com/aaronchen2k/deeptest/pkg/lib/comm"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	uuid "github.com/satori/go.uuid"
	"time"
)

type ProcessorTimer struct {
	ID uint `json:"id" yaml:"id"`
	ProcessorEntityBase

	SleepTime int `json:"sleepTime" yaml:"sleepTime"`
}

func (entity ProcessorTimer) Run(processor *Processor, session *ExecSession) (err error) {
	defer func() {
		if errX := recover(); errX != nil {
			processor.Error(session, errX)
		}
	}()
	logUtils.Infof("timer entity")

	startTime := time.Now()
	processor.Result = &agentExecDomain.ScenarioExecResult{
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

	processor.Result.Summary = fmt.Sprintf("等待\"%d\"秒。", entity.SleepTime)
	processor.AddResultToParent()
	detail := map[string]interface{}{"name": entity.Name, "sleepTime": entity.SleepTime}
	processor.Result.Detail = commonUtils.JsonEncode(detail)
	execUtils.SendExecMsg(*processor.Result, consts.Processor, session.WsMsg)

	<-time.After(time.Duration(entity.SleepTime) * time.Second)

	endTime := time.Now()
	processor.Result.EndTime = &endTime

	return
}

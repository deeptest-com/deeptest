package agentExec

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/domain"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/utils/exec"
	"time"
)

type ProcessorTimer struct {
	ID uint `json:"id" yaml:"id"`
	ProcessorEntityBase

	SleepTime int `json:"sleepTime" yaml:"sleepTime"`
}

func (entity ProcessorTimer) Run(processor *Processor, session *Session) (log domain.Result, err error) {
	processor.Result = domain.Result{
		ID:                entity.ProcessorID,
		Name:              entity.Name,
		ProcessorCategory: entity.ProcessorCategory,
		ProcessorType:     entity.ProcessorType,
		ParentId:          entity.ParentID,
	}

	processor.Result.Summary = fmt.Sprintf("等待\"%d\"秒。", entity.SleepTime)
	exec.SendExecMsg(processor.Result, session.WsMsg)

	<-time.After(time.Duration(entity.SleepTime) * time.Second)

	return
}

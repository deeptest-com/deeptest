package agentExec

import (
	"github.com/aaronchen2k/deeptest/internal/agent/exec/domain"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/utils/exec"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	uuid "github.com/satori/go.uuid"
	"time"
)

type ProcessorRoot struct {
	ID uint `json:"id" yaml:"id"`
	ProcessorEntityBase
}

func (entity ProcessorRoot) Run(processor *Processor, session *ExecSession) (err error) {

	logUtils.Infof("root entity")

	startTime := time.Now()
	processor.Result = &agentExecDomain.ScenarioExecResult{
		ID:                int(entity.ProcessorID),
		Name:              "root",
		ProcessorCategory: entity.ProcessorCategory,
		ProcessorType:     entity.ProcessorType,
		StartTime:         &startTime,
		ParentId:          int(entity.ParentID),
		ScenarioId:        processor.ScenarioId,
		ProcessorId:       processor.ID,
		LogId:             uuid.NewV4(),
		//ParentLogId:       ,
	}

	execUtils.SendExecMsg(*processor.Result, consts.Processor, session.WsMsg)

	for _, child := range processor.Children {
		if GetForceStopExec(session.ExecUuid) {
			break
		}
		if child.Disable {
			continue
		}

		child.Run(session)
	}

	endTime := time.Now()
	processor.Result.EndTime = &endTime

	return
}

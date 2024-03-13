package agentExec

import (
	"github.com/aaronchen2k/deeptest/internal/agent/exec/domain"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/utils/exec"
	ptconsts "github.com/aaronchen2k/deeptest/internal/performance/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	commonUtils "github.com/aaronchen2k/deeptest/pkg/lib/comm"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	uuid "github.com/satori/go.uuid"
	"time"
)

type ProcessorPerformanceScenario struct {
	ID uint `json:"id" yaml:"id"`
	ProcessorEntityBase

	GenerateType ptconsts.GenerateType `json:"generateType"`

	Target int `json:"target"`

	Goal         ptconsts.GoalType `json:"goal"`
	Duration     int               `json:"duration"`
	Loop         int               `json:"loop"`
	ResponseTime float32           `json:"responseTime"`
	Qps          float32           `json:"qps"`
	FailRate     float32           `json:"failRate"`

	Stages       []ProcessorPerformanceStage `gorm:"-" json:"stages"`
	RunnerIds    []int                       `gorm:"-" json:"runnerIds"`
	RunnerIdsRaw string                      `json:"runnerIdsRaw"`
}
type ProcessorPerformanceStage struct {
	Duration int `json:"duration"`
	Target   int `json:"target"`

	ScenarioId uint `json:"scenarioId"`
}

func (entity ProcessorPerformanceScenario) Run(processor *Processor, session *ExecSession) (err error) {
	defer func() {
		if errX := recover(); errX != nil {
			processor.Error(session, errX)
		}
	}()
	logUtils.Infof("performance scenario entity")

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

	processor.AddResultToParent()
	detail := map[string]interface{}{"name": entity.Name}
	processor.Result.Detail = commonUtils.JsonEncode(detail)
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

package agentExec

import (
	"github.com/aaronchen2k/deeptest/internal/agent/exec/domain"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/utils/exec"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	commonUtils "github.com/aaronchen2k/deeptest/pkg/lib/comm"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	uuid "github.com/satori/go.uuid"
	"time"
)

type ProcessorLogic struct {
	ID uint `json:"id" yaml:"id"`
	ProcessorEntityBase

	Expression string `json:"expression" yaml:"expression"`
}

func (entity ProcessorLogic) Run(processor *Processor, session *Session) (err error) {
	logUtils.Infof("logic entity")

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

	typ := entity.ProcessorType
	pass := false
	detail := map[string]interface{}{"表达式": entity.Expression}
	if typ == consts.ProcessorLogicIf {
		var result interface{}
		result, err = EvaluateGovaluateExpressionByProcessorScope(entity.Expression, entity.ProcessorID)
		if err != nil {
			pass = false
		} else {
			var ok bool
			pass, ok = result.(bool)
			if !ok {
				pass = false
			}
		}
		detail["结果"] = pass
		processor.Result.Detail = commonUtils.JsonEncode(detail)
	} else if typ == consts.ProcessorLogicElse {
		brother, ok := getPreviousBrother(*processor)
		if ok && brother.Result.ResultStatus != consts.Pass {
			pass = true
		}
	}

	processor.Result.ResultStatus, processor.Result.Summary = getResultStatus(pass)
	processor.AddResultToParent()
	execUtils.SendExecMsg(*processor.Result, session.WsMsg)

	if pass {
		for _, child := range processor.Children {
			if child.Disable {
				continue
			}

			child.Run(session)
		}
	}

	endTime := time.Now()
	processor.Result.EndTime = &endTime

	return
}

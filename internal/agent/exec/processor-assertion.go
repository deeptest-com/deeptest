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

type ProcessorAssertion struct {
	ID uint `json:"id" yaml:"id"`
	ProcessorEntityBase

	Expression string `json:"expression" yaml:"expression"`
}

func (entity ProcessorAssertion) Run(processor *Processor, session *Session) (err error) {
	defer func() {
		if errX := recover(); errX != nil {
			processor.Error(session, errX)
		}
	}()
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
		LogId:             uuid.NewV4(),
		ParentLogId:       processor.Parent.Result.LogId,
		Round:             processor.Round,
	}

	expr := ReplaceDatapoolVariInGovaluateExpress(entity.Expression)
	ret, params, err := EvaluateGovaluateExpressionByProcessorScope(expr, processor.ID)

	pass, _ := ret.(bool)

	var status string
	processor.Result.ResultStatus, status = getResultStatus(pass)

	//processor.Result.Summary = fmt.Sprintf("断言\"%s\"结果为\"%s\"。", entity.Expression, status)
	processor.Result.Summary = fmt.Sprintf("结果为\"%s\"。", status)

	detail := map[string]interface{}{
		"name":       entity.Name,
		"expression": entity.Expression,
		"result":     pass,
		"actual":     fmt.Sprintf("%t", pass),
		"variables":  params,
	}
	processor.Result.Detail = commonUtils.JsonEncode(detail)

	processor.AddResultToParent()
	execUtils.SendExecMsg(*processor.Result, consts.Processor, session.WsMsg)

	endTime := time.Now()
	processor.Result.EndTime = &endTime

	return
}

package agentExec

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/domain"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/utils/exec"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"time"
)

type ProcessorVariable struct {
	ID uint `json:"id" yaml:"id"`
	ProcessorEntityBase

	VariableName string `json:"variableName" yaml:"variableName"`
	Expression   string `json:"expression" yaml:"expression"`
}

func (entity ProcessorVariable) Run(processor *Processor, session *Session) (err error) {
	logUtils.Infof("variable entity")

	startTime := time.Now()
	processor.Result = &agentDomain.ScenarioExecResult{
		ID:                int(entity.ProcessorID),
		Name:              entity.Name,
		ProcessorCategory: entity.ProcessorCategory,
		ProcessorType:     entity.ProcessorType,
		StartTime:         &startTime,
		ParentId:          int(entity.ParentID),
	}

	if entity.ProcessorType == consts.ProcessorVariableSet {
		var variableValue interface{}
		variableValue, err = EvaluateGovaluateExpressionByScope(entity.Expression, processor.ID)

		if err != nil {
			variableValue = err.Error()
		}

		SetVariable(processor.ParentId, entity.VariableName, variableValue, consts.Public) // set in parent scope
		processor.Result.Summary = fmt.Sprintf("\"%s\"为\"%v\"。", entity.VariableName, variableValue)

	} else if entity.ProcessorType == consts.ProcessorVariableClear {
		ClearVariable(processor.ID, entity.VariableName)
		processor.Result.Summary = fmt.Sprintf("\"%s\"成功。", entity.VariableName)
	}

	processor.AddResultToParent()
	execUtils.SendExecMsg(*processor.Result, session.WsMsg)

	endTime := time.Now()
	processor.Result.EndTime = &endTime

	return
}

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

type ProcessorVariable struct {
	ID uint `json:"id" yaml:"id"`
	ProcessorEntityBase

	VariableName string `json:"variableName" yaml:"variableName"`
	Expression   string `json:"expression" yaml:"expression"`
}

func (entity ProcessorVariable) Run(processor *Processor, session *ExecSession) (err error) {
	defer func() {
		if errX := recover(); errX != nil {
			processor.Error(session, errX)
		}
	}()
	logUtils.Infof("variable entity")

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
	detail := map[string]interface{}{"name": entity.Name, "variableName": entity.VariableName}
	if entity.ProcessorType == consts.ProcessorVariableSet {
		var variableValue interface{}

		variableValue, _ = NewGojaSimple().ExecJsFuncSimple(entity.Expression, session, true)

		SetVariable(processor.ParentId, entity.VariableName, variableValue, consts.ExtractorResultTypeString,
			consts.Public, session) // set in parent scope
		processor.Result.Summary = fmt.Sprintf("\"%s\"为\"%v\"。", entity.VariableName, variableValue)
		detail["variableValue"] = variableValue

	} else if entity.ProcessorType == consts.ProcessorVariableClear {
		ClearVariable(processor.ParentId, entity.VariableName, session)
		processor.Result.Summary = fmt.Sprintf("\"%s\"成功。", entity.VariableName)
	}

	processor.AddResultToParent()
	processor.Result.Detail = commonUtils.JsonEncode(detail)
	execUtils.SendExecMsg(*processor.Result, consts.Processor, session.ScenarioDebug.WsMsg)

	endTime := time.Now()
	processor.Result.EndTime = &endTime

	return
}

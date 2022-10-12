package agentExec

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/domain"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/utils/exec"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
)

type ProcessorVariable struct {
	ID uint `json:"id" yaml:"id"`
	ProcessorEntityBase

	VariableName string `json:"variableName" yaml:"variableName"`
	RightValue   string `json:"rightValue" yaml:"rightValue"`
}

func (entity ProcessorVariable) Run(processor *Processor, session *Session) (log domain.Result, err error) {
	logUtils.Infof("variable entity")

	processor.Result = domain.Result{
		ID:                entity.ProcessorID,
		Name:              entity.Name,
		ProcessorCategory: entity.ProcessorCategory,
		ProcessorType:     entity.ProcessorType,
		ParentId:          entity.ParentID,
	}

	if entity.ProcessorType == consts.ProcessorVariableSet {
		var variableValue interface{}
		variableValue, err = EvaluateGovaluateExpressionByScope(entity.RightValue, processor.ID)

		if err != nil {
			//	entity.Result.ResultStatus = consts.Fail
			// entity.Result.Summary = fmt.Sprintf("计算表达式\"%s\"错误，\"%s\"。", entity.RightValue, err.Error())
			//
			//	exec.SendErrorMsg(entity.Result, session.WsMsg)
			//	return
		}

		SetVariable(processor.ID, entity.VariableName, variableValue, consts.Local) // set in parent scope

		processor.Result.Summary = fmt.Sprintf("\"%s\"为\"%v\"。", entity.VariableName, variableValue)
		exec.SendExecMsg(processor.Result, session.WsMsg)

	} else if entity.ProcessorType == consts.ProcessorVariableClear {
		ClearVariable(processor.ID, entity.VariableName)

		processor.Result.Summary = fmt.Sprintf("\"%s\"成功。", entity.VariableName)
		exec.SendExecMsg(processor.Result, session.WsMsg)
	}

	return
}

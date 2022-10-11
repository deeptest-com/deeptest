package agentExec

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/domain"
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

	if entity.ProcessorType == consts.ProcessorVariableSet {
		var variableValue interface{}
		variableValue, err = EvaluateGovaluateExpressionByScope(entity.RightValue, entity.ID)

		// TODO: deal with the issue of "http://" in GovaluateExpression
		//if err != nil {
		//	entity.Result.ResultStatus = consts.Fail
		//	entity.Result.Summary = fmt.Sprintf("计算表达式\"%s\"错误 \"%s\"。", entity.RightValue, err.Error())
		//
		//	exec.SendErrorMsg(entity.Result, s.WsMsg)
		//	return
		//}

		entity.Result.Summary = fmt.Sprintf("\"%s\"为\"%v\"。", entity.VariableName, variableValue)
		SetVariable(entity.ID, entity.VariableName, variableValue, consts.Local) // set in parent scope

	} else if entity.ProcessorType == consts.ProcessorVariableClear {
		entity.Result.Summary = fmt.Sprintf("清除变量。")
		ClearVariable(entity.ID, entity.VariableName)
	}

	return
}

package agentExec

import (
	"errors"
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"strings"
)

func getDynamicVariableFromScope(processorId uint, propExpression string, execUuid string) (ret domain.ExecVariable, err error) {
	allValidIds := GetValidScopeIds(processorId, execUuid)

	if allValidIds != nil {
		for _, id := range *allValidIds {
			for _, item := range GetScopedVariables(execUuid)[id] {
				var ok bool
				if ret, ok = evaluateVariablePropExpressionValue(item, propExpression); ok {
					goto LABEL
				}
			}
		}
	}

	if ret.Name == "" { // not found
		err = errors.New(fmt.Sprintf("找不到变量\"%s\"", propExpression))
	}

LABEL:
	return
}

func getVariableFromShareVar(name string, execUuid string) (ret domain.ExecVariable, err error) {
	execScene := GetExecScene(execUuid)

	ret, err = getVariableFromList(name, execScene.ShareVars)

	return
}

func getVariableFromEnvVar(name string, execUuid string) (ret domain.ExecVariable, err error) {
	execScene := GetExecScene(execUuid)

	envId := execScene.DebugInterfaceToEnvMap[GetCurrDebugInterfaceId(execUuid)]

	vars := execScene.EnvToVariables[envId]

	ret, err = getVariableFromList(name, vars)

	return
}
func getVariableFromGlobalVar(name string, execUuid string) (ret domain.ExecVariable, err error) {
	ret, err = getVariableFromList(name, GetExecScene(execUuid).GlobalVars)

	return
}

func getVariableFromList(name string, list []domain.GlobalVar) (ret domain.ExecVariable, err error) {
	for _, v := range list {
		if v.Name == name {
			ret.Name = v.Name
			ret.Value = v.LocalValue
			break
		}
	}

	return
}

// like {name.prop}
func evaluateVariablePropExpressionValue(variable domain.ExecVariable, propExpression string) (
	ret domain.ExecVariable, ok bool) {
	arr := strings.Split(propExpression, ".")
	variableName := arr[0]

	if variable.Name == variableName {
		ret = variable
		ret.Name = propExpression // set name from item to item.a

		if len(arr) > 1 {
			variableProp := arr[1]
			ret.Value = variable.Value.(domain.VarKeyValuePair)[variableProp]
		}

		ok = true
	}

	return
}

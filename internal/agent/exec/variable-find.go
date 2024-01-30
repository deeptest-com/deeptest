package agentExec

import (
	"errors"
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
)

func getDynamicVariableFromScope(processorId uint, propExpression string, execUuid string) (ret domain.ExecVariable, err error) {
	allValidIds := GetValidScopeIds(processorId, execUuid)

	if allValidIds != nil {
		for _, id := range *allValidIds {
			for _, item := range GetScopedVariables(execUuid)[id] {
				if !(item.Scope == consts.Public || (item.Scope == consts.Private && id == processorId)) {
					continue
				}

				var ok bool
				ret, ok = EvaluateVariablePropExpressionValue(item, propExpression)

				if ok {
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

	ret, err = GetVariableFromList(name, execScene.ShareVars)

	return
}

func getVariableFromEnvVar(name string, execUuid string) (ret domain.ExecVariable, err error) {
	execScene := GetExecScene(execUuid)

	envId := execScene.DebugInterfaceToEnvMap[GetCurrDebugInterfaceId(execUuid)]

	vars := execScene.EnvToVariables[envId]

	ret, err = GetVariableFromList(name, vars)

	return
}
func getVariableFromGlobalVar(name, execUuid string) (ret domain.ExecVariable, err error) {
	execScene := GetExecScene(execUuid)

	ret, err = GetVariableFromList(name, execScene.GlobalVars)

	return
}

func GetVariableFromList(name string, list []domain.GlobalVar) (ret domain.ExecVariable, err error) {
	for _, v := range list {
		if v.Name == name {
			ret.Name = v.Name
			ret.Value = v.LocalValue
			break
		}
	}

	return
}

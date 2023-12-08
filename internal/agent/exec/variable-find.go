package agentExec

import (
	"errors"
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
)

func getDynamicVariableFromScope(processorId uint, propExpression string) (ret domain.ExecVariable, err error) {
	allValidIds := GetValidScopeIds(processorId)

	if allValidIds != nil {
		for _, id := range *allValidIds {
			for _, item := range ScopedVariables[id] {
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

func getVariableFromShareVar(name string) (ret domain.ExecVariable, err error) {
	ret, err = getVariableFromList(name, ExecScene.ShareVars)

	return
}

func getVariableFromEnvVar(name string) (ret domain.ExecVariable, err error) {
	envId := ExecScene.DebugInterfaceToEnvMap[CurrDebugInterfaceId]

	vars := ExecScene.EnvToVariables[envId]

	ret, err = getVariableFromList(name, vars)

	return
}
func getVariableFromGlobalVar(name string) (ret domain.ExecVariable, err error) {
	ret, err = getVariableFromList(name, ExecScene.GlobalVars)

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

package agentExec

import (
	"errors"
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
)

func getDynamicVariableFromScope(session *ExecSession, scopeId uint, propExpression string) (ret domain.ExecVariable, err error) {
	allValidIds := GetValidScopeIds(session, scopeId)

	if allValidIds != nil {
		for _, id := range *allValidIds {
			for _, item := range session.ScopedVariables[id] {
				if !(item.Scope == consts.Public || (item.Scope == consts.Private && id == scopeId)) {
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

func getVariableFromShareVar(session *ExecSession, name string) (ret domain.ExecVariable, err error) {
	execScene := session.ExecScene

	ret, err = GetVariableFromList(name, execScene.ShareVars)

	return
}

func getVariableFromEnvVar(session *ExecSession, name string) (ret domain.ExecVariable, err error) {
	execScene := session.ExecScene

	envId := uint(session.CurrEnvironmentId)
	if envId == 0 {
		envId = execScene.DebugInterfaceToEnvMap[session.CurrDebugInterfaceId]
	}

	vars := execScene.EnvToVariables[envId]

	ret, err = GetVariableFromList(name, vars)

	return
}
func getVariableFromGlobalVar(session *ExecSession, name string) (ret domain.ExecVariable, err error) {
	execScene := session.ExecScene

	ret, err = GetVariableFromList(name, execScene.GlobalVars)

	return
}

func GetVariableFromList(name string, list []domain.GlobalVar) (ret domain.ExecVariable, err error) {
	for _, v := range list {
		if v.Name == name {
			ret.Name = v.Name

			if v.LocalValue != "" {
				ret.Value = v.LocalValue
			} else if v.RemoteValue != "" {
				ret.Value = v.RemoteValue
			}

			break
		}
	}

	return
}

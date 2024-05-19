package agentExec

import (
	"errors"
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
)

func getDynamicVariableFromScope(propExpression string, processorId uint, session *ExecSession) (ret domain.ExecVariable, err error) {
	allValidIds := GetValidScopeIds(processorId, session)

	if allValidIds != nil {
		for _, id := range *allValidIds {
			for _, item := range session.ScenarioDebug.ScopedVariables[id] {
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

func getVariableFromShareVar(name string, session *ExecSession) (ret domain.ExecVariable, err error) {
	execScene := session.ExecScene

	ret, err = GetVariableFromList(name, execScene.ShareVars)

	return
}

func getVariableFromEnvVar(name string, session *ExecSession) (ret domain.ExecVariable, err error) {
	execScene := session.ExecScene

	envId := session.EnvironmentId
	if envId == 0 {
		envId = execScene.DebugInterfaceToEnvMap[session.InterfaceDebug.DebugInterfaceId]
	}

	vars := execScene.EnvToVariables[envId]

	ret, err = GetVariableFromList(name, vars)

	return
}
func getVariableFromGlobalVar(name string, session *ExecSession) (ret domain.ExecVariable, err error) {
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

package agentExec

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
)

func GetAllValidVariables(session *ExecSession) (ret map[string]domain.ExecVariable) {
	ret = map[string]domain.ExecVariable{}

	// global variables
	globalVars := session.ExecScene.GlobalVars
	popValidGlobalVariables(globalVars, &ret)

	// env variables
	execScene := session.ExecScene
	envId := session.EnvironmentId
	if envId == 0 {
		envId = execScene.DebugInterfaceToEnvMap[session.GetCurrDebugInterfaceId()]
	}
	envVars := execScene.EnvToVariables[envId]
	popValidGlobalVariables(envVars, &ret)

	// share variables
	shareVars := session.ExecScene.ShareVars
	popValidGlobalVariables(shareVars, &ret)

	// share variables
	allValidIds := GetValidScopeIds(session.GetCurrScenarioProcessorId(), session)
	if allValidIds != nil {
		for _, id := range *allValidIds {
			for _, item := range session.ScenarioDebug.ScopedVariables[id] {
				if !(item.Scope == consts.Public ||
					(item.Scope == consts.Private && id == session.GetCurrScenarioProcessorId())) {
					continue
				}

				addValidExecVariables(item, &ret)
			}
		}
	}

	return
}

func popValidGlobalVariables(srcVars []domain.GlobalVar, ret *map[string]domain.ExecVariable) {
	for _, v := range srcVars {
		if v.Name == "" {
			continue
		}

		val := ""
		if v.LocalValue != "" {
			val = v.LocalValue
		} else if v.RemoteValue != "" {
			val = v.RemoteValue
		}

		(*ret)[v.Name] = domain.ExecVariable{
			Name:  v.Name,
			Value: val,
		}
	}

	return
}

func addValidExecVariables(srcVar domain.ExecVariable, ret *map[string]domain.ExecVariable) {
	if srcVar.Name == "" {
		return
	}

	(*ret)[srcVar.Name] = domain.ExecVariable{
		Name:  srcVar.Name,
		Value: srcVar.Value,
	}

	return
}

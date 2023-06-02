package agentExec

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	stringUtils "github.com/aaronchen2k/deeptest/pkg/lib/string"
)

func getVariableValue(name string) (ret string) {
	// priority 1: A. shared vars generated by Debug Endpoint Interface in same serve
	//			   B. shared vars generated by Extractor and Processor in scenario
	ret = getValueFromShareVar(name)
	if ret != "" {
		return
	}

	// priority 2: environment vars on project's serve settings
	ret = getValueFromEnvVar(name)
	if ret != "" {
		return
	}

	// priority 3: global vars on project level
	ret = getValueFromGlobalVar(name)
	if ret != "" {
		return
	}

	return
}

func getValueFromShareVar(name string) (ret string) {
	if CurrProcessorId == 0 { // endpoint interface debug
		// try to find in vars that set by pre-condition scripts
		vars := listCachedVariable(0)
		for _, v := range vars {
			if v.Name == name {
				return stringUtils.InterfToStr(v.Value)
			}
		}

		// find in vars in scene
		ret = getValueFromList(name, ExecScene.ShareVars)

	} else { // run scenario
		//每次都更新缓存变量
		//if CachedShareVarByProcessorForRead[CurrProcessorId] == nil {
		CachedShareVarByProcessorForRead[CurrProcessorId] = GetCachedVariableMapInContext(CurrProcessorId)
		//}

		if CachedShareVarByProcessorForRead[CurrProcessorId][name] == nil {
			return ""
		}
		ret = fmt.Sprintf("%v", CachedShareVarByProcessorForRead[CurrProcessorId][name])
		fmt.Println(name, ret, "+++")
	}

	return
}
func getValueFromEnvVar(name string) (ret string) {
	envId := ExecScene.InterfaceToEnvMap[CurrInterfaceId]

	vars := ExecScene.EnvToVariables[envId]

	ret = getValueFromList(name, vars)

	return
}
func getValueFromGlobalVar(name string) (ret string) {
	ret = getValueFromList(name, ExecScene.GlobalVars)

	return
}

func getValueFromList(name string, list []domain.GlobalVar) (ret string) {
	for _, v := range list {
		if v.Name == name {
			ret = v.LocalValue
			break
		}
	}

	return
}

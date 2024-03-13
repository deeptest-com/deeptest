package service

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"github.com/kataras/iris/v12"
)

func UpdateLocalValues(execScene *domain.ExecScene, localVarsCache iris.Map) {
	for envId, varsObj := range localVarsCache {
		varMap := varsObj.(map[string]interface{})

		if envId == "0" { // global variable
			for index, globalVar := range execScene.GlobalVars {
				valValue, ok := varMap[globalVar.Name]
				if ok {
					execScene.GlobalVars[index].LocalValue = valValue.(string)
				}
			}

			// also replace the env variables under env key 0 if exist
			for envId2, envToVars := range execScene.EnvToVariables {
				if envId2 != 0 {
					continue
				}

				updateLocalValue(envId, envId2, varMap, &envToVars)
			}

		} else if envId != "" {
			for envId2, envToVars := range execScene.EnvToVariables {
				updateLocalValue(envId, envId2, varMap, &envToVars)
			}
		}
	}

	return
}

func updateLocalValue(envId string, envId2 uint, varMap map[string]interface{}, envToVars *[]domain.GlobalVar) {
	envId2Str := fmt.Sprintf("%v", envId2)
	if envId2Str != envId {
		return
	}

	for index, envVar := range *envToVars {
		valValue, ok := varMap[envVar.Name]
		if ok {
			(*envToVars)[index].LocalValue = valValue.(string)
		}
	}
}

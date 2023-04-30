package agentExec

import (
	"errors"
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"strings"
)

//func ImportVariables(processorId uint, variables domain.VarKeyValuePair, scope consts.ExtractorScope) (err error) {
//	for key, val := range variables {
//		newVariable := domain.ExecVariable{
//			Name:  key,
//			Value: val,
//			Scope: scope,
//		}
//
//		found := false
//		for i := 0; i < len(ScopedVariables[processorId]); i++ {
//			if ScopedVariables[processorId][i].Name == key {
//				ScopedVariables[processorId][i] = newVariable
//
//				found = true
//				break
//			}
//		}
//
//		if !found {
//			ScopedVariables[processorId] = append(ScopedVariables[processorId], newVariable)
//		}
//	}
//
//	return
//}

func GetVariable(processorId uint, variablePath string) (variable domain.ExecVariable, err error) {
	allValidIds := ScopeHierarchy[processorId]
	if allValidIds != nil {
		for _, id := range *allValidIds {
			for _, item := range ScopedVariables[id] {
				var ok bool
				if variable, ok = EvaluateVariableExpressionValue(item, variablePath); ok {
					goto LABEL
				}
			}
		}
	}

	if variable.Name == "" { // not found
		err = errors.New(fmt.Sprintf("找不到变量\"%s\"", variablePath))
	}

LABEL:

	return
}

func SetVariable(processorId uint, variableName string, variableValue interface{}, scope consts.ExtractorScope) (
	err error) {

	found := false

	newVariable := domain.ExecVariable{
		Name:  variableName,
		Value: variableValue,
		Scope: scope,
	}

	allValidIds := ScopeHierarchy[processorId]
	if allValidIds != nil {
		for _, id := range *allValidIds {
			for i := 0; i < len(ScopedVariables[id]); i++ {
				if ScopedVariables[id][i].Name == variableName {
					ScopedVariables[id][i] = newVariable

					found = true
					break
				}
			}
		}
	}

	if !found {
		ScopedVariables[processorId] = append(ScopedVariables[processorId], newVariable)
	}

	return
}

func ClearVariable(processorId uint, variableName string) (err error) {
	deleteIndex := -1

	targetScopeId := uint(0)

	allValidIds := ScopeHierarchy[processorId]
	if allValidIds != nil {
		for _, id := range *ScopeHierarchy[processorId] {
			for index, item := range ScopedVariables[id] {
				if item.Name == variableName {
					deleteIndex = index
					targetScopeId = id
					break
				}
			}
		}
	}

	if deleteIndex > -1 {
		if len(ScopedVariables[targetScopeId]) == deleteIndex+1 {
			ScopedVariables[targetScopeId] = make([]domain.ExecVariable, 0)
		} else {
			ScopedVariables[targetScopeId] = append(
				ScopedVariables[targetScopeId][:deleteIndex], ScopedVariables[targetScopeId][(deleteIndex+1):]...)
		}
	}

	return
}

func ReplaceVariableValue(value string) (ret string) {
	variablePlaceholders := GetVariablesInVariablePlaceholder(value)
	ret = value

	for _, placeholder := range variablePlaceholders {
		variablePlaceholder := fmt.Sprintf("${%s}", placeholder)

		oldVal := variablePlaceholder
		newVal := getPlaceholderValue(placeholder)

		ret = strings.ReplaceAll(ret, oldVal, newVal)
	}

	return
}

func getPlaceholderValue(placeholder string) (ret string) {
	typ := getPlaceholderType(placeholder)

	if typ == consts.PlaceholderTypeVariable {
		ret = getVariableValue(placeholder)

	} else if typ == consts.PlaceholderTypeDatapool {
		ret = getDatapoolValue(placeholder)

	} else if typ == consts.PlaceholderTypeFunction {
	}

	return
}

func getVariableValue(placeholder string) (ret string) {
	// 1. global vars of project
	// TODO:

	// 2. environment vars on serve
	// TODO: InterfaceToEnvMap + EnvToVariablesMap

	// 3. shared vars in scenario processors
	cache := CachedVariablesByProcessor[CurrProcessorId]
	if cache == nil {
		cache = GetCachedVariableMapInContext(CurrProcessorId)
	}

	ret = fmt.Sprintf("%v", cache[placeholder])

	return
}

func getPlaceholderType(placeholder string) (ret consts.PlaceholderType) {
	if strings.HasPrefix(placeholder, consts.PlaceholderPrefixDatapool.String()) {
		return consts.PlaceholderTypeDatapool
	} else if strings.HasPrefix(placeholder, consts.PlaceholderPrefixFunction.String()) {
		return consts.PlaceholderTypeFunction
	}

	return consts.PlaceholderTypeVariable
}

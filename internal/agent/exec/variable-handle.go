package agentExec

import (
	"errors"
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	commUtils "github.com/aaronchen2k/deeptest/internal/pkg/utils"
	"strings"
)

func GetVariableInScope(processorId uint, variablePath string) (variable domain.ExecVariable, err error) {
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

	allValidIds := &[]uint{uint(0)}
	if processorId > 0 {
		allValidIds = ScopeHierarchy[processorId]
	}

	for _, id := range *allValidIds {
		for i := 0; i < len(ScopedVariables[id]); i++ {
			if ScopedVariables[id][i].Name == variableName {
				ScopedVariables[id][i] = newVariable

				found = true
				break
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
	variablePlaceholders := commUtils.GetVariablesInExpressionPlaceholder(value)
	ret = value

	for _, placeholder := range variablePlaceholders {
		oldVal := fmt.Sprintf("${%s}", placeholder)
		newVal := getPlaceholderVariableValue(placeholder)

		ret = strings.ReplaceAll(ret, oldVal, newVal)
	}

	return
}

func getPlaceholderVariableValue(name string) (ret string) {
	typ := getPlaceholderType(name)

	if typ == consts.PlaceholderTypeVariable {
		ret = getVariableValue(name)
	} else if typ == consts.PlaceholderTypeDatapool {
		ret = getDatapoolValue(name)
	}
	//else if typ == consts.PlaceholderTypeFunction {
	//}

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

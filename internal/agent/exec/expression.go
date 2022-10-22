package agentExec

import (
	"fmt"
	"github.com/Knetic/govaluate"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/domain"
	commUtils "github.com/aaronchen2k/deeptest/internal/pkg/utils"
	_stringUtils "github.com/aaronchen2k/deeptest/pkg/lib/string"
	"regexp"
)

var (
	GovaluateFunctions = map[string]govaluate.ExpressionFunction{
		"length": func(args ...interface{}) (interface{}, error) {
			length := len(args[0].(string))
			return (float64)(length), nil
		},
		"match": func(args ...interface{}) (interface{}, error) {
			str := args[0].(string)
			regx := regexp.MustCompile("args[1].(string)")

			ret := regx.MatchString(str)
			return ret, nil
		},
		"uuid": func(args ...interface{}) (interface{}, error) {
			ret := _stringUtils.Uuid()
			return ret, nil
		},
	}
)

func EvaluateGovaluateExpressionByScope(expression string, scopeId uint) (ret interface{}, err error) {
	expr := commUtils.RemoveLeftVariableSymbol(expression)

	valueExpression, err := govaluate.NewEvaluableExpressionWithFunctions(expr, GovaluateFunctions)
	if err != nil {
		ret = expression
		return
	}

	parameters, err := generateGovaluateParamsByScope(expression, scopeId)
	if err != nil {
		return
	}

	ret, err = valueExpression.Evaluate(parameters)

	return
}
func EvaluateGovaluateExpressionWithVariables(expression string, variables map[string]interface{}) (ret interface{}, err error) {
	expr := commUtils.RemoveLeftVariableSymbol(expression)

	govaluateExpression, err := govaluate.NewEvaluableExpressionWithFunctions(expr, GovaluateFunctions)
	if err != nil {
		return
	}

	parameters, err := generateGovaluateParamsWithVariables(expression, variables)
	if err != nil {
		return
	}

	ret, err = govaluateExpression.Evaluate(parameters)

	return
}

func generateGovaluateParamsByScope(expression string, scopeId uint) (ret map[string]interface{}, err error) {
	ret = make(map[string]interface{}, 8)

	variables := GetVariablesInVariablePlaceholder(expression)

	for _, variableName := range variables {
		var vari domain.ExecVariable
		vari, err = GetVariable(scopeId, variableName)
		if err == nil {
			ret[variableName] = vari.Value
		}
	}

	return
}

func generateGovaluateParamsWithVariables(expression string, variableMap map[string]interface{}) (ret map[string]interface{}, err error) {
	ret = make(map[string]interface{}, 0)

	variables := GetVariablesInVariablePlaceholder(expression)

	for _, variableName := range variables {
		temp := fmt.Sprintf("${%s}", variableName)
		if val, ok := variableMap[temp]; ok {
			ret[variableName] = val
		}
	}

	return
}

func GetVariablesInVariablePlaceholder(expression string) (ret []string) {
	re := regexp.MustCompile("(?siU)\\${(.*)}")
	matchResultArr := re.FindAllStringSubmatch(expression, -1)

	for _, childArr := range matchResultArr {
		variableName := childArr[1]
		ret = append(ret, variableName)
	}

	return
}

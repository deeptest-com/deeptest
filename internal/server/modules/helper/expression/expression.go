package expressionHelper

import (
	"fmt"
	"github.com/Knetic/govaluate"
	commUtils "github.com/aaronchen2k/deeptest/internal/pkg/utils"
	"github.com/aaronchen2k/deeptest/internal/server/modules/helper/exec"
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

func EvaluateGovaluateExpression(expression string, variables map[string]interface{}) (ret interface{}, err error) {
	expr := commUtils.RemoveLeftVariableSymbol(expression)

	govaluateExpression, err := govaluate.NewEvaluableExpressionWithFunctions(expr, GovaluateFunctions)
	if err != nil {
		return
	}

	parameters, err := GenGovaluateParams(expression, variables)
	if err != nil {
		return
	}

	ret, err = govaluateExpression.Evaluate(parameters)

	return
}

func GenGovaluateParams(expression string, variableMap map[string]interface{}) (ret map[string]interface{}, err error) {
	ret = make(map[string]interface{}, 0)

	variables := execHelper.GetVariablesInVariablePlaceholder(expression)

	for _, variableName := range variables {
		temp := fmt.Sprintf("${%s}", variableName)
		if val, ok := variableMap[temp]; ok {
			ret[variableName] = val
		}
	}

	return
}

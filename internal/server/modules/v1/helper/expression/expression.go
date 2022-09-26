package expressionHelper

import (
	"fmt"
	"github.com/Knetic/govaluate"
	commUtils "github.com/aaronchen2k/deeptest/internal/pkg/utils"
	execHelper "github.com/aaronchen2k/deeptest/internal/server/modules/v1/helper/exec"
	_stringUtils "github.com/aaronchen2k/deeptest/pkg/lib/string"
	"regexp"
)

var (
	Functions = map[string]govaluate.ExpressionFunction{
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

func EvaluateGovaluateExpression(expression string, variables [][]interface{}) (ret interface{}, err error) {
	expr := commUtils.RemoveLeftVariableSymbol(expression)

	valueExpression, err := govaluate.NewEvaluableExpressionWithFunctions(expr, Functions)
	if err != nil {
		return
	}

	parameters, err := GenerateParamsFromArr(expression, variables)
	if err != nil {
		return
	}

	ret, err = valueExpression.Evaluate(parameters)

	return
}

func GenerateParamsFromArr(expression string, variableArr [][]interface{}) (ret map[string]interface{}, err error) {
	ret = make(map[string]interface{}, 0)

	variableMap := map[string]interface{}{}
	for _, item := range variableArr {
		variableMap[item[0].(string)] = item[1]
	}

	variables := execHelper.GetVariables(expression)

	for _, variableName := range variables {
		temp := fmt.Sprintf("${%s}", variableName)
		if val, ok := variableMap[temp]; ok {
			ret[variableName] = val
		}
	}

	return
}

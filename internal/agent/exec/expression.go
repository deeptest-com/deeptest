package agentExec

import (
	"github.com/Knetic/govaluate"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/utils"
	"github.com/aaronchen2k/deeptest/pkg/lib/string"
	"regexp"
	"strings"
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

// called by server checkpoint service
func EvaluateGovaluateExpressionWithVariables(expression string, varMap map[string]interface{}, datapools domain.Datapools) (
	ret interface{}, err error) {

	expr := commUtils.RemoveLeftVariableSymbol(expression)

	govaluateExpression, err := govaluate.NewEvaluableExpressionWithFunctions(expr, GovaluateFunctions)
	if err != nil {
		return
	}

	// 1
	paramValMap, err := generateGovaluateParamsWithVariables(expression, varMap, datapools)
	if err != nil {
		return
	}

	ret, err = govaluateExpression.Evaluate(paramValMap)

	return
}

// called by agent processor interface
func EvaluateGovaluateExpressionByScope(expression string, scopeId uint) (ret interface{}, err error) {
	expr := commUtils.RemoveLeftVariableSymbol(expression)

	valueExpression, err := govaluate.NewEvaluableExpressionWithFunctions(expr, GovaluateFunctions)
	if err != nil {
		ret = expression
		return
	}

	// 1
	parameters, err := generateGovaluateParamsByScope(expression, scopeId)
	if err != nil {
		return
	}

	ret, err = valueExpression.Evaluate(parameters)

	return
}

func generateGovaluateParamsByScope(expression string, scopeId uint) (ret domain.VarKeyValuePair, err error) {
	ret = make(map[string]interface{}, 8)

	variables := GetVariablesInExpressionPlaceholder(expression)

	for _, variableName := range variables {
		var vari domain.ExecVariable
		vari, err = GetVariableInScope(scopeId, variableName)
		if err == nil {
			ret[variableName] = vari.Value
		}
	}

	return
}

func generateGovaluateParamsWithVariables(expression string, variableMap map[string]interface{}, datapools domain.Datapools) (
	govaluateParams map[string]interface{}, err error) {

	govaluateParams = make(map[string]interface{}, 0)

	varsInExpression := GetVariablesInExpressionPlaceholder(expression)

	for _, varName := range varsInExpression {
		varNameWithoutPlus := strings.TrimLeft(varName, "+")

		val, ok := variableMap[varNameWithoutPlus]
		if ok {
			if varNameWithoutPlus != varName {
				val = _stringUtils.StrToInt(_stringUtils.InterfToStr(val))
			}
			govaluateParams[varNameWithoutPlus] = val
		}
	}

	return
}

func GetVariablesInExpressionPlaceholder(expression string) (ret []string) {
	re := regexp.MustCompile("(?siU)\\${(.*)}")
	matchResultArr := re.FindAllStringSubmatch(expression, -1)

	for _, childArr := range matchResultArr {
		variableName := childArr[1]
		ret = append(ret, variableName)
	}

	return
}

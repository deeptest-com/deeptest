package agentExec

import (
	"fmt"
	"github.com/Knetic/govaluate"
	valueUtils "github.com/aaronchen2k/deeptest/internal/agent/exec/utils/value"
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

// called by checkpoint
func EvaluateGovaluateExpressionWithDebugVariables(expression string) (ret interface{}, err error) {
	// 1
	params, err := generateGovaluateParamsWithVariables(expression)
	if err != nil {
		return
	}
	expr := commUtils.RemoveLeftVariableSymbol(expression)

	convertParams, convertExpr := convertGovaluateParamAndExpressionForProcessor(params, expr)

	govaluateExpression, err := govaluate.NewEvaluableExpressionWithFunctions(convertExpr, GovaluateFunctions)
	if err != nil {
		return
	}

	ret, err = govaluateExpression.Evaluate(convertParams)

	return
}

// called by agent processor interface
func EvaluateGovaluateExpressionByProcessorScope(expression string, scopeId uint) (ret interface{}, err error) {
	// 1
	params, err := generateGovaluateParamsByScope(expression, scopeId)
	if err != nil {
		return
	}
	expr := commUtils.RemoveLeftVariableSymbol(expression)

	convertParams, convertExpr := convertGovaluateParamAndExpressionForProcessor(params, expr)

	valueExpression, err := govaluate.NewEvaluableExpressionWithFunctions(convertExpr, GovaluateFunctions)
	if err != nil {
		ret = expression + " with error " + err.Error()
		return
	}

	ret, err = valueExpression.Evaluate(convertParams)

	return
}

func convertGovaluateParamAndExpressionForProcessor(params domain.VarKeyValuePair, expr string) (
	convertParams domain.VarKeyValuePair, convertExpr string) {

	convertParams = map[string]interface{}{}
	convertExpr = expr

	paramIndex := 0

	for key, val := range params {
		newKey := fmt.Sprintf("p___%d", paramIndex)

		convertParams[newKey] = val
		convertExpr = strings.ReplaceAll(convertExpr, key, newKey)

		paramIndex += 1
	}

	return
}

// a.1
func generateGovaluateParamsByScope(expression string, scopeId uint) (ret domain.VarKeyValuePair, err error) {
	ret = domain.VarKeyValuePair{}

	variables := commUtils.GetVariablesInExpressionPlaceholder(expression)

	for _, varName := range variables {
		varNameWithoutPlus := strings.TrimLeft(varName, "+")

		var vari domain.ExecVariable
		vari, err = GetVariable(scopeId, varNameWithoutPlus)
		variValueStr := valueUtils.InterfaceToStr(vari.Value)

		if err == nil {
			var val interface{}
			if strings.Index(varName, "+") == 0 { // is a number like ${+id}
				val = _stringUtils.ParseInt(variValueStr)
			} else {
				val = variValueStr
			}

			ret[varNameWithoutPlus] = val
		}
	}

	return
}

// a.2
func generateGovaluateParamsWithVariables(expression string) (ret domain.VarKeyValuePair, err error) {
	ret = domain.VarKeyValuePair{}

	variNames := commUtils.GetVariablesInExpressionPlaceholder(expression)

	for _, varName := range variNames {
		variNameWithoutPlus := strings.TrimLeft(varName, "+")

		vari, _ := GetVariable(CurrScenarioProcessorId, variNameWithoutPlus)
		variValueStr := commUtils.ObjectToStrAsVariValue(vari.Value)

		var val interface{}
		if strings.Index(varName, "+") == 0 { // is a number like ${+id}
			val = _stringUtils.ParseInt(variValueStr)
		} else {
			val = variValueStr
		}

		ret[variNameWithoutPlus] = val
	}

	return
}

func ReplaceDatapoolVariInGovaluateExpress(expression string) (ret string) {
	ret = expression
	variablePlaceholders := commUtils.GetVariablesInExpressionPlaceholder(expression)

	for _, placeholder := range variablePlaceholders {
		if strings.Index(placeholder, "_dp") != 0 && strings.Index(placeholder, "_dp") != 1 {
			continue
		}

		oldVal := fmt.Sprintf("${%s}", placeholder)

		placeholderWithoutPlus := strings.TrimLeft(placeholder, "+")
		newVal := getPlaceholderVariableValue(placeholderWithoutPlus)
		if strings.Index(placeholder, "+") != 0 {
			newVal = "'" + newVal + "'"
		}

		ret = strings.ReplaceAll(ret, oldVal, newVal)
	}

	// add space to replace a==-1 to a== -1
	ret = strings.ReplaceAll(ret, "==-", "== -")

	return
}

package agentExec

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	mockData "github.com/aaronchen2k/deeptest/internal/pkg/helper/openapi-mock/openapi/generator/data"
	commUtils "github.com/aaronchen2k/deeptest/internal/pkg/utils"
	_stringUtils "github.com/aaronchen2k/deeptest/pkg/lib/string"
	"regexp"
	"strings"
)

func ReplaceVariableValueInBody(value string, tenantId consts.TenantId, projectId uint, execUuid string) (ret string) {
	// add a plus to set field vale as a number
	// {"id": "${+dev_env_var1}"} => {"id": 2}

	variablePlaceholders := commUtils.GetVariablesInExpressionPlaceholder(value)
	ret = value

	for _, placeholder := range variablePlaceholders {
		oldVal := fmt.Sprintf("${%s}", placeholder)
		if strings.Index(placeholder, "+") == 0 { // replace it with a number, if has prefix +
			oldVal = "\"" + oldVal + "\""
		}

		placeholderWithoutPlus := strings.TrimLeft(placeholder, "+")
		newVal := getPlaceholderVariableValue(placeholderWithoutPlus, execUuid)

		ret = strings.ReplaceAll(ret, oldVal, newVal)
	}

	return
}

func ReplaceVariableValue(value string, tenantId consts.TenantId, projectId uint, execUuid string) (ret string) {
	ret = replaceVariableToken(value)
	variablePlaceholders := commUtils.GetVariablesInExpressionPlaceholder(ret)

	for _, token := range variablePlaceholders {
		isFunc, isExpr, name, params, variables := parseSingleVariableToken(token)
		if name == "" { // not a valid token
			continue
		}

		if name == "_mock" { // mock
			if len(params) > 0 {
				result, err := (&mockData.MockjsGenerator{}).GenerateByMockJsExpression(params[0].Name, "string")
				if err == nil {
					newVal := _stringUtils.InterfToStr(result)

					oldVal := fmt.Sprintf("${%s}", token)
					ret = strings.Replace(ret, oldVal, newVal, -1)
				}
			}

		} else if isFunc {
			if strings.HasPrefix(name, "_") { // buildin func
				expr := strings.TrimPrefix(token, "_")

				result := ExecJsFuncSimple(expr, tenantId, projectId, execUuid, variables, false)
				newVal := _stringUtils.InterfToStr(result)

				oldVal := fmt.Sprintf("${%s}", token)
				ret = strings.Replace(ret, oldVal, newVal, -1)

			} else if strings.Contains(name, ".") { // custom func
				result := ExecJsFuncSimple(token, tenantId, projectId, execUuid, variables, true)
				newVal := _stringUtils.InterfToStr(result)

				oldVal := fmt.Sprintf("${%s}", token)
				ret = strings.Replace(ret, oldVal, newVal, -1)
			}

		} else if isExpr { // expression
			result := ExecJsFuncSimple(token, tenantId, projectId, execUuid, variables, true)
			newVal := _stringUtils.InterfToStr(result)

			oldVal := fmt.Sprintf("${%s}", token)
			ret = strings.Replace(ret, oldVal, newVal, -1)

		} else { // variable ref like ${var_name}
			newVal := getPlaceholderVariableValue(strings.TrimLeft(name, "+"), execUuid)

			oldVal := fmt.Sprintf("${%s}", name)
			ret = strings.ReplaceAll(ret, oldVal, newVal)
		}
	}

	return
}

func getPlaceholderVariableValue(name string, execUuid string) (ret string) {
	typ := getPlaceholderType(name)

	if typ == consts.PlaceholderTypeVariable {
		variable, _ := GetVariable(GetCurrScenarioProcessorId(execUuid), name, execUuid)
		ret, _ = commUtils.ConvertValueForPersistence(variable.Value)

	} else if typ == consts.PlaceholderTypeDatapool {
		ret = getDatapoolValue(name, execUuid)
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

func parseSingleVariableToken(token string) (isFunc, isExpr bool, name string, params []TokenParam, variables []string) {
	regx2 := regexp.MustCompile(`^((?U).+)\((\+?.*)\)$`)
	arr2 := regx2.FindAllStringSubmatch(token, -1)

	if len(arr2) == 0 { // is not func
		name = token

		if !regexp.MustCompile(`"[_A-Za-z][_A-Za-z0-9]*"`).MatchString(token) { // is expression
			isExpr = true

			regx0 := regexp.MustCompile(`#\[(\+?[_A-Za-z0-9]+)\]`)
			arr0 := regx0.FindAllStringSubmatch(token, -1)

			if len(arr0) > 0 { // has variable in it
				for _, item := range arr0 {
					variables = append(variables, item[1])
				}
			}
		}

		return
	}

	// is func
	isFunc = true
	name = strings.TrimSpace(arr2[0][1])
	arr3 := strings.Split(arr2[0][2], ",")
	for _, item := range arr3 {
		paramName := strings.TrimSpace(item)
		if paramName == "" {
			continue
		}

		paramType := TokenParamTypeString

		if name == "_mock" {
			paramType = "mock"
			paramName = strings.Trim(paramName, `"'`)

			params = append(params, TokenParam{
				Name: paramName,
				Type: paramType,
			})

		} else {
			regx4 := regexp.MustCompile(`#\[(\+?[_A-Za-z0-9]+)\]`)
			arr4 := regx4.FindAllStringSubmatch(item, -1)

			if len(arr4) == 0 { // has no variable in it
				paramNameTrim := strings.Trim(paramName, `"'`)
				if paramNameTrim == paramName { // number
					paramType = TokenParamTypeNumber
				}

				params = append(params, TokenParam{
					Name: paramName,
					Type: paramType,
				})

				continue
			}

			for _, item := range arr4 {
				variables = append(variables, item[1])
			}
		}
	}

	return
}

func replaceVariableToken(str string) (ret string) {
	if len(str) > 2 && str[:2] == "${" && str[len(str)-1] == '}' {
		part := str[2 : len(str)-1]

		reg := regexp.MustCompile(`(?U)\${(\+?[_A-Za-z0-9]+)}`)
		part = reg.ReplaceAllString(part, "#[$1]")

		ret = fmt.Sprintf("${%s}", part)

	} else {
		ret = str

		reg := regexp.MustCompile(`(?U)\${(\+?[_A-Za-z0-9]+)}`)
		ret = reg.ReplaceAllString(ret, "#[$1]")
	}

	return
}

type TokenParam struct {
	Name string
	Type TokenParamType
}

type TokenParamType string

const (
	TokenParamTypeVariable TokenParamType = "variable"
	TokenParamTypeString   TokenParamType = "string"
	TokenParamTypeNumber   TokenParamType = "number"
)

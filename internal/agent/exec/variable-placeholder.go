package agentExec

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	commUtils "github.com/aaronchen2k/deeptest/internal/pkg/utils"
	"regexp"
	"strings"
)

func ReplaceVariableValueInBody(value string, execUuid string) (ret string) {
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

func ReplaceVariableValue(value string, execUuid string) (ret string) {
	ret = value

	value1 := replaceVariableToken(value)
	variablePlaceholders := commUtils.GetVariablesInExpressionPlaceholder(value1)

	for _, placeholder := range variablePlaceholders {
		isFunc, name, _ := parseSingleVariableToken(placeholder)
		if name == "" { // not a valid token
			continue
		}

		if isFunc {
			if name == "_mock" { // mock func

			} else if strings.HasPrefix(name, "_") { // buildin func

			} else { // custom func

			}

		} else { // variable ref like ${var_name}
			oldVal := fmt.Sprintf("${%s}", name)
			newVal := getPlaceholderVariableValue(strings.TrimLeft(name, "+"), execUuid)

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

func parseSingleVariableToken(str string) (isFunc bool, name string, params []TokenParam) {

	regx1 := regexp.MustCompile(`\${(\+?.+)}`)
	arr1 := regx1.FindAllStringSubmatch(str, -1)
	if len(arr1) == 0 {
		return
	}

	token := strings.TrimSpace(arr1[0][1])

	regx2 := regexp.MustCompile(`^((?U).+)\((\+?.+)\)$`)
	arr2 := regx2.FindAllStringSubmatch(token, -1)

	if len(arr2) == 0 { // is not func
		name = token
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
		} else {
			regx4 := regexp.MustCompile(`#\[(\+?[A-Za-z0-9]+)\]`)
			arr4 := regx4.FindAllStringSubmatch(item, -1)

			if len(arr4) > 0 {
				paramType = TokenParamTypeVariable
				paramName = arr4[0][1]
			} else {
				paramNameTrim := strings.Trim(paramName, `"'`)
				if paramNameTrim == paramName { // number
					paramType = TokenParamTypeNumber
				}
			}
		}

		p := TokenParam{
			Name: paramName,
			Type: paramType,
		}

		params = append(params, p)
	}

	return
}

func replaceVariableToken(str string) (ret string) {
	ret = str

	regx0 := regexp.MustCompile(`(?U)\${(\+?[A-Za-z0-9]+)}`)

	arr0 := regx0.FindAllStringSubmatch(str, -1)
	if len(arr0) > 0 {
		for _, item := range arr0 {
			ret = strings.Replace(ret,
				fmt.Sprintf("${%s}", item[1]),
				fmt.Sprintf("#[%s]", item[1]), 1)
		}
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

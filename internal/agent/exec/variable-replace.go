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

func ReplaceVariableValue(value string, session *ExecSession) (ret string) {
	ret = value
	placeholders := parseStatement(value)

	for _, placeholder := range placeholders {
		// variable
		if placeholder.Type == PlaceholderTypeVariable {
			placeholderWithoutPlus := strings.TrimLeft(placeholder.Content, "+")

			newVal := getPlaceholderVariableValue(placeholderWithoutPlus, session)
			oldVal := placeholder.Whole
			ret = strings.ReplaceAll(ret, oldVal, newVal)

			return
		}

		// mock
		if placeholder.Type == PlaceholderTypeMock {
			result, err := (&mockData.MockjsGenerator{}).GenerateByMockJsExpression(placeholder.Content, "string")
			if err == nil {
				newVal := _stringUtils.InterfToStr(result)

				oldVal := placeholder.Whole
				ret = strings.Replace(ret, oldVal, newVal, -1)
			}

			return
		}

		// expression
		result := ExecJsFuncSimple(placeholder.Whole, session, true)
		newVal := _stringUtils.InterfToStr(result)

		oldVal := placeholder.Whole
		ret = strings.Replace(ret, oldVal, newVal, -1)
	}

	return
}

func ReplaceVariableValueInBody(value string, session *ExecSession) (ret string) {
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
		newVal := getPlaceholderVariableValue(placeholderWithoutPlus, session)

		ret = strings.ReplaceAll(ret, oldVal, newVal)
	}

	return
}

func parseStatement(statement string) (ret []Placeholder) {
	reg := regexp.MustCompile(`\${(\+?[_A-Za-z][_A-Za-z0-9]*)}|(?U:{{(.+)}})`)

	arr := reg.FindAllStringSubmatch(statement, -1)
	for _, items := range arr {
		placeholder := Placeholder{
			Whole: items[0],
		}

		if items[1] != "" {
			placeholder.Type = PlaceholderTypeVariable
			placeholder.Content = items[1]

			ret = append(ret, placeholder)

			return
		}

		reg1 := regexp.MustCompile(`_mock\("(.+)"\)`)
		arr1 := reg1.FindAllStringSubmatch(items[2], -1)

		if len(arr1) > 0 {
			placeholder.Type = PlaceholderTypeMock
		} else {
			placeholder.Type = PlaceholderExpression
		}

		placeholder.Content = items[2]
	}

	return
}

func getPlaceholderVariableValue(name string, session *ExecSession) (ret string) {
	typ := getPlaceholderType(name)

	if typ == consts.PlaceholderTypeVariable {
		variable, _ := GetVariable(name, session.ScenarioDebug.CurrProcessorId, session)
		ret, _ = commUtils.ConvertValueForPersistence(variable.Value)

	} else if typ == consts.PlaceholderTypeDatapool {
		ret = getDatapoolValue(name, session)
	}

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

type Placeholder struct {
	Whole   string
	Content string
	Type    PlaceholderType
}

type PlaceholderType string

const (
	PlaceholderExpression   PlaceholderType = "expression"
	PlaceholderTypeMock     PlaceholderType = "mock"
	PlaceholderTypeVariable PlaceholderType = "variable"
)

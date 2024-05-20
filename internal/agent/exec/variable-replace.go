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
		// mock
		if placeholder.Type == PlaceholderTypeMock {
			result, err := (&mockData.MockjsGenerator{}).GenerateByMockJsExpression(placeholder.Content, "string")
			if err == nil {
				newVal := _stringUtils.InterfToStr(result)

				oldVal := placeholder.Whole
				ret = strings.Replace(ret, oldVal, newVal, -1)
			}

			continue
		}

		// expression
		if placeholder.Type == PlaceholderTypeExpression || placeholder.Type == PlaceholderTypeVariable {
			en := GojaSimple{}
			result := en.ExecJsFuncSimple(placeholder.Content, session, true)

			newVal := _stringUtils.InterfToStr(result)

			oldVal := placeholder.Whole
			ret = strings.Replace(ret, oldVal, newVal, -1)

			continue
		}
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
	reg := regexp.MustCompile(`(?U)\$\{(.+)}`)
	arr := reg.FindAllStringSubmatch(statement, -1)

	for _, items := range arr {
		placeholder := Placeholder{
			Whole:   items[0],
			Content: items[1],
		}

		if isVariable(placeholder.Content) {
			placeholder.Type = PlaceholderTypeVariable
			ret = append(ret, placeholder)

			continue
		}

		reg1 := regexp.MustCompile(`_mock\("(.+)"\)`)
		arr1 := reg1.FindAllStringSubmatch(placeholder.Content, -1)

		if len(arr1) > 0 {
			placeholder.Type = PlaceholderTypeMock
			placeholder.Content = arr1[0][1]
		} else {
			placeholder.Type = PlaceholderTypeExpression
		}

		ret = append(ret, placeholder)
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
	}

	return consts.PlaceholderTypeVariable
}

func isVariable(str string) (ret bool) {
	reg := regexp.MustCompile(`^\+?[_A-Za-z][_A-Za-z0-9]*$`)
	arr := reg.FindAllStringSubmatch(str, -1)

	if len(arr) > 0 {
		ret = true
	}

	return
}

type Placeholder struct {
	Whole   string
	Content string
	Type    PlaceholderType
}

type PlaceholderType string

const (
	PlaceholderTypeExpression PlaceholderType = "expression"
	PlaceholderTypeMock       PlaceholderType = "mock"
	PlaceholderTypeVariable   PlaceholderType = "variable"
)

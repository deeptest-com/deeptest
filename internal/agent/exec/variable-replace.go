package agentExec

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	mockData "github.com/aaronchen2k/deeptest/internal/pkg/helper/openapi-mock/openapi/generator/data"
	commUtils "github.com/aaronchen2k/deeptest/internal/pkg/utils"
	_commUtils "github.com/aaronchen2k/deeptest/pkg/lib/comm"
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
			result, _, _ := NewGojaSimple().ExecJsFuncSimple(placeholder.Content, session, true)

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
	//var x interface{}
	var data interface{}
	_commUtils.JsonDecode(value, &data)
	data = execJsFuncSimple(data, session)
	ret = _commUtils.JsonEncode(data)
	return
	/*
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
	*/

	return
}

func execJsFuncSimple(data interface{}, session *ExecSession) interface{} {
	if v, ok := data.(string); ok {
		return ReplaceVariableValue(v, session)
		/*
			if strings.HasPrefix(v, "${") && strings.HasSuffix(v, "}") {
				expression := v[2 : len(v)-1]
				result, _, _ := NewGojaSimple().ExecJsFuncSimple(expression, session, true)
				return _stringUtils.InterfToStr(result)
			}
		*/
	}
	if v, ok := data.(map[string]interface{}); ok {
		for key, item := range v {
			v[key] = execJsFuncSimple(item, session)
		}
	}

	if v, ok := data.([]interface{}); ok {
		for i, item := range v {
			v[i] = execJsFuncSimple(item, session)
		}

	}

	return data
}

func parseStatement(statement string) (ret []Placeholder) {
	// 前缀 ${任意长度的JS表达式} 后缀，可包含闭包函数function(){return 1}
	//reg := regexp.MustCompile(`(?U)\$\{(.+)}`)
	//arr := reg.FindAllStringSubmatch(statement, -1)

	//str := "xxxx${_mock('@string(pool, 1, 10)')}+==1222{_mock('@string(pool, 1, 10)')}-------"

	// 编译正则表达式
	// 注意：这里的正则表达式使用了非贪婪匹配（*?）来确保只匹配到最近的'}'
	re := regexp.MustCompile(`\$\{(.*?)\}`)

	// 查找所有匹配的子串
	matches := re.FindAllStringSubmatch(statement, -1)

	var arrOrArr [][]string
	// 打印所有匹配的子串（每个子串的第一个元素是完整匹配，第二个元素是括号内的匹配内容）
	for _, match := range matches {
		arrOrArr = append(arrOrArr, match)
	}

	/*
		placeholderStart := -1
		countLeftBrace := 0
		countRightBrace := 0
		arrOrArr := [][]string{}

		for index, char := range statement {
			if char == '{' {
				countLeftBrace++

				if index > 0 && statement[index-1] == '$' { // found ${
					placeholderStart = index
					arrOrArr = append(arrOrArr, []string{"${", ""})

					continue
				} else {
					arrOrArr[len(arrOrArr)-1][0] += string(char)
					arrOrArr[len(arrOrArr)-1][1] += string(char)
				}

				continue
			}

			if char == '}' {
				arrOrArr[len(arrOrArr)-1][0] += string(char)
				countRightBrace++

				if countLeftBrace > 0 && countLeftBrace == countRightBrace { // expression finish
					placeholderStart = -1
					countLeftBrace = 0
					countRightBrace = 0
				} else {
					arrOrArr[len(arrOrArr)-1][1] += string(char)
				}

				continue
			}

			if placeholderStart > 0 {
				arrOrArr[len(arrOrArr)-1][0] += string(char)
				arrOrArr[len(arrOrArr)-1][1] += string(char)
			}
		}
	*/

	for _, arr := range arrOrArr {
		placeholder := Placeholder{
			Whole:   arr[0],
			Content: arr[1],
		}

		if isVariable(placeholder.Content) {
			placeholder.Type = PlaceholderTypeVariable
			ret = append(ret, placeholder)

			continue
		}

		reg1 := regexp.MustCompile(`_mock\("(.+)"\)`)
		arr1 := reg1.FindAllStringSubmatch(placeholder.Content, -1)

		if len(arr1) > 0 { // is mock
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
		variable, _ := GetVariable(name, session.GetCurrScenarioProcessorId(), session)
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

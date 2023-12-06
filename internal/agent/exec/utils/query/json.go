package queryUtils

import (
	"encoding/json"
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	_stringUtils "github.com/aaronchen2k/deeptest/pkg/lib/string"
	"github.com/antchfx/jsonquery"
	"github.com/antchfx/xpath"
	"github.com/oliveagle/jsonpath"
	"strings"
)

func JsonPath(content string, expression string) (result string, resultType consts.ExtractorResultType) {
	var jsonData interface{}
	json.Unmarshal([]byte(content), &jsonData)

	obj, err := jsonpath.JsonPathLookup(jsonData, expression)

	if err != nil || obj == nil {
		result = consts.ExtractorErr
		return
	}

	switch obj.(type) {
	case string:
		result = obj.(string)
		resultType = consts.ExtractorResultTypeString

	case float64:
		result = fmt.Sprintf("%d", obj)
		resultType = consts.ExtractorResultTypeNumber

	case bool:
		result = fmt.Sprintf("%t", obj)
		resultType = consts.ExtractorResultTypeBool

	default:
		result = _stringUtils.JsonWithoutHtmlEscaped(obj)
		resultType = consts.ExtractorResultTypeObject
	}

	return
}

func JsonQuery(content string, expression string) (result string, resultType consts.ExtractorResultType) {
	doc, err := jsonquery.Parse(strings.NewReader(content))
	if err != nil {
		result = consts.ContentErr
		return
	}

	if isEvaluate(expression) {
		expr, _ := xpath.Compile(expression)
		float := expr.Evaluate(jsonquery.CreateXPathNavigator(doc))
		result = fmt.Sprintf("%v", float)

		return
	}

	elem, err := jsonquery.Query(doc, expression)

	if err != nil || elem == nil {
		result = consts.ExtractorErr
		return
	}

	obj := elem.Value()

	switch obj.(type) {
	case string:
		result = obj.(string)
		resultType = consts.ExtractorResultTypeString

	case float64:
		result = fmt.Sprintf("%d", obj)
		resultType = consts.ExtractorResultTypeNumber

	case bool:
		result = fmt.Sprintf("%t", obj)
		resultType = consts.ExtractorResultTypeBool

	default:
		result = _stringUtils.JsonWithoutHtmlEscaped(obj)
		resultType = consts.ExtractorResultTypeObject
	}

	return
}

func JsonQueryWithType(content string, expression string) (result interface{}) {
	doc, err := jsonquery.Parse(strings.NewReader(content))
	if err != nil {
		result = consts.ContentErr
		return
	}

	elem, err := jsonquery.Query(doc, expression)
	if err != nil || elem == nil {
		result = consts.ExtractorErr
		return
	}

	result = elem.Value()

	//isFloat64 := fmt.Sprintf("%T", result) == "float64"

	return
}

func isEvaluate(expression string) (ret bool) {
	ret = strings.Index(expression, "count") == 0

	return
}

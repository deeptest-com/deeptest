package queryUtils

import (
	"encoding/json"
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/antchfx/jsonquery"
	"github.com/antchfx/xpath"
	"strings"
)

func JsonQuery(content string, expression string) (result string) {
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
	default:
		bytes, err := json.Marshal(obj)
		if err != nil {
			result = err.Error()
		} else {
			result = string(bytes)
		}
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

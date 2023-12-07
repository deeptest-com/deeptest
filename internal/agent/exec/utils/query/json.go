package queryUtils

import (
	"encoding/json"
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	commUtils "github.com/aaronchen2k/deeptest/internal/pkg/utils"
	"github.com/antchfx/jsonquery"
	"github.com/antchfx/xpath"
	"github.com/oliveagle/jsonpath"
	"strings"
)

func JsonPath(content string, expression string) (result interface{}, resultType consts.ExtractorResultType, err error) {
	if content == "" || expression == "" {
		result = consts.ExtractorErr
		return
	}

	var jsonData interface{}
	err = json.Unmarshal([]byte(content), &jsonData)
	if err != nil {
		result = consts.ExtractorErr
		return
	}

	obj, err := jsonpath.JsonPathLookup(jsonData, expression)

	if err != nil || obj == nil {
		result = consts.ExtractorErr
		return
	}

	result, resultType = commUtils.GetValueInfo(obj)

	return
}

func JsonQuery(content string, expression string) (result interface{}, resultType consts.ExtractorResultType) {
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

	result, resultType = commUtils.GetValueInfo(obj)

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

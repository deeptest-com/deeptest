package queryHelper

import (
	"fmt"
	"github.com/antchfx/jsonquery"
	"strings"
)

func JsonQuery(content string, expression string) (result string) {
	doc, err := jsonquery.Parse(strings.NewReader(content))
	if err != nil {
		result = "ContentErr"
		return
	}

	expression, propName := GetExpressionForCssSelector(expression)
	elem, err := jsonquery.Query(doc, expression)
	if err != nil || elem == nil {
		result = "QueryErr"
		return
	}

	var obj interface{}
	if propName != "" {
		mp, ok := elem.Value().(map[string]interface{})
		if ok {
			obj = mp[propName]
		}
	} else {
		obj = elem.Value()
	}

	result = fmt.Sprintf("%v", obj)

	return
}

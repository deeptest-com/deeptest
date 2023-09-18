package queryUtils

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	_stringUtils "github.com/aaronchen2k/deeptest/pkg/lib/string"
	"github.com/antchfx/jsonquery"
	"strings"
)

func JsonQuery(content string, expression string) (result string) {
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

	obj := elem.Value()

	switch obj.(type) {
	case string:
		result = obj.(string)
	default:
		result = _stringUtils.JsonWithoutHtmlEscaped(obj)
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

package queryHelper

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/antchfx/htmlquery"
	"strings"
)

func HtmlQuery(content string, expression string) (result string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	doc, err := htmlquery.Parse(strings.NewReader(content))
	if err != nil {
		result = consts.ContentErr
		return
	}

	expression, propName := GetExpressionForCssSelector(expression)
	elem, err := htmlquery.Query(doc, expression)
	if err != nil || elem == nil {
		result = consts.ExtractorErr
		return
	}

	obj := ""
	if propName != "" {
		obj = htmlquery.SelectAttr(elem, propName)
	} else {
		obj = htmlquery.InnerText(elem)
		if obj == "" {
			obj = htmlquery.OutputHTML(elem, true)
		}
	}

	result = fmt.Sprintf("%v", obj)

	return
}

func GetExpressionForCssSelector(str string) (expression, propName string) {
	arr := strings.Split(str, "/")
	lastSection := arr[len(arr)-1]

	if strings.Index(lastSection, "@") == 0 {
		propName = strings.TrimPrefix(lastSection, "@")
		expression = strings.TrimSuffix(str, "/"+lastSection)
	} else {
		expression = str
	}

	return
}

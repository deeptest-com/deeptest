package queryHelper

import (
	"fmt"
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
		result = "ContentErr"
		return
	}

	expression, propName := GetExpressionForCssSelector(expression)
	list, err := htmlquery.QueryAll(doc, expression)
	if err != nil {
		result = "QueryErr"
		return
	}

	results := make([]string, 0)
	for _, item := range list {
		result := ""
		if propName != "" {
			result = htmlquery.SelectAttr(item, propName)
		} else {
			result = htmlquery.InnerText(item)
		}

		results = append(results, result)
	}

	result = strings.Join(results, ", ")

	return
}

func GetExpressionForCssSelector(str string) (expression, propName string) {
	arr := strings.Split(str, "/")
	lastSection := arr[len(arr)-1]

	if strings.Index(lastSection, "@") == 0 {
		propName = strings.TrimLeft(lastSection, "@")
		expression = strings.TrimRight(str, "/"+lastSection)
	} else {
		expression = str
	}

	return
}

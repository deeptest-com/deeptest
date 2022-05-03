package queryHelper

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	"github.com/antchfx/htmlquery"
	"strings"
)

func HtmlQuery(content string, extractor *model.InterfaceExtractor) {
	doc, err := htmlquery.Parse(strings.NewReader(content))
	if err != nil {
		extractor.Result = "ContentErr"
		return
	}

	expression, propName := getExpressionForCssSelector(extractor.Expression)
	list, err := htmlquery.QueryAll(doc, expression)
	if err != nil {
		extractor.Result = "QueryErr"
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

	extractor.Result = strings.Join(results, ", ")
}

func getExpressionForCssSelector(str string) (expression, propName string) {
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

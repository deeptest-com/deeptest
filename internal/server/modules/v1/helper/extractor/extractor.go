package extractorHelper

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	"github.com/antchfx/htmlquery"
	"github.com/antchfx/jsonquery"
	"github.com/antchfx/xmlquery"
	"strings"
)

func ParserXPath(content string, extractor *model.InterfaceExtractor) {
	doc, err := xmlquery.Parse(strings.NewReader(content))
	if err != nil {
		return
	}

	list := xmlquery.Find(doc, extractor.Expression)

	results := make([]string, 0)
	for _, item := range list {
		result := item.InnerText()
		results = append(results, result)
	}

	extractor.Result = strings.Join(results, ", ")
}

func ParserCssSelector(content string, extractor *model.InterfaceExtractor) {
	doc, err := htmlquery.Parse(strings.NewReader(content))
	if err != nil {
		return
	}

	expression, propName := getExpressionForCssSelector(extractor.Expression)
	list := htmlquery.Find(doc, expression)

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
	}

	return
}

func ParserJsonPath(content string, extractor *model.InterfaceExtractor) {
	doc, _ := jsonquery.Parse(strings.NewReader(content))
	list := jsonquery.Find(doc, extractor.Expression)

	results := make([]string, 0)
	for _, item := range list {
		result := item.InnerText()
		results = append(results, result)
	}

	extractor.Result = strings.Join(results, ", ")
}

func ParserBoundary(content string, extractor *model.InterfaceExtractor) {
	doc, err := htmlquery.Parse(strings.NewReader(content))
	if err != nil {
		return
	}

	elems := htmlquery.Find(doc, extractor.Expression)

	results := make([]string, 0)
	for _, elem := range elems {
		result := ""
		if extractor.Prop == "text" || extractor.Prop == "" {
			result = htmlquery.InnerText(elem)
		} else {
			result = htmlquery.SelectAttr(elem, extractor.Prop)
		}

		results = append(results, result)
	}

	extractor.Result = strings.Join(results, ", ")
}

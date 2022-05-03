package extractorHelper

import (
	"encoding/json"
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	"github.com/antchfx/htmlquery"
	"github.com/antchfx/xmlquery"
	"github.com/oliveagle/jsonpath"
	"strings"
)

func ParserJsonPath(content string, extractor *model.InterfaceExtractor) {
	var jsonData interface{}
	json.Unmarshal([]byte(content), &jsonData)

	pat, _ := jsonpath.Compile(extractor.Expression)
	result, err := pat.Lookup(jsonData)

	if err == nil && result != nil {
		extractor.Result = fmt.Sprintf("%v", result)
	}
}

func ParserXPath(content string, extractor *model.InterfaceExtractor) {
	doc, err := xmlquery.Parse(strings.NewReader(content))
	if err != nil {
		return
	}

	elems := xmlquery.Find(doc, extractor.Expression)

	results := make([]string, 0)
	for _, elem := range elems {
		result := ""
		if extractor.Prop == "text" || extractor.Prop == "" {
			result = elem.InnerText()
		} else {
			result = elem.SelectAttr(extractor.Prop)
		}

		results = append(results, result)
	}

	extractor.Result = strings.Join(results, ", ")
}

func ParserCssSelector(content string, extractor *model.InterfaceExtractor) {
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

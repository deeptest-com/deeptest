package queryHelper

import (
	"github.com/antchfx/xmlquery"
	"strings"
)

func XmlQuery(content string, expression string) (result string) {
	doc, err := xmlquery.Parse(strings.NewReader(content))
	if err != nil {
		result = "ContentErr"
		return
	}

	list, err := xmlquery.QueryAll(doc, expression)
	if err != nil {
		result = "QueryErr"
		return
	}

	results := make([]string, 0)
	for _, item := range list {
		result := item.InnerText()
		results = append(results, result)
	}

	result = strings.Join(results, ", ")

	return
}

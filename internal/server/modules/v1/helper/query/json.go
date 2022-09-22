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

	list, err := jsonquery.QueryAll(doc, expression)
	if err != nil {
		result = "QueryErr"
		return
	}

	results := make([]string, 0)
	for _, item := range list {
		result := fmt.Sprintf("%v", item.Value())
		results = append(results, result)
	}

	result = strings.Join(results, ", ")

	return
}

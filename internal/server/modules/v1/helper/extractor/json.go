package extractorHelper

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	"github.com/antchfx/jsonquery"
	"strings"
)

func JsonQuery(content string, extractor *model.InterfaceExtractor) {
	doc, _ := jsonquery.Parse(strings.NewReader(content))
	list, err := jsonquery.QueryAll(doc, extractor.Expression)
	if err != nil {
		extractor.Result = ""
		return
	}

	results := make([]string, 0)
	for _, item := range list {
		result := item.InnerText()
		results = append(results, result)
	}

	extractor.Result = strings.Join(results, ", ")
}

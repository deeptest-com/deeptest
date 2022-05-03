package extractorHelper

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	"github.com/antchfx/xmlquery"
	"strings"
)

func XmlQuery(content string, extractor *model.InterfaceExtractor) {
	doc, err := xmlquery.Parse(strings.NewReader(content))
	if err != nil {
		extractor.Result = ""
		return
	}

	list, err := xmlquery.QueryAll(doc, extractor.Expression)
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

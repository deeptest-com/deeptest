package queryHelper

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/antchfx/xmlquery"
	"strings"
)

func XmlQuery(content string, expression string) (result string) {
	doc, err := xmlquery.Parse(strings.NewReader(content))
	if err != nil {
		result = consts.ContentErr
		return
	}

	elemOrAttr, err := xmlquery.Query(doc, expression)
	if err != nil {
		result = consts.ExtractorErr
		return
	}

	result = elemOrAttr.InnerText()

	return
}

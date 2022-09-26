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

	elemOrAttr, err := xmlquery.Query(doc, expression)
	if err != nil {
		result = "QueryErr"
		return
	}

	result = elemOrAttr.InnerText()

	return
}

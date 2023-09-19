package queryUtils

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/antchfx/xmlquery"
	"github.com/antchfx/xpath"
	"strings"
)

func XmlQuery(content string, expression string) (result string) {
	doc, err := xmlquery.Parse(strings.NewReader(content))
	if err != nil {
		result = consts.ContentErr
		return
	}

	if isEvaluate(expression) {
		expr, _ := xpath.Compile(expression)
		float := expr.Evaluate(xmlquery.CreateXPathNavigator(doc))
		result = fmt.Sprintf("%v", float)

		return
	}

	expression, propName := GetExpressionForXpathSelector(expression)
	elem, err := xmlquery.Query(doc, expression)
	if err != nil || elem == nil {
		result = consts.ExtractorErr
		return
	}

	if propName != "" {
		result = fmt.Sprintf("%v", getXmlAttr(elem, propName))
	} else {
		result = fmt.Sprintf("%v", elem.InnerText())
	}

	return
}

func getXmlAttr(node *xmlquery.Node, name string) (ret string) {
	for _, attr := range node.Attr {
		if attr.Name.Local == name {
			return attr.Value
		}
	}
	return
}

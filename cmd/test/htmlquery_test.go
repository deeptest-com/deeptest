package test

import (
	zapLog "github.com/aaronchen2k/deeptest/internal/pkg/log"
	queryHelper "github.com/aaronchen2k/deeptest/internal/server/modules/v1/helper/query"
	fileUtils "github.com/aaronchen2k/deeptest/pkg/lib/file"
	"github.com/antchfx/htmlquery"
	"strings"
	"testing"
)

const (
	expression = "//form[@id=1]/input[@id=\"kw\"]/@class"
)

func TestA(t *testing.T) {
	zapLog.Init("server")

	//html := mockHelper.GetHtmlData()
	//html, _ := _httpUtils.Get("https://baidu.com")
	html := fileUtils.ReadFile("baidu.html")

	xpath, propName := queryHelper.GetExpressionForCssSelector(expression)

	doc, err := htmlquery.Parse(strings.NewReader(string(html)))
	list, err := htmlquery.QueryAll(doc, xpath)

	result := ""
	if propName != "" {
		result = htmlquery.SelectAttr(list[0], propName)
	} else {
		result = htmlquery.InnerText(list[0])
	}

	t.Logf("%v, %v", result, err)
}

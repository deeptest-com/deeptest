package test

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	"os"
	"testing"
)

func TestHtml(t *testing.T) {
	f, err := os.Open("./baidu.html")
	if err != nil {
		panic(err)
	}
	doc, err := htmlquery.Parse(f)
	if err != nil {
		panic(err)
	}

	expression := `//form[@id="form"]//input[@id="kw"]`
	elem, _ := htmlquery.Query(doc, expression)
	fmt.Println(htmlquery.SelectAttr(elem, "class"))
}

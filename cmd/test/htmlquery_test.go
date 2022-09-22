package test

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	"os"
	"testing"
)

func TestA(t *testing.T) {
	f, err := os.Open("./baidu.html")
	if err != nil {
		panic(err)
	}
	doc, err := htmlquery.Parse(f)
	if err != nil {
		panic(err)
	}
	//  "//form[@id=1]/input[@id=\"kw\"]/@class" is invalid. changed to @id="1",
	expression := `//form[@id="form"]//input[@id="kw"]`
	elem, _ := htmlquery.Query(doc, expression)
	fmt.Println(htmlquery.SelectAttr(elem, "class"))
}

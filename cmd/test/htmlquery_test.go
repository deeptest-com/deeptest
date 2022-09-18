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
	expression := `//form[@id="form"]/input[@id="kw"]`
	list, err := htmlquery.QueryAll(doc, expression)
	if err != nil {
		panic(err)
	}
	fmt.Println(len(list))
}

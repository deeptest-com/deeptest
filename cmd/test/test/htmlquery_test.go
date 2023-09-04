package test

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	"log"
	"os"
	"testing"
)

func TestHtml(t *testing.T) {
	f, err := os.Open("./baidu.html")
	if err != nil {
		log.Println(err)
	}
	doc, err := htmlquery.Parse(f)
	if err != nil {
		log.Println(err)
	}

	expression := `//form[@id="form"]//input[@id="kw"]`
	elem, _ := htmlquery.Query(doc, expression)
	fmt.Println(htmlquery.SelectAttr(elem, "class"))
}

package test

import (
	"fmt"
	"github.com/antchfx/jsonquery"
	"os"
	"testing"
)

func TestJson(t *testing.T) {
	f, _ := os.Open("./test.json")
	doc, _ := jsonquery.Parse(f)

	expression := "//*/text()[contains(.,'refer')]"

	elem, err := jsonquery.Query(doc, expression)
	fmt.Println(elem, err)
}

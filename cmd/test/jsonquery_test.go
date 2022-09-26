package test

import (
	"fmt"
	"github.com/antchfx/jsonquery"
	"os"
	"testing"
)

func TestJson(t *testing.T) {
	f, _ := os.Open("./products.json")
	doc, _ := jsonquery.Parse(f)

	expression := `//products/*[last()]`
	elem, _ := jsonquery.Query(doc, expression)
	fmt.Println(elem)
}

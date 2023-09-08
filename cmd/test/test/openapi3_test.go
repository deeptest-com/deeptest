package test

import (
	"context"
	"github.com/getkin/kin-openapi/openapi3"
	"log"
	"testing"
)

func TestOpenAPIV3(t *testing.T) {
	pth := "/Users/aaron/rd/project/gudi/deeptest/xdoc/openapi/openapi3/my-other-openapi.json"

	ctx := context.Background()
	loader := &openapi3.Loader{Context: ctx, IsExternalRefsAllowed: true}

	doc3, err := loader.LoadFromFile(pth)
	if err != nil {
	}

	log.Print(doc3)
}

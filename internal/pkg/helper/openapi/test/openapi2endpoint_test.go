package test

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/helper/openapi"
	"github.com/getkin/kin-openapi/openapi3"
	"testing"
)

func TestOpenapi2endpoint(t *testing.T) {
	doc := new(openapi3.T)
	endpoint := openapi.NewOpenapi2endpoint(doc).Convert()
	fmt.Println(endpoint)
	return

}

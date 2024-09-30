package casesHelper

import (
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	"github.com/getkin/kin-openapi/openapi3"
)

func GetApiPathItem(doc3 *openapi3.T) (pathItem *openapi3.PathItem, err error) {
	for key, _ := range doc3.Paths {
		pathItem = doc3.Paths[key] // get one
		return
	}

	return
}
func GetApiOperation(method consts.HttpMethod, pathItem *openapi3.PathItem) (operation *openapi3.Operation, err error) {
	if method == consts.GET {
		operation = pathItem.Get
	} else if method == consts.POST {
		operation = pathItem.Post
	} else if method == consts.PUT {
		operation = pathItem.Put
	} else if method == consts.DELETE {
		operation = pathItem.Delete
	} else if method == consts.PATCH {
		operation = pathItem.Patch
	} else if method == consts.HEAD {
		operation = pathItem.Head
	} else if method == consts.CONNECT {
		operation = pathItem.Connect
	} else if method == consts.OPTIONS {
		operation = pathItem.Options
	} else if method == consts.TRACE {
		operation = pathItem.Trace
	}

	return
}

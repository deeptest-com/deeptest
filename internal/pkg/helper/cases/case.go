package cases

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/jinzhu/copier"
)

func GenerateAlternativeCase(basicDebugData domain.DebugData, apiOperation *openapi3.Operation) (
	alternativeCase []domain.DebugData, err error) {

	cases, err := GenerateByQueryParams(basicDebugData, apiOperation.Parameters)
	alternativeCase = append(alternativeCase, cases...)

	return
}

func clone(basic domain.DebugData) (ret domain.DebugData) {
	copier.CopyWithOption(&ret, basic, copier.Option{DeepCopy: true})
	return
}

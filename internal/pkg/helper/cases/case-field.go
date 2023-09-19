package casesHelper

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/getkin/kin-openapi/openapi3"
)

func addPropCase(propVal *openapi3.Schema, parent *AlternativeCase) {
	if propVal.Type == OasFieldTypeArray.String() {
		itemVal := propVal.Items.Value

		arrCase := &AlternativeCase{
			Title:    "items",
			Category: consts.AlternativeCaseProp,
			IsDir:    true,
		}

		addPropCase(itemVal, arrCase)

		return

	} else if propVal.Type == OasFieldTypeObject.String() {

		return
	}

	propCase := &AlternativeCase{
		Sample: ExampleEmpty,

		Category:      consts.AlternativeCaseCase,
		Type:          consts.AlternativeCaseRequired,
		FieldRequired: true,
		IsDir:         false,
	}

	parent.Children = append(parent.Children, propCase)
}

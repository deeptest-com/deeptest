package casesHelper

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/getkin/kin-openapi/openapi3"
)

func addPropCase(propVal *openapi3.Schema, parent *AlternativeCase) {
	if propVal.Type == OasFieldTypeArray.String() {
		arrCase := &AlternativeCase{
			Title:    "items",
			Category: consts.AlternativeCaseProp,
			IsDir:    true,
		}

		addPropCase(propVal.Items.Value, arrCase)

		parent.Children = append(parent.Children, arrCase)

		return

	} else if propVal.Type == OasFieldTypeObject.String() {
		objCase := &AlternativeCase{
			Title:    "object",
			Category: consts.AlternativeCaseObject,
			IsDir:    true,
		}

		for _, propRef := range propVal.Properties {
			addPropCase(propRef.Value, objCase)
		}

		parent.Children = append(parent.Children, objCase)

		return
	}

	propCase := &AlternativeCase{
		Sample: propVal.Title,

		Category:      consts.AlternativeCaseCase,
		Type:          consts.AlternativeCaseRequired,
		FieldRequired: true,
		IsDir:         false,
	}

	parent.Children = append(parent.Children, propCase)
}

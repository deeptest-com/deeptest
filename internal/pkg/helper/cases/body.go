package casesHelper

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/getkin/kin-openapi/openapi3"
)

func LoadForBody(body *openapi3.RequestBodyRef) (category *AlternativeCase) {
	category = &AlternativeCase{
		Title:    "请求体",
		Category: consts.AlternativeCaseCategory,
		IsDir:    true,
	}

	for mediaType, mediaObj := range body.Value.Content {
		mediaCase := &AlternativeCase{
			Title:    mediaType,
			Category: consts.AlternativeCaseObject,
			IsDir:    true,
		}

		schema := mediaObj.Schema.Value
		props := schema.Properties
		requires := schema.Required

		for propName, propRef := range props {
			propVal := propRef.Value

			propCase := &AlternativeCase{
				Title:    propName,
				Category: consts.AlternativeCaseProp,
				IsDir:    true,
			}

			addPropCase(propName, propVal, requires, propCase)

			if len(propCase.Children) > 0 {
				mediaCase.Children = append(mediaCase.Children, propCase)
			}
		}

		category.Children = append(category.Children, mediaCase)
	}

	return
}

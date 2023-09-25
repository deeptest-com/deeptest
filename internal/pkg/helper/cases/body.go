package casesHelper

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	_stringUtils "github.com/aaronchen2k/deeptest/pkg/lib/string"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/kataras/iris/v12"
)

func LoadForBody(body *openapi3.RequestBodyRef) (category *AlternativeCase) {
	category = &AlternativeCase{
		Title:    "请求体",
		Category: consts.AlternativeCaseCategory,
		IsDir:    true,
		Key:      _stringUtils.Uuid(),
		Slots:    iris.Map{"icon": "icon"},
	}

	for mediaType, mediaObj := range body.Value.Content {
		mediaCase := &AlternativeCase{
			Title:    mediaType,
			Category: consts.AlternativeCaseObject,
			IsDir:    true,
			Key:      _stringUtils.Uuid(),
			Slots:    iris.Map{"icon": "icon"},
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
				Key:      _stringUtils.Uuid(),
				Slots:    iris.Map{"icon": "icon"},
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

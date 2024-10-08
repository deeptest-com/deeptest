package casesHelper

import (
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	_stringUtils "github.com/deeptest-com/deeptest/pkg/lib/string"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/kataras/iris/v12"
	"path"
	"strings"
)

func LoadForBody(body *openapi3.RequestBodyRef, doc3 *openapi3.T) (category *AlternativeCase) {
	category = &AlternativeCase{
		Title:    Category["body"],
		Path:     AddFix("body"),
		Category: consts.AlternativeCaseCategory,
		IsDir:    true,
		Key:      _stringUtils.Uuid(),
		Slots:    iris.Map{"icon": "icon"},
	}

	for mediaType, mediaObj := range body.Value.Content {
		mediaCase := &AlternativeCase{
			Title:    mediaType,
			Path:     path.Join(category.Path, AddFix(strings.ReplaceAll(mediaType, "/", "-"))),
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
			if propVal == nil {
				propVal = getRef(propRef.Ref, doc3)
			}

			if propVal == nil {
				continue
			}

			propCase := &AlternativeCase{
				Title:    propName,
				Path:     path.Join(mediaCase.Path, propName),
				Category: consts.AlternativeCaseProp,
				IsDir:    true,
				Key:      _stringUtils.Uuid(),
				Slots:    iris.Map{"icon": "icon"},
			}

			addPropCase(propName, propVal, requires, propCase, propCase.Path, doc3)

			if len(propCase.Children) > 0 {
				mediaCase.Children = append(mediaCase.Children, propCase)
			}
		}

		category.Children = append(category.Children, mediaCase)
	}

	return
}

func getRef(ref string, doc3 *openapi3.T) (ret *openapi3.Schema) {
	for name, item := range doc3.Components.Schemas {
		if "#/components/schemas/"+name == ref {
			ret = item.Value
			ret.Type = "object"
			break
		}
	}

	return
}

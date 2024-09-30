package casesHelper

import (
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	_stringUtils "github.com/deeptest-com/deeptest/pkg/lib/string"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/kataras/iris/v12"
	"path"
)

func LoadForQueryParams(params openapi3.Parameters) (category *AlternativeCase) {
	category = &AlternativeCase{
		Title:    Category["query"],
		Path:     AddFix(consts.ParamInQuery.String()),
		Category: consts.AlternativeCaseCategory,
		IsDir:    true,
		Key:      _stringUtils.Uuid(),
		Slots:    iris.Map{"icon": "icon"},
	}

	for _, paramRef := range params {
		paramVal := paramRef.Value
		if paramVal.In != consts.ParamInQuery.String() {
			continue
		}

		paramCase := &AlternativeCase{
			Title:    paramVal.Name,
			Path:     path.Join(category.Path, paramVal.Name),
			Category: consts.AlternativeCaseParam,
			IsDir:    true,
			Key:      _stringUtils.Uuid(),
			Slots:    iris.Map{"icon": "icon"},
		}

		addParamRequiredCase(paramVal, paramCase)
		addParamTypeCase(paramVal, paramCase)
		addParamEnumCase(paramVal, paramCase)
		addParamFormatCase(paramVal, paramCase)
		addParamRuleCase(paramVal, paramCase)

		if len(paramCase.Children) > 0 {
			category.Children = append(category.Children, paramCase)
		}
	}

	return
}

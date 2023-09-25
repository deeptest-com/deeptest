package casesHelper

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	_stringUtils "github.com/aaronchen2k/deeptest/pkg/lib/string"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/kataras/iris/v12"
)

func LoadForPathParams(params openapi3.Parameters) (category *AlternativeCase) {
	category = &AlternativeCase{
		Title:    "路径参数",
		Category: consts.AlternativeCaseCategory,
		IsDir:    true,
		Key:      _stringUtils.Uuid(),
		Slots:    iris.Map{"icon": "icon"},
	}

	for _, param := range params {
		if param.Value.In != consts.ParamInPath.String() {
			continue
		}

		paramCase := &AlternativeCase{
			Title:    param.Value.Name,
			Category: consts.AlternativeCaseParam,
			IsDir:    true,
			Key:      _stringUtils.Uuid(),
			Slots:    iris.Map{"icon": "icon"},
		}

		addParamRequiredCase(param.Value, paramCase)
		addParamTypeCase(param.Value, paramCase)
		addParamEnumCase(param.Value, paramCase)
		addParamFormatCase(param.Value, paramCase)
		addParamRuleCase(param.Value, paramCase)

		if len(paramCase.Children) > 0 {
			category.Children = append(category.Children, paramCase)
		}
	}

	return
}

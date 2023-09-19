package casesHelper

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/getkin/kin-openapi/openapi3"
)

func LoadForPathParams(params openapi3.Parameters) (category *AlternativeCase) {
	category = &AlternativeCase{
		Title:    "路径参数",
		Category: consts.AlternativeCaseGroup,
		IsDir:    true,
	}

	for _, param := range params {
		if param.Value.In != consts.ParamInPath.String() {
			continue
		}

		paramCase := &AlternativeCase{
			Title:    param.Value.Name,
			Category: consts.AlternativeCaseParam,
			IsDir:    true,
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

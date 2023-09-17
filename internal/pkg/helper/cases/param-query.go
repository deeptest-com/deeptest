package casesHelper

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"github.com/getkin/kin-openapi/openapi3"
)

func LoadForQueryParams(params openapi3.Parameters) (
	category *domain.AlternativeCase) {

	category = &domain.AlternativeCase{
		Category: consts.AlternativeCaseGroup,
		IsDir:    true,
	}

	for _, param := range params {
		field := &domain.AlternativeCase{
			Category: consts.AlternativeCaseField,
			IsDir:    true,
		}

		addRequiredCase(param, category)
		addTypeCase(param, category)

		if len(field.Children) > 0 {
			category.Children = append(category.Children, field)
		}
	}

	return
}

func addRequiredCase(param *openapi3.ParameterRef, parent *domain.AlternativeCase) {
	if !param.Value.Required {
		return
	}

	required := &domain.AlternativeCase{
		Category:      consts.AlternativeCaseCase,
		Type:          consts.AlternativeCaseRequired,
		FieldRequired: true,
		IsDir:         false,
	}

	parent.Children = append(parent.Children, required)
}

func addTypeCase(param *openapi3.ParameterRef, parent *domain.AlternativeCase) {
	schema := param.Value.Schema.Value
	typ := FieldType(schema.Type)

	if typ == FieldTypeString {
		return
	}

	typeCase := &domain.AlternativeCase{
		Category:  consts.AlternativeCaseCase,
		Type:      consts.AlternativeCaseTyped,
		FieldType: typ,
		IsDir:     false,
	}

	parent.Children = append(parent.Children, typeCase)
}

func GenerateByQueryParams(basic domain.DebugData, params openapi3.Parameters) (ret []domain.DebugData, err error) {
	for _, param := range params {
		cases, _ := GenerateByQueryParam(basic, param)

		ret = append(ret, cases...)
	}

	return
}

func GenerateByQueryParam(basic domain.DebugData, param *openapi3.ParameterRef) (
	ret []domain.DebugData, err error) {

	casesByParamRequired := generateByQueryParamRequired(basic, param)
	ret = append(ret, casesByParamRequired...)

	casesByParamType := generateByQueryParamType(basic, param)
	ret = append(ret, casesByParamType...)

	return
}

func generateByQueryParamRequired(basic domain.DebugData, param *openapi3.ParameterRef) (ret []domain.DebugData) {
	if param.Value.Required {
		cs, err := updateQueryParam(basic, param.Value.Name, "")
		if err == nil {
			ret = append(ret, cs)
		}
	}

	return
}

func generateByQueryParamType(basic domain.DebugData, param *openapi3.ParameterRef) (ret []domain.DebugData) {
	schema := param.Value.Schema.Value

	typ := schema.Type
	exceptionStringVal := ""

	// no string
	if typ == FieldTypeInteger.String() {
		exceptionStringVal = ExampleString
	} else if typ == FieldTypeNumber.String() {
		exceptionStringVal = ExampleString
	} else if typ == FieldTypeBoolean.String() {
		exceptionStringVal = ExampleString
	}
	//else if typ == FieldTypeString.String() {
	//
	//} else if typ == FieldTypeArray.String() {
	//
	//} else if typ == FieldTypeObject.String() {
	//
	//} else if typ == FieldTypeNull.String() {
	//
	//}

	if exceptionStringVal != "" {
		cs, err := updateQueryParam(basic, param.Value.Name, exceptionStringVal)
		if err == nil {
			ret = append(ret, cs)
		}
	}

	// integer
	if typ == FieldTypeInteger.String() {
		exceptionStringVal = ExampleStringFloat
		cs, err := updateQueryParam(basic, param.Value.Name, exceptionStringVal)
		if err == nil {
			ret = append(ret, cs)
		}
	}

	return
}

func updateQueryParam(basic domain.DebugData, key, value string) (ret domain.DebugData, err error) {
	ret = clone(basic)

	for index, param := range ret.QueryParams {
		if param.Name == key {
			ret.QueryParams[index].Value = value
		}
	}

	return
}

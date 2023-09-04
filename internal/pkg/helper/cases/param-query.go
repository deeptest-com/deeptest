package cases

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"github.com/getkin/kin-openapi/openapi3"
)

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
	if true { // param.Value.Required, away test empty
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

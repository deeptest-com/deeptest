package casesHelper

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/getkin/kin-openapi/openapi3"
)

func LoadForQueryParams(params openapi3.Parameters) (category *AlternativeCase) {
	category = &AlternativeCase{
		Title:    "查询参数",
		Category: consts.AlternativeCaseGroup,
		IsDir:    true,
	}

	for _, paramRef := range params {
		paramVal := paramRef.Value
		if paramVal.In != consts.ParamInQuery.String() {
			continue
		}

		paramCase := &AlternativeCase{
			Title:    paramVal.Name,
			Category: consts.AlternativeCaseParam,
			IsDir:    true,
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

//func GenerateByQueryParams(basic DebugData, params openapi3.Parameters) (ret []DebugData, err error) {
//	for _, param := range params {
//		cases, _ := GenerateByQueryParam(basic, param)
//
//		ret = append(ret, cases...)
//	}
//
//	return
//}
//
//func GenerateByQueryParam(basic DebugData, param *openapi3.ParameterRef) (
//	ret []DebugData, err error) {
//
//	casesByParamRequired := generateByQueryParamRequired(basic, param)
//	ret = append(ret, casesByParamRequired...)
//
//	casesByParamType := generateByQueryParamType(basic, param)
//	ret = append(ret, casesByParamType...)
//
//	return
//}
//
//func generateByQueryParamRequired(basic DebugData, param *openapi3.ParameterRef) (ret []DebugData) {
//	if param.Value.Required {
//		cs, err := updateQueryParam(basic, param.Value.Name, "")
//		if err == nil {
//			ret = append(ret, cs)
//		}
//	}
//
//	return
//}
//
//func generateByQueryParamType(basic DebugData, param *openapi3.ParameterRef) (ret []DebugData) {
//	schema := param.Value.Schema.Value
//
//	typ := schema.Type
//	var exceptionStringVal interface{}
//
//	// no string
//	if typ == OasFieldTypeInteger.String() {
//		exceptionStringVal = RandStr()
//	} else if typ == OasFieldTypeNumber.String() {
//		exceptionStringVal = RandStr()
//	} else if typ == OasFieldTypeBoolean.String() {
//		exceptionStringVal = RandStr()
//	}
//	//else if typ == FieldTypeString.String() {
//	//
//	//} else if typ == FieldTypeArray.String() {
//	//
//	//} else if typ == FieldTypeObject.String() {
//	//
//	//} else if typ == FieldTypeNull.String() {
//	//
//	//}
//
//	if exceptionStringVal != "" {
//		cs, err := updateQueryParam(basic, param.Value.Name, exceptionStringVal)
//		if err == nil {
//			ret = append(ret, cs)
//		}
//	}
//
//	// integer
//	if typ == OasFieldTypeInteger.String() {
//		exceptionStringVal = RandFloat32()
//		cs, err := updateQueryParam(basic, param.Value.Name, exceptionStringVal)
//		if err == nil {
//			ret = append(ret, cs)
//		}
//	}
//
//	return
//}
//
//func updateQueryParam(basic DebugData, key, value string) (ret DebugData, err error) {
//	ret = clone(basic)
//
//	for index, param := range ret.QueryParams {
//		if param.Name == key {
//			ret.QueryParams[index].Value = value
//		}
//	}
//
//	return
//}

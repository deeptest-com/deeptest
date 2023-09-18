package casesHelper

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/getkin/kin-openapi/openapi3"
	"math"
)

func LoadForQueryParams(params openapi3.Parameters) (
	category *AlternativeCase) {

	category = &AlternativeCase{
		Category: consts.AlternativeCaseGroup,
		IsDir:    true,
	}

	for _, param := range params {
		field := &AlternativeCase{
			Category: consts.AlternativeCaseField,
			IsDir:    true,
		}

		addRequiredCase(param, category)
		addTypeCase(param, category)
		addEnumCase(param, category)
		addFormatCase(param, category)
		addRuleCase(param, category)

		if len(field.Children) > 0 {
			category.Children = append(category.Children, field)
		}
	}

	return
}

func addRequiredCase(param *openapi3.ParameterRef, parent *AlternativeCase) {
	if !param.Value.Required {
		return
	}

	required := &AlternativeCase{
		Sample: ExampleEmpty,

		Category:      consts.AlternativeCaseCase,
		Type:          consts.AlternativeCaseRequired,
		FieldRequired: true,
		IsDir:         false,
	}

	parent.Children = append(parent.Children, required)
}

func addTypeCase(param *openapi3.ParameterRef, parent *AlternativeCase) {
	schema := param.Value.Schema.Value
	typ := OasFieldType(schema.Type)

	if typ == OasFieldTypeAny || typ == OasFieldTypeString {
		return
	}

	var sample interface{}
	if typ == OasFieldTypeBoolean || typ == OasFieldTypeNumber || typ == OasFieldTypeArray {
		sample = RandStr()
	} else if typ == OasFieldTypeInteger {
		sample = RandFloat32()
	}

	typeCase := &AlternativeCase{
		Sample: sample,

		Category:  consts.AlternativeCaseCase,
		Type:      consts.AlternativeCaseTyped,
		FieldType: typ,
		IsDir:     false,
	}

	parent.Children = append(parent.Children, typeCase)
}

func addEnumCase(param *openapi3.ParameterRef, parent *AlternativeCase) {
	schema := param.Value.Schema.Value
	enum := schema.Enum

	if enum == nil {
		return
	}

	typeCase := &AlternativeCase{
		Sample: RandStr(),

		Category:  consts.AlternativeCaseCase,
		Type:      consts.AlternativeCaseEnum,
		FieldType: OasFieldType(schema.Type),
		IsDir:     false,
	}

	parent.Children = append(parent.Children, typeCase)
}

func addFormatCase(param *openapi3.ParameterRef, parent *AlternativeCase) {
	schema := param.Value.Schema.Value
	typ := OasFieldType(schema.Type)
	format := OasFieldFormat(schema.Format)

	if format == "" {
		return
	}

	var sample interface{}
	if typ == OasFieldTypeInteger {
		if format == OasFieldFormatInt32 {
			sample = RandInt64()
		} else if format == OasFieldFormatInt64 {
			sample = RandStr()
		}
	} else if typ == OasFieldTypeNumber {
		if format == OasFieldFormatFloat {
			sample = RandFloat64()
		} else if format == OasFieldFormatDouble {
			sample = RandStr()
		}
	} else if typ == OasFieldTypeString {
		sample = RandStr()
	}

	typeCase := &AlternativeCase{
		Sample: sample,

		Category:  consts.AlternativeCaseCase,
		Type:      consts.AlternativeCaseEnum,
		FieldType: OasFieldType(schema.Type),
		IsDir:     false,
	}

	parent.Children = append(parent.Children, typeCase)
}

func addRuleCase(param *openapi3.ParameterRef, parent *AlternativeCase) {
	schema := param.Value.Schema.Value
	typ := OasFieldType(schema.Type)

	var sample interface{}
	if typ == OasFieldTypeInteger || typ == OasFieldTypeNumber {
		if *(schema.Min) != 0 {
			sample = *schema.Min - 1
		} else if *(schema.Max) != 0 {
			sample = *schema.Max + 1
		} else if *schema.MaxLength > 0 {
			if typ == OasFieldTypeInteger {
				sample = 1 * math.Pow(10, float64(*schema.MaxLength))
			} else {
				sample = 1 / math.Pow(10, float64(*schema.MaxLength-1))
			}
		} else if schema.MinLength > 0 {
			sample = 1
		} else if *schema.MultipleOf != 0 {
			if typ == OasFieldTypeInteger {
				sample = *schema.MultipleOf + 1
			} else {
				sample = *schema.MultipleOf + *schema.MultipleOf*0.1
			}
		} else if schema.ExclusiveMin {
			sample = *schema.Min
		} else if schema.ExclusiveMax {
			sample = *schema.Max
		}

	} else {
		if schema.Pattern != "" {
			sample = RandStrSpecial()
		} else if *(schema.MaxLength) > 0 {
			sample = RandStrWithLen(int(*(schema.MaxLength) + 1))
		} else if schema.MinLength > 0 {
			sample = RandStrWithLen(int(schema.MinLength - 1))
		}
	}

	typeCase := &AlternativeCase{
		Sample: sample,

		Category:  consts.AlternativeCaseCase,
		Type:      consts.AlternativeCaseEnum,
		FieldType: OasFieldType(schema.Type),
		IsDir:     false,
	}

	parent.Children = append(parent.Children, typeCase)
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

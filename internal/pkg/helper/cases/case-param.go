package casesHelper

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/getkin/kin-openapi/openapi3"
	"math"
)

func addParamRequiredCase(paramVal *openapi3.Parameter, parent *AlternativeCase) {
	if !paramVal.Required {
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

func addParamTypeCase(paramVal *openapi3.Parameter, parent *AlternativeCase) {
	schema := paramVal.Schema.Value
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

func addParamEnumCase(paramVal *openapi3.Parameter, parent *AlternativeCase) {
	schema := paramVal.Schema.Value
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

func addParamFormatCase(paramVal *openapi3.Parameter, parent *AlternativeCase) {
	schema := paramVal.Schema.Value
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

func addParamRuleCase(paramVal *openapi3.Parameter, parent *AlternativeCase) {
	schema := paramVal.Schema.Value
	typ := OasFieldType(schema.Type)

	var sample interface{}
	if typ == OasFieldTypeInteger || typ == OasFieldTypeNumber {
		if schema.Min != nil && *schema.Min != 0 {
			sample = *schema.Min - 1
			addRuleCase(paramVal.Name, sample, schema.Type, consts.AlternativeCaseRulesMin, parent)
		}

		if schema.Max != nil && *schema.Max != 0 {
			sample = *schema.Max + 1
			addRuleCase(paramVal.Name, sample, schema.Type, consts.AlternativeCaseRulesMax, parent)
		}

		if schema.MaxLength != nil && *schema.MaxLength > 0 {
			if typ == OasFieldTypeInteger {
				sample = 1 * math.Pow(10, float64(*schema.MaxLength))
			} else {
				sample = 1 / math.Pow(10, float64(*schema.MaxLength-1))
			}
			addRuleCase(paramVal.Name, sample, schema.Type, consts.AlternativeCaseRulesMaxLength, parent)
		}

		if schema.MinLength > 0 {
			sample = 1
			addRuleCase(paramVal.Name, sample, schema.Type, consts.AlternativeCaseRulesMinLength, parent)
		}

		if schema.MultipleOf != nil && *schema.MultipleOf != 0 {
			if typ == OasFieldTypeInteger {
				sample = *schema.MultipleOf + 1
			} else {
				sample = *schema.MultipleOf + *schema.MultipleOf*0.1
			}
			addRuleCase(paramVal.Name, sample, schema.Type, consts.AlternativeCaseRulesMultipleOf, parent)
		}

		if schema.ExclusiveMin {
			sample = *schema.Min
			addRuleCase(paramVal.Name, sample, schema.Type, consts.AlternativeCaseRulesExclusiveMin, parent)
		}

		if schema.ExclusiveMax {
			sample = *schema.Max
			addRuleCase(paramVal.Name, sample, schema.Type, consts.AlternativeCaseRulesExclusiveMax, parent)
		}

	} else {
		if schema.Pattern != "" {
			sample = RandStrSpecial()
			addRuleCase(paramVal.Name, sample, schema.Type, consts.AlternativeCaseRulesPattern, parent)
		}

		if schema.MaxLength != nil && *(schema.MaxLength) > 0 {
			sample = RandStrWithLen(int(*(schema.MaxLength) + 1))
			addRuleCase(paramVal.Name, sample, schema.Type, consts.AlternativeCaseRulesMaxLength, parent)
		}

		if schema.MinLength > 0 {
			sample = RandStrWithLen(int(schema.MinLength - 1))
			addRuleCase(paramVal.Name, sample, schema.Type, consts.AlternativeCaseRulesMinLength, parent)
		}
	}
}

func addRuleCase(name string, sample interface{}, typ string, rule consts.AlternativeCaseRules, parent *AlternativeCase) {
	ruleCase := &AlternativeCase{
		Title:  name,
		Sample: sample,

		Category:  consts.AlternativeCaseCase,
		Type:      consts.AlternativeCaseRule,
		Rule:      rule,
		FieldType: OasFieldType(typ),
		IsDir:     false,
	}

	parent.Children = append(parent.Children, ruleCase)
}

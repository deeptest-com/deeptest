package casesHelper

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	_stringUtils "github.com/aaronchen2k/deeptest/pkg/lib/string"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/kataras/iris/v12"
	"math"
)

func addPropCase(propName string, propVal *openapi3.Schema, requires []string, parent *AlternativeCase) {
	if propVal.Type == OasFieldTypeArray.String() {
		arrCase := &AlternativeCase{
			Title:    "数组",
			Category: consts.AlternativeCaseArray,
			IsDir:    true,
			Key:      _stringUtils.Uuid(),
			Slots:    iris.Map{"icon": "icon"},
		}

		addPropCase(propName, propVal.Items.Value, nil, arrCase)

		parent.Children = append(parent.Children, arrCase)

		return

	} else if propVal.Type == OasFieldTypeObject.String() {
		objCase := &AlternativeCase{
			Title:    "对象",
			Category: consts.AlternativeCaseObject,
			IsDir:    true,
			Key:      _stringUtils.Uuid(),
			Slots:    iris.Map{"icon": "icon"},
		}

		for propName, propRef := range propVal.Properties {
			addPropCase(propName, propRef.Value, propVal.Required, objCase)
		}

		parent.Children = append(parent.Children, objCase)

		return
	}

	addPropRequiredCase(propName, propVal, requires, parent)
	addPropTypeCase(propName, propVal, parent)
	addPropEnumCase(propName, propVal, parent)
	addPropFormatCase(propName, propVal, parent)
	addPropRuleCase(propName, propVal, parent)
}

func addPropRequiredCase(propName string, schemaVal *openapi3.Schema, requires []string, parent *AlternativeCase) {
	if !_stringUtils.StrInArr(propName, requires) {
		return
	}

	required := &AlternativeCase{
		Sample: ExampleEmpty,

		Category:      consts.AlternativeCaseCase,
		Type:          consts.AlternativeCaseRequired,
		FieldRequired: true,
		IsDir:         false,
		Key:           _stringUtils.Uuid(),
		Slots:         iris.Map{"icon": "icon"},
	}

	parent.Children = append(parent.Children, required)
}

func addPropTypeCase(name string, schema *openapi3.Schema, parent *AlternativeCase) {
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

	if sample == nil {
		return
	}

	typeCase := &AlternativeCase{
		Title:  fmt.Sprintf("%v", typ),
		Sample: sample,

		Category:  consts.AlternativeCaseCase,
		Type:      consts.AlternativeCaseTyped,
		FieldType: typ,
		IsDir:     false,
		Key:       _stringUtils.Uuid(),
		Slots:     iris.Map{"icon": "icon"},
	}

	parent.Children = append(parent.Children, typeCase)
}

func addPropEnumCase(name string, schema *openapi3.Schema, parent *AlternativeCase) {
	enum := schema.Enum

	if enum == nil {
		return
	}

	enumCase := &AlternativeCase{
		Title:  fmt.Sprintf("enum %v", enum),
		Sample: RandStr(),

		Category:  consts.AlternativeCaseCase,
		Type:      consts.AlternativeCaseEnum,
		FieldType: OasFieldType(schema.Type),
		IsDir:     false,
		Key:       _stringUtils.Uuid(),
		Slots:     iris.Map{"icon": "icon"},
	}

	parent.Children = append(parent.Children, enumCase)
}

func addPropFormatCase(name string, schema *openapi3.Schema, parent *AlternativeCase) {
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
	} else {
		sample = RandStr()
	}

	formatCase := &AlternativeCase{
		Title:  fmt.Sprintf("format (%s)", format),
		Sample: sample,

		Category:  consts.AlternativeCaseCase,
		Type:      consts.AlternativeCaseFormat,
		FieldType: OasFieldType(schema.Type),
		IsDir:     false,
		Key:       _stringUtils.Uuid(),
		Slots:     iris.Map{"icon": "icon"},
	}

	parent.Children = append(parent.Children, formatCase)
}

func addPropRuleCase(name string, schema *openapi3.Schema, parent *AlternativeCase) {
	typ := OasFieldType(schema.Type)

	var sample interface{}
	if typ == OasFieldTypeInteger || typ == OasFieldTypeNumber {
		if schema.Min != nil && *schema.Min != 0 {
			sample = *schema.Min - 1

			if schema.ExclusiveMin {
				sample = *schema.Min
			}

			addRuleCase(name, sample, typ, *schema.Min, consts.AlternativeCaseRulesMin, parent)
		}

		if schema.Max != nil && *schema.Max != 0 {
			sample = *schema.Max + 1

			if schema.ExclusiveMax {
				sample = *schema.Max
			}

			addRuleCase(name, sample, typ, *schema.Max, consts.AlternativeCaseRulesMax, parent)
		}

		if schema.MaxLength != nil && *schema.MaxLength > 0 {
			if typ == OasFieldTypeInteger {
				sample = 1 * math.Pow(10, float64(*schema.MaxLength))
			} else {
				sample = 1 / math.Pow(10, float64(*schema.MaxLength-1))
			}

			addRuleCase(name, sample, typ, *schema.MaxLength, consts.AlternativeCaseRulesMaxLength, parent)
		}

		if schema.MinLength > 0 {
			sample = 1
			addRuleCase(name, sample, typ, schema.MinLength, consts.AlternativeCaseRulesMinLength, parent)
		}

		if schema.MultipleOf != nil && *schema.MultipleOf != 0 {
			if typ == OasFieldTypeInteger {
				sample = *schema.MultipleOf + 1
			} else {
				sample = *schema.MultipleOf + *schema.MultipleOf*0.1
			}

			addRuleCase(name, sample, typ, *schema.MultipleOf, consts.AlternativeCaseRulesMultipleOf, parent)
		}

	} else if typ == OasFieldTypeByte {
		if schema.Min != nil && *schema.Min != 0 {
			sample = fmt.Sprintf("%c", rune(*schema.Min-1))
			addRuleCase(name, sample, typ, *schema.Min, consts.AlternativeCaseRulesMin, parent)
		}

		if schema.Max != nil && *schema.Max != 0 {
			sample = fmt.Sprintf("%c", rune(*schema.Max-1))
			addRuleCase(name, sample, typ, *schema.Max, consts.AlternativeCaseRulesMax, parent)
		}
	} else {
		if schema.Pattern != "" {
			sample = RandStrSpecial()
			addRuleCase(name, sample, typ, schema.Pattern, consts.AlternativeCaseRulesPattern, parent)
		}

		if schema.MaxLength != nil && *(schema.MaxLength) > 0 {
			sample = RandStrWithLen(int(*schema.MaxLength + 1))
			addRuleCase(name, sample, typ, *schema.MaxLength, consts.AlternativeCaseRulesMaxLength, parent)
		}

		if schema.MinLength > 0 {
			sample = RandStrWithLen(int(schema.MinLength - 1))
			addRuleCase(name, sample, typ, schema.MinLength, consts.AlternativeCaseRulesMinLength, parent)
		}
	}
}

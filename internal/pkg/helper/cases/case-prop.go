package casesHelper

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	_stringUtils "github.com/aaronchen2k/deeptest/pkg/lib/string"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/kataras/iris/v12"
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

	sample := getRequiredSample()

	required := &AlternativeCase{
		Title:  fmt.Sprintf("required"),
		Sample: sample,

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

	sample := getTypeSample(typ)
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

	sample := getEnumSample()

	enumCase := &AlternativeCase{
		Title:  fmt.Sprintf("enum %v", enum),
		Sample: sample,

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

	sample := getFormatSample(format, typ)

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
	arr := getRuleSamples(schema, name)

	for _, item := range arr {
		name := item[0].(string)
		sample := item[1]
		typ := item[2].(OasFieldType)
		tag := item[3]
		rule := item[4].(consts.AlternativeCaseRules)

		addRuleCase(name, sample, typ, tag, rule, parent)
	}
}

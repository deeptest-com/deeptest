package casesHelper

import (
	"fmt"
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	_stringUtils "github.com/deeptest-com/deeptest/pkg/lib/string"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/kataras/iris/v12"
	"path"
)

func addPropCase(propName string, propVal *openapi3.Schema, requires []string, parent *AlternativeCase, pth string,
	doc3 *openapi3.T) {
	if propVal.Type == OasFieldTypeArray.String() {
		arrCase := &AlternativeCase{
			Title:    "数组",
			Path:     path.Join(pth, AddFix("arr")),
			Category: consts.AlternativeCaseArray,
			IsDir:    true,
			Key:      _stringUtils.Uuid(),
			Slots:    iris.Map{"icon": "icon"},
		}

		itemsValue := propVal.Items.Value
		if itemsValue == nil {
			arrCase.Title = fmt.Sprintf("数组[%s]", propVal.Items.Ref)
			itemsValue = getRef(propVal.Items.Ref, doc3)
		}

		addPropCase(propName, itemsValue, nil, arrCase, arrCase.Path, doc3)

		parent.Children = append(parent.Children, arrCase)

		return

	} else if propVal.Type == OasFieldTypeObject.String() {
		for propName, propRef := range propVal.Properties {
			propCase := &AlternativeCase{
				Title:    propName,
				Path:     path.Join(pth, propName),
				Category: consts.AlternativeCaseProp,
				IsDir:    true,
				Key:      _stringUtils.Uuid(),
				Slots:    iris.Map{"icon": "icon"},
			}

			propVal := propRef.Value
			if propVal == nil {
				propVal = getRef(propRef.Ref, doc3)
			}

			addPropCase(propName, propVal, propVal.Required, propCase, propCase.Path, doc3)

			parent.Children = append(parent.Children, propCase)
		}

		return
	}

	addPropRequiredCase(propName, propVal, requires, parent, pth)
	addPropTypeCase(propName, propVal, parent, pth)
	addPropEnumCase(propName, propVal, parent, pth)
	addPropFormatCase(propName, propVal, parent, pth)
	addPropRuleCase(propName, propVal, parent, pth)
}

func addPropRequiredCase(propName string, schemaVal *openapi3.Schema, requires []string, parent *AlternativeCase, pth string) {
	if !_stringUtils.StrInArr(propName, requires) {
		return
	}

	sample := getRequiredSample()

	required := &AlternativeCase{
		Title:  fmt.Sprintf("required"),
		Sample: sample,
		Path:   path.Join(pth, AddFix("required")),

		Category:      consts.AlternativeCaseCase,
		Type:          consts.AlternativeCaseRequired,
		FieldRequired: true,
		IsDir:         false,
		Key:           _stringUtils.Uuid(),
		Slots:         iris.Map{"icon": "icon"},
	}

	parent.Children = append(parent.Children, required)
}

func addPropTypeCase(name string, schema *openapi3.Schema, parent *AlternativeCase, pth string) {
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
		Path:   path.Join(pth, AddFix("type")),

		Category:  consts.AlternativeCaseCase,
		Type:      consts.AlternativeCaseTyped,
		FieldType: typ,
		IsDir:     false,
		Key:       _stringUtils.Uuid(),
		Slots:     iris.Map{"icon": "icon"},
	}

	parent.Children = append(parent.Children, typeCase)
}

func addPropEnumCase(name string, schema *openapi3.Schema, parent *AlternativeCase, pth string) {
	enum := schema.Enum

	if enum == nil || len(enum) == 0 {
		return
	}

	sample := getEnumSample()

	enumCase := &AlternativeCase{
		Title:  fmt.Sprintf("enum %v", enum),
		Sample: sample,
		Path:   path.Join(pth, AddFix("enum")),

		Category:  consts.AlternativeCaseCase,
		Type:      consts.AlternativeCaseEnum,
		FieldType: OasFieldType(schema.Type),
		IsDir:     false,
		Key:       _stringUtils.Uuid(),
		Slots:     iris.Map{"icon": "icon"},
	}

	parent.Children = append(parent.Children, enumCase)
}

func addPropFormatCase(name string, schema *openapi3.Schema, parent *AlternativeCase, pth string) {
	typ := OasFieldType(schema.Type)
	format := OasFieldFormat(schema.Format)

	if format == "" {
		return
	}

	sample, ok := getFormatSample(format, typ)
	if !ok {
		return
	}

	formatCase := &AlternativeCase{
		Title:  fmt.Sprintf("format (%s)", format),
		Sample: sample,
		Path:   path.Join(pth, AddFix("format")),

		Category:  consts.AlternativeCaseCase,
		Type:      consts.AlternativeCaseFormat,
		FieldType: OasFieldType(schema.Type),
		IsDir:     false,
		Key:       _stringUtils.Uuid(),
		Slots:     iris.Map{"icon": "icon"},
	}

	parent.Children = append(parent.Children, formatCase)
}

func addPropRuleCase(name string, schema *openapi3.Schema, parent *AlternativeCase, pth string) {
	arr := getRuleSamples(schema, name)

	for _, item := range arr {
		name := item[0].(string)
		sample := item[1]
		typ := item[2].(OasFieldType)
		tag := item[3]
		rule := item[4].(consts.AlternativeCaseRules)

		addRuleCase(name, sample, typ, tag, rule, parent, path.Join(pth, AddFix("rule")))
	}
}

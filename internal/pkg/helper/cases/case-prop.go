package casesHelper

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	_stringUtils "github.com/aaronchen2k/deeptest/pkg/lib/string"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/kataras/iris/v12"
	"path"
)

func addPropCase(propName string, propVal *openapi3.Schema, requires []string, parent *AlternativeCase, pth string) {
	if propVal.Type == OasFieldTypeArray.String() {
		arrCase := &AlternativeCase{
			Title:    "数组",
			Path:     path.Join(pth, "arr"),
			Category: consts.AlternativeCaseArray,
			IsDir:    true,
			Key:      _stringUtils.Uuid(),
			Slots:    iris.Map{"icon": "icon"},
		}

		addPropCase(propName, propVal.Items.Value, nil, arrCase, arrCase.Path)

		parent.Children = append(parent.Children, arrCase)

		return

	} else if propVal.Type == OasFieldTypeObject.String() {
		objCase := &AlternativeCase{
			Title:    "对象",
			Path:     path.Join(pth, "object"),
			Category: consts.AlternativeCaseObject,
			IsDir:    true,
			Key:      _stringUtils.Uuid(),
			Slots:    iris.Map{"icon": "icon"},
		}

		for propName, propRef := range propVal.Properties {
			temp := path.Join(objCase.Path, propName)
			addPropCase(propName, propRef.Value, propVal.Required, objCase, temp)
		}

		parent.Children = append(parent.Children, objCase)

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
		Path:   path.Join(pth, "required"),

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
		Path:   path.Join(pth, "type"),

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

	if enum == nil {
		return
	}

	sample := getEnumSample()

	enumCase := &AlternativeCase{
		Title:  fmt.Sprintf("enum %v", enum),
		Sample: sample,
		Path:   path.Join(pth, "enum"),

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

	sample := getFormatSample(format, typ)

	formatCase := &AlternativeCase{
		Title:  fmt.Sprintf("format (%s)", format),
		Sample: sample,
		Path:   path.Join(pth, "format"),

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

		addRuleCase(name, sample, typ, tag, rule, parent, path.Join(pth, "rule"))
	}
}

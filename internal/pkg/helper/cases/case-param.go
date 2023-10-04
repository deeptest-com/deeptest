package casesHelper

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	_stringUtils "github.com/aaronchen2k/deeptest/pkg/lib/string"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/kataras/iris/v12"
	"path"
)

func addParamRequiredCase(paramVal *openapi3.Parameter, parent *AlternativeCase) {
	if !paramVal.Required {
		return
	}

	sample := getRequiredSample()
	required := &AlternativeCase{
		Title:  "required",
		Sample: sample,
		Path:   path.Join(parent.Path, "required"),

		Category:      consts.AlternativeCaseCase,
		Type:          consts.AlternativeCaseRequired,
		FieldRequired: true,
		IsDir:         false,
		Key:           _stringUtils.Uuid(),
		Slots:         iris.Map{"icon": "icon"},
	}

	parent.Children = append(parent.Children, required)
}

func addParamTypeCase(paramVal *openapi3.Parameter, parent *AlternativeCase) {
	schema := paramVal.Schema.Value
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
		Path:   path.Join(parent.Path, "type"),

		Category:  consts.AlternativeCaseCase,
		Type:      consts.AlternativeCaseTyped,
		FieldType: typ,
		IsDir:     false,
		Key:       _stringUtils.Uuid(),
		Slots:     iris.Map{"icon": "icon"},
	}

	parent.Children = append(parent.Children, typeCase)
}

func addParamEnumCase(paramVal *openapi3.Parameter, parent *AlternativeCase) {
	schema := paramVal.Schema.Value
	enum := schema.Enum

	if enum == nil {
		return
	}

	sample := getEnumSample()
	typeCase := &AlternativeCase{
		Title:  fmt.Sprintf("enum %v", enum),
		Sample: sample,
		Path:   path.Join(parent.Path, "enum"),

		Category:  consts.AlternativeCaseCase,
		Type:      consts.AlternativeCaseEnum,
		FieldType: OasFieldType(schema.Type),
		IsDir:     false,
		Key:       _stringUtils.Uuid(),
		Slots:     iris.Map{"icon": "icon"},
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

	sample := getFormatSample(format, typ)

	formatCase := &AlternativeCase{
		Title:  fmt.Sprintf("format (%s)", format),
		Sample: sample,
		Path:   path.Join(parent.Path, "format"),

		Category:  consts.AlternativeCaseCase,
		Type:      consts.AlternativeCaseEnum,
		FieldType: OasFieldType(schema.Type),
		IsDir:     false,
		Key:       _stringUtils.Uuid(),
		Slots:     iris.Map{"icon": "icon"},
	}

	parent.Children = append(parent.Children, formatCase)
}

func addParamRuleCase(paramVal *openapi3.Parameter, parent *AlternativeCase) {
	schema := paramVal.Schema.Value

	arr := getRuleSamples(schema, paramVal.Name)

	for _, item := range arr {
		name := item[0].(string)
		sample := item[1]
		typ := item[2].(OasFieldType)
		tag := item[3]
		rule := item[4].(consts.AlternativeCaseRules)

		temp := path.Join("param", paramVal.In, name, "rule")
		addRuleCase(name, sample, typ, tag, rule, parent, temp)
	}
}

func addRuleCase(name string, sample interface{}, typ OasFieldType, tag interface{},
	rule consts.AlternativeCaseRules, parent *AlternativeCase, pth string) {

	ruleCase := &AlternativeCase{
		Title:  fmt.Sprintf("%s (%v)", rule.String(), tag),
		Sample: sample,
		Path:   path.Join(pth, rule.String()),

		Category:  consts.AlternativeCaseCase,
		Type:      consts.AlternativeCaseRule,
		Rule:      rule,
		FieldType: typ,
		IsDir:     false,
		Key:       _stringUtils.Uuid(),
		Slots:     iris.Map{"icon": "icon"},
	}

	parent.Children = append(parent.Children, ruleCase)
}

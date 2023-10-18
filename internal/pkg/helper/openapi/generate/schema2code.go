package generate

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/helper/openapi/generate/fields"
	"github.com/aaronchen2k/deeptest/internal/pkg/helper/openapi/generate/template"
	schemaHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/schema"
	commonUtils "github.com/aaronchen2k/deeptest/pkg/lib/comm"
	"github.com/getkin/kin-openapi/openapi3"
	"strings"
)

type Schema2Code struct {
	schemaHelper.Schema2conv
	langType template.LangType
	nameRule template.NameRule
	sets     map[string]int64
	varCount int
}

func NewSchema2Code(langType template.LangType, nameRule template.NameRule) *Schema2Code {
	obj := &Schema2Code{}
	obj.langType = langType
	obj.nameRule = nameRule
	obj.sets = make(map[string]int64)
	return obj
}

func (s *Schema2Code) schema2Fields(name string, schema schemaHelper.SchemaRef) *fields.Field {
	ref := schema.Ref
	refName := ""
	if component, ok := s.Components[schema.Ref]; ok {
		s.sets[ref]++
		schema = *component
		refName = s.getRefName(ref)
		name = refName
	}

	s.CombineSchemas(&schema)

	field := &fields.Field{FieldName: fields.FieldName(name), FieldRefName: fields.FieldName(refName), Description: schema.Value.Description}

	switch schema.Value.Type {
	case openapi3.TypeObject:
		if s.sets[ref] > 1 {
			return field
		}

		field.FieldType = openapi3.TypeObject
		if field.FieldName == "" {
			field.FieldName = fields.FieldName(s.getVarName("var"))
		}

		for key, property := range schema.Value.Properties {
			subField := s.schema2Fields(s.format(key), *property)
			subField.IsProperty = true
			field.Properties = append(field.Properties, subField)
		}
	case openapi3.TypeArray:
		if s.sets[ref] > 1 {
			return field
		}
		field.FieldType = openapi3.TypeArray
		field.SubField = s.schema2Fields("", *schema.Value.Items)
	default:
		//field.FieldName = ""
		field.FieldRefName = "" //非数组和对象类型取本身类型
		field.FieldType = fields.FieldType(schema.Value.Type)

	}
	return field
}

func (s *Schema2Code) Convert(schema schemaHelper.SchemaRef) string {
	field := s.schema2Fields(s.format("response"), schema)
	fmt.Println(commonUtils.JsonEncode(field))
	t := template.NewTemplate(s.langType, field.ToArray())
	return t.CreateCode()
}

func (s *Schema2Code) getVarName(name string) string {
	s.varCount++
	ret := fmt.Sprintf("%s%d", name, s.varCount)
	return s.format(ret)
}

func (s *Schema2Code) getRefName(ref string) string {
	refArr := strings.Split(ref, "/")
	refEnd := refArr[len(refArr)-1]
	name := strings.Split(refEnd, ".")
	ret := name[len(name)-1]
	return s.format(ret)
}

func (s *Schema2Code) format(ret string) string {
	ret = commonUtils.Case2Camel(ret)
	if s.nameRule == template.UpperCase {
		ret = strings.ToUpper(ret[:1]) + ret[1:]
	} else if s.nameRule == template.LowerCase {
		ret = strings.ToLower(ret[:1]) + ret[1:]
	} else {
		ret = commonUtils.Camel2Case(ret)
	}
	return ret
}

package generate

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/helper/openapi/generate/fields"
	"github.com/aaronchen2k/deeptest/internal/pkg/helper/openapi/generate/template"
	schemaHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/schema"
	commonUtils "github.com/aaronchen2k/deeptest/pkg/lib/comm"
	"github.com/getkin/kin-openapi/openapi3"
)

type Schema2Code struct {
	schemaHelper.Schema2conv
	langType template.LangType
	nameRule string
	sets     map[string]int64
	varCount int
}

func NewSchema2Code(langType template.LangType, nameRule string) *Schema2Code {
	obj := &Schema2Code{}
	obj.langType = langType
	obj.nameRule = nameRule
	obj.sets = make(map[string]int64)
	return obj
}

func (s *Schema2Code) schema2Fields(name string, schema schemaHelper.SchemaRef) *fields.Field {
	ref := schema.Ref
	if component, ok := s.Components[schema.Ref]; ok {
		s.sets[ref]++
		//name = ref
		schema = *component
	}

	s.CombineSchemas(&schema)

	field := &fields.Field{FieldName: fields.FieldName(name)}
	switch schema.Value.Type {
	case openapi3.TypeObject:
		if s.sets[ref] > 1 {
			return field
		}
		field.FieldType = openapi3.TypeObject
		for key, property := range schema.Value.Properties {
			subField := s.schema2Fields(key, *property)
			subField.IsProperty = true
			if subField.FieldType != "" {
				field.Properties = append(field.Properties, subField)
			}
		}
	case openapi3.TypeArray:
		if s.sets[ref] > 1 {
			return field
		}
		field.FieldType = openapi3.TypeArray
		field.SubField = s.schema2Fields(s.varName("var"), *schema.Value.Items)
	default:
		field.FieldType = fields.FieldType(schema.Value.Type)

	}
	return field
}

func (s *Schema2Code) Convert(schema schemaHelper.SchemaRef) string {
	field := s.schema2Fields("response", schema)
	fmt.Println(commonUtils.JsonEncode(field))
	t := template.NewTemplate(s.langType, field.ToArray())
	return t.CreateCode()
}

func (s *Schema2Code) varName(name string) string {
	s.varCount++
	return fmt.Sprintf("%s%d", name, s.varCount)
}

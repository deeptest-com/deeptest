package schemaHelper

import (
	"encoding/json"
	serverDomain "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	mockjsHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/mockjs"
	mockData "github.com/aaronchen2k/deeptest/internal/pkg/helper/openapi-mock/openapi/generator/data"
	"github.com/getkin/kin-openapi/openapi3"
	"math/rand"
	"reflect"
	"time"
)

type SchemaRef struct {
	Ref    string  `json:"ref,omitempty" yaml:"ref,omitempty"`
	RefExt string  `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	Value  *Schema `json:"value,omitempty" yaml:"value,omitempty"`
}
type SchemaRefs []*SchemaRef
type Schemas map[string]*SchemaRef

type Schema struct {
	openapi3.ExtensionProps
	Type        string     `json:"type,omitempty" yaml:"type,omitempty"`
	XMockType   string     `json:"x-mock-type,omitempty" yaml:"x-mock-type,omitempty"`
	Items       *SchemaRef `json:"items,omitempty" yaml:"items,omitempty"`
	Properties  Schemas    `json:"properties,omitempty" yaml:"properties,omitempty"`
	AllOf       SchemaRefs `json:"allOf,omitempty" yaml:"allOf,omitempty"`
	OneOf       SchemaRefs `json:"oneOf,omitempty" yaml:"allOf,omitempty"`
	AnyOf       SchemaRefs `json:"anyOf,omitempty" yaml:"allOf,omitempty"`
	Ref         string     `json:"ref,omitempty" yaml:"ref,omitempty"`
	RefExt      string     `json:"$ref,omitempty" yaml:"ref,omitempty"`
	Description string     `json:"description,omitempty" yaml:"description,omitempty"`
}

func (schemaRef *SchemaRef) MarshalJSON() (res []byte, err error) {
	schema := Schema{}
	if schemaRef.Ref == "" {
		schemaRef.Ref = schemaRef.RefExt
		//schemaRef.RefExt = ""
	}
	if schemaRef.Ref != "" {
		schema.Ref = schemaRef.Ref
	} else {
		if schemaRef.Value != nil {
			schema = *schemaRef.Value
		}
	}

	res, err = json.Marshal(schema)
	if err != nil {
		return
	}

	return
}

func (schemaRef *SchemaRef) UnmarshalJSON(data []byte) error {
	var schema Schema

	err := json.Unmarshal(data, &schema)

	if err != nil {
		return err
	}

	if schema.Ref == "" {
		schema.Ref = schema.RefExt
	}
	schemaRef.Ref = schema.Ref
	if schemaRef.Ref == "" {
		schemaRef.Value = &schema
	}

	return nil
}

type Components map[string]*SchemaRef

type Schema2conv struct {
	Components Components
	sets       map[string]int64
	generator  mockData.MockjsGenerator
}

func NewSchema2conv() *Schema2conv {
	obj := new(Schema2conv)
	obj.sets = map[string]int64{}
	obj.generator = mockData.MockjsGenerator{}
	return obj
}

func (s *Schema2conv) Example2Schema(object interface{}, schema *Schema) (err error) {
	V := reflect.ValueOf(object)

	switch V.Kind() {
	case reflect.Map:
		schema.Type = openapi3.TypeObject
		schema.Properties = Schemas{}
		iter := V.MapRange()
		for iter.Next() {
			key := iter.Key().String()
			schema.Properties[key] = new(SchemaRef)
			schema.Properties[key].Value = new(Schema)
			s.Example2Schema(iter.Value().Interface(), schema.Properties[key].Value)
		}
	case reflect.Slice:
		schema.Type = openapi3.TypeArray
		schema.Items = new(SchemaRef)
		schema.Items.Value = new(Schema)
		if V.Len() != 0 {
			s.Example2Schema(V.Index(0).Interface(), schema.Items.Value)
		}
	case reflect.String:
		schema.Type = openapi3.TypeString
	case reflect.Int64:
		schema.Type = openapi3.TypeInteger
	case reflect.Bool:
		schema.Type = openapi3.TypeBoolean
	case reflect.Float64:
		schema.Type = openapi3.TypeNumber
	default:
		schema.Type = openapi3.TypeObject
	}
	return
}

func (s *Schema2conv) Schema2Example(schema SchemaRef) (object interface{}) {
	ref := schema.Ref
	if component, ok := s.Components[schema.Ref]; ok {
		s.sets[ref]++
		schema = *component
	}

	s.CombineSchemas(&schema)

	switch schema.Value.Type {
	case openapi3.TypeObject:
		object = map[string]interface{}{}
		if s.sets[ref] > 1 {
			return
		}
		for key, property := range schema.Value.Properties {
			object.(map[string]interface{})[key] = s.Schema2Example(*property)
		}
	case openapi3.TypeArray:
		object = make([]interface{}, 1)
		if s.sets[ref] > 1 {
			return
		}
		object.([]interface{})[0] = s.Schema2Example(*schema.Value.Items)
	case openapi3.TypeString:
		if schema.Value.XMockType == "" {
			schema.Value.XMockType = "@word()"
		}
		object = s.mock(schema.Value.XMockType, schema.Value.Type)
	case openapi3.TypeBoolean:
		object = true
	case openapi3.TypeInteger:
		if schema.Value.XMockType == "" {
			schema.Value.XMockType = "@integer(1,100)"
		}
		object = s.mock(schema.Value.XMockType, schema.Value.Type)
	case openapi3.TypeNumber:
		schema.Value.XMockType = "@float(1, 10, 2, 5)"
		object = s.mock(schema.Value.XMockType, schema.Value.Type)

	}
	return
}

func (s *Schema2conv) mock(expr string, typ string) interface{} {
	req := serverDomain.MockJsExpression{
		Expression: expr,
	}

	ret, err := mockjsHelper.EvaluateExpression(req)
	if err != nil {
		return nil
	}

	return mockjsHelper.ConvertData(ret.Result, typ)
}

func (s *Schema2conv) CombineSchemas(schema *SchemaRef) {
	var combineSchemas SchemaRefs

	if len(schema.Value.AllOf) >= 1 {
		combineSchemas = schema.Value.AllOf
		//	fmt.Println(combineSchemas)
	} else {
		if len(schema.Value.AnyOf) >= 1 {
			rand.Seed(time.Now().UnixNano())
			n := rand.Intn(len(schema.Value.AnyOf)-1) + 1
			combineSchemas = s.anyMoreSchemas(schema.Value.AnyOf, n)

		} else if len(schema.Value.OneOf) >= 1 {
			combineSchemas = s.anyMoreSchemas(schema.Value.OneOf, 1)

		}
	}

	for _, item := range combineSchemas {

		if item.Ref != "" {
			if component, ok := s.Components[item.Ref]; ok {
				item = component
			}
		}
		schema.Value.Type = item.Value.Type

		if item.Value.Type != openapi3.TypeObject {
			item.Value.Properties = Schemas{}
			schema.Value = item.Value
			continue
		}

		for key, property := range item.Value.Properties {
			if property.Value == nil {
				continue
			}
			if property.Value.Type == "" {
				//continue
			}

			s.CombineSchemas(property)

			if schema.Value.Properties == nil {
				schema.Value.Properties = Schemas{}
			}
			schema.Value.Properties[key] = property

		}

	}

	schema.Value.AllOf = nil
	schema.Value.AnyOf = nil
	schema.Value.OneOf = nil

}

func (s *Schema2conv) anyMoreSchemas(schemas SchemaRefs, n int) (combineSchemas SchemaRefs) {
	if n > len(schemas) {
		return
	}

	for ; n > 0; n-- {
		rand.Seed(time.Now().UnixNano())
		index := rand.Intn(len(schemas))
		combineSchemas = append(combineSchemas, schemas[index])
		schemas = append(schemas[:index], schemas[index+1:]...)
	}

	return combineSchemas
}

func (s *Schema2conv) AssertDataForSchema(schema *SchemaRef, data interface{}) bool {
	dataSchema := new(Schema)
	err := s.Example2Schema(data, dataSchema)
	if err != nil {
		return false
	}
	schema1 := new(SchemaRef)
	schema1.Value = dataSchema
	return s.Equal(schema, schema1)
}

func (s *Schema2conv) Equal(schema1, schema2 *SchemaRef) (ret bool) {

	if component, ok := s.Components[schema1.Ref]; ok {
		//s.sets[ref1]++
		schema1 = component
	}

	s.CombineSchemas(schema1)

	//类型object 并且没有属性的，是空对象，匹配任何类型
	if schema2.Value.Type == openapi3.TypeObject && len(schema2.Value.Properties) == 0 {
		return true
	}

	//golang 会统一把数字类型反射成float64，默认TypeNumber，所以只要是TypeNumber，TypeInteger，就认为类型是相同的
	if (schema1.Value.Type == openapi3.TypeNumber || schema1.Value.Type == openapi3.TypeInteger) &&
		(schema2.Value.Type == openapi3.TypeNumber || schema2.Value.Type == openapi3.TypeInteger) {
		return true
	}

	if schema1.Value.Type != schema2.Value.Type {
		return false
	}

	switch schema1.Value.Type {
	case openapi3.TypeObject:

		return s.objectEqual(schema1.Value, schema2.Value)

	case openapi3.TypeArray:

		return s.arrayEqual(schema1.Value, schema2.Value)
	}

	return true

}

func (s *Schema2conv) objectEqual(schema1 *Schema, schema2 *Schema) (ret bool) {
	if len(schema1.Properties) != len(schema2.Properties) {
		return false
	}
	for key, property := range schema1.Properties {
		if item, ok := schema2.Properties[key]; ok {
			if !s.Equal(property, item) {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

func (s *Schema2conv) arrayEqual(schema1 *Schema, schema2 *Schema) (ret bool) {
	return s.Equal(schema1.Items, schema2.Items)

}

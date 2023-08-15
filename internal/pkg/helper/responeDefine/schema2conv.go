package responseDefineHelpper

import (
	"encoding/json"
	"github.com/getkin/kin-openapi/openapi3"
	"math/rand"
	"reflect"
	"time"
)

type SchemaRef struct {
	Ref   string
	Value *Schema
}
type SchemaRefs []*SchemaRef
type Schemas map[string]*SchemaRef

type Schema struct {
	openapi3.ExtensionProps
	Type        string     `json:"type,omitempty" yaml:"type,omitempty"`
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
	schema = *schemaRef.Value
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
	schemaRef.Value = &schema
	return nil
}

type Components map[string]SchemaRef

type schema2conv struct {
	Components Components
	sets       map[string]int64
}

func NewSchema2conv() *schema2conv {
	obj := new(schema2conv)
	obj.sets = map[string]int64{}
	return obj
}

func (s *schema2conv) Example2Schema(object interface{}, schema *Schema) (err error) {
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

func (s *schema2conv) Schema2Example(schema SchemaRef) (object interface{}) {
	ref := schema.Ref
	if component, ok := s.Components[schema.Ref]; ok {
		s.sets[ref]++
		schema = component
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
		object = "string"
	case openapi3.TypeBoolean:
		object = true
	case openapi3.TypeInteger:
		object = 0
	case openapi3.TypeNumber:
		object = 0.0
	}
	return
}

func (s *schema2conv) CombineSchemas(schema *SchemaRef) {
	var combineSchemas SchemaRefs

	if len(schema.Value.AllOf) >= 1 {
		combineSchemas = schema.Value.AllOf
		schema.Value.AllOf = nil
	} else {
		if len(schema.Value.AnyOf) >= 1 {
			rand.Seed(time.Now().UnixNano())
			n := rand.Intn(len(schema.Value.AnyOf)-1) + 1
			combineSchemas = s.anyMoreSchemas(schema.Value.AnyOf, n)
			schema.Value.AnyOf = nil
		} else if len(schema.Value.OneOf) >= 1 {
			combineSchemas = s.anyMoreSchemas(schema.Value.OneOf, 1)
			schema.Value.OneOf = nil
		}
	}

	for _, item := range combineSchemas {

		if item.Ref != "" {
			if component, ok := s.Components[item.Ref]; ok {
				item = &component
			}
		}
		schema.Value.Type = item.Value.Type

		if item.Value.Type != openapi3.TypeObject {
			item.Value.Properties = Schemas{}
			schema.Value = item.Value
			continue
		}

		for key, property := range item.Value.Properties {
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

}

func (s *schema2conv) anyMoreSchemas(schemas SchemaRefs, n int) (combineSchemas SchemaRefs) {
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

func (s *schema2conv) AssertDataForSchema(schema *SchemaRef, data interface{}) bool {
	dataSchema := new(Schema)
	err := s.Example2Schema(data, dataSchema)
	if err != nil {
		return false
	}
	schema1 := new(SchemaRef)
	schema1.Value = dataSchema
	return s.Equal(schema, schema1)
}

func (s *schema2conv) Equal(schema1, schema2 *SchemaRef) (ret bool) {
	ref1 := schema1.Ref
	if component, ok := s.Components[schema1.Ref]; ok {
		s.sets[ref1]++
		schema1 = &component
	}

	ref2 := schema2.Ref
	if component, ok := s.Components[schema2.Ref]; ok {
		s.sets[ref2]++
		schema2 = &component
	}

	s.CombineSchemas(schema1)
	s.CombineSchemas(schema2)

	if schema1.Value.Type != schema2.Value.Type {
		return false
	}

	switch schema1.Value.Type {
	case openapi3.TypeObject:
		if s.sets[ref1] > 2 {
			return true
		}
		return s.objectEqual(schema1.Value, schema2.Value)

	case openapi3.TypeArray:
		if s.sets[ref1] > 2 {
			return true
		}
		return s.arrayEqual(schema1.Value, schema2.Value)
	}

	return
}

func (s *schema2conv) objectEqual(schema1 *Schema, schema2 *Schema) (ret bool) {
	for key, property := range schema1.Properties {
		if item, ok := schema2.Properties[key]; ok {
			return s.Equal(property, item)
		} else {
			return false
		}
	}
	return true
}

func (s *schema2conv) arrayEqual(schema1 *Schema, schema2 *Schema) (ret bool) {
	return s.Equal(schema1.Items, schema2.Items)

}

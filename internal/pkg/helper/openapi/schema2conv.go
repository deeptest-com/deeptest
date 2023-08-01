package openapi

import (
	"encoding/json"
	"github.com/getkin/kin-openapi/openapi3"
	"reflect"
)

type SchemaRef struct {
	Ref   string
	Value *Schema
}
type SchemaRefs []*SchemaRef
type Schemas map[string]*SchemaRef

type Schema struct {
	openapi3.ExtensionProps
	Type       string     `json:"type,omitempty" yaml:"type,omitempty"`
	Items      *SchemaRef `json:"items,omitempty" yaml:"items,omitempty"`
	Properties Schemas    `json:"properties,omitempty" yaml:"properties,omitempty"`
	AllOf      SchemaRefs `json:"allOf,omitempty" yaml:"allOf,omitempty"`
	Ref        string     `json:"ref,omitempty" yaml:"ref,omitempty"`
	RefExt     string     `json:"$ref,omitempty" yaml:"ref,omitempty"`
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

	s.AllOfConv(&schema)

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

func (s *schema2conv) AllOfConv(schema *SchemaRef) {
	if len(schema.Value.AllOf) > 1 {

		//schema = SchemaRef{}
		//schema.Value = new(Schema)
		schema.Value.Type = openapi3.TypeObject
		schema.Value.Properties = Schemas{}
		for _, item := range schema.Value.AllOf {
			if item.Ref != "" {
				if component, ok := s.Components[item.Ref]; ok {
					item = &component
				}
			}

			if item.Value.Type != openapi3.TypeObject {
				continue
			}

			for key, property := range item.Value.Properties {
				if property.Value.Type == "" {
					//continue
				}

				s.AllOfConv(property)

				schema.Value.Properties[key] = property

			}

		}
		schema.Value.AllOf = nil
	}
}

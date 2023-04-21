package openapi

import (
	"encoding/json"
	"fmt"
	"github.com/getkin/kin-openapi/openapi3"
	"reflect"
)

type SchemaRef struct {
	Ref   string
	Value *Schema
}

type Schemas map[string]*SchemaRef

type Schema struct {
	openapi3.ExtensionProps
	Type       string     `json:"type,omitempty" yaml:"type,omitempty"`
	Items      *SchemaRef `json:"items,omitempty" yaml:"items,omitempty"`
	Properties Schemas    `json:"properties,omitempty" yaml:"properties,omitempty"`
	Ref        string     `json:"ref,omitempty" yaml:"ref,omitempty"`
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
	schemaRef.Value = &schema
	return nil
}

type Components map[string]Schema

type schema2conv struct {
	Components Components
	sets       map[string]int64
}

func NewSchema2conv() *schema2conv {
	obj := new(schema2conv)
	obj.sets = map[string]int64{}
	return obj
}

func (s *schema2conv) Example2Schema(object interface{}, schema *openapi3.Schema) (err error) {
	V := reflect.ValueOf(object)
	switch V.Kind() {
	case reflect.Map:
		schema.Type = openapi3.TypeObject
		schema.Properties = openapi3.Schemas{}
		iter := V.MapRange()
		for iter.Next() {
			key := iter.Key().String()
			schema.Properties[key] = new(openapi3.SchemaRef)
			schema.Properties[key].Value = new(openapi3.Schema)
			s.Example2Schema(iter.Value().Interface(), schema.Properties[key].Value)
		}
	case reflect.Slice:
		schema.Type = openapi3.TypeArray
		schema.Items = new(openapi3.SchemaRef)
		schema.Items.Value = new(openapi3.Schema)
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
	}
	return
}

func (s *schema2conv) Schema2Example(schema Schema) (object interface{}) {
	if ref, ok := s.Components[schema.Ref]; ok {
		schema = ref
		s.sets[schema.Ref]++
	}
	fmt.Println(schema, "+++++=")
	switch schema.Type {
	case openapi3.TypeObject:
		object = map[string]interface{}{}
		if s.sets[schema.Ref] > 1 {
			return
		}
		for key, property := range schema.Properties {
			object.(map[string]interface{})[key] = s.Schema2Example(*property.Value)
		}
	case openapi3.TypeArray:
		object = make([]interface{}, 1)
		if s.sets[schema.Ref] > 1 {
			return
		}
		object.([]interface{})[0] = s.Schema2Example(*schema.Items.Value)
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

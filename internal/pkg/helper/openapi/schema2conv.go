package openapi

import (
	"github.com/getkin/kin-openapi/openapi3"
	"reflect"
)

type schema2conv struct {
}

func NewSchema2conv() *schema2conv {
	return new(schema2conv)
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
	//fmt.Println(V.Kind(), "++++++++", schema)
	return
}

func (s *schema2conv) Schema2Example(schema openapi3.Schema) (object interface{}) {
	switch schema.Type {
	case openapi3.TypeObject:
		object = map[string]interface{}{}
		for key, property := range schema.Properties {
			object.(map[string]interface{})[key] = s.Schema2Example(*property.Value)
		}

	case openapi3.TypeArray:
		object = make([]interface{}, 1)
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

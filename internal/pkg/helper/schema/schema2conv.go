package schemaHelper

import (
	"encoding/json"
	"fmt"
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
	RefId  uint    `json:"refId,omitempty" yaml:"$refId,omitempty"`
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
	RefId       uint       `json:"refId,omitempty" yaml:"$refId,omitempty"`
	Description string     `json:"description,omitempty" yaml:"description,omitempty"`
	Format      string     `json:"format,omitempty" yaml:"format,omitempty"`

	Enum    []interface{} `json:"enum,omitempty" yaml:"enum,omitempty"`
	Default interface{}   `json:"default,omitempty" yaml:"default,omitempty"`
	Example interface{}   `json:"example,omitempty" yaml:"example,omitempty"`

	// Array-related, here for struct compactness
	UniqueItems bool `json:"uniqueItems,omitempty" yaml:"uniqueItems,omitempty"`
	// Number-related, here for struct compactness
	ExclusiveMin bool `json:"exclusiveMinimum,omitempty" yaml:"exclusiveMinimum,omitempty"`
	ExclusiveMax bool `json:"exclusiveMaximum,omitempty" yaml:"exclusiveMaximum,omitempty"`
	// Properties
	Nullable        bool `json:"nullable,omitempty" yaml:"nullable,omitempty"`
	ReadOnly        bool `json:"readOnly,omitempty" yaml:"readOnly,omitempty"`
	WriteOnly       bool `json:"writeOnly,omitempty" yaml:"writeOnly,omitempty"`
	AllowEmptyValue bool `json:"allowEmptyValue,omitempty" yaml:"allowEmptyValue,omitempty"`
	Deprecated      bool `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`

	// Number
	Min        *float64 `json:"minimum,omitempty" yaml:"minimum,omitempty"`
	Max        *float64 `json:"maximum,omitempty" yaml:"maximum,omitempty"`
	MultipleOf *float64 `json:"multipleOf,omitempty" yaml:"multipleOf,omitempty"`

	// String
	MinLength uint64  `json:"minLength,omitempty" yaml:"minLength,omitempty"`
	MaxLength *uint64 `json:"maxLength,omitempty" yaml:"maxLength,omitempty"`
	Pattern   string  `json:"pattern,omitempty" yaml:"pattern,omitempty"`

	// Array
	MinItems uint64  `json:"minItems,omitempty" yaml:"minItems,omitempty"`
	MaxItems *uint64 `json:"maxItems,omitempty" yaml:"maxItems,omitempty"`

	// Object
	Required []string `json:"required,omitempty" yaml:"required,omitempty"`
	MinProps uint64   `json:"minProperties,omitempty" yaml:"minProperties,omitempty"`
	MaxProps *uint64  `json:"maxProperties,omitempty" yaml:"maxProperties,omitempty"`
}

func (schemaRef *SchemaRef) MarshalJSON() (res []byte, err error) {
	schema := Schema{}
	if schemaRef.Ref == "" {
		schemaRef.Ref = schemaRef.RefExt
		//schemaRef.RefExt = ""
	}
	if schemaRef.Ref != "" {
		schema.Ref = schemaRef.Ref
		schema.RefId = schemaRef.RefId
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
	schemaRef.RefId = schema.RefId
	if schemaRef.Ref == "" {
		schemaRef.Value = &schema
	}

	return nil
}

type Components struct {
	s map[uint]*component
	m map[string]*component
}

type component struct {
	refId  uint
	ref    string
	schema *SchemaRef
}

func NewComponents() (c *Components) {
	c = new(Components)
	c.m = map[string]*component{}
	c.s = map[uint]*component{}
	return c
}

func (c *Components) Add(refId uint, ref string, schema *SchemaRef) {
	data := &component{refId: refId, ref: ref, schema: schema}
	if refId != 0 {
		c.s[refId] = data
	}
	if ref != "" {
		c.m[ref] = data
	}

}

func (c *Components) Component(schema *SchemaRef) (*SchemaRef, uint) {
	var ok bool
	var ret *component
	if ret, ok = c.s[schema.RefId]; ok {
		return ret.schema, ret.refId
	} else if ret, ok = c.m[schema.Ref]; ok {
		return ret.schema, ret.refId
	}

	return nil, 0
}

type Schema2conv struct {
	Components *Components
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
	if component, _ := s.Components.Component(&schema); component != nil {
		s.sets[ref]++
		schema = *component
	}

	s.CombineSchemas(&schema)

	if schema.Value == nil {
		return
	}
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

func (s *Schema2conv) SchemaComponents(schema *SchemaRef, components *Components) {

	if component, refId := s.Components.Component(schema); refId != 0 {
		components.Add(refId, "", component)
	} else {
		return
	}

	if schema.Value == nil {
		return
	}

	for _, item := range schema.Value.AnyOf {
		s.SchemaComponents(item, components)
	}

	for _, item := range schema.Value.OneOf {
		s.SchemaComponents(item, components)
	}

	for _, item := range schema.Value.AllOf {
		s.SchemaComponents(item, components)
	}

	//s.CombineSchemas(&schema)

	switch schema.Value.Type {
	case openapi3.TypeObject:
		for _, property := range schema.Value.Properties {
			s.SchemaComponents(property, components)
		}
	case openapi3.TypeArray:
		s.SchemaComponents(schema.Value.Items, components)

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

	if schema.Value == nil {
		return
	}

	if len(schema.Value.AllOf) >= 1 {
		combineSchemas = schema.Value.AllOf
		fmt.Println(combineSchemas)
	} else {
		if len(schema.Value.AnyOf) >= 1 {
			rand.Seed(time.Now().UnixNano())
			n := rand.Intn(len(schema.Value.AnyOf)-1) + 1
			combineSchemas = s.anyMoreSchemas(schema.Value.AnyOf, n)

		} else if len(schema.Value.OneOf) >= 1 {
			combineSchemas = s.anyMoreSchemas(schema.Value.OneOf, 1)

		}
	}
	//fmt.Println(combineSchemas)
	for _, item := range combineSchemas {

		if item.Ref != "" {
			if component, _ := s.Components.Component(item); component != nil {
				item = component
			}
			if item.Value == nil {
				continue
			}
		}
		schema.Value.Type = item.Value.Type

		if item.Value.Type != openapi3.TypeObject {
			item.Value.Properties = Schemas{}
			schema.Value = item.Value
			continue
		}

		for key, property := range item.Value.Properties {
			/*
				if property.Value == nil {
					//			continue
				}

				if property.Value.Type == "" {
					//continue
				}
			*/

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

	if component, _ := s.Components.Component(schema1); component != nil {
		//s.sets[ref1]++
		schema1 = component
	}

	if schema1.Value == nil {
		return false
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

func (s *Schema2conv) GetRefIds(schema *SchemaRef) (ret []interface{}) {
	if schema.RefId != 0 {
		return append(ret, schema.RefId)
	}

	if schema.Value == nil {
		return nil
	}

	if schema.Value.Type == openapi3.TypeArray {
		ret = s.GetRefIds(schema.Value.Items)
	}

	if schema.Value.Type == openapi3.TypeObject {
		for _, schemaRef := range schema.Value.Properties {
			ret = append(ret, s.GetRefIds(schemaRef)...)
		}
	}

	return ret

}

func (s *Schema2conv) GetRefs(schema *SchemaRef) (ret []interface{}) {
	if schema.Ref != "" {
		return append(ret, schema.Ref)
	}

	if schema.Value == nil {
		return nil
	}

	if schema.Value.Type == openapi3.TypeArray {
		ret = s.GetRefs(schema.Value.Items)
	}

	if schema.Value.Type == openapi3.TypeObject {
		for _, schemaRef := range schema.Value.Properties {
			ret = append(ret, s.GetRefs(schemaRef)...)
		}
	}

	return ret

}

func (s *Schema2conv) FillRefId(schema *SchemaRef) {
	if schema.RefId != 0 {
		return
	}

	if schema.Ref != "" {
		if res, refId := s.Components.Component(schema); res != nil {
			schema.RefId = refId
		}
		return
	}

	if schema.Value == nil {
		return
	}

	if len(schema.Value.AllOf) > 0 {
		for _, item := range schema.Value.AllOf {
			s.FillRefId(item)
		}
	}

	if len(schema.Value.AnyOf) > 0 {
		for _, item := range schema.Value.AllOf {
			s.FillRefId(item)
		}
	}

	if len(schema.Value.OneOf) > 0 {
		for _, item := range schema.Value.AllOf {
			s.FillRefId(item)
		}
	}

	if schema.Value.Type == openapi3.TypeArray {
		s.FillRefId(schema.Value.Items)
	}

	if schema.Value.Type == openapi3.TypeObject {
		for _, schemaRef := range schema.Value.Properties {
			s.FillRefId(schemaRef)
		}
	}

}

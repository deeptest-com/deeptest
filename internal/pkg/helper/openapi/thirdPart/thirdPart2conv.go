package thirdPart

import (
	"fmt"
	"github.com/getkin/kin-openapi/openapi3"
)

type thirdPart2conv struct {
}

var typeRef map[string]string

func init() {
	typeRef = map[string]string{
		"array":         openapi3.TypeArray,
		"object":        openapi3.TypeObject,
		"string":        openapi3.TypeString,
		"text":          openapi3.TypeString,
		"multipartFile": openapi3.TypeString,
		"number":        openapi3.TypeNumber,
		"char":          openapi3.TypeString,
		"datetime":      openapi3.TypeString,
		"boolean":       openapi3.TypeBoolean,
		"integer":       openapi3.TypeInteger,
		"date":          openapi3.TypeString,
		"double":        openapi3.TypeNumber,
	}
}

func NewThirdPart2conv() *thirdPart2conv {
	return new(thirdPart2conv)
}

func (t *thirdPart2conv) typeConvert(_type string) string {

	if ret, ok := typeRef[_type]; ok {
		return ret
	}
	panic(fmt.Errorf("the type %s Convert faild", _type))
}

func (t *thirdPart2conv) Convert(schemas Schemas) (ret *openapi3.SchemaRef) {
	ret = new(openapi3.SchemaRef)
	ret.Value = new(openapi3.Schema)
	ret.Value.Type = openapi3.TypeObject
	ret.Value.Properties = openapi3.Schemas{}
	for key, schema := range schemas {
		ret.Value.Properties[key] = t.schemaConvert(schema)
	}

	return

}

func (t *thirdPart2conv) schemaConvert(schema *Schema) (schemaRef *openapi3.SchemaRef) {
	if schema == nil {
		return
	}
	schemaRef = new(openapi3.SchemaRef)
	schemaRef.Value = new(openapi3.Schema)
	schemaRef.Value.Type = t.typeConvert(schema.Type)
	schemaRef.Value.Description = schema.Description
	schemaRef.Value.Title = schema.FiledName
	if schema.Type == openapi3.TypeObject {
		schemaRef.Value.Properties = openapi3.Schemas{}
		for key, property := range schema.Properties {
			schemaRef.Value.Properties[key] = t.schemaConvert(property)
		}
	} else if schema.Type == openapi3.TypeArray {
		if schema.Items != nil {
			schemaRef.Value.Items = t.schemaConvert(schema.Items)
		} else {
			schemaRef.Value.Items = t.schemaConvert(schema.Properties["items"])
		}

	}

	return
}

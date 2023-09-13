package fields

import "github.com/getkin/kin-openapi/openapi3"

type Field struct {
	FieldName  FieldName
	FieldType  FieldType
	SubField   *Field
	Properties []*Field
	IsProperty bool
}

type FieldType string

type FieldName string

type Fields map[FieldName]FieldType

type FieldArray []Field

func (f *FieldArray) Add(field Field) {
	*f = append([]Field{field}, *f...)
}

func (f *Field) ToArray() (arr *FieldArray) {
	if f.FieldType == openapi3.TypeObject || !f.IsProperty {
		arr.Add(*f)
	}
	for _, field := range f.Properties {
		arr.Add(*field)
	}
	return arr
}

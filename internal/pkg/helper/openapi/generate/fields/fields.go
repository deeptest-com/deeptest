package fields

import "github.com/getkin/kin-openapi/openapi3"

type Field struct {
	FieldName    FieldName `json:"fieldName"`
	FieldRefName FieldName `json:"fieldRefName"`
	FieldType    FieldType `json:"fieldType"`
	SubField     *Field    `json:"subField"`
	Properties   []*Field  `json:"properties"`
	IsProperty   bool      `json:"isProperty"`
	Description  string    `json:"description"`
}

const (
	Array  FieldType = "array"
	Object FieldType = "object"
	number FieldType = "number"
)

type FieldType string

type FieldName string

type Fields map[FieldName]FieldType

type FieldArray []Field

func (f *FieldArray) Add(field Field) {
	*f = append([]Field{field}, *f...)
}

func (f *Field) ToArray() (arr *FieldArray) {
	arr = new(FieldArray)
	f.setArray(f, arr)
	return arr
}

func (f *Field) setArray(field *Field, arr *FieldArray) {
	if field.FieldType == openapi3.TypeObject || !field.IsProperty {
		arr.Add(*field)
	}
	if field.SubField != nil {
		f.setArray(field.SubField, arr)
	}
	for _, item := range field.Properties {
		f.setArray(item, arr)
	}
	//return arr
}

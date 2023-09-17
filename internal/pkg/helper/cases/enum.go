package casesHelper

type FieldType string

const (
	FieldTypeInteger FieldType = "integer"
	FieldTypeNumber  FieldType = "number"
	FieldTypeBoolean FieldType = "boolean"
	FieldTypeString  FieldType = "string"
	FieldTypeArray   FieldType = "array"
	FieldTypeObject  FieldType = "object"
	FieldTypeNull    FieldType = "null"
)

func (e FieldType) String() string {
	return string(e)
}

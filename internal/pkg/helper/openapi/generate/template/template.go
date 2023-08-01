package template

type Interface interface {
	Field(name, _type string)
	Class(name string)
}

type field struct {
	fieldName fieldName
	fieldType fieldType
}

type fieldType string

type fieldName string

type fields map[fieldName]fieldType

type filedArray []field

type class string

type Base struct {
	fields     fields
	class      class
	fieldArray filedArray
}

func (t *Base) AddField(field field) {
	if t.fields == nil {
		t.fields = make(fields)
	}
	if _, ok := t.fields[field.fieldName]; !ok {
		t.fieldArray = append(t.fieldArray, field)
	}
}

func (t *Base) SetClass(class class) {
	t.class = class
}

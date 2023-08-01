package template

import (
	"strings"
)

type TypeScript struct {
	Base
}

func (t *TypeScript) fieldTypeConv(fieldType fieldType) (newType fieldType) {
	return
}

func (t *TypeScript) fieldString(field field) (ret string) {
	ret = strings.ReplaceAll("${_name_}: ${_type_},", "${_name_}", string(field.fieldName))
	ret = strings.ReplaceAll(ret, "${_type_}", string(field.fieldType))
	return
}

func (t *TypeScript) classString(name class) string {
	return strings.ReplaceAll("export interface ${_name_} { ${_content_} }", "${_name_}", string(name))
}

func (t *TypeScript) AddField(fieldName fieldName, fieldType fieldType) {
	field := field{fieldName: fieldName}
	field.fieldType = t.fieldTypeConv(fieldType)
	t.Base.AddField(field)
}

func (t *TypeScript) Content() (ret string) {
	ret = t.classString(t.class)
	fieldString := ""

	for _, field := range t.fieldArray {
		fieldString += t.fieldString(field)
	}

	return strings.ReplaceAll(ret, "{_content_}", fieldString)
}

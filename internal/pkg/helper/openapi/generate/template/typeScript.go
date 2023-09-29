package template

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/helper/openapi/generate/fields"
	"strings"
)

type typeScript struct {
	fieldTpl string
	classTpl string
}

var typeScriptTypeMap map[fields.FieldType]string

func init() {
	typeScriptTypeMap = map[fields.FieldType]string{
		"number": "number",
		"string": "string",
	}
}

func newTypeScript() (ret *typeScript) {
	ret = new(typeScript)
	ret.fieldTpl = ""
	return
}

func (t *typeScript) typeConvert(fieldType fields.FieldType) (newType string) {
	var ok bool
	if newType, ok = typeScriptTypeMap[fieldType]; !ok {
		panic(fmt.Errorf("%s can't convert", fieldType))
	}
	return newType
}

func (t *typeScript) CreateCode(field fields.Field) string {
	switch field.FieldType {
	case fields.Array:
		return t.array(field)
	case fields.Object:
		return t.object(field)

	}
	return ""
}

func (t *typeScript) array(field fields.Field) string {
	return fmt.Sprintf("type %s = %s[]", field.FieldName, t.typeConvert(field.SubField.FieldType))
}

func (t *typeScript) object(field fields.Field) string {
	code := fmt.Sprintf("export interface %s {", field.FieldName)
	var properties []string
	for _, property := range field.Properties {
		properties = append(properties, fmt.Sprintf("%s:%s", property.FieldName, t.typeConvert(property.FieldType)))
	}

	return fmt.Sprintf("%s \n %s \n }", code, strings.Join(properties, "\n"))
}

/*

func (t *typeScript) fieldTypeConv(fieldType ieldType) (newType fieldType) {
	return
}

func (t *typeScript) fieldString(field field) (ret string) {
	ret = strings.ReplaceAll("${_name_}: ${_type_},", "${_name_}", string(field.fieldName))
	ret = strings.ReplaceAll(ret, "${_type_}", string(field.fieldType))
	return
}

func (t *typeScript) classString(class class) (ret string) {
	return strings.ReplaceAll("export interface ${_name_} { ${_content_} }", "${_name_}", string(class.name))
}

/*
func (t *typeScript) AddField(fieldName fieldName, fieldType fieldType) {
	field := field{fieldName: fieldName}
	field.fieldType = t.fieldTypeConv(fieldType)
	t.Base.AddField(field)
}
*/

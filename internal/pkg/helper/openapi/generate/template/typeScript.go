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
		"number":  "number",
		"string":  "string",
		"object":  "object",
		"integer": "number",
		"":        "null",
		"any":     "any",
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
		//panic(fmt.Errorf("%s can't convert", fieldType))
		newType = string(fieldType)
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
	//if field.SubField.FieldType ==
	varName := field.SubField.FieldName
	if varName == "" {
		varName = fields.FieldName(field.SubField.FieldType)
	}
	ret := fmt.Sprintf("type %s = %s[]", field.FieldName, t.typeConvert(fields.FieldType(varName)))
	if field.IsProperty {
		ret = fmt.Sprintf("	%s: %s[]", field.FieldName, t.typeConvert(fields.FieldType(varName)))
	}
	ret = t.addDescription(After, &field, ret)
	return ret
}

func (t *typeScript) object(field fields.Field) string {
	code := fmt.Sprintf("export interface %s {", field.FieldName)
	var properties []string
	for _, property := range field.Properties {
		var propertyStr string
		if property.FieldType == fields.Array {
			propertyStr = t.array(*property)
		} else {
			if property.FieldRefName != "" {
				propertyStr = fmt.Sprintf("	%s: %s", property.FieldName, property.FieldRefName)
			} else {
				propertyStr = fmt.Sprintf("	%s: %s", property.FieldName, t.typeConvert(property.FieldType))
			}
			propertyStr = t.addDescription(After, property, propertyStr)
		}
		properties = append(properties, propertyStr)
	}
	code = t.addDescription(Before, &field, code)
	return fmt.Sprintf("%s \n %s \n }", code, strings.Join(properties, "\n"))
}

func (t *typeScript) addDescription(position Position, field *fields.Field, str string) string {
	if field.Description != "" {
		if position == After {
			str += fmt.Sprintf("	//%s", field.Description)
		} else {
			str = fmt.Sprintf("	/*\n%s\n*/\n", field.Description) + str
		}

	}
	return str
}

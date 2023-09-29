package template

type golang struct {
}

func newGolang() (ret *typeScript) {
	ret = new(typeScript)
	ret.fieldTpl = ""
	return
}

/*

func (l *golang) fieldTypeConv(fieldType FieldType) (newType fieldType) {
	return
}

func (l *golang) fieldString(field field) (ret string) {
	ret = strings.ReplaceAll("${_name_}: ${_type_},", "${_name_}", string(field.fieldName))
	ret = strings.ReplaceAll(ret, "${_type_}", string(field.fieldType))
	return
}

func (l *golang) classString(class class) string {
	return strings.ReplaceAll("export interface ${_name_} { ${_content_} }", "${_name_}", string(class.name))
}

*/

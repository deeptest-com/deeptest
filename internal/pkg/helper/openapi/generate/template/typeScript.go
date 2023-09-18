package template

type typeScript struct {
	fieldTpl string
	classTpl string
}

func newTypeScript() (ret *typeScript) {
	ret = new(typeScript)
	ret.fieldTpl = ""
	return
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

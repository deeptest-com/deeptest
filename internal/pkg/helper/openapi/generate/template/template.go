package template

type languageInter interface {
	fieldTypeConv(fieldType fieldType) (newType fieldType)
	fieldString(field field) (ret string)
	classString(class class) (ret string)
}

type field struct {
	fieldName fieldName
	fieldType fieldType
}

type fieldType string

type fieldName string

type fields map[fieldName]fieldType

type filedArray []field

type class struct {
	name string
}

type langType string

const (
	Go langType = "golang"
	TS langType = "typeScript"
)

type languages map[langType]languageInter

var language languages

var varTpl string
var fieldTpl string

func init() {
	language = languages{
		Go: newGolang(),
		TS: newTypeScript(),
	}

	varTpl = ` ${__class__} {
				${__fields__}
         };`

	fieldTpl = `${__name__}:${__type__},`
}

type template struct {
	varType    fieldType
	fields     fields
	class      class
	fieldArray filedArray
	language   languageInter
}

func newTemplate(langType langType, class class) (ret *template) {
	ret = &template{class: class}
	ret.setLanguage(langType)
	return
}

func (t *template) setLanguage(langType langType) {
	t.language = language[langType]
}

func (t *template) AddField(field field) {
	if t.fields == nil {
		t.fields = make(fields)
	}

	field.fieldType = t.fieldTypeConv(field.fieldType)
	if _, ok := t.fields[field.fieldName]; !ok {
		t.fieldArray = append(filedArray{field}, t.fieldArray...)
		t.fieldArray = append(t.fieldArray, field)
	}
}

func (t *template) SetClass(class class) {
	t.class = class
}

func (t *template) fieldTypeConv(fieldType fieldType) (newType fieldType) {
	newType = t.language.fieldTypeConv(fieldType)
	return
}

func (t *template) ClassContent() (ret string) {
	return
}

/*
func (t *template) fieldString(field field) (ret string) {
	ret = strings.ReplaceAll("${_name_}: ${_type_},", "${_name_}", string(field.fieldName))
	ret = strings.ReplaceAll(ret, "${_type_}", string(field.fieldType))
	return
}

func (t *template) classString(name class) string {
	return strings.ReplaceAll("export interface ${_name_} { ${_content_} }", "${_name_}", string(name))
}
*/
/*
func (t *template) Content() (ret string) {
	ret = t.classString(t.class)
	fieldString := ""

	for _, field := range t.fieldArray {
		fieldString += t.fieldString(field)
	}

	return strings.ReplaceAll(ret, "{_content_}", fieldString)
}

*/

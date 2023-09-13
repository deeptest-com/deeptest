package template

import "github.com/aaronchen2k/deeptest/internal/pkg/helper/openapi/generate/fields"

type languageInter interface {
	//fieldTypeConv(fieldType fields.fieldType) (newType fieldType)
	//fieldString(field field) (ret string)
}

type LangType string

const (
	Go LangType = "golang"
	TS LangType = "typeScript"
)

type languages map[LangType]languageInter

var language languages

var varTpl string
var fieldTpl string

func init() {
	language = languages{
		Go: newGolang(),
		TS: newTypeScript(),
	}
}

type template struct {
	fieldArray fields.FieldArray
	language   languageInter
}

func NewTemplate(langType LangType, fields *fields.FieldArray) (ret *template) {
	ret.setLanguage(langType)
	return
}

func (t *template) setLanguage(langType LangType) {
	t.language = language[langType]
}

func (t *template) CreateCode() string {
	return ""
}

/*
func (t *template) fieldTypeConv(fieldType fieldType) (newType fieldType) {
	newType = t.language.fieldTypeConv(fieldType)
	return
}

func (t *template) ClassContent() (ret string) {
	return
}
*/

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

package template

import (
	"fmt"
	"github.com/deeptest-com/deeptest/internal/pkg/helper/openapi/generate/fields"
	"regexp"
	"strings"
)

type languageInter interface {
	//fieldTypeConv(fieldType fields.fieldType) (newType fieldType)
	CreateCode(field fields.Field) (ret string)
}

type LangType string

const (
	Go LangType = "golang"
	TS LangType = "typeScript"
)

type NameRule string

const (
	LowerCase NameRule = "lowerCase"
	UpperCase NameRule = "upperCase"
	Underline NameRule = "underline"
)

type Position string

const (
	Before Position = "before"
	After  Position = "after"
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
	fieldArray *fields.FieldArray
	language   languageInter
}

func NewTemplate(langType LangType, fields *fields.FieldArray) (ret *template) {
	ret = new(template)
	ret.fieldArray = fields
	ret.setLanguage(langType)
	return
}

func (t *template) setLanguage(langType LangType) {
	var ok bool
	if t.language, ok = language[langType]; !ok {
		panic(fmt.Errorf("no %s langType", langType))
	}

}

func (t *template) CreateCode() (ret string) {
	var codes []string
	if t.fieldArray != nil {
		for _, field := range *t.fieldArray {
			codes = append(codes, t.language.CreateCode(field))
		}
	}
	ret = strings.Join(codes, "\n")
	re := regexp.MustCompile("\n+")
	ret = re.ReplaceAllLiteralString(ret, "\n")
	return
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

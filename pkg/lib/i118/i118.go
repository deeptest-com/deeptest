package _i118Utils

import (
	"encoding/json"
	"fmt"
	"github.com/aaronchen2k/deeptest"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"path"
)

var I118Prt *message.Printer

func Init(lang string, app string) {
	//var once sync.Once
	//once.Do(func() {

	langRes := path.Join("res", lang, "messages.json")
	fmt.Printf("path %s\n", langRes)

	bytes, _ := deeptest.ReadResData(langRes)
	fmt.Printf("content %s\n", string(bytes))

	InitResFromAsset(bytes)

	if lang == "zh" {
		I118Prt = message.NewPrinter(language.SimplifiedChinese)
	} else {
		I118Prt = message.NewPrinter(language.AmericanEnglish)
	}
	//})
}

type I18n struct {
	Language string    `json:"language"`
	Messages []Message `json:"messages"`
}

type Message struct {
	Id          string `json:"id"`
	Message     string `json:"message,omitempty"`
	Translation string `json:"translation,omitempty"`
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func InitResFromAsset(bytes []byte) {
	var i18n I18n
	json.Unmarshal(bytes, &i18n)

	msgArr := i18n.Messages
	tag := language.MustParse(i18n.Language)

	for _, e := range msgArr {
		message.SetString(tag, e.Id, e.Translation)
	}
}

func Sprintf(key message.Reference, a ...interface{}) string {
	if I118Prt == nil {
		return fmt.Sprintf("%s, %#v", key.(string), a)
	} else {
		return I118Prt.Sprintf(key, a...)
	}
}

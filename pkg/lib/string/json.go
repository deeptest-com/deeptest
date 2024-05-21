package _stringUtils

import (
	"bytes"
	"encoding/json"
	"github.com/kataras/iris/v12"
	"strings"
)

func JsonWithoutHtmlEscaped(obj interface{}) (ret string) {
	bf := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(bf)
	jsonEncoder.SetEscapeHTML(false)
	err := jsonEncoder.Encode(obj)

	if err != nil {
		ret = err.Error()
	} else {
		ret = strings.TrimSpace(bf.String())
	}

	return
}

func FormatJsonStr(str string) (ret string) {
	mp := iris.Map{}
	json.Unmarshal([]byte(str), &mp)

	bytes, err := json.MarshalIndent(mp, "", "    ")

	if err == nil {
		ret = string(bytes)
	} else {
		ret = err.Error()
	}

	return
}

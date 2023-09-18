package _stringUtils

import (
	"bytes"
	"encoding/json"
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

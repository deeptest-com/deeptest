package k6Comm

import (
	"github.com/aaronchen2k/deeptest"
	"path/filepath"
)

func GetTmpl(name string) (content string) {
	tmplFile := filepath.Join("res", "k6", name+".ftl")
	bytes, _ := deeptest.ReadResData(tmplFile)

	content = string(bytes)

	return
}

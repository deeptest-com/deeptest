package convert

import (
	"encoding/json"
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/helper/openapi/convert/postman"
	"github.com/getkin/kin-openapi/openapi3"
	"os/exec"
	"runtime"
)

type Postman struct {
	driver
	doc *postman.Doc
}

func newPostman() *Postman {
	return new(Postman)
}

func (d *Postman) toOpenapi() (doc *openapi3.T, err error) {
	system := runtime.GOOS
	cmd := fmt.Sprintf("plugins/postman2openapi/%s/postman2openapi", system)
	// 此处是windows版本
	var output []byte
	if system == "windows" {
		c := exec.Command("cmd", "/C", cmd)
		output, _ = c.CombinedOutput()
	} else {
		c := exec.Command("bash", "-c", cmd)
		output, _ = c.CombinedOutput()
	}
	err = json.Unmarshal(output, doc)

	fmt.Println(string(output))
	fmt.Println(cmd)
	return
}

func (d *Postman) Doc(data []byte) {
	/*
		err := json.Unmarshal(data, d.doc)
		if err != err {
			panic(err)
		}
	*/
}

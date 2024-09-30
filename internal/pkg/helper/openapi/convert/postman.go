package convert

import (
	"encoding/json"
	"fmt"
	"github.com/deeptest-com/deeptest/internal/pkg/helper/openapi/convert/postman"
	logUtils "github.com/deeptest-com/deeptest/pkg/lib/log"
	"github.com/getkin/kin-openapi/openapi3"
	"os"
	"os/exec"
	"runtime"
	"strings"
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
	cmd := ""
	var c *exec.Cmd
	var output []byte
	if system == "windows" {
		cmd = fmt.Sprintf("plugins/postman2openapi/%s/postman2openapi.exe -f json %s", system, d.FilePath)
		cmd = strings.ReplaceAll(cmd, "/", "\\")
		c = exec.Command("cmd", "/C", cmd)
	} else {
		cmd = fmt.Sprintf("plugins/postman2openapi/%s/postman2openapi -f json %s", system, d.FilePath)
		c = exec.Command("bash", "-c", cmd)
	}
	dir, _ := os.Getwd()
	logUtils.Info("workdir:" + dir + ",命令路径:" + cmd)
	output, err = c.CombinedOutput()
	if err != nil {
		return
	}

	doc = new(openapi3.T)
	err = json.Unmarshal(output, doc)
	if err != nil {
		return
	}

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

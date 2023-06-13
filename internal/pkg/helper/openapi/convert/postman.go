package convert

import (
	"encoding/json"
	"github.com/aaronchen2k/deeptest/internal/pkg/helper/openapi/convert/postman"
	"github.com/getkin/kin-openapi/openapi3"
)

type Postman struct {
	driver
	doc *postman.Doc
}

func newPostman() *Postman {
	return new(Postman)
}

func (d *Postman) toOpenapi() (doc *openapi3.T) {
	doc = new(openapi3.T)
	return
}

func (d *Postman) Doc(data []byte) {
	err := json.Unmarshal(data, d.doc)
	if err != err {
		panic(err)
	}
}

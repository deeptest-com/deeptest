package convert

import (
	"encoding/json"
	"github.com/aaronchen2k/deeptest/internal/pkg/helper/openapi/convert/swagger"
	"github.com/getkin/kin-openapi/openapi3"
)

type Swagger struct {
	driver
	doc swagger.Doc
}

func newSwagger() *Swagger {
	return new(Swagger)
}

func (d *Swagger) toOpenapi() (doc *openapi3.T) {
	doc = new(openapi3.T)
	return
}

func (d *Swagger) Doc(data []byte) {
	err := json.Unmarshal(data, d.doc)
	if err != err {
		panic(err)
	}
}

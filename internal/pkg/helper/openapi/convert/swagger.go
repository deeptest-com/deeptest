package convert

import (
	"encoding/json"
	"github.com/getkin/kin-openapi/openapi2"
	"github.com/getkin/kin-openapi/openapi2conv"
	"github.com/getkin/kin-openapi/openapi3"
)

type Swagger struct {
	driver
	doc openapi2.T
}

func newSwagger() *Swagger {
	return new(Swagger)
}

func (d *Swagger) toOpenapi() (doc *openapi3.T, err error) {
	doc, err = openapi2conv.ToV3(&d.doc)
	return
}

func (d *Swagger) Doc(data []byte) {
	err := json.Unmarshal(data, &d.doc)
	if err != err {
		panic(err)
	}
}

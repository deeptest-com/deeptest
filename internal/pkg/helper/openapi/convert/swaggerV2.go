package convert

import (
	"encoding/json"
	"github.com/getkin/kin-openapi/openapi2"
	"github.com/getkin/kin-openapi/openapi2conv"
	"github.com/getkin/kin-openapi/openapi3"
)

type SwaggerV2 struct {
	driver
	doc openapi2.T
}

func newSwaggerV2() *SwaggerV2 {
	return new(SwaggerV2)
}

func (d *SwaggerV2) toOpenapi() (doc *openapi3.T, err error) {
	doc, err = openapi2conv.ToV3(&d.doc)
	return
}

func (d *SwaggerV2) Doc(data []byte) {
	err := json.Unmarshal(data, &d.doc)
	if err != err {
		panic(err)
	}
}

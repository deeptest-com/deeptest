package convert

import (
	"encoding/json"
	"github.com/aaronchen2k/deeptest/internal/pkg/helper/openapi/convert/yapi"
	"github.com/getkin/kin-openapi/openapi3"
)

type YApi struct {
	driver
	doc yapi.Doc
}

func newYApi() *YApi {
	return new(YApi)
}

func (d *YApi) toOpenapi() (doc *openapi3.T) {
	doc = new(openapi3.T)
	return
}

func (d *YApi) Doc(data []byte) {
	err := json.Unmarshal(data, d.doc)
	if err != err {
		panic(err)
	}
}

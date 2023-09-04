package convert

import (
	"encoding/json"
	"github.com/getkin/kin-openapi/openapi2"
	"github.com/getkin/kin-openapi/openapi2conv"
	"github.com/getkin/kin-openapi/openapi3"
	"go.uber.org/zap"
)

var Logger *zap.Logger

type YApi struct {
	driver
	doc openapi2.T
}

func newYApi() *YApi {
	return new(YApi)
}

func (d *YApi) toOpenapi() (doc *openapi3.T, err error) {
	doc, err = openapi2conv.ToV3(&d.doc)
	return
}

func (d *YApi) Doc(data []byte) {
	err := json.Unmarshal(data, &d.doc)
	if err != err {
		panic(err)
	}
}

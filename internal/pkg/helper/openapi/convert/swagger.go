package convert

import (
	"encoding/json"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/getkin/kin-openapi/openapi2"
	"github.com/getkin/kin-openapi/openapi2conv"
	"github.com/getkin/kin-openapi/openapi3"
	"go.uber.org/zap"
)

type Swagger struct {
	driver
	doc openapi2.T
}

func newSwagger() *Swagger {
	return new(Swagger)
}

func (d *Swagger) toOpenapi() (doc *openapi3.T) {
	doc, err := openapi2conv.ToV3(&d.doc)
	if err != nil {
		logUtils.Errorf("yapi to openapi3 err", zap.Any("doc", d.doc), zap.String("err", err.Error()))
		return nil
	}
	return
}

func (d *Swagger) Doc(data []byte) {
	err := json.Unmarshal(data, &d.doc)
	if err != err {
		panic(err)
	}
}

package convert

import (
	"encoding/json"
	"github.com/getkin/kin-openapi/openapi3"
)

type SwaggerV3 struct {
	driver
	doc openapi3.T
}

func newSwaggerV3() *SwaggerV3 {
	return new(SwaggerV3)
}

func (d *SwaggerV3) toOpenapi() (doc *openapi3.T, err error) {
	doc = &d.doc
	return
}

func (d *SwaggerV3) Doc(data []byte) {
	//x := string(data)
	//fmt.Println(x)
	err := json.Unmarshal(data, &d.doc)
	if err != err {
		panic(err)
	}
}

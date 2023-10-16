package service

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/helper/openapi/generate"
	"github.com/aaronchen2k/deeptest/internal/pkg/helper/openapi/generate/template"
	schemaHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/schema"
	_commUtils "github.com/aaronchen2k/deeptest/pkg/lib/comm"
)

type EndpointCodeService struct {
	ServeService *ServeService `inject:""`
}

func (s *EndpointCodeService) Generate(langType template.LangType, serveId uint, data string) (code string) {
	if data == "" {
		return
	}

	schema2Code := generate.NewSchema2Code(langType, "")
	schema2Code.Components = s.ServeService.Components(serveId)
	//schema1 := openapi3.Schema{}
	//_commUtils.JsonDecode(data, &schema)
	//_commUtils.JsonDecode("{\"type\":\"array\",\"items\":{\"type\":\"number\"}}", &schema)
	//_commUtils.JsonDecode("{\"properties\":{\"id\":{\"type\":\"number\"},\"name\":{\"type\":\"string\"}},\"type\":\"object\"}", &schema)
	//_commUtils.JsonDecode("{\"type\":\"array\",\"items\":{\"properties\":{\"id\":{\"type\":\"number\"},\"name\":{\"type\":\"string\"}},\"type\":\"object\"}}", &schema)
	schema := schemaHelper.SchemaRef{}
	//data = "{\"type\":\"object\",\"properties\":{\"name1\":{\"type\":\"object\",\"ref\":\"#/components/schemas/user1\",\"name\":\"user1\"},\"name2\":{\"type\":\"string\"},\"name3\":{\"type\":\"string\"}}}"
	_commUtils.JsonDecode(data, &schema)
	//_commUtils.JsonDecode("{\"type\":\"array\",\"items\":{\"type\":\"number\"}}", &schema1)
	//copier.CopyWithOption(&schema, a, copier.Option{DeepCopy: true})
	//fmt.Println(schema, "+++++++++++++")

	code = schema2Code.Convert(schema)
	return

}

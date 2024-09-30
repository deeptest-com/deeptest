package test

import (
	"context"
	fileUtils "github.com/deeptest-com/deeptest/pkg/lib/file"
	"github.com/getkin/kin-openapi/openapi3"
	"testing"
)

func TestYapi(t *testing.T) {
	pth := "/Users/aaron/rd/project/gudi/deeptest/xdoc/openapi/yapi/swagger.json"
	//content := fileUtils.ReadFileBuf(pth)
	//data := domain.YapiData{}
	//json.Unmarshal(content, &data.YapiCategories)
	//log.Print(data)

	ctx := context.Background()
	loader := &openapi3.Loader{Context: ctx, IsExternalRefsAllowed: true}

	doc3, err := loader.LoadFromFile(pth)
	if err != nil {
		panic(err)
	}

	jsonBytes, _ := doc3.MarshalJSON()

	fileUtils.WriteFile("/Users/aaron/out/input.json", string(jsonBytes))

	//var doc2 openapi2.T
	//err := json.Unmarshal(content, &doc2)
	//if err != nil {
	//	panic(err)
	//}
	//
	//doc3, err := openapi2conv.ToV3(&doc2)
	//if err != nil {
	//	panic(err)
	//}
	//
	//err = doc3.Validate(context.Background())
	//if err != nil {
	//	panic(err)
	//}
	//
	//jsonBytes, err := json.Marshal(doc3)
	//if err != nil {
	//	panic(err)
	//}
	//
	//fileUtils.WriteFile("/Users/aaron/out/input.json", string(jsonBytes))
}

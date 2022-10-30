package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/helper/openapi"
	_shellUtils "github.com/aaronchen2k/deeptest/pkg/lib/shell"
	"github.com/getkin/kin-openapi/openapi2"
	"github.com/getkin/kin-openapi/openapi3"
	"log"
)

type ImportService struct {
}

func (s *ImportService) Import(content []byte, typ string, targetId int) (err error) {
	//ctx := context.Background()
	//loader := &openapi3.Loader{Context: ctx, IsExternalRefsAllowed: true}
	//
	//doc, err := loader.LoadFromData(content)
	//err = doc.Validate(loader.Context)

	var doc3 *openapi3.T

	err = json.Unmarshal(content, &doc3)

	log.Print(doc3)

	return
}

func (s *ImportService) OpenApi2To3(src []byte) (ret []byte, err error) {
	var doc2 openapi2.T
	err = json.Unmarshal(src, &doc2)

	doc3, err := openapi.ToV3(&doc2)
	err = doc3.Validate(context.Background())

	ret, err = json.Marshal(doc3)

	return
}

func (s *ImportService) PostmanToOpenApi3(pth string) (ret []byte, err error) {
	cmd := fmt.Sprintf(`p2o %s`, pth)

	out, _ := _shellUtils.ExeShell(cmd)

	ret = []byte(out)

	return
}

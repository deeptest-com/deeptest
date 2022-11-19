package service

import (
	"context"
	"fmt"
	"github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/getkin/kin-openapi/openapi3"
)

type ImportService struct {
}

func (s *ImportService) Import(req domain.InterfaceImportReq, typ string, targetId int) (err error) {
	ctx := context.Background()
	loader := &openapi3.Loader{Context: ctx, IsExternalRefsAllowed: true}

	doc3, err := loader.LoadFromFile(req.File)

	if err != nil {
		return
	}

	fmt.Println(doc3)

	return
}

//func (s *ImportService) OpenApi2To3(src []byte) (ret []byte, err error) {
//	var doc2 openapi2.T
//	err = json.Unmarshal(src, &doc2)
//
//	doc3, err := openapi.ToV3(&doc2)
//	err = doc3.Validate(context.Background())
//
//	ret, err = json.Marshal(doc3)
//
//	return
//}
//
//func (s *ImportService) PostmanToOpenApi3(pth string) (ret []byte, err error) {
//	cmd := fmt.Sprintf(`p2o %s`, pth)
//
//	out, _ := _shellUtils.ExeShell(cmd)
//
//	ret = []byte(out)
//
//	return
//}

package service

import (
	"context"
	"encoding/json"
	"github.com/getkin/kin-openapi/openapi3"
	"net/url"
	"strings"
)

type SpecService struct {
}

func (s *SpecService) Load(pathOrUrl, typ string) (content string, err error) {
	var doc3 *openapi3.T

	//if typ == "openapi2" {
	//	v2Content := _fileUtils.ReadFileBuf(pathOrUrl)
	//
	//	var doc2 openapi2.T
	//	err = json.Unmarshal(v2Content, &doc2)
	//	if err != nil {
	//		return
	//	}
	//
	//	doc3, err = openapi.ToV3(&doc2)
	//	if err != nil {
	//		return
	//	}
	//
	//	//err = doc3.Validate(context.Background())
	//
	//	var bytes []byte
	//	bytes, err = json.MarshalIndent(doc3, "", "\t")
	//	if err != nil {
	//		return
	//	}
	//
	//	content = string(bytes)
	//} else if typ == "openapi3" {
	ctx := context.Background()
	loader := &openapi3.Loader{Context: ctx, IsExternalRefsAllowed: true}

	if strings.HasPrefix(pathOrUrl, "http") {
		var u *url.URL
		u, err = url.Parse(pathOrUrl)
		if err != nil {
			return
		}

		doc3, err = loader.LoadFromURI(u)
		if err != nil {
			return
		}
	} else {
		doc3, err = loader.LoadFromFile(pathOrUrl)
	}

	if err != nil {
		return
	}

	//_ := doc3.Validate(ctx)

	var bytes []byte
	bytes, err = json.Marshal(doc3)
	if err != nil {
		return
	}

	content = string(bytes)
	//}

	return
}

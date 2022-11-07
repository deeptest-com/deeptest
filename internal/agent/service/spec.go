package service

import (
	"context"
	"fmt"
	"github.com/aaronchen2k/deeptest/cmd/agent/v1/domain"
	"github.com/getkin/kin-openapi/openapi3"
	"gopkg.in/yaml.v3"
)

type SpecService struct {
}

func (s *SpecService) SubmitSpec(req domain.SubmitSpecReq) (doc3 *openapi3.T, err error) {
	ctx := context.Background()
	loader := &openapi3.Loader{Context: ctx, IsExternalRefsAllowed: true}

	doc3, err = loader.LoadFromFile(req.File)

	c, _ := yaml.Marshal(doc3)

	fmt.Sprintln(c)

	if err != nil {
		return
	}

	return
}

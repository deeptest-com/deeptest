package service

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/getkin/kin-openapi/openapi3"
)

type InterfaceService struct {
}

func (s *InterfaceService) Generate(doc *openapi3.T) (err error) {

	return
}

func (s *InterfaceService) ConvertPath(pth *openapi3.PathItem) (doc3 *model.Interface, err error) {

	return
}

package service

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/helper/openapi"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	"github.com/getkin/kin-openapi/openapi3"
)

type InterfaceService struct {
	InterfaceRepo *repo.InterfaceRepo `inject:""`
}

func (s *InterfaceService) Generate(doc *openapi3.T, projectId uint) (err error) {
	interfaces, err := openapi.ConvertPathsToInterfaces(doc)
	if err != nil {
		return
	}

	for _, interf := range interfaces {
		fmt.Sprintln(interf)
	}

	return
}

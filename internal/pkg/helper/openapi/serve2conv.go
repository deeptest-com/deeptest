package openapi

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/getkin/kin-openapi/openapi3"
)

type Serve2conv struct {
	serve    model.Serve
	endpoint []model.Endpoint
	doc      *openapi3.T
}

func NewServe2conv(serve model.Serve, endpoint []model.Endpoint) *Serve2conv {
	return &Serve2conv{serve: serve, endpoint: endpoint}
}

func (s *Serve2conv) ToV3() *openapi3.T {
	return s.doc
}

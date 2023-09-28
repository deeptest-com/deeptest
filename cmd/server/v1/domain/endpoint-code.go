package serverDomain

import "github.com/aaronchen2k/deeptest/internal/pkg/helper/openapi/generate/template"

type EndpointCodeReq struct {
	ServeId  uint
	Data     string
	LangType template.LangType
}

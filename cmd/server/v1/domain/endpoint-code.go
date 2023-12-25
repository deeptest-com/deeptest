package serverDomain

import "github.com/aaronchen2k/deeptest/internal/pkg/helper/openapi/generate/template"

type EndpointCodeReq struct {
	ProjectId uint              `json:"projectId"`
	Data      string            `json:"data"`
	LangType  template.LangType `json:"langType"`
	NameRule  template.NameRule `json:"nameRule"`
}

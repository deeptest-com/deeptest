package serverDomain

import (
	"github.com/aaronchen2k/deeptest/pkg/domain"
)

type ScenarioReqPaginate struct {
	_domain.PaginateReq
	Keywords string `json:"keywords"`
	Enabled  string `json:"enabled"`
}

package serverDomain

import (
	"github.com/aaronchen2k/deeptest/pkg/domain"
)

type DatapoolReqPaginate struct {
	_domain.PaginateReq
	ProjectId int64  `json:"projectId"`
	Name      string `json:"name"`
}

type DatapoolReq struct {
	_domain.Model

	Name string `json:"name"`
	Desc string `json:"desc"`
	Data string `json:"data"`

	ProjectId uint `json:"projectId"`
}

type DatapoolUploadResp struct {
	Path string          `json:"path"`
	Data [][]interface{} `json:"data"`
}

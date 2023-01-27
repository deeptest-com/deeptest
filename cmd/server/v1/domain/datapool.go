package domain

import (
	"github.com/aaronchen2k/deeptest/pkg/domain"
)

type DatapoolReq struct {
	_domain.Model

	Name string `json:"name"`
	Desc string `json:"desc"`
	Data string `json:"data"`

	ProjectId uint `json:"projectId"`
}

type DatapoolUploadResp struct {
	Path string `json:"path"`

	Headers []string   `json:"headers"`
	Rows    [][]string `json:"rows"`
}

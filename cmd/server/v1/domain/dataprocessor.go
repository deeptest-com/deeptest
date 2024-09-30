package serverDomain

import (
	"github.com/deeptest-com/deeptest/pkg/domain"
)

type ProcessorDataUploadReq struct {
	_domain.Model

	Name string `json:"name"`
	Desc string `json:"desc"`
	Data string `json:"data"`

	ProjectId uint `json:"projectId"`
}

type ProcessorDataUploadResp struct {
	Path string `json:"path"`
	Data string `json:"data"`
}

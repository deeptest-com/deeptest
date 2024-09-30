package serverDomain

import (
	"github.com/deeptest-com/deeptest/pkg/domain"
)

type DbConnReq struct {
	_domain.Model

	Name string `json:"name"`

	CreateUser string `json:"createUser"`
	UpdateUser string `json:"updateUser"`

	ProjectId uint `json:"projectId"`
}

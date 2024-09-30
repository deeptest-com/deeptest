package serverDomain

import (
	"github.com/deeptest-com/deeptest/pkg/domain"
)

type AgentReq struct {
	_domain.Model

	Name string `json:"name"`

	CreateUser string `json:"createUser"`
	UpdateUser string `json:"updateUser"`
}

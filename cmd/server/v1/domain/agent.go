package serverDomain

import (
	"github.com/aaronchen2k/deeptest/pkg/domain"
)

type AgentReq struct {
	_domain.Model

	Name string `json:"name"`

	CreateUser string `json:"createUser"`
	UpdateUser string `json:"updateUser"`
}

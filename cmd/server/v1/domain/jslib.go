package serverDomain

import (
	"github.com/deeptest-com/deeptest/pkg/domain"
)

type JslibReq struct {
	_domain.Model

	Name       string `json:"name"`
	ScriptFile string `json:"scriptFile"`
	TypesFile  string `json:"typesFile"`

	CreateUser string `json:"createUser"`
	UpdateUser string `json:"updateUser"`

	ProjectId uint `json:"projectId"`
}

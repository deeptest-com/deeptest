package model

type SysAgent struct {
	BaseModel

	Name string `json:"name" validate:"required"`
	Code string `json:"code" validate:"required"`
	Url  string `json:"url" validate:"required"`
	Desc string `json:"desc"`

	CreateUser string `json:"createUser"`
	UpdateUser string `json:"updateUser"`
}

func (SysAgent) TableName() string {
	return "sys_agent"
}

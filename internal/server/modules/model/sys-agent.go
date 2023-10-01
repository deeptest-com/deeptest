package model

type SysAgent struct {
	BaseModel

	Name string `json:"name" validate:"required"`
	Url  string `json:"url" validate:"required"`
	Desc string `json:"desc"`
}

func (SysAgent) TableName() string {
	return "sys_agent"
}

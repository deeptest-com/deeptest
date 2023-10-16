package model

type SysJslib struct {
	BaseModel

	Name       string `json:"name"`
	ScriptFile string `json:"scriptFile" gorm:"type:text" validate:"required"`
	TypesFile  string `json:"typesFile" gorm:"type:text"`

	ProjectId uint `json:"projectId"`

	CreateUser string `json:"createUser"`
	UpdateUser string `json:"updateUser"`
}

func (SysJslib) TableName() string {
	return "sys_jslib"
}

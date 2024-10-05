package model

type Jslib struct {
	BaseModel

	Name string `json:"name"`

	ScriptFile string `json:"scriptFile" gorm:"type:text" validate:"required"`
	scriptName string `json:"scriptName" gorm:"type:text" validate:"required"`

	TypesFile string `json:"typesFile" gorm:"type:text"`
	typesName string `json:"typesName" gorm:"type:text"`

	ProjectId uint `json:"projectId"`

	CreateUser string `json:"createUser"`
	UpdateUser string `json:"updateUser"`
}

func (Jslib) TableName() string {
	return "biz_jslib"
}

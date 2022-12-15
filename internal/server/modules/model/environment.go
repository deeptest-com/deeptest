package model

type Environment struct {
	BaseModel

	Name string           `json:"name"`
	Vars []EnvironmentVar `gorm:"-" json:"vars"`
}

type EnvironmentVar struct {
	BaseModel

	Name       string `json:"name"`
	RightValue string `json:"rightValue"`

	EnvironmentId uint `json:"environmentId"`
}

func (Environment) TableName() string {
	return "biz_environment"
}
func (EnvironmentVar) TableName() string {
	return "biz_environment_var"
}

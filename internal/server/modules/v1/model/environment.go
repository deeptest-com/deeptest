package model

type Environment struct {
	BaseModel

	Name string `json:"name"`

	Vars      []EnvironmentVar `gorm:"-" json:"vars"`
	ProjectId uint             `json:"projectId"`
}

type EnvironmentVar struct {
	BaseModel

	Name string `json:"name"`
	Val  string `json:"val"`

	EnvironmentId uint `json:"environmentId"`
}

func (Environment) TableName() string {
	return "biz_environment"
}
func (EnvironmentVar) TableName() string {
	return "biz_environment_var"
}

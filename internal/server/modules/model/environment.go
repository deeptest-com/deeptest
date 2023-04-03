package model

type Environment struct {
	BaseModel
	ProjectId    uint             `json:"projectId"`
	Name         string           `json:"name"`
	Vars         []EnvironmentVar `gorm:"-" json:"vars"`
	ServeServers []ServeServer    `gorm:"-" json:"serveServers"`
	Sort         uint             `json:"sort"`
}

type EnvironmentVar struct {
	BaseModel
	Name          string `json:"name"`
	RightValue    string `json:"rightValue"`
	LocalValue    string `json:"localValue"`
	RemoteValue   string `json:"remoteValue"`
	EnvironmentId uint   `json:"environmentId"`
	ProjectId     uint   `json:"projectId"`
}

type EnvironmentParam struct {
	BaseModel
	Name         string `json:"name"`
	Type         string `json:"type"`
	Required     bool   `json:"required"`
	DefaultValue string `json:"defaultValue"`
	Description  string `json:"description"`
	In           string `json:"in"`
	ProjectId    uint   `json:"projectId"`
}

func (Environment) TableName() string {
	return "biz_environment"
}
func (EnvironmentVar) TableName() string {
	return "biz_environment_var"
}

func (EnvironmentParam) TableName() string {
	return "biz_environment_param"
}

package model

type ServeServer struct {
	BaseModel
	EnvironmentId   uint   `json:"environmentId" gorm:"uniqueIndex:idx_environment_id_serve_id"`
	ServeId         uint   `json:"serveId" gorm:"uniqueIndex:idx_environment_id_serve_id"`
	Url             string `json:"url"`
	Description     string `json:"description"`
	ServeName       string `gorm:"-" json:"serveName"`
	EnvironmentName string `gorm:"-" json:"environmentName"`
}

func (ServeServer) TableName() string {
	return "biz_project_serve_server"
}

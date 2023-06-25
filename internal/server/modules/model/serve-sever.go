package model

type ServeServer struct {
	BaseModel

	ServeId       uint `json:"serveId" gorm:"uniqueIndex:idx_environment_id_serve_id"`
	EnvironmentId uint `json:"environmentId" gorm:"uniqueIndex:idx_environment_id_serve_id"`

	Url             string `json:"url"`
	Description     string `json:"description"`
	ServeName       string `gorm:"-" json:"serveName"`
	EnvironmentName string `gorm:"-" json:"environmentName"`
	Sort            uint   `gorm:"-" json:"-"`
}

func (ServeServer) TableName() string {
	return "biz_project_serve_server"
}

type ServeServerArr []ServeServer

func (a ServeServerArr) Len() int {
	return len(a)
}
func (a ServeServerArr) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a ServeServerArr) Less(i, j int) bool {
	return a[i].Sort < a[j].Sort
}

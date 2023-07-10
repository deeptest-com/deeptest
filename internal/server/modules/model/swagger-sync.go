package model

type SwaggerSync struct {
	BaseModel
	Switch     bool   `json:"switch"`
	SyncType   uint   `json:"syncType"`
	CategoryId uint   `json:"categoryId"`
	Url        string `json:"url"`
	Cron       string `json:"cron"`
}

func (SwaggerSync) TableName() string {
	return "biz_project_serve_swagger_sync"
}

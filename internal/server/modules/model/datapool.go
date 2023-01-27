package model

type Datapool struct {
	BaseModel

	Name string `json:"name"`
	Desc string `json:"desc"`

	Path string `json:"path"`
	Data string `json:"data"`

	ProjectId uint `json:"projectId"`
}

func (Datapool) TableName() string {
	return "biz_datapool"
}

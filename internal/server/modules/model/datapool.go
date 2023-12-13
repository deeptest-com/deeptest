package model

type Datapool struct {
	BaseModel

	Name string `json:"name"`
	Desc string `gorm:"type:text" json:"desc"`

	Path string `json:"path"`
	Data string `gorm:"type:longtext" json:"data"`

	CreateUser string `json:"createUser"`
	ProjectId  uint   `json:"projectId"`
}

func (Datapool) TableName() string {
	return "biz_datapool"
}

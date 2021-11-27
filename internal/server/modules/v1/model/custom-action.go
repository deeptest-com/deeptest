package model

type CustomAction struct {
	BaseModel

	Name string `json:"name"`
	Desc string `json:"desc" gorm:"column:descr"`

	StatusFrom string `json:"statusFrom"`
	StatusTo   string `json:"statusTo"`

	StatusFromId uint `json:"statusFromId"`
	StatusToId   uint `json:"statusToId"`

	WorkitemId uint `json:"workitemId"`
	PageId     uint `json:"pageId"`
	OrgId      uint `json:"orgId" gorm:"comment:'所属组织'"`
}

func (CustomAction) TableName() string {
	return "custom_action"
}

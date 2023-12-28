package model

type Org struct {
	BaseModel

	Name string `json:"name"`
	Desc string `json:"desc" gorm:"column:descr;type:text"`
}

func (Org) TableName() string {
	return "biz_org"
}

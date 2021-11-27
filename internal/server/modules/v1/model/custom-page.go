package model

type CustomPage struct {
	BaseModel

	Name string `json:"name"`
	Desc string `json:"desc" gorm:"column:descr"`

	WorkitemId uint               `json:"workitemId"`
	Fields     []*CustomPageField `json:"fields" gorm:"foreignKey:page_id"`
}

func (CustomPage) TableName() string {
	return "custom_page"
}

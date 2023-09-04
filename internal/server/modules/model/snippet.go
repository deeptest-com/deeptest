package model

type Snippet struct {
	BaseModel

	Name   string `json:"name"`
	Desc   string `json:"desc"`
	Script string `json:"script" gorm:"type:text"`
}

func (Snippet) TableName() string {
	return "biz_snippet"
}

package model

type Snippet struct {
	BaseModel

	Name   string `json:"name"`
	Desc   string `json:"desc"`
	Script string `json:"script"`
}

func (Snippet) TableName() string {
	return "biz_snippet"
}

package model

import ()

type Label struct {
	BaseModel
	Name string `json:"name"`
}

func (Label) TableName() string {
	return "biz_label"
}

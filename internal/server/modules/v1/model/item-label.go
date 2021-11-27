package model

import ()

type ItemLabel struct {
	BaseModel

	LabelId  uint   `json:"labelId"`
	ItemId   uint   `json:"itemId"`
	ItemType string `json:"itemType"`
}

func (ItemLabel) TableName() string {
	return "item_label"
}

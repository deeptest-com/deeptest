package model

import ()

type ItemLink struct {
	BaseModel

	Source      uint   `json:"source"`
	Destination uint   `json:"destination"`
	LinkType    string `json:"linkType"`
	ItemType    uint   `json:"itemType"` // refer to ItemLinkType
}

func (ItemLink) TableName() string {
	return "item_link"
}

package model

import ()

type ItemLinkType struct {
	BaseModel

	Name    string `json:"name"`
	InWard  string `json:"inWard"`
	OutWard string `json:"outWard"`
	Style   string `json:"style"`
}

func (ItemLinkType) TableName() string {
	return "item_link_type"
}

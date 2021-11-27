package model

import ()

type AttachmentItem struct {
	BaseModel

	AttachmentId uint   `json:"attachmentId"`
	ItemId       uint   `json:"itemId"`
	ItemType     string `json:"itemType"`
}

func (AttachmentItem) TableName() string {
	return "biz_attachment_item"
}

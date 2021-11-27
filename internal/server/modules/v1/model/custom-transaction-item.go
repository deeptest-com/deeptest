package model

import ()

type CustomTransactionItem struct {
	BaseModel

	Name string `json:"name"`
	Desc string `json:"desc" gorm:"column:descr"`

	WorkitemId   uint `json:"workitemId"`
	SrcStatusId  uint `json:"srcStatusId"`
	DistStatusId uint `json:"distStatusId"`

	TransactionId uint `json:"transactionId"`
	OrgId         uint `json:"orgId"`
}

func (CustomTransactionItem) TableName() string {
	return "custom_transaction_item"
}

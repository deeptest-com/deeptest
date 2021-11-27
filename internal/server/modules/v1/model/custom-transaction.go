package model

import ()

type CustomTransaction struct {
	BaseModel

	Name string `json:"name"`
	Desc string `json:"desc" gorm:"column:descr"`

	Items      []*CustomTransactionItem `json:"items" gorm:"foreignKey:transaction_id"`
	WorkitemId uint                     `json:"workitemId"`
	OrgId      uint                     `json:"orgId"`
}

func (CustomTransaction) TableName() string {
	return "custom_transaction"
}

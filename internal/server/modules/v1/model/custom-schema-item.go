package model

import ()

type CustomSchemaItem struct {
	BaseModel

	Name string `json:"name"`
	Desc string `json:"desc"`

	FieldId       uint `json:"fieldId"`
	TransactionId uint `json:"transactionId"`

	SchemaId uint `json:"schemaId"`
	OrgId    uint `json:"orgId"`
}

func (CustomSchemaItem) TableName() string {
	return "custom_schema_item"
}

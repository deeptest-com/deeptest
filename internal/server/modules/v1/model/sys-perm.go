package model

import ()

type Permission struct {
	BaseModel
	BasePermission
}

type BasePermission struct {
	Name        string `gorm:"index:perm_index,unique;not null ;type:varchar(256)" json:"name" validate:"required,gte=4,lte=50"`
	DisplayName string `gorm:"type:varchar(256)" json:"displayName"`
	Description string `gorm:"type:varchar(256)" json:"description"`
	Act         string `gorm:"index:perm_index;type:varchar(256)" json:"act" validate:"required"`
}

func (Permission) TableName() string {
	return "sys_permission"
}

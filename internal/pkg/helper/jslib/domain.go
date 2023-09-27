package jslibHelper

import "time"

type Jslib struct {
	Id uint `json:"id"`

	Name   string `json:"name"`
	Desc   string `json:"desc"`
	Script string `json:"script" gorm:"type:text"`

	UpdatedAt time.Time `json:"updatedAt"`
}

type SysJslib struct {
	ID        uint       `gorm:"primary_key" sql:"type:INT(10) UNSIGNED NOT NULL" json:"id"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`

	Deleted  bool `json:"-" gorm:"default:false"`
	Disabled bool `json:"disabled,omitempty" gorm:"default:false"`

	Name       string `json:"name"`
	ScriptFile string `json:"scriptFile" gorm:"type:text" validate:"required"`
	TypesFile  string `json:"typesFile" gorm:"type:text"`

	CreateUser string `json:"createUser"`
	UpdateUser string `json:"updateUser"`
}

func (SysJslib) TableName() string {
	return "sys_jslib"
}

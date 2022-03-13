package model

import "github.com/kataras/iris/v12"

type TestInterface struct {
	BaseModel

	Name string `json:"name"`
	Desc string `json:"desc"`

	IsDir     bool `json:"isDir"`
	ParentId  uint `json:"parentId"`
	ProjectId uint `json:"projectId"`
	UseID     uint `json:"useId"`

	Ordr     int              `json:"ordr"`
	Children []*TestInterface `gorm:"-" json:"children"`

	Slots iris.Map `gorm:"-" json:"slots"`

	// config
	Method            string `gorm:"default:GET" json:"method"`
	Body              string `gorm:"default:{}" json:"body"`
	BodyType          string `gorm:"default:json" json:"bodyType"`
	AuthorizationType string `gorm:"default:''" json:"authorizationType"`
	PreRequestScript  string `gorm:"default:''" json:"preRequestScript"`
	ValidationScript  string `gorm:"default:''" json:"validationScript"`
}

func (TestInterface) TableName() string {
	return "biz_test_interface"
}

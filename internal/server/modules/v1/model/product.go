package model

import (
	"github.com/aaronchen2k/deeptest/internal/comm/consts"
)

type Product struct {
	BaseModel

	Category string               `json:"category"`
	Name     string               `json:"name"`
	Desc     string               `json:"desc" gorm:"column:descr"`
	Status   consts.ProductStatus `json:"status" gorm:"default:active"`

	ParentId uint `json:"parentId" gorm:"default:0"`
	OrgId    uint `json:"orgId"`
	//Org      Org `gorm:"foreignKey:org_id"`

	Children []*Product `json:"children" gorm:"foreignKey:parent_id"`
	Projects []*Project `json:"projects" gorm:"many2many:r_project_product;"`
}

func (Product) TableName() string {
	return "biz_product"
}

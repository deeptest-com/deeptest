package model

import serverDomain "github.com/aaronchen2k/deeptest/internal/server/modules/v1/domain"

type Project struct {
	BaseModel

	serverDomain.ProjectBase
	//Products []*Product `json:"products" gorm:"many2many:biz_project_product_r;"`
}

func (Project) TableName() string {
	return "biz_project"
}

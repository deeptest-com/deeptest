package model

type ProjectProductRel struct {
	BaseModel

	ProjectId uint `json:"projectId"`
	ProductId uint `json:"productId"`
}

func (ProjectProductRel) TableName() string {
	return "biz_integration_project_product_rel"
}

type ProjectSpaceRel struct {
	BaseModel

	ProjectId uint   `json:"projectId"`
	SpaceCode string `json:"spaceCode"`
}

func (ProjectSpaceRel) TableName() string {
	return "biz_integration_project_space_rel"
}

type ProjectEngineeringRel struct {
	BaseModel

	ProjectId uint   `json:"projectId"`
	Code      string `json:"code"`
	Name      string `json:"name"`
}

func (ProjectEngineeringRel) TableName() string {
	return "biz_integration_project_engineering_rel"
}

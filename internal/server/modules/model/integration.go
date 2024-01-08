package model

type ProjectProductRel struct {
	BaseModel

	ProjectShortName string `json:"projectShortName"`
	ProductCode      string `json:"productCode"`
}

func (ProjectProductRel) TableName() string {
	return "biz_integration_project_product_rel"
}

type ProjectSpaceRel struct {
	BaseModel

	ProjectShortName string `json:"projectShortName"`
	SpaceCode        string `json:"spaceCode"`
}

func (ProjectSpaceRel) TableName() string {
	return "biz_integration_project_space_rel"
}

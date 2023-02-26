package model

type Plan struct {
	BaseModel

	Version float64 `json:"version" yaml:"version"`
	Name    string  `json:"name" yaml:"name"`
	Desc    string  `json:"desc" yaml:"desc"`

	CategoryId uint `json:"categoryId"`
	ProjectId  uint `json:"projectId"`
}

func (Plan) TableName() string {
	return "biz_plan"
}

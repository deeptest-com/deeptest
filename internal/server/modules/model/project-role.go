package model

type ProjectRole struct {
	BaseModel

	Name string `json:"name"`
	Desc string `json:"desc" gorm:"column:descr"`
}

func (ProjectRole) TableName() string {
	return "biz_project_role"
}

package model

type ProjectRoleMenu struct {
	BaseModel
	ProjectRoleId          uint   `gorm:"index:index_project_role_menu,unique;not null" json:"project_role_id"`
	ProjectMenuName        string `gorm:"index:index_project_role_menu,unique;not null;type:varchar(100)" json:"project_menu_name"`
	ProjectMenuDescription string `gorm:"type:varchar(100)" json:"project_menu_description"`
}

func (ProjectRoleMenu) TableName() string {
	return "biz_project_role_menu"
}

package model

type ProjectRoleMenu struct {
	BaseModel
	RoleId uint `gorm:"index:index_role_menu_id,unique;not null" json:"role_id"`
	MenuId uint `gorm:"index:index_role_menu_id,unique;not null" json:"menu_id"`
}

func (ProjectRoleMenu) TableName() string {
	return "biz_project_role_menu"
}

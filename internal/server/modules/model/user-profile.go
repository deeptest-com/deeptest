package model

type SysUserProfile struct {
	BaseModel

	Phone         string `json:"phone"`
	CurrProjectId uint   `json:"currProjectId"`

	UserId uint `json:"userId"`
}

func (SysUserProfile) TableName() string {
	return "sys_user_profile"
}

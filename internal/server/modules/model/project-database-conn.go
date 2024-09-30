package model

import "github.com/deeptest-com/deeptest/internal/pkg/consts"

type DatabaseConn struct {
	BaseModel

	Name string              `json:"name"`
	Type consts.DatabaseType `json:"type"`
	Desc string              `gorm:"type:text" json:"desc"`

	Host     string `json:"host"`
	Port     string `json:"port"`
	DbName   string `json:"dbName"`
	Username string `json:"username"`
	Password string `json:"password"`

	EnvironmentId uint `json:"environmentId"`
	ProjectId     uint `json:"projectId"`

	CreateUser string `json:"createUser"`
	UpdateUser string `json:"updateUser"`
}

func (DatabaseConn) TableName() string {
	return "biz_database_conn"
}

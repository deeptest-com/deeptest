package model

import "github.com/aaronchen2k/deeptest/internal/pkg/consts"

type DatabaseConn struct {
	BaseModel

	Name string              `json:"name"`
	Type consts.DatabaseType `json:"name"`

	Host     string `json:"host"`
	Port     string `json:"port"`
	DbName   string `json:"dbName"`
	Username string `json:"username"`
	Password string `json:"password"`

	EnvironmentId uint `json:"environmentId"`
	ProjectId     uint `json:"projectId"`
}

func (DatabaseConn) TableName() string {
	return "biz_database_conn"
}

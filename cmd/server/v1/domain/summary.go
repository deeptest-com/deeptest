package domain

import (
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
)

type SummaryReqPaginate struct {
	_domain.PaginateReq
	ProjectId int64  `json:"projectId"`
	Name      string `json:"name"`
}

type ReqSummaryBugs struct {
	ProjectId     int64  `json:"project_id" validate:"required"`
	BugId         string `gorm:"type:text" json:"bug_id"`
	Source        string `gorm:"type:text" json:"source"`
	BugSeverity   string `gorm:"type:text" json:"bug_severity"`
	BugCreateDate string `gorm:"type:text" json:"bug_create_date"`
	BugClassify   string `gorm:"type:text" json:"bug_classify"`
	ProductId     string `gorm:"type:text" json:"product_id"`
	ProductName   string `gorm:"type:text" json:"product_name"`
	BugState      string `gorm:"type:text" json:"bug_state"`
}

type ReqProjectId struct {
	ProjectId int64 `json:"projectId" param:"projectId"`
}

type ResProjectId struct {
	ProjectId int64 `json:"projectId" param:"projectId"`
}

type ResSummaryBugs struct {
	Total    int64   `gorm:"default:0" json:"total"`
	Critical float64 `gorm:"default:0" json:"critical"`
	Blocker  float64 `gorm:"default:0" json:"blocker"`
	Deadly   float64 `gorm:"default:0" json:"deadly"`
	Major    float64 `gorm:"default:0" json:"major"`
	Minor    float64 `gorm:"default:0" json:"minor"`
}

type ResSummaryCard struct {
	ProjectTotal   int64   `gorm:"default:0" json:"project_total"`
	InterfaceTotal int64   `gorm:"default:0" json:"interface_total"`
	ScenarioTotal  int64   `gorm:"default:0" json:"scenario_total"`
	ExecTotal      int64   `gorm:"default:0" json:"exec_total"`
	PassRate       float64 `gorm:"default:0" json:"pass_rate"`
	Coverage       float64 `gorm:"default:0" json:"coverage"`
	InterfaceHb    float64 `gorm:"default:0" json:"interface_hb"`
	ScenarioHb     float64 `gorm:"default:0" json:"scenario_hb"`
	CoverageHb     float64 `gorm:"default:0" json:"coverage_hb"`
}

type ResSummaryDetail struct {
	_domain.PaginateReq
	ProjectTotal int64               `gorm:"default:0" json:"project_total"`
	ProjectList  []ResSummaryDetails `json:"project_list"`
}

type ResSummaryDetails struct {
	Id                 int64         `gorm:"default:0" json:"id"`
	ProjectId          int64         `gorm:"default:0" json:"project_id"`
	ProjectName        int64         `gorm:"default:0" json:"project_name"`
	ProjectDes         int64         `gorm:"default:0" json:"project_des"`
	ProjectChineseName int64         `gorm:"default:0" json:"project_chinese_name"`
	Enable             int64         `gorm:"default:0" json:"enable"`
	AdminUser          int64         `gorm:"default:0" json:"admin_user"`
	InterfaceTotal     int64         `gorm:"default:0" json:"interface_total"`
	ScenarioTotal      int64         `gorm:"default:0" json:"scenario_total"`
	CreateTime         int64         `gorm:"default:0" json:"create_time"`
	PassRate           int64         `gorm:"default:0" json:"pass_rate"`
	ExecTotal          int64         `gorm:"default:0" json:"exec_total"`
	Coverage           int64         `gorm:"default:0" json:"coverage"`
	BugTotal           int64         `gorm:"default:0" json:"bug_total"`
	UserList           []ReqUserList `json:"user_list"`
}

type ReqUserList struct {
	UserId   int64  `json:"user_id"`
	UserName string `gorm:"type:text" json:"user_name"`
}

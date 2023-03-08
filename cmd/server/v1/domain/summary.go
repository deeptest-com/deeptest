package domain

import (
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
)

type SummaryReqPaginate struct {
	_domain.PaginateReq
	ProjectId int64  `json:"projectId"`
	Name      string `json:"name"`
}

type SummaryBugsReq struct {
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
	ProjectId uint `json:"projectId" param:"projectId"`
}

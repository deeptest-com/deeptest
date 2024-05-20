package serverDomain

import (
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
)

type SummaryReqPaginate struct {
	_domain.PaginateReq
	ProjectId int64  `json:"projectId"`
	Name      string `json:"name"`
}

type ReqSummaryBugs struct {
	ProjectId     int64  `json:"projectId"`
	BugId         string `gorm:"type:text" json:"bugId"`
	Source        string `gorm:"type:text" json:"source"`
	BugSeverity   string `gorm:"type:text" json:"bugSeverity"`
	BugCreateDate string `gorm:"type:text" json:"bugCreatedAt"`
	BugClassify   string `gorm:"type:text" json:"bugClassify"`
	BugState      string `gorm:"type:text" json:"bugState"`
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
	Suggest  float64 `gorm:"default:0" json:"suggest"`
}

type ResSummaryCard struct {
	ProjectTotal   int64   `gorm:"default:0" json:"projectTotal"`
	InterfaceTotal int64   `gorm:"default:0" json:"interfaceTotal"`
	ScenarioTotal  int64   `gorm:"default:0" json:"scenarioTotal"`
	ExecTotal      int64   `gorm:"default:0" json:"execTotal"`
	UserTotal      int64   `gorm:"default:0" json:"userTotal"`
	PassRate       float64 `gorm:"default:0" json:"passRate"`
	Coverage       float64 `gorm:"default:0" json:"coverage"`
	InterfaceHb    float64 `gorm:"default:0" json:"interfaceHb"`
	ScenarioHb     float64 `gorm:"default:0" json:"scenarioHb"`
	CoverageHb     float64 `gorm:"default:0" json:"coverageHb"`
}

type ResSummaryDetail struct {
	UserProjectTotal int64               `gorm:"default:0" json:"userProjectTotal"`
	ProjectTotal     int64               `gorm:"default:0" json:"projectTotal"`
	UserProjectList  []ResSummaryDetails `json:"userProjectList"`
	ProjectList      []ResSummaryDetails `json:"projectList"`
}

type ResSummaryDetails struct {
	Id               uint               `gorm:"default:0" json:"id"`
	ProjectId        int64              `gorm:"default:0" json:"projectId"`
	ProjectName      string             `gorm:"default:" json:"projectName"`
	ProjectDescr     string             `gorm:"default:" json:"projectDescr"`
	ProjectShortName string             `gorm:"default:" json:"projectShortName"`
	ScenarioTotal    int64              `gorm:"default:0" json:"scenarioTotal"`
	InterfaceTotal   int64              `gorm:"default:0" json:"interfaceTotal"`
	ExecTotal        int64              `gorm:"default:0" json:"execTotal"`
	PassRate         float64            `gorm:"default:0" json:"passRate"`
	Coverage         float64            `gorm:"default:0" json:"coverage"`
	Disabled         bool               `gorm:"default:false" json:"disabled"`
	AdminId          int64              `gorm:"default:0" json:"adminId"`
	AdminName        string             `gorm:"default:" json:"adminName"`
	CreatedAt        string             `gorm:"default:" json:"createdAt"`
	BugTotal         int64              `gorm:"default:0" json:"bugTotal"`
	UserList         []ResUserIdAndName `json:"userList"`
	Accessible       int                `json:"accessible"`
	Products         []uint             `json:"products"`
	Spaces           []string           `json:"spaces"`
}

type ResUserIdAndName struct {
	UserId   int64  `json:"userId"`
	UserName string `json:"userName"`
}

type ResRankingList struct {
	UserRankingList []ResUserRanking `json:"userRankingList"`
}

type ResUserRanking struct {
	Sort          int64  `gorm:"default:0" json:"sort"`
	UserId        int64  `gorm:"default:0" json:"userId"`
	UserName      string `json:"userName"`
	ScenarioTotal int64  `gorm:"default:0" json:"scenarioTotal"`
	TestCaseTotal int64  `gorm:"default:0" json:"testCaseTotal"`
	Hb            int64  `gorm:"default:0" json:"hb"`
	UpdatedAt     string `json:"updatedAt"`
}
